package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init(){
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil{
		panic(err)
	}
	if viper.GetBool(`debug`){
		log.Println("This Services RUN on DEBUG Mode")
	}
}

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ping Test Connection")
	})
	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}