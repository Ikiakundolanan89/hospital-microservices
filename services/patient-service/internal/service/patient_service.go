// Business logic
// internal/service/patient_service.go
package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"patient-service/internal/domain"
	"patient-service/internal/repository"
)

type patientService struct {
	patientRepo repository.PatientRepository
}

func NewPatientService(patientRepo repository.PatientRepository) PatientService {
	return &patientService{
		patientRepo: patientRepo,
	}
}

func (s *patientService) CreatePatient(ctx context.Context, patient *domain.Patient) (*domain.Patient, error) {
	// Validate required fields
	if err := s.validatePatient(patient); err != nil {
		return nil, err
	}

	// Check if patient with NIK already exists
	existingPatient, _ := s.patientRepo.GetByNIK(ctx, patient.NIK)
	if existingPatient != nil {
		return nil, domain.NewCustomError("PATIENT_EXISTS", "Patient with this NIK already exists", "")
	}

	// Generate medical record number if not provided
	if patient.MedicalRecordNo == "" {
		patient.MedicalRecordNo = s.generateMedicalRecordNo()
	}

	// Set default values
	if patient.IsActive {
		patient.IsActive = true
	}

	// Create patient
	if err := s.patientRepo.Create(ctx, patient); err != nil {
		return nil, fmt.Errorf("failed to create patient: %w", err)
	}

	// Return created patient
	return s.patientRepo.GetByID(ctx, patient.ID)
}

func (s *patientService) GetPatient(ctx context.Context, id string) (*domain.Patient, error) {
	if id == "" {
		return nil, domain.ErrInvalidInput
	}

	patient, err := s.patientRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (s *patientService) GetPatientByNIK(ctx context.Context, nik string) (*domain.Patient, error) {
	if nik == "" {
		return nil, domain.ErrInvalidInput
	}

	patient, err := s.patientRepo.GetByNIK(ctx, nik)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (s *patientService) UpdatePatient(ctx context.Context, patient *domain.Patient) (*domain.Patient, error) {
	if patient.ID == "" {
		return nil, domain.ErrInvalidInput
	}

	// Check if patient exists
	existing, err := s.patientRepo.GetByID(ctx, patient.ID)
	if err != nil {
		return nil, err
	}

	// Validate update data
	if err := s.validatePatient(patient); err != nil {
		return nil, err
	}

	// Check if NIK is being changed and already exists
	if patient.NIK != existing.NIK {
		existingWithNIK, _ := s.patientRepo.GetByNIK(ctx, patient.NIK)
		if existingWithNIK != nil && existingWithNIK.ID != patient.ID {
			return nil, domain.NewCustomError("NIK_EXISTS", "NIK already used by another patient", "")
		}
	}

	// Update patient
	if err := s.patientRepo.Update(ctx, patient); err != nil {
		return nil, fmt.Errorf("failed to update patient: %w", err)
	}

	// Return updated patient
	return s.patientRepo.GetByID(ctx, patient.ID)
}

func (s *patientService) DeletePatient(ctx context.Context, id string) error {
	if id == "" {
		return domain.ErrInvalidInput
	}

	// Check if patient exists
	exists, err := s.patientRepo.Exists(ctx, id)
	if err != nil {
		return err
	}

	if !exists {
		return domain.ErrPatientNotFound
	}

	// Soft delete patient
	return s.patientRepo.Delete(ctx, id)
}

func (s *patientService) ListPatients(ctx context.Context, filter domain.PatientFilter) ([]*domain.Patient, int, error) {
	// Set default pagination
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Limit <= 0 {
		filter.Limit = 10
	}
	if filter.Limit > 100 {
		filter.Limit = 100 // Max limit
	}

	// Set default sorting
	if filter.Sort == "" {
		filter.Sort = "created_at"
	}
	if filter.Order == "" {
		filter.Order = "DESC"
	}

	// Validate sort field to prevent SQL injection
	allowedSortFields := map[string]bool{
		"created_at": true,
		"updated_at": true,
		"first_name": true,
		"last_name":  true,
		"nik":        true,
	}

	if !allowedSortFields[filter.Sort] {
		filter.Sort = "created_at"
	}

	// Validate order
	filter.Order = strings.ToUpper(filter.Order)
	if filter.Order != "ASC" && filter.Order != "DESC" {
		filter.Order = "DESC"
	}

	return s.patientRepo.List(ctx, filter)
}

func (s *patientService) GetPatientPublicInfo(ctx context.Context, id string) (*domain.Patient, error) {
	if id == "" {
		return nil, domain.ErrInvalidInput
	}

	patient, err := s.patientRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Return only public information
	publicInfo := &domain.Patient{
		ID:              patient.ID,
		MedicalRecordNo: patient.MedicalRecordNo,
		FirstName:       patient.FirstName,
		LastName:        patient.LastName,
		Gender:          patient.Gender,
		// Hide sensitive information
	}

	return publicInfo, nil
}

// Helper methods

func (s *patientService) validatePatient(patient *domain.Patient) error {
	if patient.NIK == "" || len(patient.NIK) != 16 {
		return domain.NewCustomError("INVALID_NIK", "NIK must be 16 characters", "")
	}

	if patient.FirstName == "" {
		return domain.NewCustomError("INVALID_NAME", "First name is required", "")
	}

	if patient.DateOfBirth.IsZero() {
		return domain.NewCustomError("INVALID_DOB", "Date of birth is required", "")
	}

	if patient.DateOfBirth.After(time.Now()) {
		return domain.NewCustomError("INVALID_DOB", "Date of birth cannot be in the future", "")
	}

	if patient.Gender != "MALE" && patient.Gender != "FEMALE" {
		return domain.NewCustomError("INVALID_GENDER", "Gender must be MALE or FEMALE", "")
	}

	if patient.Phone == "" {
		return domain.NewCustomError("INVALID_PHONE", "Phone number is required", "")
	}

	// Validate blood type if provided
	if patient.BloodType != "" {
		validBloodTypes := map[string]bool{
			"A+": true, "A-": true,
			"B+": true, "B-": true,
			"AB+": true, "AB-": true,
			"O+": true, "O-": true,
		}

		if !validBloodTypes[patient.BloodType] {
			return domain.NewCustomError("INVALID_BLOOD_TYPE", "Invalid blood type", "")
		}
	}

	return nil
}

func (s *patientService) generateMedicalRecordNo() string {
	// Generate medical record number with format: RM-YYYYMMDD-XXXXX
	timestamp := time.Now().Format("20060102")
	random := time.Now().UnixNano() % 100000
	return fmt.Sprintf("RM-%s-%05d", timestamp, random)
}
