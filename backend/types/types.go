package types

import (
	"time"

	"github.com/gorilla/websocket"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(user User) error
	GetTests() ([]Block, error)

	// Новые методы для уроков и категорий
	GetLessonsByCategory(categoryID int) ([]Lesson, error)
	GetCategoryByID(categoryID int) (*Category, error)
	GetLessonByID(lessonID int) (*Lesson, error)
	GetAllCategories() ([]Category, error)
}

type CourseStore interface {
	// returns summarized course info for the user (includes stats)
	GetCoursesByUserRelatedID(userID int) ([]CoursesInfo, error)
}

type NotificationStore interface {

}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type Teacher struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type Message struct {
	ID                int       `json:"id"`
	Text              string    `json:"text"`
	SendedAt          time.Time `json:"sendedAt"`
	SendedTo          int       `json:"sendedTo"`
	IsRecieverTeacher bool      `json:"isRecieverTeacher"`
}

type Score struct {
	ID      int       `json:"id"`
	Score   int       `json:"score"`
	Teacher int       `json:"teacher"`
	Student int       `json:"student"`
	AddedAt time.Time `json:"addedAt"`
}

type Course struct {
	ID        	int    `json:"id"`
	Name      	string `json:"name"`
	Students  	[]int `json:"students"`
	Homework	int `json:"homeworks"`
	Categories     []int `json:"categories"`
}

type Categories struct {
	ID      int    `json:"id"`
	Lessons []int `json:"lessons"`
	Name 	string `json:"name"`
}

type Homework struct {
	ID        int       `json:"id"`
	Text	  string   `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	StartsAt  time.Time `json:"startsAt"`
	EndsAt    time.Time `json:"endsAt"`
	Files     []int    `json:"files"`
}

type Lesson struct {
	ID          int       `json:"id"`
	Teacher     int       `json:"teacher"`
	Text        string   `json:"text"`
	CreatedAt   time.Time `json:"createdAt"`
	Steps      []int    `json:"steps"`
	Image       int      `json:"image"`
	Name        string    `json:"name"`
	UsersPassed []int     `json:"usersPassed"`
}
type Block struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Path *int   `json:"path"` //если блок не нуждается в хранении (например, это ссылка), то это поле - nil
}

type FilePath struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	FileType string `json:"fileType"` //мб следует реализовать что-то вроде enum
}
type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=72"`
}

type Category struct {
	ID      int    `json:"id"`
	Lessons []int  `json:"lessons"`
	Name    string `json:"name"`
}

type CategoryWithLessons struct {
	Category Category `json:"category"`
	Lessons  []Lesson `json:"lessons"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CoursesInfo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	LessonsDone int `json:"lessonsDone"`
	LessonsTotal int `json:"lessonsTotal"`
	DeadlineDate string `json:"deadlineDate"`
	NotificationsCount int `json:"notificationsCount"`
	LastUpdate string `json:"lastUpdate"`
}

type Client struct {
    UserID string
    Conn   *websocket.Conn
}

type Notification struct {
	ID int `json:"id"`
	Type string `json:"type"`
	Date string `json:"date"`
	DayOfWeek string `json:"dayOfWeek"`
	Title string  `json:"title"`
}