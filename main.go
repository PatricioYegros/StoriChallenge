package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/patricioyegros/storichallenge/app"
)

const csvFileNameParam = "transactions.csv"

func main() {

	var correctEmail bool = false

	fmt.Println(("Welcome to Transaction Summary App"))
	fmt.Println(("Please insert the email where the summary will be sent: "))

	var destinationEmail string
	fmt.Scanln(&destinationEmail)

	for !correctEmail {
		if !strings.Contains(destinationEmail, "@") {
			fmt.Println("Invalid email, please insert a new email: ")
			fmt.Scanln(&destinationEmail)
		} else {
			correctEmail = true
		}
	}

	processService, err := app.NewService()
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = processService.Process(csvFileNameParam, destinationEmail)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Transactions processed successfully")
}
