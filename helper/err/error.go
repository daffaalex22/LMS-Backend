package err

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")

	ErrNotFound = errors.New("data not found")

	ErrCategoryNotFound = errors.New("category not found")

	ErrTeacherNotFound = errors.New("teacher not found")

	ErrTitleEmpty = errors.New("title are empty")

	ErrCategoryIdEmpty = errors.New("category_id are empty")

	ErrTeacherIdEmpty = errors.New("teacher_id are empty")

	ErrEnrollmentsNotFound = errors.New("category not found")

	ErrCoursesNotFound = errors.New("course not found")
)
