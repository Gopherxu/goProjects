package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Function to get user input for two dates and return them as time.Time objects
func getUserDates() (time.Time, time.Time) {
	var str1, str2 string
	fmt.Println(`Enter 1st Date in Format (YYYY-MM-DD) : `)
	fmt.Scanf("%v", &str1)
	fmt.Println("Enter 2nd Date in Format (YYYY-MM-DD) : ")
	fmt.Scanf("%v", &str2)
	startDate, _ := time.Parse("2006-01-02", str1)
	endDate, _ := time.Parse("2006-01-02", str2)
	return startDate, endDate
}

// Function to get user input for a single date and return it as a time.Time object
func getUserDatesInput() time.Time {
	var str1 string
	fmt.Println(`Enter 1st Date in Format (YYYY-MM-DD) : `)
	fmt.Scanf("%v", &str1)
	startDate, err1 := time.Parse("2006-01-02", str1)
	if err1 != nil {
		log.Fatal("Error parsing date input:", err1)
	}
	return startDate
}

// Function to get user input for two dates with time and return them as time.Time objects
func getUserDatesInputDiff() (time.Time, time.Time) {
	var str1, str2 string
	fmt.Println(`Enter 1st Date in Format (YYYY-MM-DD_hh:mm:ss) : `)
	fmt.Scanf("%v", &str1)
	fmt.Println("Enter 2nd Date in Format (YYYY-MM-DD_hh:mm:ss) : ")
	fmt.Scanf("%v", &str2)
	startDate, err := time.Parse("2006-01-02_15:04:05", str1)
	if err != nil {
		log.Fatal("1st Date is not in the given Format: ", err)
	}
	endDate, err := time.Parse("2006-01-02_15:04:05", str2)
	if err != nil {
		log.Fatal("2nd Date is not in the given Format: ", err)
	}
	return startDate, endDate
}

// Function to print the current date/time formats in UTC
func UTC(currentTime time.Time) {
	fmt.Println(currentTime.Format("02 Jan 2006"))
	fmt.Println(currentTime.Format("Jan 02, 2006"))
	fmt.Println(currentTime.Format("2006-01-02"))
	fmt.Println(currentTime.Format(time.RFC3339))
	fmt.Println(currentTime.Format("Monday, 02 January 2006"))
}

// Function to print the current date/time formats in IST (Indian Standard Time)
func IST(currentTimeIST time.Time) {
	fmt.Println("In IST:")
	fmt.Println(currentTimeIST.Format("02 Jan 2006"))
	fmt.Println(currentTimeIST.Format("Jan 02, 2006"))
	fmt.Println(currentTimeIST.Format("2006-01-02"))
	fmt.Println(currentTimeIST.Format(time.RFC3339))
	fmt.Println(currentTimeIST.Format("Monday, 02 January 2006"))
}

// Function to get user input for a timezone and return the current time in that timezone
func User_timeZone(timeZone string) (string, time.Time) {
	timeZone = strings.Title(timeZone)
	zonePath := "/usr/share/zoneinfo/" + timeZone
	_, err := os.Stat(zonePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Time zone not found:", timeZone)
		}
	}
	loc, _ := time.LoadLocation(timeZone)
	currentTime := time.Now().In(loc)

	return timeZone, currentTime

}

// Function to calculate the difference between two dates
func dateDiff(startDate time.Time, endDate time.Time) (time.Duration, time.Duration, time.Duration, time.Duration, time.Duration) {

	Difference := endDate.Sub(startDate)
	fmt.Println(Difference)
	years := Difference / (365 * 24 * time.Hour)
	Difference -= years * 365 * 24 * time.Hour
	months := Difference / (30 * 24 * time.Hour)
	Difference -= months * 30 * 24 * time.Hour
	days := Difference / (24 * time.Hour)
	Difference -= days * 24 * time.Hour
	hours := Difference / time.Hour
	Difference -= hours * time.Hour
	minutes := Difference / time.Minute
	return years, months, days, hours, minutes
}

// Function to check if a given year is a leap year
func isLeapYear(year int) (boolen bool) {

	boolen = year%4 == 0 && (year%100 != 0 || year%400 == 0)
	return boolen
}

// Function to check if a given date falls in a leap year
func leapYear(date time.Time) (x string) {
	year := date.Year()
	if isLeapYear(year) {
		x = "This is a Leap Year"
	} else {
		x = "Not a Leap Year"
	}
	return
}

// Function to compare two dates and determine if they are equal, earlier or later
func equalDate(d1, d2 time.Time) (x string) {
	switch {
	case d1.Before(d2):
		x = "Date 1 is earlier than Date 2."
	case d1.After(d2):
		x = "Date 1 is later than Date 2."
	default:
		x = "Date 1 is equal to Date 2."
	}
	return
}

// Function to count the number of weekend days between two dates
func WeekendDaysCount(startDate time.Time, endDate time.Time) int {
	count := 0
	for date := startDate; date.Before(endDate); date = date.AddDate(0, 0, 1) {
		if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
			count++
		}
	}
	return count
}

