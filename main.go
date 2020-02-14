package main

import (
	"capulus/api"
	setting "capulus/config/settings"
	"capulus/database"
	"capulus/database/migrations"
	"capulus/package/logging"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Init the application
func init() {
	// config
	setting.Setup()
	logging.Setup()
	// database
	db, _ := database.Setup()
	migrations.Migrate(db)
}

// @title CAPULUS API
// @version 1.0
// @description Wake up website & ping
// @repository https://github.com/capulus
func main() {
	// gin mode
	gin.SetMode(setting.ServerSetting.RunMode)
	// server settings
	routersInit := api.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	// server http
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
