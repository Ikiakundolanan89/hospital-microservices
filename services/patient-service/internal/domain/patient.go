// Domain model
package domain

import (
	"time"
)

type Patient struct {
	ID                string    `json:"id"`
	MedicalRecordNo   string    `json:"medical_record_no"`
	NIK               string    `json:"nik"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	DateOfBirth       time.Time `json:"date_of_birth"`
	Gender            string    `json:"gender"`
	BloodType         string    `json:"blood_type"`
	Phone             string    `json:"phone"`
	Email             string    `json:"email"`
	Address           string    `json:"address"`
	City              string    `json:"city"`
	Province          string    `json:"province"`
	PostalCode        string    `json:"postal_code"`
	EmergencyContact  string    `json:"emergency_contact"`
	EmergencyPhone    string    `json:"emergency_phone"`
	InsuranceProvider string    `json:"insurance_provider"`
	InsuranceNumber   string    `json:"insurance_number"`
	Allergies         string    `json:"allergies"`
	ChronicConditions string    `json:"chronic_conditions"`
	IsActive          bool      `json:"is_active"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	CreatedBy         string    `json:"created_by"`
	UpdatedBy         string    `json:"updated_by"`
}

// PatientFilter untuk query filtering
type PatientFilter struct {
	Search   string
	City     string
	Province string
	IsActive *bool
	Page     int
	Limit    int
	Sort     string
	Order    string // ASC or DESC
}
