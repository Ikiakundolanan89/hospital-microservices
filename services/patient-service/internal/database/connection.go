// Database connection
// internal/database/connection.go
package database

import (
	"database/sql"
	"fmt"
	"time"

	"patient-service/internal/config"

	_ "github.com/denisenkom/go-mssqldb"
)

func NewConnection(cfg config.DatabaseConfig) (*sql.DB, error) {
	// Connection string untuk SQL Server
	connString := fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s;encrypt=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
	)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Create tables if not exists
	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	query := `
	IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='patients' AND xtype='U')
	CREATE TABLE patients (
		id NVARCHAR(50) PRIMARY KEY,
		medical_record_no NVARCHAR(50) UNIQUE NOT NULL,
		nik NVARCHAR(16) UNIQUE NOT NULL,
		first_name NVARCHAR(100) NOT NULL,
		last_name NVARCHAR(100),
		date_of_birth DATE NOT NULL,
		gender NVARCHAR(10) NOT NULL,
		blood_type NVARCHAR(5),
		phone NVARCHAR(20) NOT NULL,
		email NVARCHAR(100),
		address NVARCHAR(255),
		city NVARCHAR(100),
		province NVARCHAR(100),
		postal_code NVARCHAR(10),
		emergency_contact NVARCHAR(100),
		emergency_phone NVARCHAR(20),
		insurance_provider NVARCHAR(100),
		insurance_number NVARCHAR(50),
		allergies NVARCHAR(MAX),
		chronic_conditions NVARCHAR(MAX),
		is_active BIT DEFAULT 1,
		created_at DATETIME2 DEFAULT GETDATE(),
		updated_at DATETIME2 DEFAULT GETDATE(),
		created_by NVARCHAR(50),
		updated_by NVARCHAR(50)
	);

	-- Create indexes
	IF NOT EXISTS (SELECT * FROM sys.indexes WHERE name = 'idx_patients_nik')
		CREATE INDEX idx_patients_nik ON patients(nik);
	
	IF NOT EXISTS (SELECT * FROM sys.indexes WHERE name = 'idx_patients_medical_record')
		CREATE INDEX idx_patients_medical_record ON patients(medical_record_no);
	
	IF NOT EXISTS (SELECT * FROM sys.indexes WHERE name = 'idx_patients_name')
		CREATE INDEX idx_patients_name ON patients(first_name, last_name);
	`

	_, err := db.Exec(query)
	return err
}
