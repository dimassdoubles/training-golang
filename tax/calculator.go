package tax

import (
	"errors"
)

func isValidAmount(amount float64) bool {
	return amount > 0
}

type Calculator interface {
	Calculate(amount float64) (float64, error)
}

func getPrecentage(amount , precentage float64) float64 {
	return amount * precentage
}

func getPrecentageInclude(amount, precentage float64) float64 {
	return precentage * (amount / (1+precentage))
}

type CalculatorPpn10 struct {}
	func (calc CalculatorPpn10) Calculate(amount float64) (float64, error) {
		if !isValidAmount(amount) {
			return 0.0, errors.New("amount tidak valid")
		}
		return getPrecentage(amount, 0.10), nil
	}

type CalculatorPpn11 struct {}
	func (calc CalculatorPpn11) Calculate(amount float64) (float64, error) {
		if !isValidAmount(amount) {
			return 0.0, errors.New("amount tidak valid")
		}
		return getPrecentage(amount, 0.11), nil
	}

type CalculatorPpn10IncludeTax struct {}
	func (calc CalculatorPpn10IncludeTax) Calculate(amount float64) (float64, error) {
		if !isValidAmount(amount) {
			return 0.0, errors.New("amount tidak valid")
		}
		return getPrecentageInclude(amount, 0.10), nil
	}

type CalculatorPpn11IncludeTax struct {}
	func (calc CalculatorPpn11IncludeTax) Calculate(amount float64) (float64, error) {
		if !isValidAmount(amount) {
			return 0.0, errors.New("amount tidak valid")
		}
		return getPrecentageInclude(amount, 0.11), nil
	}

type CalculatorPph21 struct {}
	func (calc CalculatorPph21) Calculate(amount float64) (float64, error) {
		if !isValidAmount(amount) {
			return 0.0, errors.New("amount tidak valid")
		}
		
		juta := 1000000.0
		if amount < 40*juta {
			return 0.0, nil
		} else if amount < 50*juta {
			return getPrecentage(amount, 0.05), nil
		} else if amount < 250*juta {
			return getPrecentage(amount, 0.15), nil
		} else if amount < 500*juta {
			return getPrecentage(amount, 0.25), nil
		} else if amount >= 500*juta {
			return getPrecentage(amount, 0.30), nil
		} else {
			return -1, errors.New("error tidak diketahui di calculator pph 21 calculate")
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
