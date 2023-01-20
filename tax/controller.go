package tax

import (
	"fmt"
)

type Controller struct {}

func (controller Controller) GetTax(taxType TaxType, amount float64) {
	calc, err := CalculatorFactory{}.GetCalculator(taxType)
	if err != nil {
		fmt.Println(err)
	}

	view, err := ViewFactory{}.GetView(calc)
	if err != nil {
		fmt.Println(err)
	} 

	result, err := calc.Calculate(amount)

	if err != nil {
		fmt.Println(err)
	} else {
		view.PrintResult(amount, result)
	}
}