package calculationservice

import (
    "github.com/google/uuid"
    "fmt"
    "github.com/vjeantet/govaluate"
)


type CalculationService interface {
	CreateCalculation(req CalculationRequest) (Calculation, error)
    GetAllCalculations() ([]Calculation, error)
    GetCalculationByID(id string) (Calculation, error)
    GetCalculationsByType(calcType string) ([]Calculation, error)
    UpdateCalculation(id string, expression string) (Calculation, error)
    DeleteCalculation(id string) error
}


type calcService struct {
    repo CalculationRepository
}


func NewCalculationService(r CalculationRepository) CalculationService {
    return &calcService{repo: r}
}

func (s *calcService) calculateExpression(expression string) (string, error) {
    expr, err := govaluate.NewEvaluableExpression(expression)
    if err != nil {
        return "", err
    }
    result, err := expr.Evaluate(nil)
    if err != nil {
        return "", err
    }
    return fmt.Sprintf("%v", result), nil
}


func (s *calcService) CreateCalculation(req CalculationRequest) (Calculation, error) {
    var result string
    var err error

    switch req.Type {
    case "basic":
        result, err = CalculateBasic(req.Expression)
    //case "astro":
    //     result, err = CalculateAstro(req.Expression)
    // case "physics":
    //     result, err = CalculatePhysics(req.Expression)
    default:
        result, err = CalculateBasic(req.Expression) // чтоб не сломать старый код
    }

    if err != nil {
        return Calculation{}, err
    }

    calc := Calculation{
        ID:         uuid.NewString(),
        Expression: req.Expression,
        Type:       req.Type,
        Result:     result,
    }

    if err := s.repo.CreateCalculation(calc); err != nil {
        return Calculation{}, err
    }

    return calc, nil
}           

func (s *calcService) GetAllCalculations() ([]Calculation, error) {
    return s.repo.GetAllCalculations()
}

func (s *calcService) GetCalculationByID(id string) (Calculation, error) {
    return s.repo.GetCalculationByID(id)
}

func (s *calcService) GetCalculationsByType(calcType string) ([]Calculation, error) {
    return s.repo.GetCalculationsByType(calcType)
}

func (s *calcService) UpdateCalculation(id string, expression string) (Calculation, error) {
    calc, err := s.repo.GetCalculationByID(id)
    if err != nil {
        return Calculation{}, err
    }

    result, err := s.calculateExpression(expression)
    if err != nil {
        return Calculation{}, err
    }

    calc.Expression = expression
    calc.Result = result

    if err := s.repo.UpdateCalculation(calc); err != nil {
        return Calculation{}, err
    }

    return calc, nil
}

func (s *calcService) DeleteCalculation(id string) error {
    return s.repo.DeleteCalculation(id)
}