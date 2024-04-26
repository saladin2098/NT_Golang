package main

import (
	// "encoding"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Student struct {
	Id      int64    `json:"id"`
	Name    string   `json:"name"`
	Course  []Course `json:"course"`
	Address Address  `json:"address"`
}
type Course struct {
	CourseId  string `json:"course"`
	Grade     string `json:"grade"`
	Professor string `json:"professor"`
}
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func WriteStudentsToFile(student []Student, filename string) error {
	file, error := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if error != nil {
		log.Printf("Could not open file")
		return error
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(student)
	return nil
}
func ReadStudentsFromFile(filename string) ([]Student, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Could not open file")
		return nil, err
	}
	defer file.Close()
	students := []Student{}
	decoder := json.NewDecoder(file)
	decoder.Decode(students)
	return students, nil
	
}
func PrintStudentDetails(students []Student) {
	for _, student := range students {
		fmt.Println(student)
	}
}
func FindTopScoringStudent(filename string) (*Student, error) {
	students, err := ReadStudentsFromFile(filename)
	if err != nil {

		return nil, err
	}
	MaksAverage := 0
	baholar := map[*Student]int{}
	for _, student := range students {
		totalScore := 0
		for _, score := range student.Course {
			grade, _ := strconv.Atoi(score.Grade)
			totalScore += grade
		}
		baholar[&student] = totalScore
		if MaksAverage < totalScore {
			MaksAverage = totalScore
		}
	}
	for k, v := range baholar {
		if v == MaksAverage {
			return k, nil
		}
	}
	return nil,nil
}
func GroupStudentsByCategory(students []Student) (map[string][]Student, error) {
	mapStudent := make(map[string][]Student)
	if len(students) == 0 {
		return mapStudent,errors.New("NO students were found")
	}
	for _,student := range students {
		if _,ok := mapStudent[student.Course[0].CourseId] ; !ok { 
			mapStudent[student.Course[0].CourseId] = []Student{}
			mapStudent[student.Course[0].CourseId] = append(mapStudent[student.Course[0].CourseId], student)
		} else {
			mapStudent[student.Course[0].CourseId] = append(mapStudent[student.Course[0].CourseId], student)
		}

	}
	return mapStudent,nil
	
}

func main() {

	// courseOfStudent1 := Course{
	// 	CourseId:  "1",
	// 	Grade:     "5",
	// 	Professor: "Palonchi",
	// }
	// courseOfStudent2 := Course{
	// 	CourseId:  "2",
	// 	Grade:     "4",
	// 	Professor: "Pistonchi",
	// }
	// courses1 := []Course{courseOfStudent1}
	// courses2 := []Course{courseOfStudent2}

	// addres1 := Address{Street: "Uzar", City: "Tashlkent", Country: "Uzbekistan"}
	// student1 := Student{Id: 1, Name: "Shamsiddin", Course: courses1, Address: addres1}

	// addres2 := Address{Street: "Uzar", City: "Tashlkent", Country: "Uzbekistan"}
	// student2 := Student{Id: 2, Name: "Salohiddin", Course: courses2, Address: addres2}

	// students := []Student{student1, student2}

    // err := WriteStudentsToFile(students,"students.json")
	// if err != nil {
	// 	fmt.Println("xatolik",err)
	// }
	readStudents,err := ReadStudentsFromFile("students.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(readStudents)
}


// how to find the type of any input in golang 
