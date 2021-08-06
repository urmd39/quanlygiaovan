package main

import (
	"fmt"
	"net/http"
	"quanlygiaovan/apis"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Host: localhost:8080
// Basepath: /
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Homepage"))
	})

	r.Route("/api/v1/", func(r chi.Router) {
		// Vehicle
		r.With(paginate).Get("/vehicles", apis.GetVehicles)
		r.Get("/vehicles/{vehicleId}", apis.GetVehicle)

		// Travel History
		r.Get("/travels/{vehicleId}/distance", apis.GetDistanceTraveled)
		r.Get("/travels/{vehicleId}/filter", apis.GetTravelHistoriesVehicleWithFilter)
		r.Get("/travels/{vehicleId}/all", apis.GetTravelHistoriesOfVehicle)
		r.Post("/travels/{vehicleId}", apis.AddTravelHistory)
		r.Get("/travels/{date}", apis.GetTravelHistoriesOnDate)

		// Statistics
		r.Get("/statistics/{vehicleId}", apis.GetStatistics)

	})

	// Mount the admin sub-router
	// r.Mount("/admin", adminRouter())

	// Port 8080
	fmt.Println("Listen server at port 8080")
	http.ListenAndServe(":8080", r)
}

func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}
