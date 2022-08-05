package calculator_test

import (
	"fmt"

	"github.com/cucumber/godog"
	ca "github.com/felixgiftinfo/calculator-bbd"
)

type CalculatorSuite struct {
	*godog.ScenarioContext
	calc *ca.Calculator
}

func (cs *CalculatorSuite) calculatorIsCleared() error {
	cs.calc.Clear()
	return nil
}

func (cs *CalculatorSuite) iPress(num int) error {
	cs.calc.Press(num)
	return nil
}

func (cs *CalculatorSuite) iAdd(num int) error {
	cs.calc.Add(num)
	return nil
}

func (cs *CalculatorSuite) iSubtract(num int) error {
	cs.calc.Sub(num)
	return nil
}

// Assertions
func (cs *CalculatorSuite) theResultShouldBe(expectedResult int) error {
	result := cs.calc.GetResult()
	if result == expectedResult {
		return nil
	}
	return fmt.Errorf("%d doesn't match expectation: %d", result, expectedResult)
}

func (cs *CalculatorSuite) iMultiply(factor int) error {
	return cs.calc.Multiply(factor)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	calInstance := CalculatorSuite{
		calc:            new(ca.Calculator),
		ScenarioContext: ctx,
	}
	ctx.Step(`^calculator is cleared$`, calInstance.calculatorIsCleared)
	ctx.Step(`^i add (\d+)$`, calInstance.iAdd)
	ctx.Step(`^i multiply by (\d+)$`, calInstance.iMultiply)
	ctx.Step(`^i press (\d+)$`, calInstance.iPress)
	ctx.Step(`^i subtract (\d+)$`, calInstance.iSubtract)
	ctx.Step(`^the result should be (\d+)$`, calInstance.theResultShouldBe)
}
