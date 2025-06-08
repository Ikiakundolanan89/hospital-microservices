// HTTP handlers
// internal/handler/patient_handler.go
package handler

import (
	"math"
	"patient-service/internal/domain"
	"patient-service/internal/dto"
	"patient-service/internal/service"
	"patient-service/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PatientHandler struct {
	patientService service.PatientService
	validator      *validator.Validate
}

func NewPatientHandler(patientService service.PatientService, validator *validator.Validate) *PatientHandler {
	return &PatientHandler{
		patientService: patientService,
		validator:      validator,
	}
}

// CreatePatient godoc
// @Summary Create a new patient
// @Description Create a new patient record
// @Tags patients
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body dto.CreatePatientRequest true "Patient data"
// @Success 201 {object} dto.PatientResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/patients [post]
func (h *PatientHandler) CreatePatient(c *fiber.Ctx) error {
	var req dto.CreatePatientRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "INVALID_REQUEST", "Invalid request body", err.Error())
	}

	// Validate request
	if err := h.validator.Struct(&req); err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	// Get user info from JWT context
	userID := c.Locals("userID").(string)

	// Convert to domain model
	patient := dto.ToPatientDomain(&req)
	patient.CreatedBy = userID
	patient.UpdatedBy = userID

	// Create patient
	createdPatient, err := h.patientService.CreatePatient(c.Context(), patient)
	if err != nil {
		if customErr, ok := err.(*domain.CustomError); ok {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, customErr.Code, customErr.Message, customErr.Details)
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "CREATE_FAILED", "Failed to create patient", err.Error())
	}

	// Return response
	return c.Status(fiber.StatusCreated).JSON(dto.ToPatientResponse(createdPatient))
}

// GetPatient godoc
// @Summary Get patient by ID
// @Description Get patient details by ID
// @Tags patients
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Patient ID"
// @Success 200 {object} dto.PatientResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/patients/{id} [get]
func (h *PatientHandler) GetPatient(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "INVALID_ID", "Patient ID is required", "")
	}

	patient, err := h.patientService.GetPatient(c.Context(), id)
	if err != nil {
		if err == domain.ErrPatientNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound, "NOT_FOUND", "Patient not found", "")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "GET_FAILED", "Failed to get patient", err.Error())
	}

	return c.JSON(dto.ToPatientResponse(patient))
}

// UpdatePatient godoc
// @Summary Update patient
// @Description Update patient details
// @Tags patients
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Patient ID"
// @Param request body dto.UpdatePatientRequest true "Patient data"
// @Success 200 {object} dto.PatientResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/patients/{id} [put]
func (h *PatientHandler) UpdatePatient(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "INVALID_ID", "Patient ID is required", "")
	}

	var req dto.UpdatePatientRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "INVALID_REQUEST", "Invalid request body", err.Error())
	}

	// Validate request
	if err := h.validator.Struct(&req); err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	// Get user info from JWT context
	userID := c.Locals("userID").(string)

	// Convert to domain model
	patient := dto.ToUpdatePatientDomain(id, &req)
	patient.UpdatedBy = userID

	// Update patient
	updatedPatient, err := h.patientService.UpdatePatient(c.Context(), patient)
	if err != nil {
		if err == domain.ErrPatientNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound, "NOT_FOUND", "Patient not found", "")
		}
		if customErr, ok := err.(*domain.CustomError); ok {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, customErr.Code, customErr.Message, customErr.Details)
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "UPDATE_FAILED", "Failed to update patient", err.Error())
	}

	return c.JSON(dto.ToPatientResponse(updatedPatient))
}

// DeletePatient godoc
// @Summary Delete patient
// @Description Soft delete patient
// @Tags patients
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Patient ID"
// @Success 200 {object} dto.SuccessResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/patients/{id} [delete]
func (h *PatientHandler) DeletePatient(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "INVALID_ID", "Patient ID is required", "")
	}

	err := h.patientService.DeletePatient(c.Context(), id)
	if err != nil {
		if err == domain.ErrPatientNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound, "NOT_FOUND", "Patient not found", "")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "DELETE_FAILED", "Failed to delete patient", err.Error())
	}

	return c.JSON(dto.SuccessResponse{
		Message: "Patient deleted successfully",
	})
}

// ListPatients godoc
// @Summary List patients
// @Description Get list of patients with pagination and filtering
// @Tags patients
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param search query string false "Search by name, NIK, or medical record number"
// @Param city query string false "Filter by city"
// @Param province query string false "Filter by province"
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Items per page (default: 10, max: 100)"
// @Param sort query string false "Sort field (created_at, updated_at, first_name, last_name, nik)"
// @Param order query string false "Sort order (ASC, DESC)"
// @Success 200 {object} dto.ListPatientsResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/patients [get]
func (h *PatientHandler) ListPatients(c *fiber.Ctx) error {
	var req dto.ListPatientsRequest

	// Parse query parameters
	if err := c.QueryParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "INVALID_REQUEST", "Invalid query parameters", err.Error())
	}

	// Set defaults
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 10
	}

	// Validate request
	if err := h.validator.Struct(&req); err != nil {
		return utils.ValidationErrorResponse(c, err)
	}

	// Create filter
	filter := domain.PatientFilter{
		Search:   req.Search,
		City:     req.City,
		Province: req.Province,
		IsActive: req.IsActive,
		Page:     req.Page,
		Limit:    req.Limit,
		Sort:     req.Sort,
		Order:    req.Order,
	}

	// Get patients
	patients, total, err := h.patientService.ListPatients(c.Context(), filter)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "LIST_FAILED", "Failed to list patients", err.Error())
	}

	// Convert to response
	patientResponses := make([]*dto.PatientResponse, len(patients))
	for i, patient := range patients {
		patientResponses[i] = dto.ToPatientResponse(patient)
	}

	// Create pagination response
	totalPages := int(math.Ceil(float64(total) / float64(req.Limit)))

	response := dto.ListPatientsResponse{
		Data: patientResponses,
		Pagination: dto.PaginationResponse{
			Page:       req.Page,
			Limit:      req.Limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}

	return c.JSON(response)
}

// GetPatientPublicInfo - Handler function tanpa struct
func GetPatientPublicInfo(patientService service.PatientService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "INVALID_ID", "Patient ID is required", "")
		}

		patient, err := patientService.GetPatientPublicInfo(c.Context(), id)
		if err != nil {
			if err == domain.ErrPatientNotFound {
				return utils.ErrorResponse(c, fiber.StatusNotFound, "NOT_FOUND", "Patient not found", "")
			}
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "GET_FAILED", "Failed to get patient info", err.Error())
		}

		return c.JSON(patient)
	}
}
