package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type EmailValidate struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Token     string    `json:"-" db:"-"`
	TokenHash string    `json:"token_hash" db:"token_hash"`
}

// String is not required by pop and may be deleted
func (e EmailValidate) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// EmailValidates is not required by pop and may be deleted
type EmailValidates []EmailValidate

// String is not required by pop and may be deleted
func (e EmailValidates) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Create wraps up the pattern of encrypting the password and
// running validations. Useful when writing tests.
func (ev *EmailValidate) Create(tx *pop.Connection) (*validate.Errors, error) {

	if strings.Trim(ev.Token, " ") != "" {
		ev.TokenHash = hmacl.hash(ev.Token)
		ev.Token = ""
	}

	all, err := tx.ValidateAndCreate(ev)
	if err != nil {
		return all, err
	}
	return &validate.Errors{}, nil
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *EmailValidate) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: e.UserID, Name: "UserID"},
		&validators.StringIsPresent{Field: e.TokenHash, Name: "TokenHash"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *EmailValidate) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *EmailValidate) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
