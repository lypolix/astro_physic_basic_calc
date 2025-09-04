package calculationservice

import (
    "fmt"
    "github.com/vjeantet/govaluate"
)



func CalculateBasic(expression string) (string, error) {
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