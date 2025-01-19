package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURI = "mongodb://localhost:27017"

const (
	databaseName   = "myDatabase"
	collectionName = "myCollection"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/data", getAllDataHandler).Methods("GET")

	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getAllDataHandler(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		http.Error(w, "Failed to create MongoDB client", http.StatusInternalServerError)
		log.Println("Error creating MongoDB client:", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		http.Error(w, "Failed to connect to MongoDB", http.StatusInternalServerError)
		log.Println("Error connecting to MongoDB:", err)
		return
	}
	defer client.Disconnect(ctx)

	collection := client.Database(databaseName).Collection(collectionName)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		log.Println("Error retrieving data:", err)
		return
	}
	defer cursor.Close(ctx)

	var results []map[string]interface{}
	if err = cursor.All(ctx, &results); err != nil {
		http.Error(w, "Failed to decode data", http.StatusInternalServerError)
		log.Println("Error decoding data:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
