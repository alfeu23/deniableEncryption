package models

import (
	"time"
)

type Candidate string

const (
	CandidateA Candidate = "Jair Alfeu"
	CandidateB Candidate = "Luiz Ferreira"
)

type Vote struct {
	ID        string    `json:"id"`
	Candidate Candidate `json:"candidate"`
	Timestamp time.Time `json:"timestamp"`
	VoterID   string    `json:"voter_id"`
}

type Receipt struct {
	ID           string    `json:"id"`
	VoteID       string    `json:"vote_id"`
	Candidate    Candidate `json:"candidate"`
	Timestamp    time.Time `json:"timestamp"`
	Signature    string    `json:"signature"`
	DeniableData string    `json:"deniable_data"`
}

type DeniableEncryption struct {
	CiphertextA []byte // Encryption that could decrypt to Candidate A
	CiphertextB []byte // Encryption that could decrypt to Candidate B
	KeyA        []byte // Key that reveals vote for A
	KeyB        []byte // Key that reveals vote for B
	ActualKey   []byte // The real key used
	FakeKey     []byte // The fake key for deniability
}
