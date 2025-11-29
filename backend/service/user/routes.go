package user

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/Jeno7u/studybud/config"
	"github.com/Jeno7u/studybud/service/auth"
	"github.com/Jeno7u/studybud/types"
	"github.com/Jeno7u/studybud/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store        types.UserStore
	dataProvider utils.DataProvider
}

func NewHandler(store types.UserStore, dataProvider utils.DataProvider) *Handler {
	return &Handler{store: store, dataProvider: dataProvider}
}

func (h *Handler) RegisterRoutes(router gin.IRouter) {
	// Аутентификация
	router.POST("/login", h.handleLogin)
	router.POST("/register", h.handleRegister)
	router.GET("/me", auth.AuthMiddleware(), h.handleMe)

	// Тесты
	router.GET("/get_tests", h.getTests)

	// Файлы - защищенные эндпоинты
	files := router.Group("/files", auth.AuthMiddleware())
	{
		// Домашние задания
		files.POST("/homework/:id/submit", h.submitHomework)
		files.GET("/homework/:id/submissions", h.getHomeworkSubmissions)
	}

	// Курсы - защищенные эндпоинты
	courses := router.Group("/courses", auth.AuthMiddleware())
	{
		courses.POST("/:id/image", h.uploadCourseImage)
		courses.GET("/:id/image", h.getCourseImage)
	}

	// Уроки - защищенные эндпоинты
	lessons := router.Group("/lessons", auth.AuthMiddleware())
	{
		lessons.POST("/:id/image", h.uploadLessonImage)
		lessons.POST("/:id/materials", h.uploadLessonMaterial)
		lessons.GET("/:id/materials", h.getLessonMaterials)
	}

	// Общие файловые операции
	router.GET("/files/:id/download", h.downloadFile)
	router.DELETE("/files/:id", auth.AuthMiddleware(), h.deleteFile)
}

func (h *Handler) handleLogin(c *gin.Context) {
	// get JSON payload
	var payload types.LoginUserPayload
	if err := c.BindJSON(&payload); err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("error during JSON converting, %v", err))
		return
	}

	// validate the payload
	if err := utils.Vaildate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// check if the user exists
	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("not found, invalid email or passowrd"))
		return
	}

	if !auth.ComparePassowrds([]byte(u.Password), []byte(payload.Password)) {
		log.Println(string(u.Password), string(payload.Password))
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("not found, invalid email or passowrd"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"auth_token",
		token, 
		86400,     
		"/",       
		"",       
		false,  
		true,    
	)

  	c.JSON(http.StatusCreated, gin.H{"status": "ok"})
}

func (h *Handler) handleRegister(c *gin.Context) {
	// get JSON payload
	var payload types.RegisterUserPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Vaildate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// check if the user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	// if it doesnt we create the new user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(c.Writer, http.StatusCreated, map[string]string{"token": token})

	c.JSON(http.StatusCreated, gin.H{"status": "ok"})
}

func (h *Handler) getTests(c *gin.Context) {
	tests, err := h.store.GetTests()
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(c.Writer, http.StatusOK, tests)
}

func (h *Handler) submitHomework(c *gin.Context) {
	userID := c.GetInt("user_id")
	homeworkID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid homework ID"))
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("failed to parse form: %v", err))
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("no files provided"))
		return
	}

	var fileIDs []int

	// Сохраняем каждый файл
	for _, file := range files {

		record, err := h.dataProvider.SaveHomeworkSubmission(file, homeworkID, userID)
		if err != nil {
			utils.WriteError(c.Writer, http.StatusInternalServerError, err)
			return
		}
		fileIDs = append(fileIDs, record.ID)
	}

	// Регистрируем отправку в БД
	submission, err := h.dataProvider.CreateHomeworkSubmission(userID, homeworkID, fileIDs)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Homework submitted successfully",
		"submission": submission,
	})
}

func (h *Handler) getHomeworkSubmissions(c *gin.Context) {
	homeworkID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid homework ID"))
		return
	}

	// TODO: Проверить что пользователь имеет доступ к этим данным
	// (учитель курса или администратор)

	submissions, err := h.dataProvider.GetHomeworkSubmissions(homeworkID)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"submissions": submissions,
	})
}

