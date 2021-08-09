package control

import (
	"fmt"
	"log"
	"math"
	"quanlygiaovan/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// string-time convert type
const YMD = "2006-01-02"

// TODO: Get list vehicles
func GetVehicles() (list []bson.M) {
	client, ctx := Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("quanlygiaovan").Collection("vehicle_info")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var vehicle bson.M
		if err = cursor.Decode(&vehicle); err != nil {
			log.Fatal(err)
		}
		list = append(list, vehicle)
	}
	return list
}

// TODO: Get vehicle by id
func GetVehicle(idStr string) (vehicle bson.M) {
	client, ctx := Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("quanlygiaovan").Collection("vehicle_info")

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&vehicle); err != nil {
		log.Fatal(err)
	}
	return vehicle
}

// TODO: Get list travel histories of vehicle by vehicle id
func GetTravelHistoriesOfVehicle(vehicleId string) (list []bson.M) {
	client, ctx := Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("quanlygiaovan").Collection("travel_history")

	id, err := primitive.ObjectIDFromHex(vehicleId)
	if err != nil {
		log.Fatal(err)
	}
	cursor, err := collection.Find(ctx, bson.M{"vehicleId": id})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var vehicle bson.M
		if err = cursor.Decode(&vehicle); err != nil {
			log.Fatal(err)
		}
		list = append(list, vehicle)
	}
	return list
}

// TODO: Get list all travel histories on date (YYYY-MM-DD)
func GetTravelHistoriesOnDate(date string) (list []entities.TravelHistory) {
	client, ctx := Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("quanlygiaovan").Collection("travel_history")

	findOptions := options.Find().SetSort(bson.D{{"updatedAt", 1}})
	d, _ := time.Parse(YMD, date)
	fmt.Println(d)
	cursor, err := collection.Find(ctx, bson.M{"updatedAt": bson.M{"$gte": primitive.NewDateTimeFromTime(d),
		"$lte": primitive.NewDateTimeFromTime(d.AddDate(0, 0, 1))}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var travelHistory entities.TravelHistory
		if err = cursor.Decode(&travelHistory); err != nil {
			log.Fatal(err)
		}
		list = append(list, travelHistory)
	}
	return list
}

// TODO: Get list travel histories of vehicle on date (YYYY-MM-DD)
func GetTravelHistoriesVehicleOnDate(vehicleId string, date string) (list []entities.TravelHistory) {
	client, ctx := Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("quanlygiaovan").Collection("travel_history")

	id, err := primitive.ObjectIDFromHex(vehicleId)
	if err != nil {
		log.Fatal(err)
	}

	findOptions := options.Find().SetSort(bson.D{{"updatedAt", 1}})
	d, _ := time.Parse(YMD, date)
	cursor, err := collection.Find(ctx, bson.M{"vehicleId": id, "updatedAt": bson.M{"$gte": primitive.NewDateTimeFromTime(d),
		"$lte": primitive.NewDateTimeFromTime(d.AddDate(0, 0, 1))}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var travelHistory entities.TravelHistory
		if err = cursor.Decode(&travelHistory); err != nil {
			log.Fatal(err)
		}
		list = append(list, travelHistory)
	}
	return list
}

// TODO: Get list travel histories of vehicle on month (YYYY-MM)
func GetTravelHistoriesVehicleOnMonth(vehicleId string, month string) (list []entities.TravelHistory) {
	client, ctx := Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("quanlygiaovan").Collection("travel_history")

	id, err := primitive.ObjectIDFromHex(vehicleId)
	if err != nil {
		log.Fatal(err)
	}

	findOptions := options.Find().SetSort(bson.D{{"updatedAt", 1}})
	d, _ := time.Parse(YMD, month+"-01")
	cursor, err := collection.Find(ctx, bson.M{"vehicleId": id,
		"updatedAt": bson.M{"$gte": primitive.NewDateTimeFromTime(d),
			"$lt": primitive.NewDateTimeFromTime(d.AddDate(0, 1, 0))}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var travelHistory entities.TravelHistory
		if err = cursor.Decode(&travelHistory); err != nil {
			log.Fatal(err)
		}
		list = append(list, travelHistory)
	}
	return list
}

// TODO: Get distance traveled of vehicle on time (date: YYYY-MM-DD / month: YYYY-MM)
func GetDistanceTraveledOnTime(vehicleId string, time string) (distance float64) {
	var list []entities.TravelHistory
	l := len(time)
	if l == 7 {
		list = GetTravelHistoriesVehicleOnMonth(vehicleId, time)
	} else {
		list = GetTravelHistoriesVehicleOnDate(vehicleId, time)
	}

	if len(list) > 1 {
		for i := 0; i < len(list)-1; i++ {
			if list[i].TransportID == list[i+1].TransportID {
				// Get Location
				loc1 := list[i].Location
				loc2 := list[i+1].Location

				// Get Latitude, Longitude
				lon1 := loc1.Coordinates[0]
				lon2 := loc2.Coordinates[0]
				lat1 := loc1.Coordinates[1]
				lat2 := loc2.Coordinates[1]
				distance += GetDistance(lon1, lat1, lon2, lat2)
			}
		}
	}
	return distance
}

// TODO: func to calculate distance traveled
func GetDistance(longitude1 float64, latitude1 float64, longitude2 float64,
	latitude2 float64) (distance float64) {
	const R = 6378
	lon1 := longitude1 * (math.Pi / 180)
	lat1 := latitude1 * (math.Pi / 180)
	lon2 := longitude2 * (math.Pi / 180)
	lat2 := latitude2 * (math.Pi / 180)
	distance = R * math.Acos((math.Sin(lat1)*math.Sin(lat2))+math.Cos(lat1)*math.Cos(lat2)*math.Cos(lon2-lon1))
	return distance
}

// TODO: Save travel history record of vehicle to Database
func AddTravelHistory(vehicleId string, travelHistory bson.M) bson.M {
	client, ctx := Connected()
	defer client.Disconnect(ctx)
	collection := client.Database("quanlygiaovan").Collection("travel_history")

	id_vehicle, err := primitive.ObjectIDFromHex(vehicleId)
	if err != nil {
		log.Fatal(err)
	}
	travelHistory["vehicleId"] = id_vehicle
	addResult, err := collection.InsertOne(ctx, travelHistory)
	if err != nil {
		log.Fatal(err)
	}
	travelHistory["_id"] = addResult.InsertedID
	return travelHistory
}

// TODO: Get statistics of vehicle on time (date: YYYY-MM-DD / month: YYYY-MM)
func GetStatistics(vehicleId string, date string, month string) (statistics entities.Statistics) {
	var list []entities.TravelHistory
	if month != "" {
		list = GetTravelHistoriesVehicleOnMonth(vehicleId, month)
		statistics.DistanceTraveled = GetDistanceTraveledOnTime(vehicleId, month)
	} else if date != "" {
		list = GetTravelHistoriesVehicleOnDate(vehicleId, date)
		statistics.DistanceTraveled = GetDistanceTraveledOnTime(vehicleId, date)
	}

	statistics.TransportNumber = 1
	var t time.Duration
	if len(list) > 1 {
		for i := 0; i < len(list)-1; i++ {
			if list[i].TransportID == list[i+1].TransportID {
				t += list[i+1].UpdatedAt.Sub(list[i].UpdatedAt)
			} else {
				statistics.TransportNumber++
			}
		}
	}
	statistics.TotalTime = t.String()
	fmt.Println(statistics)
	return statistics
}
