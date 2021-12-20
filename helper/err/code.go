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
func ErrorEnrollmentCheck(thisError error) int {
	if errors.Is(thisError, ErrEnrollmentsNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}
