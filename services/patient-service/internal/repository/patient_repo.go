// Patient repository implementation
// internal/repository/patient_repo.go
package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"patient-service/internal/domain"

	"github.com/google/uuid"
)

type patientRepository struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) PatientRepository {
	return &patientRepository{db: db}
}

func (r *patientRepository) Create(ctx context.Context, patient *domain.Patient) error {
	patient.ID = uuid.New().String()
	patient.CreatedAt = time.Now()
	patient.UpdatedAt = time.Now()

	query := `
		INSERT INTO patients (
			id, medical_record_no, nik, first_name, last_name,
			date_of_birth, gender, blood_type, phone, email,
			address, city, province, postal_code,
			emergency_contact, emergency_phone,
			insurance_provider, insurance_number,
			allergies, chronic_conditions,
			is_active, created_at, updated_at, created_by, updated_by
		) VALUES (
			@p1, @p2, @p3, @p4, @p5,
			@p6, @p7, @p8, @p9, @p10,
			@p11, @p12, @p13, @p14,
			@p15, @p16, @p17, @p18,
			@p19, @p20, @p21, @p22, @p23, @p24, @p25
		)
	`

	_, err := r.db.ExecContext(ctx, query,
		patient.ID, patient.MedicalRecordNo, patient.NIK, patient.FirstName, patient.LastName,
		patient.DateOfBirth, patient.Gender, patient.BloodType, patient.Phone, patient.Email,
		patient.Address, patient.City, patient.Province, patient.PostalCode,
		patient.EmergencyContact, patient.EmergencyPhone,
		patient.InsuranceProvider, patient.InsuranceNumber,
		patient.Allergies, patient.ChronicConditions,
		patient.IsActive, patient.CreatedAt, patient.UpdatedAt, patient.CreatedBy, patient.UpdatedBy,
	)

	return err
}

func (r *patientRepository) GetByID(ctx context.Context, id string) (*domain.Patient, error) {
	query := `
		SELECT 
			id, medical_record_no, nik, first_name, last_name,
			date_of_birth, gender, blood_type, phone, email,
			address, city, province, postal_code,
			emergency_contact, emergency_phone,
			insurance_provider, insurance_number,
			allergies, chronic_conditions,
			is_active, created_at, updated_at, created_by, updated_by
		FROM patients
		WHERE id = @p1 AND is_active = 1
	`

	patient := &domain.Patient{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&patient.ID, &patient.MedicalRecordNo, &patient.NIK, &patient.FirstName, &patient.LastName,
		&patient.DateOfBirth, &patient.Gender, &patient.BloodType, &patient.Phone, &patient.Email,
		&patient.Address, &patient.City, &patient.Province, &patient.PostalCode,
		&patient.EmergencyContact, &patient.EmergencyPhone,
		&patient.InsuranceProvider, &patient.InsuranceNumber,
		&patient.Allergies, &patient.ChronicConditions,
		&patient.IsActive, &patient.CreatedAt, &patient.UpdatedAt, &patient.CreatedBy, &patient.UpdatedBy,
	)

	if err == sql.ErrNoRows {
		return nil, domain.ErrPatientNotFound
	}

	return patient, err
}

func (r *patientRepository) GetByNIK(ctx context.Context, nik string) (*domain.Patient, error) {
	query := `
		SELECT 
			id, medical_record_no, nik, first_name, last_name,
			date_of_birth, gender, blood_type, phone, email,
			address, city, province, postal_code,
			emergency_contact, emergency_phone,
			insurance_provider, insurance_number,
			allergies, chronic_conditions,
			is_active, created_at, updated_at, created_by, updated_by
		FROM patients
		WHERE nik = @p1 AND is_active = 1
	`

	patient := &domain.Patient{}
	err := r.db.QueryRowContext(ctx, query, nik).Scan(
		&patient.ID, &patient.MedicalRecordNo, &patient.NIK, &patient.FirstName, &patient.LastName,
		&patient.DateOfBirth, &patient.Gender, &patient.BloodType, &patient.Phone, &patient.Email,
		&patient.Address, &patient.City, &patient.Province, &patient.PostalCode,
		&patient.EmergencyContact, &patient.EmergencyPhone,
		&patient.InsuranceProvider, &patient.InsuranceNumber,
		&patient.Allergies, &patient.ChronicConditions,
		&patient.IsActive, &patient.CreatedAt, &patient.UpdatedAt, &patient.CreatedBy, &patient.UpdatedBy,
	)

	if err == sql.ErrNoRows {
		return nil, domain.ErrPatientNotFound
	}

	return patient, err
}

func (r *patientRepository) GetByMedicalRecordNo(ctx context.Context, mrNo string) (*domain.Patient, error) {
	query := `
		SELECT 
			id, medical_record_no, nik, first_name, last_name,
			date_of_birth, gender, blood_type, phone, email,
			address, city, province, postal_code,
			emergency_contact, emergency_phone,
			insurance_provider, insurance_number,
			allergies, chronic_conditions,
			is_active, created_at, updated_at, created_by, updated_by
		FROM patients
		WHERE medical_record_no = @p1 AND is_active = 1
	`

	patient := &domain.Patient{}
	err := r.db.QueryRowContext(ctx, query, mrNo).Scan(
		&patient.ID, &patient.MedicalRecordNo, &patient.NIK, &patient.FirstName, &patient.LastName,
		&patient.DateOfBirth, &patient.Gender, &patient.BloodType, &patient.Phone, &patient.Email,
		&patient.Address, &patient.City, &patient.Province, &patient.PostalCode,
		&patient.EmergencyContact, &patient.EmergencyPhone,
		&patient.InsuranceProvider, &patient.InsuranceNumber,
		&patient.Allergies, &patient.ChronicConditions,
		&patient.IsActive, &patient.CreatedAt, &patient.UpdatedAt, &patient.CreatedBy, &patient.UpdatedBy,
	)

	if err == sql.ErrNoRows {
		return nil, domain.ErrPatientNotFound
	}

	return patient, err
}