// эндпоинты курсов
func (h *Handler) uploadCourseImage(c *gin.Context) {
	courseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid course ID"))
		return
	}

	// TODO: Проверить что пользователь - учитель этого курса

	file, err := c.FormFile("image")
	if err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("failed to get image file: %v", err))
		return
	}

	// Проверяем что это изображение
	if !isImageFile(file.Filename) {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("only image files are allowed"))
		return
	}

	record, err := h.dataProvider.SaveCourseImage(file, courseID)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Course image uploaded successfully",
		"fileId":  record.ID,
		"url":     fmt.Sprintf("/api/v1/files/%d/download", record.ID),
	})
}

func (h *Handler) getCourseImage(c *gin.Context) {
	courseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid course ID"))
		return
	}

	record, err := h.dataProvider.GetCourseImage(courseID)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusNotFound, fmt.Errorf("course image not found"))
		return
	}

	fileBytes, err := utils.ReadFile(record.Path)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	c.Data(http.StatusOK, "image/jpeg", fileBytes)
}

// эндпоинты уроков
func (h *Handler) uploadLessonImage(c *gin.Context) {
	lessonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid lesson ID"))
		return
	}

	// TODO: Проверить права доступа

	file, err := c.FormFile("image")
	if err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("failed to get image file: %v", err))
		return
	}

	if !isImageFile(file.Filename) {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("only image files are allowed"))
		return
	}

	record, err := h.dataProvider.SaveLessonImage(file, lessonID)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lesson image uploaded successfully",
		"fileId":  record.ID,
	})
}

func (h *Handler) uploadLessonMaterial(c *gin.Context) {
	lessonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid lesson ID"))
		return
	}

	// TODO: Проверить права доступа

	form, err := c.MultipartForm()
	if err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("failed to parse form: %v", err))
		return
	}

	files := form.File["materials"]
	if len(files) == 0 {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("no files provided"))
		return
	}

	var results []gin.H
	for _, file := range files {
		record, err := h.dataProvider.SaveLessonContent(file, lessonID)
		if err != nil {
			utils.WriteError(c.Writer, http.StatusInternalServerError, err)
			return
		}

		results = append(results, gin.H{
			"filename": file.Filename,
			"fileId":   record.ID,
			"type":     record.FileType,
			"url":      fmt.Sprintf("/api/v1/files/%d/download", record.ID),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lesson materials uploaded successfully",
		"files":   results,
	})
}

func (h *Handler) getLessonMaterials(c *gin.Context) {
	lessonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid lesson ID"))
		return
	}

	files, err := h.dataProvider.GetEntityFiles("lesson_content", lessonID)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"materials": files,
	})
}

// общие файлы и эндпоинты
func (h *Handler) downloadFile(c *gin.Context) {
	fileID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid file ID"))
		return
	}

	record, fileBytes, err := h.dataProvider.GetFileForDownload(fileID)
	if err != nil {
		utils.WriteError(c.Writer, http.StatusNotFound, err)
		return
	}

	// Устанавливаем правильный Content-Type и заголовки для скачивания
	contentType := getContentType(record.FileType)
	c.Header("Content-Type", contentType)
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", record.Name))
	c.Header("Content-Length", fmt.Sprintf("%d", len(fileBytes)))

	c.Data(http.StatusOK, contentType, fileBytes)
}

func (h *Handler) deleteFile(c *gin.Context) {
	fileID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.WriteError(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid file ID"))
		return
	}

	// TODO: Проверить что пользователь имеет права на удаление файла

	if err := h.dataProvider.DeleteFile(fileID); err != nil {
		utils.WriteError(c.Writer, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File deleted successfully",
	})
}

func isImageFile(filename string) bool {
	ext := filepath.Ext(filename)
	imageExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
		".webp": true,
	}
	return imageExts[ext]
}

func getContentType(fileType string) string {
	switch fileType {
	case "image", "avatar":
		return "image/jpeg"
	case "pdf":
		return "application/pdf"
	case "document", "text":
		return "application/octet-stream"
	case "video":
		return "video/mp4"
	case "audio":
		return "audio/mpeg"
	case "submission":
		return "application/octet-stream"
	default:
		return "application/octet-stream"
	}
}

// func getPayload(c *gin.Context) (*types.RegisterUserPayload, error) {
// 	// get JSON payload
// 	var payload types.RegisterUserPayload
// 	if err := c.ShouldBindJSON(&payload); err != nil {
// 		utils.WriteError(c.Writer, http.StatusBadRequest, err)
// 		return nil, err

// 	return payload, nil
// }
