package err

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")

<<<<<<< HEAD
	//notfound
	ErrNotFound            = errors.New("data not found")
	ErrCategoryNotFound    = errors.New("category not found")
	ErrEnrollNotFound      = errors.New("enrollment not found")
	ErrCourseNotFound      = errors.New("course not found")
	ErrTeacherNotFound     = errors.New("teacher not found")
	ErrStudentNotFound     = errors.New("student not found")
	ErrModulesNotFound     = errors.New("modules not found")
	ErrEnrollmentsNotFound = errors.New("category not found")

	//empty
	ErrIdEmpty              = errors.New("id are empty")
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

	//relasiproblem
	ErrIdStudent = errors.New("id student not working")
	ErrIdCourse  = errors.New("id course not working")

	//others
	ErrWrongPassword      = errors.New("wrong password")
	ErrEmailHasApplied    = errors.New("email has applied")
	ErrValidationPassword = errors.New("password must same with confirm password")
	ErrEmailNotExist      = errors.New("email not exist")
	ErrConvertId          = errors.New("convert id error")
=======
	ErrNotFound = errors.New("data not found")

	ErrCategoryNotFound = errors.New("category not found")

	ErrTeacherNotFound = errors.New("teacher not found")

	ErrConvertId = errors.New("convert id error")

	ErrIdEmpty = errors.New("id are empty")

	ErrTitleEmpty = errors.New("title are empty")

	ErrCategoryIdEmpty = errors.New("category_id are empty")

	ErrTeacherIdEmpty = errors.New("teacher_id are empty")

	ErrEnrollmentsNotFound = errors.New("category not found")

	ErrCoursesNotFound = errors.New("course not found")

	ErrCourseNotFound = errors.New("course not found")
>>>>>>> cda27e6e36708927f057cb1500be1c99c819ee99
)
