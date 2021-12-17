package main

import (
	"log"
	"net/http"
	"time"

	"backend/app/routes"

	_mysqldriver "backend/drivers/mysql"

	_categoriesUsecase "backend/business/categories"
	_categoriesController "backend/controllers/categories"
	_categoriesdb "backend/drivers/database/categories"

	_courseUsecase "backend/business/course"
	_courseController "backend/controllers/courses"
	_coursedb "backend/drivers/database/course"

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

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&_categoriesdb.Category{})
	db.AutoMigrate(&_coursedb.Course{})
}

func main() {
	configDB := _mysqldriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	Conn := configDB.InitialDB()
	DBMigrate(Conn)

	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ping Test Connection")
	})

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	//categories
	categoriesRepository := _categoriesdb.NewMysqlCategoryRepository(Conn)
	categoriesUseCase := _categoriesUsecase.NewCategoryUsecase(timeoutContext, categoriesRepository)
	CategoriesController := _categoriesController.NewCategoriesController(categoriesUseCase)

	//course
	courseRepository := _coursedb.NewMysqlCategoryRepository(Conn)
	courseUseCase := _courseUsecase.NewCourseUsecase(timeoutContext, courseRepository)
	CourseController := _courseController.NewCourseController(courseUseCase)

	routesInit := routes.ControllerList{
		CategoryController: *CategoriesController,
		CourseController:   *CourseController,
	}

	routesInit.CategoriesRouteRegister(e, timeoutContext)
	routesInit.CourseRouteRegister(e, timeoutContext)

	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}
