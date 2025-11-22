package main

import (
	"math/big"
	"servicoEmprestimo/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type SimulateRequest struct {
	Value        float64 `json:"value"`
	Rate         float64 `json:"rate"`
	Term         int     `json:"term"`
	FirstDueDate string  `json:"first_due_date"`
}

type Installment struct {
	Number  int     `json:"number"`
	DueDate string  `json:"due_date"`
	Amount  float64 `json:"amount"`
}

type SimulateResponse struct {
	InstallmentValue float64       `json:"installment_value"`
	TotalAmount      float64       `json:"total_amount"`
	Schedule         []Installment `json:"schedule"`
}

func main() {
	app := fiber.New()

	app.Post("/simular", func(c *fiber.Ctx) error {
		var req SimulateRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
		}

		precision := uint(128)

		pv := new(big.Float).SetPrec(precision).SetFloat64(req.Value)
		rate := new(big.Float).SetPrec(precision).SetFloat64(req.Rate)

		installment := service.MonthlyPaymentFloat(pv, rate, req.Term)

		// Convertendo para float64
		installmentFloat, _ := installment.Float64()

		// Montando cronograma
		firstDate, _ := time.Parse("2006-01-02", req.FirstDueDate)
		schedule := make([]Installment, 0, req.Term)

		for i := 1; i <= req.Term; i++ {
			dueDate := firstDate.AddDate(0, i-1, 0)
			schedule = append(schedule, Installment{
				Number:  i,
				DueDate: dueDate.Format("2006-01-02"),
				Amount:  installmentFloat,
			})
		}

		total := installmentFloat * float64(req.Term)

		resp := SimulateResponse{
			InstallmentValue: installmentFloat,
			TotalAmount:      total,
			Schedule:         schedule,
		}

		return c.JSON(resp)
	})

	app.Listen(":8080")
}
