// Sample test file: internal/service/patient_service_test.go
package service

import (
	"context"
	"testing"
	"time"

	"patient-service/internal/domain"
	"patient-service/internal/repository"
)

// MockPatientRepository for testing
type mockPatientRepository struct {
	patients map[string]*domain.Patient
}

func NewMockPatientRepository() repository.PatientRepository {
	return &mockPatientRepository{
		patients: make(map[string]*domain.Patient),
	}
}

func (m *mockPatientRepository) Create(ctx context.Context, patient *domain.Patient) error {
	m.patients[patient.ID] = patient
	return nil
}

func (m *mockPatientRepository) GetByID(ctx context.Context, id string) (*domain.Patient, error) {
	patient, exists := m.patients[id]
	if !exists {
		return nil, domain.ErrPatientNotFound
	}
	return patient, nil
}

func (m *mockPatientRepository) GetByNIK(ctx context.Context, nik string) (*domain.Patient, error) {
	for _, patient := range m.patients {
		if patient.NIK == nik {
			return patient, nil
		}
	}
	return nil, domain.ErrPatientNotFound
}

func (m *mockPatientRepository) GetByMedicalRecordNo(ctx context.Context, mrNo string) (*domain.Patient, error) {
	for _, patient := range m.patients {
		if patient.MedicalRecordNo == mrNo {
			return patient, nil
		}
	}
	return nil, domain.ErrPatientNotFound
}

func (m *mockPatientRepository) Update(ctx context.Context, patient *domain.Patient) error {
	if _, exists := m.patients[patient.ID]; !exists {
		return domain.ErrPatientNotFound
	}
	m.patients[patient.ID] = patient
	return nil
}

func (m *mockPatientRepository) Delete(ctx context.Context, id string) error {
	if _, exists := m.patients[id]; !exists {
		return domain.ErrPatientNotFound
	}
	delete(m.patients, id)
	return nil
}

func (m *mockPatientRepository) List(ctx context.Context, filter domain.PatientFilter) ([]*domain.Patient, int, error) {
	var result []*domain.Patient
	for _, patient := range m.patients {
		result = append(result, patient)
	}
	return result, len(result), nil
}

func (m *mockPatientRepository) Exists(ctx context.Context, id string) (bool, error) {
	_, exists := m.patients[id]
	return exists, nil
}

func TestCreatePatient(t *testing.T) {
	repo := NewMockPatientRepository()
	service := NewPatientService(repo)

	patient := &domain.Patient{
		NIK:         "1234567890123456",
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Gender:      "MALE",
		Phone:       "081234567890",
		IsActive:    true,
		CreatedBy:   "test-user",
		UpdatedBy:   "test-user",
	}

	createdPatient, err := service.CreatePatient(context.Background(), patient)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if createdPatient.NIK != patient.NIK {
		t.Errorf("Expected NIK %s, got %s", patient.NIK, createdPatient.NIK)
	}

	if createdPatient.MedicalRecordNo == "" {
		t.Error("Expected medical record number to be generated")
	}
}
