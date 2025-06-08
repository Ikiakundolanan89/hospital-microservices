// Service interfaces
// internal/service/interfaces.go
package service

import (
	"context"
	"patient-service/internal/domain"
)

type PatientService interface {
	CreatePatient(ctx context.Context, patient *domain.Patient) (*domain.Patient, error)
	GetPatient(ctx context.Context, id string) (*domain.Patient, error)
	GetPatientByNIK(ctx context.Context, nik string) (*domain.Patient, error)
	UpdatePatient(ctx context.Context, patient *domain.Patient) (*domain.Patient, error)
	DeletePatient(ctx context.Context, id string) error
	ListPatients(ctx context.Context, filter domain.PatientFilter) ([]*domain.Patient, int, error)
	GetPatientPublicInfo(ctx context.Context, id string) (*domain.Patient, error)
}
