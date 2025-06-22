// 1. Have gr8 than 100 records from the csv file.
// 2. Convert Height and Weight column as BMI ( 0 - normal, otherwise 1)
// 3. Convert no of surgery as 0 or 1 instead of count.
// 4. write converted output into csv file
// 5. Read the final out file and give different possible queries.
//      a) average premium for give age
//      b) health indicator sum of all indicators
//      c) query average premium by indicator 

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strconv"

	"os"
)

// Define a struct to represent each row of csv data
type Person struct {
	Age                     uint8
	Diabetes                bool
	BloodPressure           bool
	AnyTransplants          bool
	AnyChronicDisease       bool
	Height                  uint16
	Weight                  uint16
	KnownAllergies          bool
	HistoryOfCancerInFamily bool
	NumberOfMajorSurgeries  uint8
	PremiumPrice            uint64
}


// New Structure to Store Data in Required Format
type NewPerson struct {
	Age            uint8
	Diabetes       bool
	BloodPressure  bool
	Transplant     bool
	ChronicDisease bool
	BMI            bool
	Allergie       bool
	CancerInFamily bool
	Surgerie       bool
	Premium        uint64
}

// Function to read csv file and store its records
func readCsv(csv_file *os.File) [][]string {
	//Create csv reader
	reader := csv.NewReader((csv_file))

	// Read the header row performing this to skip headers
	_, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading header row:", err)
	}

	//Read all Records from csv File
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

// To create instrance of People Structure by inserting values of records
func fileData() []Person {

	//open Medicalpremium.csv file to Preform operations on it
	fileName := "Medicalpremium.csv"
	csv_file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// Store all Data of csv file in records
	records := readCsv(csv_file)

	//Create a varible of type Person to store records
	// As it is rawData data yet to be processed
	var rawData []Person

	//Store records in rawdata instance of  structure Person
	rawData = insertRecords(rawData, records)

	return rawData
}

// Function to insert records in Person Struct
func insertRecords(structure []Person, records [][]string) []Person {

	for _, record := range records {

		//Parse from string to Required data Type
		//Parse age and Check if there is any Error
		age, err := strconv.ParseUint(record[0], 10, 16)
		if err != nil {
			log.Fatal("The age values have some Error !!", err)
		}

		//Parse Diabetes and Check if there is any Error
		diabetes, err := strconv.ParseBool(record[1])
		if err != nil {
			log.Fatal("The Diabetes values have some Error !!", err)
		}

		//Parse the value and Check if there is any Error
		bloodPressure, err := strconv.ParseBool(record[2])
		if err != nil {
			log.Fatal("The Blood Pressure values have some Error !!", err)
		}

		//Parse the value and Check if there is any Error
		transplants, err := strconv.ParseBool(record[3])
		if err != nil {
			log.Fatal("The Transplant values have some Error !!", err)
		}

		//Parse the value and Check if there is any Error
		chronicDiseases, err := strconv.ParseBool(record[4])
		if err != nil {
			log.Fatal("The Chronic Diseases values have some Error !!", err)
		}

		//Parse the value and Check if there is any Error
		height, err := strconv.ParseUint(record[5], 10, 16)
		if err != nil {
			log.Fatal("The Height values have some Error !!", err)
		}

		//Parse the value and Check if there is any Error
		weight, err := strconv.ParseUint(record[6], 10, 16)
		if err != nil {
			log.Fatal("The Weight values have some Error !!", err)
		}

		//Parse the value and Check if there is any Error
		allergies, err := strconv.ParseBool(record[7])
		if err != nil {
			log.Fatal("The Allergies values have some Error !!", err)
		}

		//Parse the value and Check if there is any Error
		cancerHistory, err := strconv.ParseBool(record[8])
		if err != nil {
			log.Fatal("The Cancer History values have some Error !!", err)
		}

		//Parse the value and Check if there is any Error
		surgeries, err := strconv.ParseUint(record[9], 10, 8)
		if err != nil {
			log.Fatal("The Surgeries values have some Error !!", err)
		}

		//Parse the value and Check if there is any Error
		premiumPrice, err := strconv.ParseUint(record[10], 10, 64)
		if err != nil {
			log.Fatal("The Premium Price values have some Error !!", err)
		}

		// Append the Parsed and data directly to the people Slice
		structure = append(structure, Person{

			Age:                     uint8(age),
			Diabetes:                diabetes,
			BloodPressure:           bloodPressure,
			AnyTransplants:          transplants,
			AnyChronicDisease:       chronicDiseases,
			Height:                  uint16(height),
			Weight:                  uint16(weight),
			KnownAllergies:          allergies,
			HistoryOfCancerInFamily: cancerHistory,
			NumberOfMajorSurgeries:  uint8(surgeries),
			PremiumPrice:            premiumPrice,
		})
	}
	return structure

}

