package calculationservice

import "time"

type Calculation struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Type       string `json:"type"`
	Result     string `json:"result"`
}

type CalculationRequest struct {
	Type       string `json:"type"`
	Expression string `json:"expression"`
}

type AstroCalculation struct {
    ID        string    `json:"id" gorm:"primaryKey"`
    Operation string    `json:"operation"`   // напр "moon_phase"
    Input     string    `json:"input"`       // Например, "2025-08-14"
    Result    string    `json:"result"`
    CreatedAt time.Time `json:"created_at"`
}

// Запрос от пользователя
type AstroCalculationRequest struct {
    Operation string `json:"operation"`    // Например, "moon_phase"
    Input     string `json:"input"`        // Например, "2025-08-17"
}