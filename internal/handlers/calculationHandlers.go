package handlers

import (
    calculationservice "calculatorback/internal/calculationService/basic"
    "net/http"
    "github.com/labstack/echo/v4"
)

type CalculationHandler struct {
    service calculationservice.CalculationService
}

func NewCalculationHandler(s calculationservice.CalculationService) *CalculationHandler {
    return &CalculationHandler{service: s}
}

// GET /calculations (все расчёты)
func (h *CalculationHandler) GetCalculations(c echo.Context) error {
    calculations, err := h.service.GetAllCalculations()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get calculations"})
    }
    return c.JSON(http.StatusOK, calculations)
}

// POST /calculations (создание нового расчёта)
func (h *CalculationHandler) PostCalculations(c echo.Context) error {
    var req calculationservice.CalculationRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }
    // Важно! Передаём весь req в сервис
    calc, err := h.service.CreateCalculation(req)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not create calculation"})
    }
    return c.JSON(http.StatusCreated, calc)
}

// PATCH /calculations/:id (обновить выражение)
func (h *CalculationHandler) PatchCalculations(c echo.Context) error {
    id := c.Param("id")
    var req calculationservice.CalculationRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }
    calc, err := h.service.UpdateCalculation(id, req.Expression)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not update calculation"})
    }
    return c.JSON(http.StatusOK, calc)
}

// DELETE /calculations/:id
func (h *CalculationHandler) DeleteCalculations(c echo.Context) error {
    id := c.Param("id")
    if err := h.service.DeleteCalculation(id); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete calculation"})
    }
    return c.NoContent(http.StatusNoContent)
}