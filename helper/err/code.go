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

func ErrorDifficultyCheck(thisError error) int {
	if errors.Is(thisError, ErrDifficultyNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

func ErrorCreateCourse(thisError error) int {
	if errors.Is(thisError, ErrTitleEmpty) || errors.Is(thisError, ErrCategoryIdEmpty) || errors.Is(thisError, ErrTeacherIdEmpty) || errors.Is(thisError, ErrTeacherNotFound) || errors.Is(thisError, ErrCourseNotFound) {
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

func ErrorAddModulesCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) || errors.Is(thisError, ErrIdCourse) || errors.Is(thisError, ErrTitleEmpty) || errors.Is(thisError, ErrOrderEmpty) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorAddReadingsCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) || errors.Is(thisError, ErrIdModule) || errors.Is(thisError, ErrTitleEmpty) || errors.Is(thisError, ErrOrderEmpty) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorUpdateModulesCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) || errors.Is(thisError, ErrIdCourse) || errors.Is(thisError, ErrTitleEmpty) || errors.Is(thisError, ErrOrderEmpty) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorUpdateEnrollmentCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) || errors.Is(thisError, ErrIdCourse) || errors.Is(thisError, ErrIdStudent) || errors.Is(thisError, ErrRatingEmpty) || errors.Is(thisError, ErrReviewEmpty) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorGetByCourseIdModulesCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorGetByCourseIdEnrollCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorDeleteModulesCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) {
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

func ErrorGetAllCourse(thisError error) int {
	if errors.Is(thisError, ErrCourseNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

func ErrorGetCourseById(thisError error) (int, string) {
	if errors.Is(thisError, ErrIdEmpty) {
		return http.StatusBadRequest, "error request"
	}
	if errors.Is(thisError, ErrConvertId) {
		return http.StatusServiceUnavailable, "error convert param"
	}
	if errors.Is(thisError, ErrCourseNotFound) {
		return http.StatusServiceUnavailable, "error data not found"
	}
	return http.StatusInternalServerError, "server error"
}

func ErrorUpdateCourseCheck(thisError error) (int, string) {
	if errors.Is(thisError, ErrTitleEmpty) || errors.Is(thisError, ErrCategoryIdEmpty) || errors.Is(thisError, ErrTeacherIdEmpty) || errors.Is(thisError, ErrTeacherNotFound) || errors.Is(thisError, ErrCourseNotFound) {
		return http.StatusBadRequest, "error request"
	}
	if errors.Is(thisError, ErrCourseNotFound) {
		return http.StatusServiceUnavailable, "error database"
	}
	return http.StatusInternalServerError, "server error"
}

func ErrorDeleteCourse(thisError error) (int, string) {
	if errors.Is(thisError, ErrIdEmpty) {
		return http.StatusBadRequest, "error request"
	}
	if errors.Is(thisError, ErrCourseNotFound) {
		return http.StatusServiceUnavailable, "service unavaliable"
	}
	return http.StatusInternalServerError, "internal server error"
}

func ErrorUpdateReadingsCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) || errors.Is(thisError, ErrIdModule) || errors.Is(thisError, ErrTitleEmpty) || errors.Is(thisError, ErrOrderEmpty) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
func ErrorGetByModuleIdReadingsCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorDeleteReadingsCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorReadingsCheck(thisError error) int {
	if errors.Is(thisError, ErrReadingsNotFound) {
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

func ErrorGetCourseByStudentId(thisError error) (int, string) {
	if errors.Is(thisError, ErrCourseNotFound) {
		return http.StatusServiceUnavailable, "error data not found"
	}
	return http.StatusInternalServerError, "server error"
}

func ErrorGetCourseByTeacherId(thisError error) (int, string) {
	if errors.Is(thisError, ErrCourseNotFound) {
		return http.StatusServiceUnavailable, "error data not found"
	}
	return http.StatusInternalServerError, "server error"
}

func ErrorAddVideosCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) || errors.Is(thisError, ErrIdModule) || errors.Is(thisError, ErrTitleEmpty) || errors.Is(thisError, ErrOrderEmpty) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorUpdateVideosCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) || errors.Is(thisError, ErrIdModule) || errors.Is(thisError, ErrTitleEmpty) || errors.Is(thisError, ErrOrderEmpty) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
func ErrorGetByModuleIdVideosCheck(thisError error) int {
	if errors.Is(thisError, ErrNotFound) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
