package models

import (
	"fmt"
	"time"
)

const (
	NivelFacil         = "facil"
	NivelIntermediario = "intermediario"
	NivelAvancado      = "avançado"
)

type User struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	Email          string `gorm:"unique"`
	Password       string
	RoleID         uint
	PersonalDataID uint
	AddressID      uint
	IsActive       bool `gorm:"default:true"`
	LastLogin      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Role struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	LevelAccess int `gorm:"default:1"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Project struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Nivel       string
	StartDate   time.Time
	EndDate     time.Time
	OwnerID     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PersonalData struct {
	ID        uint `gorm:"primaryKey"`
	Phone     string
	Birthdate time.Time
	Gender    string
	CPF       string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Address struct {
	ID         uint `gorm:"primaryKey"`
	Street     string
	Number     string
	Complement string
	City       string
	State      string
	Country    string
	PostalCode string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (p *Project) SetNivel(nivel string) error {
	switch nivel {
	case NivelFacil, NivelIntermediario, NivelAvancado:
		p.Nivel = nivel
		return nil
	default:
		return fmt.Errorf("invalid nivel: %s", nivel)
	}
}