// TO check if BMI is Normal or not
func calculateBMI(height uint16, weight uint16) (normal bool) {
	// bmi as float varible
	var bmi float32
	var heightInmetersSquare float32

	// using formula to calculat bmi weight/metre^2
	heightInmetersSquare = (float32(height*height) / 10000)
	bmi = float32(weight) / heightInmetersSquare

	if bmi >= 18.5 && bmi <= 24.9 {
		normal = false
	} else {
		normal = true
	}

	return
}

// Function to process Data
func PreparePersonData(persons []Person) []NewPerson {

	var structuredData []NewPerson

	for _, person := range persons {

		//isNormalBMI stores if bmi is normal or not
		isNormalBMI := calculateBMI(person.Height, person.Weight)

		//Convert the number of surgeries to a boolean value
		hasSurgeries := person.NumberOfMajorSurgeries > 0

		structuredData = append(structuredData, NewPerson{

			Age:            person.Age,
			Diabetes:       person.Diabetes,
			BloodPressure:  person.BloodPressure,
			Transplant:     person.AnyTransplants,
			ChronicDisease: person.AnyChronicDisease,
			BMI:            isNormalBMI,
			Allergie:       person.KnownAllergies,
			CancerInFamily: person.HistoryOfCancerInFamily,
			Surgerie:       hasSurgeries,
			Premium:        person.PremiumPrice,
		})

	}

	return structuredData
}

// Functon to write processed Data into Structure
func writeCSV(filename string, structuredData []NewPerson) {

	//Create file to store structuredData
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	//csv writer to write into csv file
	writer := csv.NewWriter(file)
	defer writer.Flush()

	//Write csv header
	header := []string{"Age", "Diabetes", "BloodPressure", "Transplant", "BMI", "ChronicDisease", "Allergie", "CancerInFamily", "Surgerie", "Premium"}
	writer.Write(header)

	//Conver bool data to integer to Store in file
	for _, person := range structuredData {
		var diabetes int8 = 0
		if person.Diabetes {
			diabetes = 1
		}
		var bloodPressure int8 = 0
		if person.BloodPressure {
			bloodPressure = 1
		}
		var transplant int8 = 0
		if person.Transplant {
			transplant = 1
		}
		var BMI int8 = 0
		if person.BMI {
			BMI = 1
		}
		var chronicDisease int8 = 0
		if person.ChronicDisease {
			chronicDisease = 1
		}
		var allergie int8 = 0
		if person.Allergie {
			allergie = 1
		}
		var cancerInFamily int8 = 0
		if person.CancerInFamily {
			cancerInFamily = 1
		}
		var surgerie int8 = 0
		if person.Surgerie {
			surgerie = 1
		}

		record := []string{
			//write data to CSV file as it is in integer
			strconv.FormatUint(uint64(person.Age), 10),
			strconv.Itoa(int(diabetes)),
			strconv.Itoa(int(bloodPressure)),
			strconv.Itoa(int(transplant)),
			strconv.Itoa(int(BMI)),
			strconv.Itoa(int(chronicDisease)),
			strconv.Itoa(int(allergie)),
			strconv.Itoa(int(cancerInFamily)),
			strconv.Itoa(int(surgerie)),
			strconv.FormatUint(person.Premium, 10),
		}
		//this will Write each record into the file
		writer.Write(record)

	}

}

// Average Premium for number of active Health indicators
func averagePremiumIndicator(indicatorCount uint8, structuredData []NewPerson) float64 {

	var Premium uint = 0
	var count uint16 = 0
	var averagePremium float64 = 0

	for _, person := range structuredData {

		var icount uint8 = 0

		//This if blocks increase the count if the given indicator is active
		if person.Diabetes {
			icount++
		}
		if person.BloodPressure {
			icount++
		}
		if person.Transplant {
			icount++
		}
		if person.ChronicDisease {
			icount++
		}
		if person.BMI {
			icount++
		}
		if person.Allergie {
			icount++
		}
		if person.CancerInFamily {
			icount++

		}
		if person.Surgerie {
			icount++
		}

		//check if both counts are equal to calcukate Premium
		if icount == indicatorCount {

			Premium += uint(person.Premium)
			count++

		}

	}

	//Calculate Average Premium
	averagePremium = float64(Premium) / float64(count)
	fmt.Printf("Count of People Having Health Indicator %v is : %v\n", indicatorCount, count)
	return averagePremium

}

