package main

import (
	"log"
	"net/http"
	"time"

	_middleware "backend/app/middleware"
	"backend/app/routes"
	_categoriesUsecase "backend/business/categories"
	enrollmentsUseCase "backend/business/enrollments"
	modulesUseCase "backend/business/modules"
	readingsUseCase "backend/business/readings"
	studentUseCase "backend/business/student"
	teacherUseCase "backend/business/teacher"
	_categoriesController "backend/controllers/categories"
	enrollmentsController "backend/controllers/enrollments"
	modulesController "backend/controllers/modules"
	readingsController "backend/controllers/readings"
	studentController "backend/controllers/student"
	teacherController "backend/controllers/teacher"
	_categoriesdb "backend/drivers/database/categories"
	enrollmentsRepo "backend/drivers/database/enrollments"
	modulesRepo "backend/drivers/database/modules"
	"backend/drivers/database/mysql"
	readingsRepo "backend/drivers/database/readings"
	studentRepo "backend/drivers/database/student"
	teacherRepo "backend/drivers/database/teacher"

	_courseUsecase "backend/business/course"
	_courseController "backend/controllers/courses"
	_coursedb "backend/drivers/database/course"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("This Services RUN on DEBUG Mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&studentRepo.Student{})
	db.AutoMigrate(&teacherRepo.Teacher{})
	db.AutoMigrate(&_categoriesdb.Category{})
	db.AutoMigrate(&_coursedb.Course{})
	db.AutoMigrate(&enrollmentsRepo.Enrollments{})
	db.AutoMigrate(&modulesRepo.Modules{})
	db.AutoMigrate(&readingsRepo.Readings{})
}

func main() {
	configDb := mysql.ConfigDB{
		DB_Username: viper.GetString("database.user"),
		DB_Password: viper.GetString("database.pass"),
		DB_Host:     viper.GetString("database.host"),
		DB_Port:     viper.GetString("database.port"),
		DB_Database: viper.GetString("database.name"),
	}
	db := configDb.InitialDb()
	dbMigrate(db)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	jwt := _middleware.ConfigJWT{
		SecretJWT: viper.GetString("jwt.secretStudent"),
	}
	jwtTch := _middleware.ConfigsJWT{
		SecretJWTch: viper.GetString("jwt.secretTeacher"),
	}

	//student
	studentRepoInterface := studentRepo.NewStudentRepository(db, &jwt)
	studentUseCaseInterface := studentUseCase.NewUseCase(studentRepoInterface, timeoutContext)
	studentUseControllerInterface := studentController.NewStudentController(studentUseCaseInterface)

	//teacher
	teacherRepoInterface := teacherRepo.NewTeacherRepository(db, &jwtTch)
	teacherUseCaseInterface := teacherUseCase.NewUseCase(teacherRepoInterface, timeoutContext)
	teacherUseControllerInterface := teacherController.NewTeacherController(teacherUseCaseInterface)

	//categories
	categoriesRepository := _categoriesdb.NewMysqlCategoryRepository(db)
	categoriesUseCase := _categoriesUsecase.NewCategoryUsecase(timeoutContext, categoriesRepository)
	CategoriesController := _categoriesController.NewCategoriesController(categoriesUseCase)

	//teacher
	enrollmentsRepoInterface := enrollmentsRepo.NewEnrollmentsRepository(db)
	enrollmentsUseCaseInterface := enrollmentsUseCase.NewUseCase(enrollmentsRepoInterface, timeoutContext)
	enrollmentsUseControllerInterface := enrollmentsController.NewEnrollmentsController(enrollmentsUseCaseInterface)

	//modules
	modulesRepoInterface := modulesRepo.NewModulesRepository(db)
	modulesUseCaseInterface := modulesUseCase.NewUseCase(modulesRepoInterface, timeoutContext)
	modulesUseControllerInterface := modulesController.NewModulesController(modulesUseCaseInterface)

	//readings
	readingsRepoInterface := readingsRepo.NewReadingsRepository(db)
	readingsUseCaseInterface := readingsUseCase.NewUseCase(readingsRepoInterface, timeoutContext)
	readingsUseControllerInterface := readingsController.NewReadingsController(readingsUseCaseInterface)

	//course
	courseRepository := _coursedb.NewMysqlCategoryRepository(db)
	courseUseCase := _courseUsecase.NewCourseUsecase(timeoutContext, courseRepository)
	CourseController := _courseController.NewCourseController(courseUseCase)

	routesInit := routes.RouteControllerList{
		StudentController:     *studentUseControllerInterface,
		JWTConfig:             jwt.Init(),
		TeacherController:     *teacherUseControllerInterface,
		JWTConfigs:            jwtTch.Init1(),
		CategoryController:    *CategoriesController,
		CourseController:      *CourseController,
		EnrollmentsController: *enrollmentsUseControllerInterface,
		ModulesController:     *modulesUseControllerInterface,
		ReadingsController:    *readingsUseControllerInterface,
	}

	routesInit.CourseRouteRegister(e, timeoutContext)
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
