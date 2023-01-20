package tax

import (
	"errors"
	"fmt"
)

const labelJenisPajak string  = "Jenis Pajak  :"
const labelNilaiAmount string = "Nilai Amount :"
const labelPajak string       = "Pajak        :"

func printTitle() {
	fmt.Println("Hasil Perhitungan Pajak")
	fmt.Println("-----------------------")
}

func printBody(amount, result float64, taxType string) {
	fmt.Println(labelJenisPajak, taxType)
	fmt.Printf("%v %.2f\n", labelNilaiAmount, amount)
	fmt.Printf("%v %.2f\n", labelPajak, result)
}

type View interface {
	PrintResult(amount, result float64)
}

type Ppn10View struct {}
	func (view Ppn10View) PrintResult(amount, result float64) {
		printTitle()
		printBody(amount, result, "PPN 10% / Harga Exclude Tax")
	} 

type Ppn11View struct {}
	func (view Ppn11View) PrintResult(amount, result float64) {
		printTitle()
		printBody(amount, result, "PPN 11% / Harga Exclude Tax")
	}

type Ppn10IncludeTaxView struct {}
	func (view Ppn10IncludeTaxView) PrintResult(amount, result float64) {
		printTitle()
		printBody(amount, result, "PPN 10% / Harga Include Tax")
	}

type Ppn11IncludeTaxView struct {}
	func (view Ppn11IncludeTaxView) PrintResult(amount, result float64) {
		printTitle()
		printBody(amount, result, "PPN 11% / Harga Include Tax")
	}

type Pph21View struct {}
	func (view Pph21View) PrintResult(amount, result float64) {
		printTitle()
		printBody(amount, result, "PPH 21")
	}

type ViewFactory struct {}
	func (factory ViewFactory) GetView(calculator Calculator) (View, error) {
		switch calculator.(type) {
			case CalculatorPpn10: return Ppn10View{}, nil
			case CalculatorPpn11: return Ppn11View{}, nil
			case CalculatorPpn10IncludeTax: return Ppn10IncludeTaxView{}, nil
			case CalculatorPpn11IncludeTax: return Ppn11IncludeTaxView{}, nil
			case CalculatorPph21: return Pph21View{}, nil
			default: return nil, errors.New("kalkulator view tidak ditemukan")
		}
	}
