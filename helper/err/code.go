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

func ErrorStudentRegisterCheck(thisError error) int {
	if errors.Is(thisError, ErrNameEmpty) || errors.Is(thisError, ErrEmailEmpty) || errors.Is(thisError, ErrPasswordEmpty) || errors.Is(thisError, ErrEmailHasApplied) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorStudentLoginCheck(thisError error) int {
	if errors.Is(thisError, ErrEmailEmpty) || errors.Is(thisError, ErrPasswordEmpty) || errors.Is(thisError, ErrEmailNotExist) || errors.Is(thisError, ErrStudentNotFound) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorStudentUpdateCheck(thisError error) int {
	if errors.Is(thisError, ErrNameEmpty) || errors.Is(thisError, ErrEmailEmpty) || errors.Is(thisError, ErrAvatarEmpty) || errors.Is(thisError, ErrPhoneEmpty) || errors.Is(thisError, ErrAddressEmpty) || errors.Is(thisError, ErrPasswordEmpty) || errors.Is(thisError, ErrConfirmPasswordEmpty) || errors.Is(thisError, ErrValidationPassword) || errors.Is(thisError, ErrNotFound) || errors.Is(thisError, ErrEmailHasApplied) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorStudentProfileCheck(thisError error) int {
	if errors.Is(thisError, ErrEmailEmpty) || errors.Is(thisError, ErrPasswordEmpty) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorTeacherRegisterCheck(thisError error) int {
	if errors.Is(thisError, ErrNameEmpty) || errors.Is(thisError, ErrEmailEmpty) || errors.Is(thisError, ErrPasswordEmpty) || errors.Is(thisError, ErrEmailHasApplied) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorTeacherLoginCheck(thisError error) int {
	if errors.Is(thisError, ErrEmailEmpty) || errors.Is(thisError, ErrPasswordEmpty) || errors.Is(thisError, ErrEmailNotExist) || errors.Is(thisError, ErrTeacherNotFound) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorTeacherUpdateCheck(thisError error) int {
	if errors.Is(thisError, ErrNameEmpty) || errors.Is(thisError, ErrEmailEmpty) || errors.Is(thisError, ErrAvatarEmpty) || errors.Is(thisError, ErrPhoneEmpty) || errors.Is(thisError, ErrAddressEmpty) || errors.Is(thisError, ErrPasswordEmpty) || errors.Is(thisError, ErrConfirmPasswordEmpty) || errors.Is(thisError, ErrValidationPassword) || errors.Is(thisError, ErrNotFound) || errors.Is(thisError, ErrEmailHasApplied) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorTeacherProfileCheck(thisError error) int {
	if errors.Is(thisError, ErrEmailEmpty) || errors.Is(thisError, ErrPasswordEmpty) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorAddEnrollCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) || errors.Is(thisError, ErrIdStudent) || errors.Is(thisError, ErrIdCourse) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorModulesCheck(thisError error) int {
	if errors.Is(thisError, ErrModulesNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

func ErrorEnrollmentCheck(thisError error) int {
	if errors.Is(thisError, ErrEnrollNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

func ErrorModulesCheck(thisError error) int {
	if errors.Is(thisError, ErrModulesNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

func ErrorGetAllCourse(thisError error) int {
	if errors.Is(thisError, ErrCourseNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}
