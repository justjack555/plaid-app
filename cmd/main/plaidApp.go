package main

import (
	"github.com/justjack555/plaid-app/pkg/apply"
	"log"
)

func main(){
	log.Println("Running main...")
	app := apply.Create()

	app.Apply()
}