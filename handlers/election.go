package handlers

import (
	"encoding/json"
	"net/http"

	"deniableEncryption/models"
)

func ElectionResultsHandler(w http.ResponseWriter, r *http.Request) {
	candidateAVotes := 0
	candidateBVotes := 0
	totalVotes := 0

	for _, vote := range votes {
		totalVotes++
		switch vote.Candidate {
		case models.CandidateA:
			candidateAVotes++
		case models.CandidateB:
			candidateBVotes++
		}
	}
	results := map[string]interface{}{
		"total_votes": totalVotes,
		"results": map[string]interface{}{
			string(models.CandidateA): candidateAVotes,
			string(models.CandidateB): candidateBVotes,
		},
		"winner": determineWinner(candidateAVotes, candidateBVotes),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func determineWinner(candidateAVotes, candidateBVotes int) string {
	if candidateAVotes > candidateBVotes {
		return string(models.CandidateA)
	} else if candidateBVotes > candidateAVotes {
		return string(models.CandidateB)
	} else {
		return "tie"
	}
}
