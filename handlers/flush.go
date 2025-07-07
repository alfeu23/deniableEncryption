package handlers

import (
	"deniableEncryption/models"
	"encoding/json"
	"net/http"
)

func FlushCacheHandler(w http.ResponseWriter, r *http.Request) {
	votes = make(map[string]models.Vote)
	receipts = make(map[string]models.Receipt)
	fakeReceipts = make(map[string]models.Receipt)
	deniableData = make(map[string]models.DeniableEncryption)
	voterRegistry = make(map[string]bool)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Cache flushed successfully",
		"status":  "success",
	})
}
