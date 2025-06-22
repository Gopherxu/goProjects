package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	var date string
	var date1 string
	fmt.Println("Take 2 Date inputs : ")
	fmt.Println("Enter Dates 1st (YYYY MM DD_hh:mm:ss) :")

	fmt.Scanf("%v", &date)
	fmt.Println(date)
	startdate, err := time.Parse("2006-01-02_15:04:05", date)
	if err != nil {
		log.Fatal("An err occured while converting to date", err)
	}

	fmt.Println("Enter Dates 2st (YYYY MM DD_hh:mm:ss):")
	fmt.Scanf("%v", &date1)
	enddate, err := time.Parse("2006-01-02_15:04:05", date1)
	if err != nil {
		log.Fatal("An err occured while converting to date ", err)
	}

	// //now to find difference betn two dates
	// difference := enddate.Sub(startdate)

	// fmt.Println(difference)

	// year := int((difference.Hours() / 24) / 365)
	// difference -= time.Duration(year) * 365 * 24 * time.Hour
	// month := int((difference.Hours() / 24) / 30)
	//month := difference / (30 * 24 * time.Hour)
	// //difference -= month * 30 * 24 * time.Hour
	// days := difference / (24 * time.Hour)
	// difference -= days * 24 * time.Hour
	// hours := difference / time.Hour
	// difference -= hours * time.Hour
	// minutes := difference / time.Minute

	//fmt.Printf("years : %d and months : %d ", year, month)
	// year := startdate.Year()
	// fmt.Println(year)
	// if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
	// 	fmt.Println("Yup you got lleapp year")
	// } else {
	// 	fmt.Println("Nope try again")
	// }
	if startdate.Before(enddate) {
		fmt.Println("The startdate is before end date")

	} else {
		fmt.Println("Start date is after")
	}

}
