package crons

import (
	"capulus/database/models"
	"fmt"
	"log"
	"net/http"

	"github.com/jasonlvhit/gocron"
)

// Run all crons
func Run() {
	s := gocron.NewScheduler()
	s.Every(20).Minutes().Do(WakeUpHost)
	<-s.Start()
}

func WakeUpHost() {
	RequestHost()
	PingHost()
}

func PingHost() {}

func RequestHost() {
	length := len(ReturnHosts())
	for i := 0; i < length; i++ {
		response, err := http.Get(ReturnHosts()[i].Host)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n The response of the host (%v) is: %v \n", ReturnHosts()[i].Host, response)
	}
}

func ReturnHosts() []models.Host {
	hosts, err := models.GetHosts()
	if err != nil {
		return nil
	}
	return hosts
}