// Function to print the count of weekend days between two dates
func weekEnd(startDate time.Time, endDate time.Time) {

	weekendDays := WeekendDaysCount(startDate, endDate)
	fmt.Println("Weekend days are : ", weekendDays)
}

// Function to get user input for a list of holidays
func getHolidayList() []time.Time {
	var str string
	var arr []time.Time
	fmt.Println(`Enter Dates Having Holiday in Format YYYY-MM-DD (Enter "Done" once completed) : `)
	for {
		fmt.Scanf("%v", &str)
		if strings.ToLower(str) == "done" {
			break
		}
		date, _ := time.Parse("2006-01-02", str)
		arr = append(arr, date)
	}
	return arr
}

// Function to check if a given date is a holiday
func isHoliday(holidays []time.Time, date time.Time) (x bool) {
	for _, holiday := range holidays {
		if holiday.Equal(date) {
			x = true
		}
	}
	return
}

// Function to check if a given date is a working day
func calculateWorkingDay(holidays []time.Time, date time.Time) (x bool) {

	if date.Weekday() != time.Saturday && date.Weekday() != time.Sunday && !isHoliday(holidays, date) {
		x = true
	}
	return
}

// Function to calculate the number of working days between two dates
func calculateWorkingDays(holidays []time.Time, startDate, endDate time.Time) int {
	count := 0
	for date := startDate; date.Before(endDate); date = date.AddDate(0, 0, 1) {
		if date.Weekday() != time.Saturday && date.Weekday() != time.Sunday && !isHoliday(holidays, date) {
			count++
		}
	}
	return count
}

// Function to calculate a date relative to an input date based on a specified number of business days
func calculateBusinessDate(holidays []time.Time, inputDate time.Time, businessDays int) time.Time {
	date := inputDate
	for i := 0; i < abs(businessDays); {
		if businessDays > 0 {
			date = date.AddDate(0, 0, 1)
		} else {
			date = date.AddDate(0, 0, -1)
		}
		if date.Weekday() != time.Saturday && date.Weekday() != time.Sunday && !isHoliday(holidays, date) {
			i++
		}
	}
	return date
}

// Function to return the absolute value of a number
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var num int
	fmt.Println(`Which Operation do you Want to Perform :
	1) Print current date/time formats in UTC (Q14) .
	2) Print current date/time formats in IST (Q15) .
	3) To accept a valid timezone as input and print time as per the time zone selected by an end user. (Q16) .
	4) To print a difference between Two Dates Q(17):
	5) To Find date falls in a leap year or not (Q18):
	6) To print an output whether date1 and date2 are equal or not (Q19) :
	7) To Print the count of weekend days (Q20) :
	8) A program to manage a list of holidays and determine if a given date is a working day or not (Q21) .
	9) To print the number of working days between two dates, considering both dates in the calculation (Q22) .
	10) To calculate the date relative to an input date based on a specified number of business days (Q23).

	Enter The Operation number : `)
	fmt.Scanf("%v", &num)
	switch num {
	case 1:
		currentTime := time.Now().UTC()
		UTC(currentTime)

	case 2:
		currentTime := time.Now().UTC()
		istLocation, _ := time.LoadLocation("Asia/Kolkata")
		currentTimeIST := currentTime.In(istLocation)
		IST(currentTimeIST)

	case 3:
		var timeZone string
		fmt.Println("Enter The TimeZone : ")
		fmt.Scanf("%v", &timeZone)
		timezone, currentTime := User_timeZone(timeZone)
		fmt.Printf("For timezone %v Current Time is %v", timezone, currentTime)

	case 4:
		str1, str2 := getUserDatesInputDiff()
		years, months, days, hours, minutes := dateDiff(str1, str2)
		fmt.Printf("Difference: %d years, %d months, %d days, %d hours, %d minutes", years, months, days, hours, minutes)

	case 5:
		str := getUserDatesInput()
		value := leapYear(str)
		fmt.Println(value)

	case 6:
		str1, str2 := getUserDates()
		value := equalDate(str1, str2)
		fmt.Println(value)
	case 7:
		str1, str2 := getUserDates()
		weekEnd(str1, str2)

	case 8:

		holidays := getHolidayList()
		date := getUserDatesInput()

		fmt.Println("Is a working day: ", calculateWorkingDay(holidays, date))

	case 9:
		startDate, endDate := getUserDates()
		holidays := getHolidayList()
		workingDays := calculateWorkingDays(holidays, startDate, endDate)
		fmt.Println("Number of working days:", workingDays)

	case 10:

		inputDate := getUserDatesInput()
		holidays := getHolidayList()
		fmt.Println("Enter number of business days:")
		var businessDays int
		fmt.Scanln(&businessDays)
		resultDate := calculateBusinessDate(holidays, inputDate, businessDays)
		fmt.Println("Result date:", resultDate.Format("2006-01-02"))

	}
}
