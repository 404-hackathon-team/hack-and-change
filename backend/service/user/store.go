package user

import (
	"database/sql"
	"fmt"

	"github.com/Jeno7u/studybud/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func scanRowIntoUser(rows *sql.Row) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	var u *types.User

	row := s.db.QueryRow("SELECT * FROM users WHERE email = $1", email)

	u, err := scanRowIntoUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return u, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	var u *types.User

	row := s.db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	u, err := scanRowIntoUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return u, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec(`INSERT INTO users (firstName, lastName, email, password)
	VALUES ($1, $2, $3, $4)`, user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetTests() ([]types.Block, error) {
	testsRows, err := s.db.Query("SELECT * FROM blocks WHERE type='quiz'")
	if err != nil {
		return nil, err
	}
	defer testsRows.Close()

	var blocks []types.Block

	for testsRows.Next() {
		var block types.Block
		err := testsRows.Scan(&block.ID, &block.Type, &block.Path)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, block)
	}

	return blocks, nil
}

// Получение урока по ID
func (s *Store) GetLessonByID(lessonID int) (*types.Lesson, error) {
	var lesson types.Lesson
	err := s.db.QueryRow(`
        SELECT id, teacher, text, createdAt, steps, image, name, "usersPassed"
        FROM lessons
        WHERE id = $1
    `, lessonID).Scan(
		&lesson.ID,
		&lesson.Teacher,
		&lesson.Text,
		&lesson.CreatedAt,
		&lesson.Image,
		&lesson.Name,
		&lesson.UsersPassed,
	)

	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

// Получение всех категорий
func (s *Store) GetAllCategories() ([]types.Category, error) {
	rows, err := s.db.Query(`
        SELECT id, lessons, name
        FROM categories
        ORDER BY name
    `)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []types.Category
	for rows.Next() {
		var category types.Category
		err := rows.Scan(&category.ID, &category.Lessons, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (s *Store) GetLessonsByCategory(categoryID int) ([]types.Lesson, error) {
	rows, err := s.db.Query(`
        SELECT id, teacher, text, "createdAt", steps, image, name, "usersPassed"
        FROM lessons 
        WHERE id IN (
            SELECT unnest(lessons) FROM categories WHERE id = $1
        )
        ORDER BY "createdAt" DESC
    `, categoryID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []types.Lesson
	for rows.Next() {
		var lesson types.Lesson
		err := rows.Scan(
			&lesson.ID,
			&lesson.Teacher,
			&lesson.Text,
			&lesson.CreatedAt,
			&lesson.Image,
			&lesson.Name,
			&lesson.UsersPassed,
		)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}

	return lessons, nil
}

// Получение категории по ID
func (s *Store) GetCategoryByID(categoryID int) (*types.Category, error) {
	var category types.Category
	err := s.db.QueryRow(`
        SELECT id, lessons, name
        FROM categories
        WHERE id = $1
    `, categoryID).Scan(&category.ID, &category.Lessons, &category.Name)

	if err != nil {
		return nil, err
	}

	return &category, nil
}
