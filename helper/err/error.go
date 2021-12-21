package err

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")

	//notfound
	ErrNotFound         = errors.New("data not found")
	ErrCategoryNotFound = errors.New("category not found")
	ErrEnrollNotFound   = errors.New("enrollment not found")
	ErrCourseNotFound   = errors.New("course not found")
	ErrTeacherNotFound  = errors.New("teacher not found")
	ErrStudentNotFound  = errors.New("student not found")
	ErrModulesNotFound  = errors.New("modules not found")

	//empty
	ErrNameEmpty            = errors.New("name are empty")
	ErrEmailEmpty           = errors.New("email are empty")
	ErrPasswordEmpty        = errors.New("password are empty")
	ErrConfirmPasswordEmpty = errors.New("confirm password are empty")
	ErrAvatarEmpty          = errors.New("avatar are empty")
	ErrPhoneEmpty           = errors.New("phone are empty")
	ErrAddressEmpty         = errors.New("address are empty")
	ErrTitleEmpty           = errors.New("title are empty")
	ErrCategoryIdEmpty      = errors.New("category_id are empty")
	ErrTeacherIdEmpty       = errors.New("teacher_id are empty")
	ErrStudentIdEmpty       = errors.New("student_id are empty")
	ErrCourseIdEmpty        = errors.New("course_id are empty")

	//others
	ErrWrongPassword      = errors.New("wrong password")
	ErrEmailHasApplied    = errors.New("email has applied")
	ErrValidationPassword = errors.New("password must same with confirm password")
	ErrEmailNotExist      = errors.New("email not exist")
)
