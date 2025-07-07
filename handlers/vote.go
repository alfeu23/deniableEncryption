package handlers

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"deniableEncryption/models"
)

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	var vote models.Vote

	if err := json.NewDecoder(r.Body).Decode(&vote); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if _, voted := voterRegistry[vote.VoterID]; voted {
		http.Error(w, "You have already voted", http.StatusForbidden)
		return
	}

	vote.ID = generateID()
	vote.Timestamp = time.Now()

	deniable := createDeniableEncryption(vote.Candidate)
	deniableData[vote.ID] = deniable

	votes[vote.ID] = vote

	realReceipt := models.Receipt{
		ID:           generateID(),
		VoteID:       vote.ID,
		Candidate:    vote.Candidate,
		Timestamp:    vote.Timestamp,
		Signature:    generateDeniableSignature(vote.ID, string(vote.Candidate), deniable.KeyA),
		DeniableData: hex.EncodeToString(deniable.CiphertextA),
	}

	receipts[realReceipt.ID] = realReceipt

	var otherCandidate models.Candidate
	if vote.Candidate == models.CandidateA {
		otherCandidate = models.CandidateB
	} else {
		otherCandidate = models.CandidateA
	}

	fakeReceipt := models.Receipt{
		ID:           generateID(),
		VoteID:       vote.ID,
		Candidate:    otherCandidate,
		Timestamp:    vote.Timestamp,
		Signature:    generateDeniableSignature(vote.ID, string(otherCandidate), deniable.KeyB),
		DeniableData: hex.EncodeToString(deniable.CiphertextB),
	}

	fakeReceipts[fakeReceipt.ID] = fakeReceipt

	voterRegistry[vote.VoterID] = true

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"vote_id":         vote.ID,
		"real_receipt_id": realReceipt.ID,
		"fake_receipt_id": fakeReceipt.ID,
		"message":         "Vote recorded successfully with both receipts",
	})
}
