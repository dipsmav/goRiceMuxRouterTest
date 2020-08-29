package main

import (
	"fmt"
	"time"
)

func main() {
	utc := time.Now().UTC()
	fmt.Println(utc)
	local := utc
	location, err := time.LoadLocation("Asia/Kolkata")
	if err == nil {
		local = local.In(location)
	}
	fmt.Println("UTC", utc.Format("15:04"), local.Location(), local.Format("15:04"))
	local = utc
	location, err = time.LoadLocation("America/Los_Angeles")
	if err == nil {
		local = local.In(location)
	}
	fmt.Println("UTC", utc.Format("15:04"), local.Location(), local.Format("15:04"))
}
