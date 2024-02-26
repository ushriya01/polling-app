package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"poll-app/internal/models"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// ListPollsHandler handles the request to list all polls
func ListPollsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	polls, err := models.ListPolls(r.Context())
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(polls)
}

// CreatePollHandler handles the request to create a new poll
func CreatePollHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.Context().Value("user_id").(string)
	var poll models.Poll
	if err := json.NewDecoder(r.Body).Decode(&poll); err != nil {
		log.Println("Error:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	var options []string
	if len(poll.PollOptions) > 0 {
		for _, option := range poll.PollOptions {
			options = append(options, option.Text)
		}
	}
	createdPoll, err := models.CreatePoll(r.Context(), poll.Question, userID, options)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdPoll)
}

// GetPollHandler handles the request to retrieve a poll by ID
func GetPollHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pollID, err := strconv.Atoi(ps.ByName("pollID"))
	if err != nil {
		http.Error(w, "Invalid poll ID", http.StatusBadRequest)
		return
	}
	poll, err := models.GetPoll(r.Context(), pollID)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Poll not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(poll)
}

// DeletePollHandler handles the request to delete a poll by ID
func DeletePollHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Context().Value("user_id").(string)
	pollID, err := strconv.Atoi(ps.ByName("pollID"))
	if err != nil {
		http.Error(w, "Invalid poll ID", http.StatusBadRequest)
		return
	}
	poll, err := models.GetPoll(r.Context(), pollID)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Poll not found", http.StatusNotFound)
		return
	}
	if poll.CreatedBy != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	err = models.DeletePoll(r.Context(), pollID)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// UpdatePollHandler handles the request to update a poll by ID
func UpdatePollHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Context().Value("user_id").(string)
	pollID, err := strconv.Atoi(ps.ByName("pollID"))
	if err != nil {
		http.Error(w, "Invalid poll ID", http.StatusBadRequest)
		return
	}
	var updatedPoll models.Poll
	if err := json.NewDecoder(r.Body).Decode(&updatedPoll); err != nil {
		log.Println("Error:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	poll, err := models.GetPoll(r.Context(), pollID)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Poll not found", http.StatusNotFound)
		return
	}
	if poll.CreatedBy != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	err = models.UpdatePoll(r.Context(), pollID, updatedPoll)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
