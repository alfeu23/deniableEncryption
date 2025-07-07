package handlers

import (
	"encoding/hex"
	"encoding/json"
	"net/http"

	"deniableEncryption/models"

	"github.com/gorilla/mux"
)

func ReceiptHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	receiptID := vars["id"]

	if receipt, exists := receipts[receiptID]; exists {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(receipt)
		return
	}

	if receipt, exists := fakeReceipts[receiptID]; exists {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(receipt)
		return
	}

	http.Error(w, "Receipt not found", http.StatusNotFound)
}

func VerifyReceiptHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	receiptID := vars["id"]

	var receipt models.Receipt
	var isReal bool
	var exists bool

	if receipt, exists = receipts[receiptID]; exists {
		isReal = true
	} else if receipt, exists = fakeReceipts[receiptID]; exists {
		isReal = false
	} else {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	deniable, deniableExists := deniableData[receipt.VoteID]
	if !deniableExists {
		http.Error(w, "Vote data not found", http.StatusNotFound)
		return
	}

	var proof map[string]interface{}

	if isReal {
		if receipt.Candidate == models.CandidateA {
			decrypted := doDecryption(deniable.CiphertextA, deniable.CiphertextB, deniable.KeyA)
			proof = map[string]interface{}{
				"C1":        hex.EncodeToString(deniable.CiphertextA),
				"C2":        hex.EncodeToString(deniable.CiphertextB),
				"key":       hex.EncodeToString(deniable.KeyA),
				"decrypted": decrypted,
				"message":   "This proves the vote was for " + decrypted,
				"type":      "real_receipt",
			}
		} else {
			decrypted := doDecryption(deniable.CiphertextA, deniable.CiphertextB, deniable.KeyA)
			proof = map[string]interface{}{
				"C1":        hex.EncodeToString(deniable.CiphertextA),
				"C2":        hex.EncodeToString(deniable.CiphertextB),
				"key":       hex.EncodeToString(deniable.KeyA),
				"decrypted": decrypted,
				"message":   "This proves the vote was for " + decrypted,
				"type":      "real_receipt",
			}
		}
	} else {
		if receipt.Candidate == models.CandidateA {
			decrypted := doDecryption(deniable.CiphertextA, deniable.CiphertextB, deniable.KeyB)
			proof = map[string]interface{}{
				"C1":        hex.EncodeToString(deniable.CiphertextA),
				"C2":        hex.EncodeToString(deniable.CiphertextB),
				"key":       hex.EncodeToString(deniable.KeyB),
				"decrypted": decrypted,
				"message":   "This proves the vote was for " + decrypted,
				"type":      "fake_receipt",
			}
		} else {
			decrypted := doDecryption(deniable.CiphertextA, deniable.CiphertextB, deniable.KeyB)
			proof = map[string]interface{}{
				"C1":        hex.EncodeToString(deniable.CiphertextA),
				"C2":        hex.EncodeToString(deniable.CiphertextB),
				"key":       hex.EncodeToString(deniable.KeyB),
				"decrypted": decrypted,
				"message":   "This proves the vote was for " + decrypted,
				"type":      "fake_receipt",
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proof)
}