func (r *patientRepository) Update(ctx context.Context, patient *domain.Patient) error {
	patient.UpdatedAt = time.Now()

	query := `
		UPDATE patients SET
			medical_record_no = @p2,
			nik = @p3,
			first_name = @p4,
			last_name = @p5,
			date_of_birth = @p6,
			gender = @p7,
			blood_type = @p8,
			phone = @p9,
			email = @p10,
			address = @p11,
			city = @p12,
			province = @p13,
			postal_code = @p14,
			emergency_contact = @p15,
			emergency_phone = @p16,
			insurance_provider = @p17,
			insurance_number = @p18,
			allergies = @p19,
			chronic_conditions = @p20,
			updated_at = @p21,
			updated_by = @p22
		WHERE id = @p1
	`

	result, err := r.db.ExecContext(ctx, query,
		patient.ID, patient.MedicalRecordNo, patient.NIK, patient.FirstName, patient.LastName,
		patient.DateOfBirth, patient.Gender, patient.BloodType, patient.Phone, patient.Email,
		patient.Address, patient.City, patient.Province, patient.PostalCode,
		patient.EmergencyContact, patient.EmergencyPhone,
		patient.InsuranceProvider, patient.InsuranceNumber,
		patient.Allergies, patient.ChronicConditions,
		patient.UpdatedAt, patient.UpdatedBy,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrPatientNotFound
	}

	return nil
}

func (r *patientRepository) Delete(ctx context.Context, id string) error {
	// Soft delete
	query := `UPDATE patients SET is_active = 0, updated_at = @p2 WHERE id = @p1`

	result, err := r.db.ExecContext(ctx, query, id, time.Now())
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrPatientNotFound
	}

	return nil
}

func (r *patientRepository) List(ctx context.Context, filter domain.PatientFilter) ([]*domain.Patient, int, error) {
	// Build dynamic query
	var conditions []string
	var args []interface{}
	argCount := 1

	baseQuery := `FROM patients WHERE is_active = 1`

	if filter.Search != "" {
		conditions = append(conditions, fmt.Sprintf(
			"(first_name LIKE @p%d OR last_name LIKE @p%d OR nik LIKE @p%d OR medical_record_no LIKE @p%d)",
			argCount, argCount, argCount, argCount,
		))
		searchParam := "%" + filter.Search + "%"
		args = append(args, searchParam)
		argCount++
	}

	if filter.City != "" {
		conditions = append(conditions, fmt.Sprintf("city = @p%d", argCount))
		args = append(args, filter.City)
		argCount++
	}

	if filter.Province != "" {
		conditions = append(conditions, fmt.Sprintf("province = @p%d", argCount))
		args = append(args, filter.Province)
		argCount++
	}

	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	// Count total records
	countQuery := "SELECT COUNT(*) " + baseQuery
	var total int
	err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Add sorting
	orderBy := " ORDER BY created_at DESC"
	if filter.Sort != "" {
		orderBy = fmt.Sprintf(" ORDER BY %s %s", filter.Sort, filter.Order)
	}

	// Add pagination
	offset := (filter.Page - 1) * filter.Limit
	paginationQuery := fmt.Sprintf(" OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", offset, filter.Limit)

	// Final query
	selectQuery := `
		SELECT 
			id, medical_record_no, nik, first_name, last_name,
			date_of_birth, gender, blood_type, phone, email,
			address, city, province, postal_code,
			emergency_contact, emergency_phone,
			insurance_provider, insurance_number,
			allergies, chronic_conditions,
			is_active, created_at, updated_at, created_by, updated_by
		` + baseQuery + orderBy + paginationQuery

	rows, err := r.db.QueryContext(ctx, selectQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var patients []*domain.Patient
	for rows.Next() {
		patient := &domain.Patient{}
		err := rows.Scan(
			&patient.ID, &patient.MedicalRecordNo, &patient.NIK, &patient.FirstName, &patient.LastName,
			&patient.DateOfBirth, &patient.Gender, &patient.BloodType, &patient.Phone, &patient.Email,
			&patient.Address, &patient.City, &patient.Province, &patient.PostalCode,
			&patient.EmergencyContact, &patient.EmergencyPhone,
			&patient.InsuranceProvider, &patient.InsuranceNumber,
			&patient.Allergies, &patient.ChronicConditions,
			&patient.IsActive, &patient.CreatedAt, &patient.UpdatedAt, &patient.CreatedBy, &patient.UpdatedBy,
		)
		if err != nil {
			return nil, 0, err
		}
		patients = append(patients, patient)
	}

	return patients, total, nil
}

func (r *patientRepository) Exists(ctx context.Context, id string) (bool, error) {
	query := `SELECT COUNT(*) FROM patients WHERE id = @p1 AND is_active = 1`

	var count int
	err := r.db.QueryRowContext(ctx, query, id).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
