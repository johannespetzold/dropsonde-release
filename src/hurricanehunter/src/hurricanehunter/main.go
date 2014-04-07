package main

import (
	"github.com/cloudfoundry-incubator/dropsonde"
	"hurricanehunter/hunter"
	"log"
	"net/http"
)

func main() {
	log.Print("Launching HurricaneHunter on port 8080 … testing dropsondes")

	handler := new(hunter.Handler)
	instrumentedHunter, err := dropsonde.InstrumentedHandler(handler, "hunter", 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(http.ListenAndServe(":8080", instrumentedHunter))
}
