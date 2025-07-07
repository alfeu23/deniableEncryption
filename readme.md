# Deniable Encryption Voting System

A demonstration of deniable encryption capabilities in a voting context for computational security research. This system allows voters to cast votes while maintaining plausible deniability about their actual choice.

## Overview

This project implements a voting system where voters can vote for one of two candidates and later generate cryptographic "proof" that they voted for either candidate, regardless of their actual vote. This demonstrates the concept of deniable encryption where the same ciphertext can be plausibly decrypted to different plaintexts using different keys.

## How Deniable Encryption Works

The system uses XOR-based deniable encryption:

1. **Single Ciphertext**: One encrypted message is created
2. **Dual Keys**: Two different keys are generated
3. **Dual Decryption**: The same ciphertext decrypts to different candidates depending on which key is used
4. **Mathematical Validity**: Both decryptions are cryptographically valid and indistinguishable

### Example:
- `ciphertext XOR key1 = "Jair Alfeu"`
- `ciphertext XOR key2 = "Luiz Ferreira"`
- Both results appear equally legitimate

## Features

- **XOR-based deniable encryption** implementation
- **REST API** for voting and receipt generation  
- **Cryptographic receipts** with verifiable signatures
- **Deniable proof generation** for either candidate
- **Educational demonstration** of privacy-preserving cryptography

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
  "receipt_id": "789ghi012jkl", 
  "message": "Vote recorded successfully"
}
```

### GET /receipt/{receipt_id}
Retrieve a voting receipt with cryptographic proof.

**Response:**
```json
{
  "id": "789ghi012jkl",
  "vote_id": "abc123def456",
  "candidate": "Jair Alfeu",
  "timestamp": "2025-07-07T10:30:00Z",
  "signature": "sha256_hash_signature",
  "deniable_data": "hex_encoded_ciphertext"
}
```

### POST /deny/{vote_id}
Generate deniable proof for a different candidate.

**Request Body:**
```json
{
  "desired_candidate": "Luiz Ferreira"
}
```

**Response:**
```json
{
  "ciphertext": "hex_encoded_ciphertext",
  "key": "hex_encoded_key",
  "decrypted": "Luiz Ferreira",
  "message": "This proves the vote was for Luiz Ferreira"
}
```

## Project Structure

```
deniableEncryption/
├── api/
│   └── cmd/
│       └── main.go          # Application entry point
├── handlers/
│   └── handler.go           # HTTP request handlers
├── models/
│   └── models.go            # Data structures
├── static/                  # Frontend files (optional)
├── go.mod                   # Go module dependencies
```

## Installation & Running

1. **Initialize the project:**
```bash
cd /Users/deliverymuch/UFSC/deniableEncryption
go mod init deniableEncryption
go mod tidy
```

2. **Install dependencies:**
```bash
go get github.com/gorilla/mux
```

3. **Run the application:**
```bash
go run api/cmd/main.go
```

4. **Access the API:**
- Server runs on `http://localhost:3000`
- Use curl, Postman, or any HTTP client to test endpoints

## Usage Example

```bash
# Cast a vote
curl -X POST http://localhost:3000/vote \
  -H "Content-Type: application/json" \
  -d '{"candidate": "Jair Alfeu", "voter_id": "voter123"}'

# Get receipt
curl http://localhost:3000/receipt/{receipt_id}

# Generate deniable proof for the other candidate
curl -X POST http://localhost:3000/deny/{vote_id} \
  -H "Content-Type: application/json" \
  -d '{"desired_candidate": "Luiz Ferreira"}'
```

## Security Properties

- **Plausible Deniability**: Impossible to prove which candidate was actually voted for
- **Cryptographic Validity**: Both "proofs" are mathematically sound
- **Coercion Resistance**: Voters can satisfy any coercer by showing fake proof
- **Privacy Protection**: True voting preferences remain hidden

## Further Reading

- [Deniable Encryption (Canetti et al.)](https://en.wikipedia.org/wiki/Deniable_encryption)
- [Coercion-Resistant Electronic Elections](https://www.usenix.org/legacy/event/sec05/tech/full_papers/juels/juels.pdf)
- [Privacy-Preserving Cryptography](https://en.wikipedia.org/wiki/Privacy-preserving_cryptography)
