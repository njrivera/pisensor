package endpoints

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-martini/martini"
	"github.com/pisensor/server/internal/types"
)

const (
	clientTimeFormat = "1/2/2006, 15:04:05 PM"
)

var (
	db types.DB
)

func RegisterReadingsEndpoints(m *martini.ClassicMartini, d types.DB) {
	db = d

	m.Group("/readings", func(r martini.Router) {
		r.Get("/betweentimes", getTempBetweenTimes)
	})
}

func getTempBetweenTimes(r *http.Request, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	args := r.URL.Query()

	serial, startArg, endArg := args.Get("serial"), args.Get("start"), args.Get("end")

	start, err := time.Parse(clientTimeFormat, startArg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error parsing start time: %s", err.Error())
		return
	}
	end, err := time.Parse(clientTimeFormat, endArg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error parsing end time: %s", err.Error())
		return
	}

	readings, err := db.GetTempsBetweenTimes(serial, start, end)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error getting temp readings: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(readings)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshaling temp readings: %s", err.Error())
		return
	}

	w.Write(js)
}
