package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"poll-app/internal/models"
	"poll-app/internal/utils"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// VoteHandler handles the request to vote on a poll
func VoteHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var vote models.Vote
	if err := json.NewDecoder(r.Body).Decode(&vote); err != nil {
		log.Println("Error:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if err := models.VoteOnPoll(r.Context(), vote.PollID, userID, vote.OptionID); err != nil {
		if strings.Contains(err.Error(), "user has already voted") {
			http.Error(w, "User has already voted on this poll", http.StatusConflict)
			return
		}
		log.Println("Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	updatedPoll, err := models.GetPoll(r.Context(), vote.PollID)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPoll)
}

// GetVotesHandler handles the request to retrieve votes for a poll option
func GetVotesHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	_, err := utils.ParseToken(tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	pollID, err := strconv.Atoi(ps.ByName("pollID"))
	if err != nil {
		http.Error(w, "Invalid poll ID", http.StatusBadRequest)
		return
	}
	optionID, err := strconv.Atoi(ps.ByName("optionID"))
	if err != nil {
		http.Error(w, "Invalid option ID", http.StatusBadRequest)
		return
	}
	votes, err := models.GetVotes(r.Context(), pollID, optionID)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, "Failed to retrieve votes", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(votes)
}
