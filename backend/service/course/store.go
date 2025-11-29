package course

import (
	"database/sql"
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
		stepsStr     sql.NullString
	)

	course := new(types.Course)

	err := rows.Scan(
		&course.ID,
		&course.Name,
		&studentsStr,
		&homeworksStr,
		&stepsStr,
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

	if steps, err := parsePGIntArray(stepsStr); err == nil {
		course.Steps = steps
	} else {
		return nil, err
	}

	return course, nil
}

func (s *Store) GetCoursesByUserRelatedID(userID int) ([]types.Course, error) {

	rows, err := s.db.Query(`
		SELECT id, name, students, homeworks, steps
		FROM courses
		WHERE $1 = ANY(students);
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []types.Course

	for rows.Next() {
		c, err := scanRowIntoCourse(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, *c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
