package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func main() {

	seconds := time.Now().Unix()
	fmt.Println("Seconds obtained ", seconds)

	//generate the random number

	rand.Seed(seconds) //seed the random number generator

	target := rand.Intn(100) + 1 // generates the random numner

	fmt.Println("Target", target)

	f := 0.1
	t := 0.3

	fmt.Printf("Value %0.2f\n", f/t) //upto 2 decimals

	//Sprintf gives the formatted String

	resultString := fmt.Sprintf("The another value is %0.2f\n", f/t)
	fmt.Printf(resultString)

	myInt := 10
	myVal(&myInt)
	pass := &myInt
	*pass = 100
	myVal(&myInt)
	fmt.Println("The value is ", myInt)

	naturalNos := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum := 0

	for _, value := range naturalNos {
		sum += value
	}
	fmt.Println(sum)

	//ceation of slice
	text := make([]int, 0, 10)
	text = append(text, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	slice := text[1:]
	slice = append(slice, 14)
	fmt.Println(slice)

	mymap := make(map[int]int)

	mymap[1] = 1
	mymap[2] = 2

	fmt.Println(mymap[1])

	gradeStatus("")

	gradesOrdered()

	date := Date{}
	date.SetDate(2000)
	fmt.Println("Date - of - year set is", date.Year)

}

func myVal(myInt *int) {
	*myInt *= 2
}

func gradeStatus(name string) {
	grades := map[string]int{"Anna": 70, "Carl": 10}
	var ok bool
	grade, ok := grades[name]
	fmt.Println("Grade is", grade)
	if !ok {
		fmt.Println("No grades defined for the name", name)
	} else if grade < 60 {
		fmt.Println("grades are low", grade)
	} else {
		fmt.Println("Passed")
	}

}

func gradesOrdered() {
	gradesMap := map[string]int{"Anna": 100, "Vijay": 200, "Carl": 90}

	// This prints unordered map
	for key, value := range gradesMap {
		fmt.Println("map", key, value)
	}

	//ordered map

	names := []string{}

	for name := range gradesMap {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Println("Sorted names", name, gradesMap[name])
	}

}

func (d *Date) SetDate(year int) {
	d.Year = year
}
