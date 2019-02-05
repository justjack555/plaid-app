package apply

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

/**
	Application request structure
 */
type Application struct {
	Name string
	Email string
	Resume string
	Github string
}

type ScanMap map[string]string

const _PLAID_ENDPOINT = "https://contact.plaid.com/jobs"

/**
	Post application to Plaid API
 */
func (app *Application) Apply() {
	rawJson, err := json.Marshal(app)
	if err != nil {
		log.Fatalln("apply.Apply(): Unable to convert application to JSON...")
	}

	resp, err := http.Post(_PLAID_ENDPOINT, "application/json", bytes.NewBuffer(rawJson))
	if err != nil {
		log.Fatalln("apply.Apply(): Unable to post request to Plaid")
	}

	log.Println("apply.APply(): Response status code is: ", resp.Status)
}

/**
	Scan field values in from standard input
	by prompting the user with the field name
 */
func scanFields(s []string) ScanMap {
	scanMap := make(ScanMap)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%s: ", s[i])

		if !scanner.Scan() {
			break
		}

		scanMap[s[i]] = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Apply.scanNext(): Unable to scan due to: ", err)
	}

	return scanMap
}

/**
	Load the application request structure
	by initializing a map of field names
	to values and assigning the appropriate
	values to application structure members
 */
func loadApplication() *Application {
	res := new(Application)
	fields := []string{
		"name", "email", "resume", "github",
	}

	scanMap := scanFields(fields)

	res.Name = scanMap["name"]
	res.Email = scanMap["email"]
	res.Resume = scanMap["resume"]
	res.Github = scanMap["github"]

	return res
}

/**
	Create simply loads application structure
 */
func Create() *Application {
	res := loadApplication()

	return res
}