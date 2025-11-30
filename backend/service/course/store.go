package course

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	"github.com/Jeno7u/studybud/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func parsePGIntArray(s sql.NullString) ([]int, error) {
	if !s.Valid {
		return nil, nil
	}
	str := strings.TrimSpace(s.String)
	if str == "" || str == "{}" {
		return nil, nil
	}
	// remove surrounding braces if present
	if strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}") {
		str = str[1 : len(str)-1]
	}
	if strings.TrimSpace(str) == "" {
		return nil, nil
	}
	parts := strings.Split(str, ",")
	out := make([]int, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		out = append(out, n)
	}
	return out, nil
}

func scanRowIntoCourse(rows *sql.Rows) (*types.Course, error) {
	var (
		studentsStr  sql.NullString
		homeworksStr sql.NullString
		categoriesStr     sql.NullString
	)

	course := new(types.Course)

	err := rows.Scan(
		&course.ID,
		&course.Name,
		&studentsStr,
		&homeworksStr,
		&categoriesStr,
	)
	if err != nil {
		return nil, err
	}


	if students, err := parsePGIntArray(studentsStr); err == nil {
		course.Students = students
	} else {
		return nil, err
	}

	if homeworks, err := parsePGIntArray(homeworksStr); err == nil {
		course.Homeworks = homeworks
	} else {
		return nil, err
	}

	if categories, err := parsePGIntArray(categoriesStr); err == nil {
		course.Categories = categories
	} else {
		return nil, err
	}

	return course, nil
}

func (s *Store) GetCoursesByUserRelatedID(userID int) ([]types.Course, error) {

	// получение основной информации о курсах принадлежащих пользователю
	rows, err := s.db.Query(`
		SELECT id, name, students, homeworks, categories
		FROM courses
		WHERE $1 = ANY(students);
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []types.Course

	for rows.Next() {
		c, err := scanRowIntoCourse(rows)
		if err != nil {
			return nil, err
		}
		courses = append(courses, *c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// получение статистики по курсу
	var result [](map[string]string )
	for _, course := range courses {
		lessonsAmount, passedAmount, err := s.amountOfLessonsInCourse(result, userID)
		if err != nil {
			return nil, err
		}
		var courseInfo map[string]string = {
			"id": course.id,
			"title": course.name,
			"lessonsDone": passedAmount,
			"lessonsTotal": lessonsAmount,
			"deadlineDate": homeworkDeadline(course.id)
			"notificatinosCount": 2
			"lastUpdate": "20.11.25"
		}
	}
	lessonsAmount, passedAmount, err := s.amountOfLessonsInCourse(result, userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// получение кол-ва уроков в курсе
func (s *Store)  amountOfLessonsInCourse(course type.Course, userID int) (int, int, error) {
	var count int
	var countPassed int

	courseID := course.id

	rows, err := s.db.Query(`
		SELECT categories FROM courses 
		WHERE id == $1
	`, courseID)
	if err != nil {
		return -1, err
	}

	var categoriesStr sql.NullString
	rows.Scan(
		&categoriesStr,
	)
	categories := parsePGIntArray(categoriesStr)

	for _, categorieID := range categories {
		// get lessonsCount
		var lessonsCount int

		err := s.db.QueryRow(`
			SELECT cardinality(lessons) FROM categories 
			WHERE id = $1;
		`, categorieID).Scan(&lessonsCount)
		if err != nil {
			return -1, -1, err
		}

		count += lessonsCount	

		// get passed lessonsPassed
		var passedLessons int

		var lessons []int

		err := s.db.QueryRow(`
			SELECT lessons FROM categories 
			WHERE id == $1 
		`, categorieID).Scan(&lessons)
		if err != nil {
			return -1, -1, err
		}

		for _, lessonID := lessons {
			var passed bool
			s.db.QueryRow(`
				SELECT id IS NOT NULL FROM lessons 
				WHERE id = $1 AND $2 = ANY(usersPassed);
			`, lessonID, userID).Scan(&passed)
			
			countPassed += int(passed)
		}
	}

	return count, countPassed, nil
}

func (s *Store) homeworkDeadline(course type.Course) (string, error) {
	var homeworkID int

	err := s.db.QueryRow(`
		SELECT homeworks FROM courses 
		WHERE id = $1;
	`, course.id).Scan(&homework)
	if err != nil {
		return nil, err
	}
	var deadlineDate string
	err := s.db.QueryRow(`
		SELECT to_char(endsAt, 'DD-MM-YY') FROM homework 
		WHERE id = $1;
	`, homeworkID).Scan(&deadlineDate)
	if err != nil {
		return nil, err
	}

	return deadlineDate, nil
}