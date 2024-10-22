package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/patricioyegros/storichallenge/app"
)

const csvFileNameParam = "txns2.csv"

func main() {

	//var correctEmail bool = false
	var err error

	fmt.Println(("Welcome to Transaction Summary App"))
	/*fmt.Println(("Please insert the email where the summary will be sent: "))

	var destinationEmail string
	fmt.Scanln(&destinationEmail)

	for !correctEmail {
		if !strings.Contains(destinationEmail, "@") {
			fmt.Println("Invalid email, please insert a new email: ")
			fmt.Scanln(&destinationEmail)
		} else {
			correctEmail = true
		}
	}*/

	db, err := app.NewDBConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}

	script, err := os.ReadFile("app/resources/db.sql")
	if err != nil {
		log.Fatalln(err.Error())
	}
	query := string(script)
	split := strings.Split(query, ";")
	for _, line := range split {
		if line != "" {
			_, err = db.Exec(line)
			if err != nil {
				log.Fatalln(err.Error())
			}
		}
	}

	processService := app.NewService(db)

	err = processService.Process(csvFileNameParam, db)
	if err != nil {
		log.Fatalln(err.Error())
	}

	db.Close()

	log.Println("Transactions processed successfully")
}
