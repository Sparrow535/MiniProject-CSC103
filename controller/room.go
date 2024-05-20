package controller

import (
	"encoding/json"
	"myapp/model"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	var room model.Room
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = room.CreateRoom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Construct the response JSON
	response := map[string]interface{}{
		"message": "Room added successfully",
		"room":    room, // Include the room details if needed
	}

	// Set the response header
	w.Header().Set("Content-Type", "application/json")
	// Set the HTTP status code
	w.WriteHeader(http.StatusCreated)
	// Encode the response as JSON and write it to the response writer
	json.NewEncoder(w).Encode(response)
}

// GetRoomHandler handles the retrieval of a room.
func GetRoom(w http.ResponseWriter, r *http.Request) {
	// Extract room number from URL path
	vars := mux.Vars(r)
	roomNo := vars["roomNo"]

	// Convert room number to appropriate type if needed (e.g., string to int)
	// Example: roomNoInt, err := strconv.Atoi(roomNo)
	// if err != nil {
	//     http.Error(w, "Invalid room number", http.StatusBadRequest)
	//     return
	// }

	// Call the method to fetch room details using the room number
	room, err := model.GetRoomByNumber(roomNo)
	if err != nil {
		// Handle error (e.g., room not found)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Respond with the retrieved room details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}
func CreateRoomG(w http.ResponseWriter, r *http.Request) {
	var room model.RoomG
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = room.CreateRoomG()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Construct the response JSON
	response := map[string]interface{}{
		"message": "Room added successfully",
		"room":    room, // Include the room details if needed
	}

	// Set the response header
	w.Header().Set("Content-Type", "application/json")
	// Set the HTTP status code
	w.WriteHeader(http.StatusCreated)
	// Encode the response as JSON and write it to the response writer
	json.NewEncoder(w).Encode(response)
}

// GetRoomHandler handles the retrieval of a room.
func GetRoomG(w http.ResponseWriter, r *http.Request) {
	// Extract room number from URL path
	vars := mux.Vars(r)
	roomNo := vars["roomNo"]

	// Convert room number to appropriate type if needed (e.g., string to int)
	// Example: roomNoInt, err := strconv.Atoi(roomNo)
	// if err != nil {
	//     http.Error(w, "Invalid room number", http.StatusBadRequest)
	//     return
	// }

	// Call the method to fetch room details using the room number
	room, err := model.GetRoomByNumberG(roomNo)
	if err != nil {
		// Handle error (e.g., room not found)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Respond with the retrieved room details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}
