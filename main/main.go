package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type Line struct {
	ID   string             `json:"id"`
	Path geojson.LineString `json:"path"`
}

func checkIntersectionsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var linestring geojson.LineString
	err := json.NewDecoder(r.Body).Decode(&linestring)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the linestring
	if len(linestring.Coordinates) < 2 {
		http.Error(w, "Invalid linestring", http.StatusBadRequest)
		return
	}

	// Simulate 50 randomly spread lines
	lines := generateRandomLines()

	// Perform intersection check
	intersections := make(map[string]geojson.Point)
	for _, line := range lines {
		if intersection, ok := intersect.LineStringLineString(&linestring, &line.Path); ok {
			intersections[line.ID] = intersection
		}
	}

	// Return the results
	if len(intersections) == 0 {
		w.Write([]byte("[]")) // No intersections
	} else {
		response, err := json.Marshal(intersections)
		if err != nil {
			http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func generateRandomLines() []Line {
	// Simulate 50 randomly spread lines
	var lines []Line
	for i := 1; i <= 50; i++ {
		line := Line{
			ID: fmt.Sprintf("L%02d", i),
			Path: geojson.LineString{
				Coordinates: []geojson.Position{
					{RandomFloat(), RandomFloat()},
					{RandomFloat(), RandomFloat()},
				},
			},
		}
		lines = append(lines, line)
	}
	return lines
}

func RandomFloat() float64 {
	// Simulate random float values between -10 and 10
	return -10 + (20 * rand.Float64())
}

func main() {
	http.HandleFunc("/checkIntersections", checkIntersectionsHandler)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
