package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DLzer/icf/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	log.Printf("Using %s Driver for DB Connection...", DbDriver)

	if DbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(DbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database -- ", DbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database -- ", DbDriver)
		}
	}
	if DbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(DbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database -- ", DbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database -- ", DbDriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.Workout{})  //database migration
	server.DB.Debug().AutoMigrate(&models.Exercise{}) //database migration
	server.DB.Debug().AutoMigrate(&models.User{})     //database migration
	server.DB.Debug().AutoMigrate(&models.Tracker{})  //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()

}

// Start the server
func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
