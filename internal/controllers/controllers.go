package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ysdinesh31/temperature-error-logs/internal/db"
	"github.com/ysdinesh31/temperature-error-logs/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

type TemperatureRequest struct {
	Data string `json:"data"`
}

func HandleTemperatureReading(w http.ResponseWriter, r *http.Request) {
	var req TemperatureRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Fatalf("Invalid request format")
		StoreErrorReading(req.Data)
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}

	parts := strings.Split(req.Data, ":")
	if len(parts) != 4 || parts[2] != "'Temperature'" {
		log.Fatalf("Invalid data format")
		StoreErrorReading(req.Data)
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}

	deviceID, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("Invalid device ID")
		StoreErrorReading(req.Data)
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}

	epochMS, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		log.Fatalf("Invalid epoch timestamp")
		StoreErrorReading(req.Data)
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}

	temperature, err := strconv.ParseFloat(parts[3], 64)
	if err != nil {
		log.Fatalf("Invalid temperature value")
		StoreErrorReading(req.Data)
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}

	if temperature >= 90 {
		formattedTime := time.Unix(0, epochMS*int64(time.Millisecond)).Format("2024/06/14 15:04:05")
		response := map[string]interface{}{
			"overtemp":       true,
			"device_id":      deviceID,
			"formatted_time": formattedTime,
		}
		json.NewEncoder(w).Encode(response)
	} else {
		response := map[string]interface{}{"overtemp": false}
		json.NewEncoder(w).Encode(response)
	}
}

func StoreErrorReading(data string) {
	collection := db.Client.Database("temperature").Collection("errors")
	_, err := collection.InsertOne(context.TODO(), models.ErrorRecord{Data: data})
	if err != nil {
		log.Fatalf("Failed to store error in MongoDB")
	}
}

func GetErrors(w http.ResponseWriter, r *http.Request) {
	collection := db.Client.Database("temperature").Collection("errors")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatalf("Failed to retrieve errors from MongoDB")
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var errorRecords []models.ErrorRecord
	if err := cursor.All(context.TODO(), &errorRecords); err != nil {
		log.Fatalf("Failed to decode error records")
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}

	errors := make([]string, len(errorRecords))
	for i, record := range errorRecords {
		errors[i] = record.Data
	}

	response := map[string]interface{}{"errors": errors}
	json.NewEncoder(w).Encode(response)
}

func DeleteErrors(w http.ResponseWriter, r *http.Request) {
	collection := db.Client.Database("temperature").Collection("errors")
	_, err := collection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		log.Fatalf("Failed to delete errors from MongoDB")
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
