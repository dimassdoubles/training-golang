package main

// import (
// 	"fmt"
// )

type TaxCalculator interface {
	Calculate(amount float64) float64 
}

func getPrecentage(amount , precentage float64) float64 {
	return amount * precentage
}

func getPrecentageInclude(amount, precentage float64) float64 {
	return precentage * (amount / (1+precentage))
}


type TaxCalculatorPpn10 struct {}
	func (calc TaxCalculatorPpn10) Calculate(amount float64) float64 {
		return getPrecentage(amount, 0.10)
	}

type TaxCalculatorPpn11 struct {}
	func (calc TaxCalculatorPpn11) Calculate(amount float64) float64 {
		return getPrecentage(amount, 0.11)
	}

type TaxCalculatorPpn10IncludeTax struct {}
	func (calc TaxCalculatorPpn10IncludeTax) Calculate(amount float64) float64 {
		return getPrecentageInclude(amount, 0.10)
	}

type TaxCalculatorPpn11IncludeTax struct {}
	func (calc TaxCalculatorPpn11IncludeTax) Calculate(amount float64) float64 {
		return getPrecentageInclude(amount, 0.11)
	}

type TaxCalculatorPph21 struct {}
	func (calc TaxCalculatorPph21) Calculate(amount float64) float64 {
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

type TaxCalculatorFactory struct {}
	func (factory TaxCalculatorFactory) getCalculator(typeName string) TaxCalculator {
		switch typeName {
			case "ppn10": return TaxCalculatorPpn10{}
			case "ppn11": return TaxCalculatorPpn11{}
			case "ppn10IncludeTax": return TaxCalculatorPpn10IncludeTax{}
			case "ppn11IncludeTax": return TaxCalculatorPpn11IncludeTax{}
			case "pph21": return TaxCalculatorPph21{}
			default: return TaxCalculatorPpn10{}
		}
	}


// func main() {
// 	calcFactory := TaxCalculatorFactory{}
// 	fmt.Println(calcFactory.getCalculator("ppn10").Calculate(100))
// 	fmt.Println(calcFactory.getCalculator("ppn11").Calculate(100))
// 	fmt.Println(calcFactory.getCalculator("ppn10IncludeTax").Calculate(110))
// 	fmt.Println(calcFactory.getCalculator("ppn11IncludeTax").Calculate(111))
// 	fmt.Println(calcFactory.getCalculator("pph21").Calculate(20000000))
// }