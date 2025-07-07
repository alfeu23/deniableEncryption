package handlers

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"deniableEncryption/models"
)

var (
	votes         = make(map[string]models.Vote)
	receipts      = make(map[string]models.Receipt)
	fakeReceipts  = make(map[string]models.Receipt)
	deniableData  = make(map[string]models.DeniableEncryption)
	voterRegistry = make(map[string]bool)
)

func generateID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func createDeniableEncryption(actualVote models.Candidate) models.DeniableEncryption {
	realKey := make([]byte, 32)
	rand.Read(realKey)

	nonce := make([]byte, 32)
	rand.Read(nonce)

	candidateABytes := padToLength([]byte(string(models.CandidateA)), 32)
	candidateBBytes := padToLength([]byte(string(models.CandidateB)), 32)

	C1 := make([]byte, 32)
	for i := 0; i < 32; i++ {
		C1[i] = nonce[i] ^ realKey[i]
	}

	var C2 []byte
	if actualVote == models.CandidateA {
		C2 = make([]byte, 32)
		for i := 0; i < 32; i++ {
			C2[i] = candidateABytes[i] ^ nonce[i]
		}
	} else {
		C2 = make([]byte, 32)
		for i := 0; i < 32; i++ {
			C2[i] = candidateBBytes[i] ^ nonce[i]
		}
	}

	var fakeKey []byte
	if actualVote == models.CandidateA {
		intermediate := make([]byte, 32)
		for i := 0; i < 32; i++ {
			intermediate[i] = C2[i] ^ candidateBBytes[i]
		}
		fakeKey = make([]byte, 32)
		for i := 0; i < 32; i++ {
			fakeKey[i] = C1[i] ^ intermediate[i]
		}
	} else {
		intermediate := make([]byte, 32)
		for i := 0; i < 32; i++ {
			intermediate[i] = C2[i] ^ candidateABytes[i]
		}
		fakeKey = make([]byte, 32)
		for i := 0; i < 32; i++ {
			fakeKey[i] = C1[i] ^ intermediate[i]
		}
	}

	return models.DeniableEncryption{
		CiphertextA: C1,
		CiphertextB: C2,
		KeyA:        realKey,
		KeyB:        fakeKey,
		ActualKey:   realKey,
		FakeKey:     fakeKey,
	}
}

func padToLength(data []byte, length int) []byte {
	if len(data) >= length {
		return data[:length]
	}

	padded := make([]byte, length)
	copy(padded, data)
	for i := len(data); i < length; i++ {
		padded[i] = 0x20
	}
	return padded
}

func doDecryption(C1, C2, key []byte) string {
	r := make([]byte, len(C1))
	for i := 0; i < len(C1); i++ {
		r[i] = C1[i] ^ key[i]
	}

	decrypted := make([]byte, len(C2))
	for i := 0; i < len(C2); i++ {
		decrypted[i] = C2[i] ^ r[i]
	}

	return string(bytes.TrimSpace(decrypted))
}

func generateDeniableSignature(voteID, candidate string, key []byte) string {
	hash := sha256.Sum256(append([]byte(voteID+candidate), key...))
	return hex.EncodeToString(hash[:])
}

func VerifySignature(voteID, candidate, signature string, key []byte) bool {
	expectedSig := generateDeniableSignature(voteID, candidate, key)
	return signature == expectedSig
}
