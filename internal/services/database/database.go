package database

import (
	"sync"

	"github.com/lvlcn-t/DevSMTP/internal/models"
)

type Database interface {
	SaveEmail(email models.Email) error
	GetEmails() ([]models.Email, error)
}

type inMemory struct {
	emails []models.Email
	mu     sync.RWMutex
}

// NewDatabase creates a new in-memory database instance
func NewDatabase() Database {
	return &inMemory{
		emails: []models.Email{},
	}
}

// SaveEmail saves an email to the in-memory database
func (db *inMemory) SaveEmail(email models.Email) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.emails = append(db.emails, email)
	return nil
}

// GetEmails retrieves all emails from the in-memory database
func (db *inMemory) GetEmails() ([]models.Email, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	// Returning a copy of the slice to avoid external modification
	ec := make([]models.Email, len(db.emails))
	copy(ec, db.emails)
	return ec, nil
}
