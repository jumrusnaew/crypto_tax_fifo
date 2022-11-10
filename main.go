package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RecordOfTaxCrypto struct {
	Activity string
	Coin     string
	Price    float64
	Amount   float64
}
type recordOfSale struct {
	line   int
	amount float64
	coin   string
}

func main() {
	record, err := os.Open("record.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer record.Close()
	scanner := bufio.NewScanner(record)
	lines := []RecordOfTaxCrypto{}

	for scanner.Scan() {
		line := scanner.Text()
		field := strings.Split(line, " ")
		price, _ := strconv.ParseFloat(field[2], 32)
		am, _ := strconv.ParseFloat(field[3], 32)
		lines = append(lines, RecordOfTaxCrypto{
			Activity: field[0],
			Coin:     field[1],
			Price:    float64(price),
			Amount:   float64(am),
		})
	}
	fmt.Printf("transection : %v \n", len(lines))
	fmt.Println("---------------------------------------------------------")
	rowIndex := 0
	SalPrice := 0.0
	Profit := 0.0
	histPrice := []recordOfSale{}
	for _, row := range lines {
		//fmt.Println(row)

		if row.Activity == "S" {
			aliasAmount := row.Amount
			priceSale := 0.0
			priceBuy := 0.0
			for o := 0; o < len(histPrice); o += 1 {
				if aliasAmount > 0 && histPrice[o].amount > 0 {
					if histPrice[o].coin == row.Coin {
						if histPrice[o].amount >= aliasAmount {

							fmt.Println("-------------------")
							fmt.Println(histPrice[o])
							fmt.Println("-------------------")
							fmt.Println("   histamount = ", histPrice[o].amount)
							fmt.Println("   alias = ", aliasAmount)
							histPrice[o].amount -= aliasAmount
							priceSale += aliasAmount * row.Price
							priceBuy += aliasAmount * lines[histPrice[o].line].Price
							fmt.Printf("   %.2f X %.2f = %.2f\n", aliasAmount, lines[histPrice[o].line].Price, (aliasAmount * lines[histPrice[o].line].Price))
							aliasAmount = 0
						} else {

							fmt.Println("-------------------")
							fmt.Println(histPrice[o])
							fmt.Println("-------------------")
							fmt.Println("   histamount = ", histPrice[o].amount)
							fmt.Println("   alias = ", aliasAmount)
							aliasAmount -= histPrice[o].amount
							priceSale += histPrice[o].amount * row.Price
							priceBuy += histPrice[o].amount * lines[histPrice[o].line].Price
							fmt.Printf("   %.2f X %.2f = %.2f\n", histPrice[o].amount, lines[histPrice[o].line].Price, (histPrice[o].amount * lines[histPrice[o].line].Price))
							histPrice[o].amount = 0
						}
					}
				}
			}

			Profit += (float64(priceSale) - priceBuy)
			SalPrice += float64(priceSale)
			fmt.Print("\nSale Amount : ", row.Amount-aliasAmount)
			fmt.Println("\nBalance :", histPrice)
			fmt.Printf("%v Sale Price : %.2f Buy Price : %.2f Profit : %.2f \n", row.Coin, priceSale, priceBuy, Profit)
			fmt.Println("---------------------------------------------------------")
		} else {
			histPrice = append(histPrice, recordOfSale{
				line:   rowIndex,
				amount: row.Amount,
				coin:   row.Coin,
			})
		}
		rowIndex += 1
	}
	fmt.Println(histPrice)
	fmt.Println("---------------------------------------------------------")
	fmt.Println("Sale Total", fmt.Sprintf("%.2f", SalPrice))
	fmt.Println("Profit Total", fmt.Sprintf("%.2f", Profit))
	fmt.Println("---------------------------------------------------------")
	fmt.Println("---------------------------------------------------------")
}
