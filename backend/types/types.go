package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
	GetTests() ([]Block, error)
}

type CourseStore interface {
	GetCoursesByUserRelatedID(userID int) ([]Course, error)
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
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Students  []int `json:"students"`
	Homeworks []int `json:"homeworks"`
	Steps     []int `json:"steps"`
}

type Step struct {
	ID      int    `json:"id"`
	Courses *[]int `json:"courses"`
	UsersAt *[]int `json:"usersAt"`
}

type Homework struct {
	ID        int       `json:"id"`
	Text      *string   `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	StartsAt  time.Time `json:"startsAt"`
	EndsAt    time.Time `json:"endsAt"`
	Files     *[]int    `json:"files"`
}

type Lesson struct {
	ID        int       `json:"id"`
	Teacher   int       `json:"teacher"`
	Text      *string   `json:"text"`
	Blocks    *[]int    `json:"blocks"`
	CreatedAt time.Time `json:"createdAt"`
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

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
