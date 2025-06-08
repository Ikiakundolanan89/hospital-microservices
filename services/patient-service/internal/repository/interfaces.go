// Repository interfaces
// internal/repository/interfaces.go
package repository

import (
	"context"
	"patient-service/internal/domain"
)

type PatientRepository interface {
	Create(ctx context.Context, patient *domain.Patient) error
	GetByID(ctx context.Context, id string) (*domain.Patient, error)
	GetByNIK(ctx context.Context, nik string) (*domain.Patient, error)
	GetByMedicalRecordNo(ctx context.Context, mrNo string) (*domain.Patient, error)
	Update(ctx context.Context, patient *domain.Patient) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter domain.PatientFilter) ([]*domain.Patient, int, error)
	Exists(ctx context.Context, id string) (bool, error)
}
