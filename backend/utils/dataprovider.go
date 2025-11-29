package utils

import (
	"database/sql"
	"fmt"
	"golang.org/x/exp/rand"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type DataProvider struct {
	db          *sql.DB
	storagePath string
}

type FileRecord struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Path       string    `json:"path"`
	FileType   string    `json:"fileType"`
	EntityID   int       `json:"entityId"`
	EntityType string    `json:"entityType"`
	CreatedAt  time.Time `json:"createdAt"`
}

type HomeworkSubmission struct {
	ID         int   `json:"id"`
	UserID     int   `json:"userId"`
	HomeworkID int   `json:"homeworkId"`
	FileIDs    []int `json:"fileIds"`
}

func NewDataProvider(db *sql.DB) *DataProvider {
	path := os.Getenv("DATA_PATH")
	if path == "" {
		path = "./data"
	}

	baseDirs := []string{"courses", "lessons", "homeworks", "users", "teachers", "shared/temp"}
	for _, dir := range baseDirs {
		os.MkdirAll(filepath.Join(path, dir), 0755)
	}

	return &DataProvider{
		db:          db,
		storagePath: path,
	}
}

func generateUniqueFileName(originalName string) string {
	ext := filepath.Ext(originalName)
	name := originalName[:len(originalName)-len(ext)]
	timestamp := time.Now().UnixNano()
	randomPart := fmt.Sprintf("%06d", rand.Intn(1000000))
	return fmt.Sprintf("%s_%d_%s%s", name, timestamp, randomPart, ext)
}

// БАЗОВЫЙ МЕТОД ДЛЯ ВСЕХ ФАЙЛОВ
func (dp *DataProvider) SaveFile(file *multipart.FileHeader, fileType, entityType string, entityID int) (*FileRecord, error) {
	entityPath := dp.getEntityPath(entityType, entityID)
	os.MkdirAll(entityPath, 0755)

	subfolder := dp.getFileSubfolder(fileType)
	finalPath := filepath.Join(entityPath, subfolder)
	os.MkdirAll(finalPath, 0755)

	uniqueName := generateUniqueFileName(file.Filename)
	filePath := filepath.Join(finalPath, uniqueName)

	// Сохраняем файл
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %v", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return nil, fmt.Errorf("failed to save file: %v", err)
	}

	// Регистрируем в БД
	var fileID int
	err = dp.db.QueryRow(`
		INSERT INTO file_pathes (name, path, "fileType", entity_type, entity_id) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id
	`, uniqueName, filePath, fileType, entityType, entityID).Scan(&fileID)

	if err != nil {
		os.Remove(filePath)
		return nil, fmt.Errorf("failed to register file in DB: %v", err)
	}

	return &FileRecord{
		ID:         fileID,
		Name:       uniqueName,
		Path:       filePath,
		FileType:   fileType,
		EntityID:   entityID,
		EntityType: entityType,
		CreatedAt:  time.Now(),
	}, nil
}

// методы для курсов

// Загрузка изображения курса
func (dp *DataProvider) SaveCourseImage(file *multipart.FileHeader, courseID int) (*FileRecord, error) {
	return dp.SaveFile(file, "image", "course", courseID)
}

// Получение изображения курса
func (dp *DataProvider) GetCourseImage(courseID int) (*FileRecord, error) {
	return dp.getFileByEntity("course", courseID, "image")
}

// методы для уроков

// Загрузка изображения урока
func (dp *DataProvider) SaveLessonImage(file *multipart.FileHeader, lessonID int) (*FileRecord, error) {
	return dp.SaveFile(file, "image", "lesson", lessonID)
}

// Загрузка материалов урока (блоков контента)
func (dp *DataProvider) SaveLessonContent(file *multipart.FileHeader, lessonID int) (*FileRecord, error) {
	return dp.SaveFile(file, "document", "lesson_content", lessonID)
}

// Получение изображения урока
func (dp *DataProvider) GetLessonImage(lessonID int) (*FileRecord, error) {
	return dp.getFileByEntity("lesson", lessonID, "image")
}

// методы для домашек

// Загрузка файлов задания (от учителя)
func (dp *DataProvider) SaveHomeworkAssignment(file *multipart.FileHeader, homeworkID int) (*FileRecord, error) {
	return dp.SaveFile(file, "document", "homework_assignment", homeworkID)
}

// Загрузка студенческой работы
func (dp *DataProvider) SaveHomeworkSubmission(file *multipart.FileHeader, homeworkID, userID int) (*FileRecord, error) {
	submissionPath := dp.getHomeworkSubmissionPath(homeworkID, userID)
	os.MkdirAll(submissionPath, 0755)

	uniqueName := generateUniqueFileName(file.Filename)
	filePath := filepath.Join(submissionPath, uniqueName)

	// Сохраняем файл
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return nil, err
	}

	// Регистрируем в БД
	var fileID int
	err = dp.db.QueryRow(`
		INSERT INTO file_pathes (name, path, "fileType", entity_type, entity_id) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id
	`, uniqueName, filePath, "submission", "homework_submission", homeworkID).Scan(&fileID)

	if err != nil {
		os.Remove(filePath)
		return nil, err
	}

	return &FileRecord{
		ID:         fileID,
		Name:       uniqueName,
		Path:       filePath,
		FileType:   "submission",
		EntityID:   homeworkID,
		EntityType: "homework_submission",
		CreatedAt:  time.Now(),
	}, nil
}

