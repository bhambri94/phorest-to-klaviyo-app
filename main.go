package main

import (
	"fmt"
	"time"

	"github.com/bhambri94/phorest-to-klaviyo-app/configs"
	"github.com/bhambri94/phorest-to-klaviyo-app/phorest"
)

func main() {
	configs.SetConfig()
	var fromDateTime string
	var toDateTime string
	currentTime := time.Now()
	HoursCount := configs.Configurations.OldDateInHours
	fromDateTime = currentTime.Add(time.Duration(-HoursCount) * time.Hour).Format("2006-01-02")
	toDateTime = currentTime.Format("2006-01-02")
	fmt.Println("Fetching results from Date: "+fromDateTime, " to "+toDateTime)

	BranchList := phorest.GetBranches()
	phorest.TrackAppointmentDetails(BranchList, fromDateTime, toDateTime)
}
