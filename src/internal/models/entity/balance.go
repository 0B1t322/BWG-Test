package entity

import "github.com/google/uuid"

type Balance struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	Balance float64
}
