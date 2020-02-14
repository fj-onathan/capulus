package setting

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
	TimeFormat      string
	RuntimeRootPath string
	ExportSavePath  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type  string
	Local string
}

var DatabaseSetting = &Database{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	// deploy port
	if len(os.Getenv("PORT")) > 0 {
		ServerSetting.HttpPort, _ = strconv.Atoi(os.Getenv("PORT"))
	}
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
