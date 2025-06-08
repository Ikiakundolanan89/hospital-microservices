// internal/dto/response.go
package dto

import (
	"patient-service/internal/domain"
	"time"
)

type PatientResponse struct {
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
}

type ListPatientsResponse struct {
	Data       []*PatientResponse `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

type PaginationResponse struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Converter functions
func ToPatientResponse(patient *domain.Patient) *PatientResponse {
	return &PatientResponse{
		ID:                patient.ID,
		MedicalRecordNo:   patient.MedicalRecordNo,
		NIK:               patient.NIK,
		FirstName:         patient.FirstName,
		LastName:          patient.LastName,
		DateOfBirth:       patient.DateOfBirth,
		Gender:            patient.Gender,
		BloodType:         patient.BloodType,
		Phone:             patient.Phone,
		Email:             patient.Email,
		Address:           patient.Address,
		City:              patient.City,
		Province:          patient.Province,
		PostalCode:        patient.PostalCode,
		EmergencyContact:  patient.EmergencyContact,
		EmergencyPhone:    patient.EmergencyPhone,
		InsuranceProvider: patient.InsuranceProvider,
		InsuranceNumber:   patient.InsuranceNumber,
		Allergies:         patient.Allergies,
		ChronicConditions: patient.ChronicConditions,
		IsActive:          patient.IsActive,
		CreatedAt:         patient.CreatedAt,
		UpdatedAt:         patient.UpdatedAt,
	}
}

func ToPatientDomain(req *CreatePatientRequest) *domain.Patient {
	return &domain.Patient{
		NIK:               req.NIK,
		FirstName:         req.FirstName,
		LastName:          req.LastName,
		DateOfBirth:       req.DateOfBirth,
		Gender:            req.Gender,
		BloodType:         req.BloodType,
		Phone:             req.Phone,
		Email:             req.Email,
		Address:           req.Address,
		City:              req.City,
		Province:          req.Province,
		PostalCode:        req.PostalCode,
		EmergencyContact:  req.EmergencyContact,
		EmergencyPhone:    req.EmergencyPhone,
		InsuranceProvider: req.InsuranceProvider,
		InsuranceNumber:   req.InsuranceNumber,
		Allergies:         req.Allergies,
		ChronicConditions: req.ChronicConditions,
		IsActive:          true,
	}
}

func ToUpdatePatientDomain(id string, req *UpdatePatientRequest) *domain.Patient {
	return &domain.Patient{
		ID:                id,
		NIK:               req.NIK,
		FirstName:         req.FirstName,
		LastName:          req.LastName,
		DateOfBirth:       req.DateOfBirth,
		Gender:            req.Gender,
		BloodType:         req.BloodType,
		Phone:             req.Phone,
		Email:             req.Email,
		Address:           req.Address,
		City:              req.City,
		Province:          req.Province,
		PostalCode:        req.PostalCode,
		EmergencyContact:  req.EmergencyContact,
		EmergencyPhone:    req.EmergencyPhone,
		InsuranceProvider: req.InsuranceProvider,
		InsuranceNumber:   req.InsuranceNumber,
		Allergies:         req.Allergies,
		ChronicConditions: req.ChronicConditions,
	}
}
