package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"quanlygiaovan/control"
	"quanlygiaovan/entities"

	"github.com/go-chi/chi"
)

// TODO: Get list vehicles
func GetVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := control.GetVehicles()
	json.NewEncoder(w).Encode(list)
}

// TODO: Get vehicle by id
func GetVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "vehicleId")
	vehicle := control.GetVehicle(id)
	json.NewEncoder(w).Encode(vehicle)
}

// TODO: Get list travel histories of vehicle by vehicle id
func GetTravelHistoriesOfVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vehicleId := chi.URLParam(r, "vehicleId")
	list := control.GetTravelHistoriesOfVehicle(vehicleId)
	json.NewEncoder(w).Encode(list)
}

// TODO: Get list travel histories on date
func GetTravelHistoriesOnDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	date := chi.URLParam(r, "date")
	list := control.GetTravelHistoriesOnDate(date)
	json.NewEncoder(w).Encode(list)
}

// TODO: Get travel histories of vehicle with time filter (YYYY-MM-DD / YYYY-MM)
func GetTravelHistoriesVehicleWithFilter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vehicleId := chi.URLParam(r, "vehicleId")
	date := r.FormValue("date")
	month := r.FormValue("month")
	var list []entities.TravelHistory
	if date != "" {
		list = control.GetTravelHistoriesVehicleOnDate(vehicleId, date)
	} else if month != "" {
		list = control.GetTravelHistoriesVehicleOnMonth(vehicleId, month)
	}
	json.NewEncoder(w).Encode(list)
}

// TODO: Get distance traveled of vehicle on time (YYYY-MM-DD / YYYY-MM)
func GetDistanceTraveled(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vehicleId := chi.URLParam(r, "vehicleId")
	date := r.FormValue("date")
	month := r.FormValue("month")
	var distance float64
	if date != "" {
		distance = control.GetDistanceTraveledOnTime(vehicleId, date)
	} else if month != "" {
		distance = control.GetDistanceTraveledOnTime(vehicleId, month)
	}
	message := fmt.Sprintf("%f km", distance)
	json.NewEncoder(w).Encode(message)
}

// TODO: Save travel history (location) of vehicle to Database
func AddTravelHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vehicleId := chi.URLParam(r, "vehicleId")
	body, _ := ioutil.ReadAll(r.Body)
	var th entities.TravelHistoryWithoutId
	json.Unmarshal(body, &th)
	travelHistory := control.AddTravelHistory(vehicleId, th)
	json.NewEncoder(w).Encode(travelHistory)
}

// TODO: Get statistics of vehicle on time (YYYY-MM-DD / YYYY-MM)
func GetStatistics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vehicleId := chi.URLParam(r, "vehicleId")
	date := r.FormValue("date")
	month := r.FormValue("month")
	statistics := control.GetStatistics(vehicleId, date, month)
	json.NewEncoder(w).Encode(statistics)
}
