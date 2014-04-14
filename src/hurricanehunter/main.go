package main

import (
	"github.com/cloudfoundry-incubator/dropsonde"
	"hurricanehunter/hunter"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("Launching HurricaneHunter on port " + os.Getenv("HH_PORT") + " … testing dropsondes")

	handler := new(hunter.Handler)
	instrumentedHunter, err := dropsonde.InstrumentedHandler(handler, "hunter", 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(http.ListenAndServe(":" + os.Getenv("HH_PORT"), instrumentedHunter))
}