// Регистрация отправки домашнего задания
func (dp *DataProvider) CreateHomeworkSubmission(userID, homeworkID int, fileIDs []int) (*HomeworkSubmission, error) {
	var submissionID int

	err := dp.db.QueryRow(`
		INSERT INTO uploaded_homework ("user", homework, files) 
		VALUES ($1, $2, $3) 
		RETURNING id
	`, userID, homeworkID, fileIDs).Scan(&submissionID)

	if err != nil {
		return nil, fmt.Errorf("failed to create homework submission: %v", err)
	}

	return &HomeworkSubmission{
		ID:         submissionID,
		UserID:     userID,
		HomeworkID: homeworkID,
		FileIDs:    fileIDs,
	}, nil
}

// Получение отправленных домашних заданий
func (dp *DataProvider) GetHomeworkSubmissions(homeworkID int) ([]HomeworkSubmission, error) {
	rows, err := dp.db.Query(`
		SELECT id, "user", homework, files 
		FROM uploaded_homework 
		WHERE homework = $1
	`, homeworkID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var submissions []HomeworkSubmission
	for rows.Next() {
		var sub HomeworkSubmission
		if err := rows.Scan(&sub.ID, &sub.UserID, &sub.HomeworkID, &sub.FileIDs); err != nil {
			continue
		}
		submissions = append(submissions, sub)
	}

	return submissions, nil
}

// Загрузка файла для блока (например, изображение/видео в шаге урока)
func (dp *DataProvider) SaveBlockFile(file *multipart.FileHeader, blockID int, fileType string) (*FileRecord, error) {
	return dp.SaveFile(file, fileType, "block", blockID)
}

// Получение файла блока
func (dp *DataProvider) GetBlockFile(blockID int) (*FileRecord, error) {
	return dp.getFileByEntity("block", blockID, "")
}

// вспомогалки

func (dp *DataProvider) getEntityPath(entityType string, entityID int) string {
	return filepath.Join(dp.storagePath, entityType+"s", strconv.Itoa(entityID))
}

func (dp *DataProvider) getFileSubfolder(fileType string) string {
	switch fileType {
	case "image", "avatar":
		return "images"
	case "document", "pdf", "text":
		return "documents"
	case "video":
		return "videos"
	case "audio":
		return "audio"
	case "submission":
		return "submissions"
	default:
		return "files"
	}
}

func (dp *DataProvider) getHomeworkSubmissionPath(homeworkID, userID int) string {
	return filepath.Join(
		dp.storagePath,
		"homeworks",
		strconv.Itoa(homeworkID),
		"submissions",
		strconv.Itoa(userID),
	)
}

// Универсальный метод поиска файла по сущности
func (dp *DataProvider) getFileByEntity(entityType string, entityID int, fileType string) (*FileRecord, error) {
	query := `
		SELECT id, name, path, "fileType", entity_type, entity_id, created_at 
		FROM file_pathes 
		WHERE entity_type = $1 AND entity_id = $2
	`
	args := []interface{}{entityType, entityID}

	if fileType != "" {
		query += " AND \"fileType\" = $3"
		args = append(args, fileType)
	}

	query += " ORDER BY created_at DESC LIMIT 1"

	var record FileRecord
	err := dp.db.QueryRow(query, args...).Scan(
		&record.ID, &record.Name, &record.Path, &record.FileType,
		&record.EntityType, &record.EntityID, &record.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("file not found for %s %d: %v", entityType, entityID, err)
	}

	return &record, nil
}

// Получение всех файлов сущности
func (dp *DataProvider) GetEntityFiles(entityType string, entityID int) ([]FileRecord, error) {
	rows, err := dp.db.Query(`
		SELECT id, name, path, "fileType", entity_type, entity_id, created_at 
		FROM file_pathes 
		WHERE entity_type = $1 AND entity_id = $2
		ORDER BY created_at DESC
	`, entityType, entityID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []FileRecord
	for rows.Next() {
		var file FileRecord
		if err := rows.Scan(&file.ID, &file.Name, &file.Path, &file.FileType,
			&file.EntityType, &file.EntityID, &file.CreatedAt); err != nil {
			continue
		}
		files = append(files, file)
	}

	return files, nil
}

// Удаление файла
func (dp *DataProvider) DeleteFile(fileID int) error {
	// Получаем информацию о файле
	var filePath string
	err := dp.db.QueryRow("SELECT path FROM file_pathes WHERE id = $1", fileID).Scan(&filePath)
	if err != nil {
		return err
	}

	// Удаляем файл с диска
	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete file from disk: %v", err)
	}

	// Удаляем запись из БД
	_, err = dp.db.Exec("DELETE FROM file_pathes WHERE id = $1", fileID)
	return err
}
