package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vehicle struct {
	VehicleID primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
}

type Location struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type TravelHistory struct {
	TravelHistoryID        primitive.ObjectID `json:"_id" bson:"_id"`
	TravelHistoryWithoutId `bson:",inline"`
}

type TravelHistoryWithoutId struct {
	VehicleID   primitive.ObjectID `json:"vehicleId" bson:"vehicleId"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
	Location    Location           `json:"location" bson:"location"`
	TransportID primitive.ObjectID `json:"transportId" bson:"transportId"`
}

type Statistics struct {
	TransportNumber  int     `json:"transportNumber"`
	DistanceTraveled float64 `json:"distanceTraveled"`
	TotalTime        string  `json:"totalTime"`
}
