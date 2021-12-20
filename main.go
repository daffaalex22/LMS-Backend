package main

import (
	"log"
	"time"

	_middleware "backend/app/middleware"
	"backend/app/routes"
	_categoriesUsecase "backend/business/categories"
	enrollmentsUseCase "backend/business/enrollments"
	studentUseCase "backend/business/student"
	teacherUseCase "backend/business/teacher"
	_categoriesController "backend/controllers/categories"
	enrollmentsController "backend/controllers/enrollments"
	studentController "backend/controllers/student"
	teacherController "backend/controllers/teacher"
	_categoriesdb "backend/drivers/database/categories"
	enrollmentsRepo "backend/drivers/database/enrollments"
	"backend/drivers/database/mysql"
	studentRepo "backend/drivers/database/student"
	teacherRepo "backend/drivers/database/teacher"

	"github.com/labstack/echo/v4"
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
	db.AutoMigrate(&enrollmentsRepo.Enrollments{})
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

	routesInit := routes.RouteControllerList{
		StudentController:     *studentUseControllerInterface,
		JWTConfig:             jwt.Init(),
		TeacherController:     *teacherUseControllerInterface,
		JWTConfigs:            jwtTch.Init1(),
		CategoryController:    *CategoriesController,
		EnrollmentsController: *enrollmentsUseControllerInterface,
	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
