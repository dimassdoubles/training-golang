package main

import (
	"errors"
	"fmt"
	"git.solusiteknologi.co.id/golang/traingo/tax"
)

func isAmountValid(amount float64) bool {
	return amount > 0
}

func inputAmount() (float64, error) {
	var amount float64;
	fmt.Print("Masukan nilai yang ingin dihitung pajaknya: ")
	_, err := fmt.Scanf("%f", &amount)

	if (err != nil) {
		return 0, err
	} else if isAmountValid(amount) {
		return amount, nil
	}
	return 0, errors.New("nilai tidak valid")
}

func inputOption() (string, error) {
	fmt.Println("Kalkulator Perhitungan Pajak")
	fmt.Println("----------------------------")

	fmt.Println("1. PPN 10%")
	fmt.Println("2. PPN 11%")
	fmt.Println("3. PPH 21%")
	fmt.Print("Masukan pilihan: ")
	
	var jenisPajak int
	_, err := fmt.Scanf("%d", &jenisPajak)

	if err != nil {
		return "", err
	}

	if jenisPajak == 1 || jenisPajak == 2 {
		fmt.Println("")
		fmt.Println("1. Harga Include Tax")
		fmt.Println("2. Harga Exclude Tax")
		fmt.Print("Masukan pilihan: ")

		var includeExclude int
		_, err = fmt.Scanf("%d", &includeExclude)
		if err != nil {
			return "", err
		}

		if includeExclude == 1 || includeExclude == 2 {
			return fmt.Sprintf("%d%d", jenisPajak, includeExclude), nil;
		}
		return "", errors.New("pilihan tidak valid")
			
	} else if jenisPajak == 3 {
		return fmt.Sprintf("%d", jenisPajak), nil;
	}
	
	return "", errors.New("pilihan tidak valid")
}

func main() {
	taxController := tax.Controller{} 

	for {
		option, err := inputOption()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("")
			amount, err := inputAmount()
			if err != nil {
				fmt.Println(err)
			} else {
				switch option {
					case "11": taxController.GetTax(tax.TypePpn10IncludeTax, amount)
					case "12": taxController.GetTax(tax.TypePpn10, amount)
					case "21": taxController.GetTax(tax.TypePpn11IncludeTax, amount)
					case "22": taxController.GetTax(tax.TypePpn11, amount)
					case "3" : taxController.GetTax(tax.TypePph21, amount)
					default  : fmt.Println("Pilihan tidak valid")
				}
				fmt.Println("")
			}
		}
		
		
	}
}