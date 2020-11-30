package database

import (
	"math/rand"
	"time"
)

type DB struct {
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) Load() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(10*n) * time.Millisecond)
	return n
}
