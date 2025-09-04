package main

import (
	"calculatorback/internal/calculationService/basic"
	"calculatorback/internal/db"
	"calculatorback/internal/handlers"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	e := echo.New()

	calcRepo := calculationservice.NewCalculationRepository(database)
	calcService := calculationservice.NewCalculationService(calcRepo)
	calcHandlers := handlers.NewCalculationHandler(calcService)


	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	
	e.GET("/calculations", calcHandlers.GetCalculations)
	e.POST("/calculations", calcHandlers.PostCalculations)
	e.PATCH("/calculations/:id", calcHandlers.PatchCalculations)
	e.DELETE("/calculations/:id", calcHandlers.DeleteCalculations)

	e.Start("localhost:8080")
}