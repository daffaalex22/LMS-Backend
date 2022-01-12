package main

import (
	"log"
	"net/http"
	"time"

	_middleware "backend/app/middleware"
	"backend/app/routes"
	categoriesUsecase "backend/business/categories"
	difficultiesUsecase "backend/business/difficulties"
	enrollmentsUseCase "backend/business/enrollments"
	modulesUseCase "backend/business/modules"
	readingsUseCase "backend/business/readings"
	requestsUseCase "backend/business/requests"
	studentUseCase "backend/business/student"
	teacherUseCase "backend/business/teacher"
	typesUsecase "backend/business/types"
	videosUseCase "backend/business/videos"
	categoriesController "backend/controllers/categories"
	difficultiesController "backend/controllers/difficulties"
	enrollmentsController "backend/controllers/enrollments"
	modulesController "backend/controllers/modules"
	readingsController "backend/controllers/readings"
	requestsController "backend/controllers/requests"
	studentController "backend/controllers/student"
	teacherController "backend/controllers/teacher"
	typesController "backend/controllers/types"
	videosController "backend/controllers/videos"
	categoriesdb "backend/drivers/database/categories"
	difficultiesRepo "backend/drivers/database/difficulties"
	enrollmentsRepo "backend/drivers/database/enrollments"
	modulesRepo "backend/drivers/database/modules"
	"backend/drivers/database/mysql"
	readingsRepo "backend/drivers/database/readings"
	requestsRepo "backend/drivers/database/requests"
	studentRepo "backend/drivers/database/student"
	teacherRepo "backend/drivers/database/teacher"
	typesRepo "backend/drivers/database/types"
	videosRepo "backend/drivers/database/videos"

	courseUsecase "backend/business/course"
	courseController "backend/controllers/courses"
	coursedb "backend/drivers/database/course"

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
	db.AutoMigrate(&categoriesdb.Category{})
	db.AutoMigrate(&coursedb.Course{})
	db.AutoMigrate(&enrollmentsRepo.Enrollments{})
	db.AutoMigrate(&requestsRepo.Requests{})
	db.AutoMigrate(&modulesRepo.Modules{})
	db.AutoMigrate(&readingsRepo.Readings{})
	db.AutoMigrate(&videosRepo.Videos{})
	db.AutoMigrate(&difficultiesRepo.Difficulty{})
	db.AutoMigrate(&typesRepo.Types{})
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
	categoriesRepository := categoriesdb.NewMysqlCategoryRepository(db)
	categoriesUseCase := categoriesUsecase.NewCategoryUsecase(timeoutContext, categoriesRepository)
	CategoriesController := categoriesController.NewCategoriesController(categoriesUseCase)

	//enrollments
	enrollmentsRepoInterface := enrollmentsRepo.NewEnrollmentsRepository(db)
	enrollmentsUseCaseInterface := enrollmentsUseCase.NewUseCase(enrollmentsRepoInterface, timeoutContext)
	enrollmentsUseControllerInterface := enrollmentsController.NewEnrollmentsController(enrollmentsUseCaseInterface)

	//enrollments
	requestsRepoInterface := requestsRepo.NewRequestsRepository(db)
	requestsUseCaseInterface := requestsUseCase.NewUseCase(requestsRepoInterface, timeoutContext)
	requestsUseControllerInterface := requestsController.NewRequestsController(requestsUseCaseInterface)

	//modules
	modulesRepoInterface := modulesRepo.NewModulesRepository(db)
	modulesUseCaseInterface := modulesUseCase.NewUseCase(modulesRepoInterface, timeoutContext)
	modulesUseControllerInterface := modulesController.NewModulesController(modulesUseCaseInterface)

	//readings
	readingsRepoInterface := readingsRepo.NewReadingsRepository(db)
	readingsUseCaseInterface := readingsUseCase.NewUseCase(readingsRepoInterface, timeoutContext)
	readingsUseControllerInterface := readingsController.NewReadingsController(readingsUseCaseInterface)

	//videos
	videosRepoInterface := videosRepo.NewVideosRepository(db)
	videosUseCaseInterface := videosUseCase.NewUseCase(videosRepoInterface, timeoutContext)
	videosUseControllerInterface := videosController.NewVideosController(videosUseCaseInterface)

	//course
	courseRepository := coursedb.NewMysqlCategoryRepository(db)
	courseUseCase := courseUsecase.NewCourseUsecase(timeoutContext, courseRepository)
	CourseController := courseController.NewCourseController(courseUseCase)

	//difficulties
	difficultiesRepository := difficultiesRepo.NewMysqlDifficultyRepository(db)
	difficultiesUseCase := difficultiesUsecase.NewDifficultyUsecase(timeoutContext, difficultiesRepository)
	difficultiesController := difficultiesController.NewDifficultiesController(difficultiesUseCase)

	//types
	typesRepository := typesRepo.NewMysqlTypeRepository(db)
	typesUsecase := typesUsecase.NewTypeUsecase(timeoutContext, typesRepository)
	typesController := typesController.NewTypesController(typesUsecase)

	routesInit := routes.RouteControllerList{
		StudentController:     *studentUseControllerInterface,
		JWTConfig:             jwt.Init(),
		TeacherController:     *teacherUseControllerInterface,
		JWTConfigs:            jwtTch.Init1(),
		CategoryController:    *CategoriesController,
		DifficultyController:  *difficultiesController,
		TypeController:        *typesController,
		CourseController:      *CourseController,
		EnrollmentsController: *enrollmentsUseControllerInterface,
		RequestsController:    *requestsUseControllerInterface,
		ModulesController:     *modulesUseControllerInterface,
		ReadingsController:    *readingsUseControllerInterface,
		VideosController:      *videosUseControllerInterface,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
