# Deniable Encryption Voting System

A demonstration of deniable encryption capabilities in a voting context for computational security research. This system allows voters to cast votes while maintaining plausible deniability about their actual choice through cryptographic receipts.

## Overview

This project implements a voting system where voters can vote for one of two candidates (`Jair Alfeu` or `Luiz Ferreira`) and receive both real and fake receipts. The system demonstrates deniable encryption where voters can produce cryptographic "proof" that they voted for either candidate, regardless of their actual vote.

## How Deniable Encryption Works

The system uses a dual-ciphertext XOR-based deniable encryption scheme:

1. **Dual Ciphertexts**: Two ciphertexts (C1, C2) are generated for each vote
2. **Dual Keys**: Two different keys are created - one real, one fake
3. **Plausible Decryption**: The same ciphertext pair can decrypt to either candidate depending on which key is used
4. **Cryptographic Receipts**: Both receipts are mathematically valid and indistinguishable

### Technical Implementation:
- `C1 = nonce ⊕ realKey`
- `C2 = candidateVote ⊕ nonce`
- `fakeKey` computed to make `C1 ⊕ fakeKey = nonce'` where `C2 ⊕ nonce' = otherCandidate`

## Features

- **XOR-based deniable encryption** with dual-key generation
- **REST API** for voting, receipt retrieval, and verification
- **Dual receipt system** (real and fake receipts per vote)
- **Cryptographic signatures** for receipt authenticity
- **Election results** aggregation
- **Voter registry** to prevent double voting
- **Administrative cache flushing**

## Project Structure

```
deniableEncryption/
├── cmd/
│   └── main.go              # Application entry point
├── handlers/
│   ├── election.go          # Election results handler
│   ├── flush.go             # Cache management
│   ├── helper.go            # Crypto utilities & storage
│   ├── receit.go            # Receipt retrieval & verification
│   └── vote.go              # Vote submission
├── models/
│   └── model.go             # Data structures
├── routes/
│   └── routes.go            # HTTP route definitions
├── coverage/                # Test coverage reports
├── go.mod                   # Go module dependencies
└── readme.md               # This file
```

## API Endpoints

### POST /vote
Submit a vote for a candidate.

**Request Body:**
```json
{
  "candidate": "Jair Alfeu",
  "voter_id": "voter123"
}
```

**Response:**
```json
{
  "vote_id": "abc123def456",
  "real_receipt_id": "real789ghi012jkl",
  "fake_receipt_id": "fake456def789abc", 
  "message": "Vote recorded successfully with both receipts"
}
```

### GET /receipt/{receipt_id}
Retrieve a voting receipt (real or fake).

**Response:**
```json
{
  "id": "789ghi012jkl",
  "vote_id": "abc123def456",
  "candidate": "Jair Alfeu",
  "timestamp": "2025-01-27T10:30:00Z",
  "signature": "sha256_hash_signature",
  "deniable_data": "hex_encoded_ciphertext"
}
```

### GET /verify/{receipt_id}
Verify a receipt and get cryptographic proof.

**Response:**
```json
{
  "C1": "hex_encoded_ciphertext1",
  "C2": "hex_encoded_ciphertext2",
  "key": "hex_encoded_key",
  "decrypted": "Jair Alfeu",
  "message": "This proves the vote was for Jair Alfeu",
  "type": "real_receipt"
}
```

### GET /election/results
Get current election results.

**Response:**
```json
{
  "total_votes": 42,
  "results": {
    "Jair Alfeu": 23,
    "Luiz Ferreira": 19
  },
  "winner": "Jair Alfeu"
}
```

### POST /admin/flush
Reset all voting data (administrative endpoint).

**Response:**
```json
{
  "message": "Cache flushed successfully",
  "status": "success"
}
```

## Installation & Running

1. **Install dependencies:**
```bash
go mod tidy
```

2. **Run the application:**
```bash
go run cmd/main.go
```

3. **Access the API:**
- Server runs on `http://localhost:3000`
- Use curl, Postman, or any HTTP client to test endpoints

## Usage Example

```bash
# Cast a vote
curl -X POST http://localhost:3000/vote \
  -H "Content-Type: application/json" \
  -d '{"candidate": "Jair Alfeu", "voter_id": "voter123"}'

# Get a receipt (real or fake)
curl http://localhost:3000/receipt/{receipt_id}

# Verify receipt and get cryptographic proof
curl http://localhost:3000/verify/{receipt_id}

# Check election results
curl http://localhost:3000/election/results

# Reset system (admin only)
curl -X POST http://localhost:3000/admin/flush
```

## Security Properties

- **Plausible Deniability**: Voters receive both real and fake receipts, making it impossible to prove which candidate they actually voted for
- **Cryptographic Validity**: Both receipts have valid signatures and can be verified
- **Coercion Resistance**: Voters can show either receipt to satisfy coercers
- **Double Voting Prevention**: Voter registry prevents multiple votes from same voter ID
- **Privacy Protection**: True voting preferences remain cryptographically hidden

## Key Components

- **[`models.DeniableEncryption`](models/model.go)**: Core encryption structure with dual keys
- **[`handlers.createDeniableEncryption`](handlers/helper.go)**: Generates deniable encryption for votes
- **[`handlers.doDecryption`](handlers/helper.go)**: Decrypts ciphertext with provided key
- **[`handlers.generateDeniableSignature`](handlers/helper.go)**: Creates cryptographic signatures
- **[`routes.SetupRoutes`](routes/routes.go)**: HTTP route configuration

## Dependencies

- **[gorilla/mux](https://github.com/gorilla/mux)**: HTTP router for REST API
- **Go standard library**: crypto/rand, crypto/sha256, encoding/hex, etc.

## Research Context

This implementation demonstrates key concepts in:
- **Deniable Encryption**: Cryptographic systems that allow plausible denial
- **Coercion-Resistant Voting**: Protecting voter privacy against coercion
- **Privacy-Preserving Cryptography**: Maintaining confidentiality while enabling verification
