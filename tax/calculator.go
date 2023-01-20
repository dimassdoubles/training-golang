package tax

import (
	"errors"
)

type Calculator interface {
	Calculate(amount float64) float64
}

func getPrecentage(amount , precentage float64) float64 {
	return amount * precentage
}

func getPrecentageInclude(amount, precentage float64) float64 {
	return precentage * (amount / (1+precentage))
}

type CalculatorPpn10 struct {}
	func (calc CalculatorPpn10) Calculate(amount float64) float64 {
		return getPrecentage(amount, 0.10)
	}

type CalculatorPpn11 struct {}
	func (calc CalculatorPpn11) Calculate(amount float64) float64 {
		return getPrecentage(amount, 0.11)
	}

type CalculatorPpn10IncludeTax struct {}
	func (calc CalculatorPpn10IncludeTax) Calculate(amount float64) float64 {
		return getPrecentageInclude(amount, 0.10)
	}

type CalculatorPpn11IncludeTax struct {}
	func (calc CalculatorPpn11IncludeTax) Calculate(amount float64) float64 {
		return getPrecentageInclude(amount, 0.11)
	}

type CalculatorPph21 struct {}
	func (calc CalculatorPph21) Calculate(amount float64) float64 {
		juta := 1000000.0
		if amount < 40*juta {
			return 0.0
		} else if amount < 50*juta {
			return getPrecentage(amount, 0.05)
		} else if amount < 250*juta {
			return getPrecentage(amount, 0.15)
		} else if amount < 500*juta {
			return getPrecentage(amount, 0.25)
		} else if amount >= 500*juta {
			return getPrecentage(amount, 0.30)
		} else {
			return -1
		}
	}

const TypePpn10 string = "ppn10"
const TypePpn11 string = "ppn11"
const TypePpn10IncludeTax string = "ppn10IncludeTax"
const TypePpn11IncludeTax string = "ppn11IncludeTax"
const TypePph21 string = "pph21"

type CalculatorFactory struct {}
	func (factory CalculatorFactory) GetCalculator(typeName string) (Calculator, error) {
		switch typeName {
			case TypePpn10: return CalculatorPpn10{}, nil
			case TypePpn11: return CalculatorPpn11{}, nil
			case TypePpn10IncludeTax: return CalculatorPpn10IncludeTax{}, nil
			case TypePpn11IncludeTax: return CalculatorPpn11IncludeTax{}, nil
			case TypePph21: return CalculatorPph21{}, nil
			default: return nil, errors.New("kalkulator tidak ditemukan")
		}
	}
