package err

import (
	"errors"
	"net/http"
)

func ErrorCategoryCheck(thisError error) int {
	if errors.Is(thisError, ErrCategoryNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

func ErrorCreateCourse(thisError error) int {
	if errors.Is(thisError, ErrTitleEmpty) || errors.Is(thisError, ErrCategoryIdEmpty) || errors.Is(thisError, ErrTeacherIdEmpty) || errors.Is(thisError, ErrTeacherNotFound) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorEnrollmentCheck(thisError error) int {
	if errors.Is(thisError, ErrEnrollmentsNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

func ErrorGetAllCourse(thisError error) int {
	if errors.Is(thisError, ErrCoursesNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

// func ErrorRequest(thisError error) (int, string) {
// 	if errors.Is(thisError, ErrIdEmpty) || errors.Is(thisError, ErrTitleEmpty) || errors.Is(thisError, ErrCategoryIdEmpty) || errors.Is(thisError, ErrTeacherIdEmpty) || errors.Is(thisError, ErrTeacherNotFound) {
// 		return http.StatusBadRequest, "error request"
// 	}
// 	return 0, ""
// }

func ErrorGetCourseById(thisError error) (int, string) {
	if errors.Is(thisError, ErrIdEmpty) {
		return http.StatusBadRequest, "error request"
	}
	if errors.Is(thisError, ErrConvertId) {
		return http.StatusServiceUnavailable, "error convert param"
	}
	return http.StatusInternalServerError, "server error"
}
