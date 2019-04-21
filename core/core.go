package core

import (
	"../db"
	"crypto/sha256"
	"encoding/hex"
	"github.com/op/go-logging"
	"sync"
)

var (
	log = logging.MustGetLogger("core")
	WG  sync.WaitGroup
)

func ProcessHash(id int, payload string, rounds int) {
	//decrease WaitGroup count
	defer WG.Done()
	log.Debug("Start hash job #", id)
	// calculate hash using hash rounds count
	hash := calculateHash(payload, rounds)
	log.Debug("Finish hash job #", id)
	// update job in db
	if err := db.Update(id, hash); err != nil {
		log.Errorf("db.Update: %v", err)
	}
}

func calculateHash(payload string, rounds int) string {
	hash := payload
	for i := 0; i < rounds; i++ {
		h := sha256.Sum256([]byte(hash))
		hash = hex.EncodeToString(h[:])
	}
	return hash
}
