package main

import (
	"log"
	"time"

	_middleware "backend/app/middleware"
	"backend/app/routes"
	_categoriesUsecase "backend/business/categories"
	studentUseCase "backend/business/student"
	_categoriesController "backend/controllers/categories"
	studentController "backend/controllers/student"
	_categoriesdb "backend/drivers/database/categories"
	"backend/drivers/database/mysql"
	studentRepo "backend/drivers/database/student"

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
	db.AutoMigrate(&_categoriesdb.Category{})
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
		SecretJWT:       viper.GetString("jwt.secret"),
		ExpiresDuration: viper.GetInt("jwt.expired"),
	}

	//student
	studentRepoInterface := studentRepo.NewStudentRepository(db)
	studentUseCaseInterface := studentUseCase.NewUseCase(studentRepoInterface, timeoutContext, &jwt)
	studentUseControllerInterface := studentController.NewStudentController(studentUseCaseInterface)

	//categories
	categoriesRepository := _categoriesdb.NewMysqlCategoryRepository(db)
	categoriesUseCase := _categoriesUsecase.NewCategoryUsecase(timeoutContext, categoriesRepository)
	CategoriesController := _categoriesController.NewCategoriesController(categoriesUseCase)

	routesInit := routes.RouteControllerList{
		StudentController:  *studentUseControllerInterface,
		CategoryController: *CategoriesController,
	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