// Average Premium for specific age
func averagePremiumAge(structuredData []NewPerson, Age uint8) float64 {

	var Premium uint = 0
	var count uint16 = 0
	var averagePremium float64 = 0
	for _, person := range structuredData {

		//Check if the age is same as enterd age to Calculate average
		if person.Age == uint8(Age) {

			//Add it to Premium to get average at last
			Premium += uint(person.Premium)
			count++
		}
	}
	if count != 0 {

		//Calculate Average Premium for given age
		averagePremium = float64(Premium) / float64(count)
	} else {
		//if count is zero no age found
		averagePremium = 0
	}

	return averagePremium
}

// To calculate Average Premium for Cancer Patient
func averagePremiumCancer(structuredData []NewPerson) float64 {

	var Premium uint = 0
	var count uint16 = 0
	var averagePremium float64 = 0
	for _, person := range structuredData {

		//Check if Patient Having Cancer or not
		if person.CancerInFamily {
			Premium += uint(person.Premium)
			//Increase the count of people Having Cancer
			count++
		}
	}
	//calculate Average Premium
	averagePremium = float64(Premium) / float64(count)

	return averagePremium
}

// To get User input Age
func getAge() uint8 {

	var age uint8
	fmt.Println("Enter Age to Get its Average Premium : ")
	fmt.Scanf("%v", &age)
	return age
}

// To get User input Indicator Count
func getIndicatorCount() uint8 {

	var indicatorCount uint8
	fmt.Println("Enter Number of  Active Indicators : ")
	fmt.Scanf("%v", &indicatorCount)
	return indicatorCount
}

// func getIndicator() uint8 {
// 	var indicatorCount uint8
// 	fmt.Println("Enter Number of Indicators which should not be considered: ")
// 	fmt.Scanf("%v", &indicatorCount)
// 	return indicatorCount
// }

// This Function Handles input Arguments and Function calling?
func SelectPremiumCalculations(structuredData []NewPerson) {

	for {
		var key uint8
		fmt.Println(`Enter which Operation You want to Perform.
	1) Get Average Premium of Perticular Age : 
	2) Get Average Premium by Health Indicator ::
	3) Get Average Premium When there is Cancer history in Family : `)
		fmt.Scanf("%v", &key)

		switch key {
		case 1:
			//To get User input for which age Average premium is to be Calculated
			Age := getAge()
			averagePremium := averagePremiumAge(structuredData, Age)

			//if Averege Premium is there means the age is valid.
			if averagePremium > 0 {
				fmt.Printf("Average Premium for Age %v : %v \n", Age, averagePremium)
			} else {
				fmt.Println("Age is not found")
			}

		case 2:
			//To get User input for Heath Average premium is to be Calculated

			indicatorCount := getIndicatorCount()
			fmt.Println("Average Premium is  : ", averagePremiumIndicator(indicatorCount, structuredData))

		case 3:
			//Average Premium for Cancer history in family
			averagePremiumWhenCancer := averagePremiumCancer(structuredData)
			fmt.Println("Average Premium For Cancer Histoy in Family is : ", averagePremiumWhenCancer)

		default:
			fmt.Println("\nEnter Key mentioned in above options !!!")

		}

		fmt.Println(`If you want Contine Enter "1" else "0": `)
		var loopAgain byte
		fmt.Scanf("%v", &loopAgain)

		if loopAgain == 0 {
			break
		}
	}

}

func main() {

	//instance of Structure NewPerson to store Structured data
	var structuredData []NewPerson

	//This fileData function returns stored data
	rawData := fileData()

	//Store structured Data in StructuredData instance of class NewPerson

	structuredData = PreparePersonData(rawData)
	fmt.Println("Modifued csv File Created Successfully !!")

	//store Structured data in csv file
	writeCSV("structured_data.csv", structuredData)

	//This Function Selects which operation to Perform
	SelectPremiumCalculations(structuredData)

}
