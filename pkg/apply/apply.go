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
	log.Println("apply.Apply(): Raw JSON is: ", string(rawJson))

	resp, err := http.Post(_PLAID_ENDPOINT, "application/json", bytes.NewBuffer(rawJson))
	if err != nil {
		log.Fatalln("apply.Apply(): Unable to post request to Plaid")
	}

	log.Println("apply.APply(): Response status code is: ", resp.Status)
}

func scanFields(s []string) ScanMap {
	scanMap := make(ScanMap)
	scanner := bufio.NewScanner(os.Stdin)

//	log.Println("ScanNext(): About to scan text...")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s: ", s[i])

		if !scanner.Scan() {
			break
		}

		scanMap[s[i]] = scanner.Text()
	}
//	log.Println("ScanNext(): Done scanning...")

	if err := scanner.Err(); err != nil {
		log.Fatalln("Apply.scanNext(): Unable to scan due to: ", err)
	}

	return scanMap
}
func loadApplication() *Application {
	res := new(Application)
	fields := []string{
		"name", "email", "resume", "github",
	}

	scanMap := scanFields(fields)

/*	for k , v := range scanMap {
		log.Println(k, ": ", v)
	}
*/
	res.Name = scanMap["name"]
	res.Email = scanMap["email"]
	res.Resume = scanMap["resume"]
	res.Github = scanMap["github"]

//	log.Println("Resulting APplication object is: ", res)

	return res
}

func Create() *Application {
	log.Println("Apply.Create():")

	res := loadApplication()

	return res
}