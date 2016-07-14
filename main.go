package main

import (
	"github.com/mannkind/dashbtn/cmd"
	"log"
)

func main() {
	if err := cmd.DashBtnCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
