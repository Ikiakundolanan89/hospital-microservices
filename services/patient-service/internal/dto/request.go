// internal/dto/request.go
package dto

import "time"

type CreatePatientRequest struct {
	NIK               string    `json:"nik" validate:"required,len=16"`
	FirstName         string    `json:"first_name" validate:"required,min=2,max=100"`
	LastName          string    `json:"last_name" validate:"max=100"`
	DateOfBirth       time.Time `json:"date_of_birth" validate:"required"`
	Gender            string    `json:"gender" validate:"required,oneof=MALE FEMALE"`
	BloodType         string    `json:"blood_type" validate:"omitempty,oneof=A+ A- B+ B- AB+ AB- O+ O-"`
	Phone             string    `json:"phone" validate:"required,min=10,max=20"`
	Email             string    `json:"email" validate:"omitempty,email"`
	Address           string    `json:"address" validate:"max=255"`
	City              string    `json:"city" validate:"max=100"`
	Province          string    `json:"province" validate:"max=100"`
	PostalCode        string    `json:"postal_code" validate:"max=10"`
	EmergencyContact  string    `json:"emergency_contact" validate:"max=100"`
	EmergencyPhone    string    `json:"emergency_phone" validate:"max=20"`
	InsuranceProvider string    `json:"insurance_provider" validate:"max=100"`
	InsuranceNumber   string    `json:"insurance_number" validate:"max=50"`
	Allergies         string    `json:"allergies"`
	ChronicConditions string    `json:"chronic_conditions"`
}

type UpdatePatientRequest struct {
	NIK               string    `json:"nik" validate:"required,len=16"`
	FirstName         string    `json:"first_name" validate:"required,min=2,max=100"`
	LastName          string    `json:"last_name" validate:"max=100"`
	DateOfBirth       time.Time `json:"date_of_birth" validate:"required"`
	Gender            string    `json:"gender" validate:"required,oneof=MALE FEMALE"`
	BloodType         string    `json:"blood_type" validate:"omitempty,oneof=A+ A- B+ B- AB+ AB- O+ O-"`
	Phone             string    `json:"phone" validate:"required,min=10,max=20"`
	Email             string    `json:"email" validate:"omitempty,email"`
	Address           string    `json:"address" validate:"max=255"`
	City              string    `json:"city" validate:"max=100"`
	Province          string    `json:"province" validate:"max=100"`
	PostalCode        string    `json:"postal_code" validate:"max=10"`
	EmergencyContact  string    `json:"emergency_contact" validate:"max=100"`
	EmergencyPhone    string    `json:"emergency_phone" validate:"max=20"`
	InsuranceProvider string    `json:"insurance_provider" validate:"max=100"`
	InsuranceNumber   string    `json:"insurance_number" validate:"max=50"`
	Allergies         string    `json:"allergies"`
	ChronicConditions string    `json:"chronic_conditions"`
}

type ListPatientsRequest struct {
	Search   string `query:"search"`
	City     string `query:"city"`
	Province string `query:"province"`
	IsActive *bool  `query:"is_active"`
	Page     int    `query:"page" validate:"min=1"`
	Limit    int    `query:"limit" validate:"min=1,max=100"`
	Sort     string `query:"sort" validate:"omitempty,oneof=created_at updated_at first_name last_name nik"`
	Order    string `query:"order" validate:"omitempty,oneof=ASC DESC asc desc"`
}
