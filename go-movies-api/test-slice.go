package main

import "fmt"

type Movie struct {
	ID       string
	title    string
	Isbn     string
	Director *Director
}

type Director struct {
	firstName string
	lastName  string
}

//define slice of the movies

var movies []Movie

func main() {

	movies = append(movies, Movie{ID: "1", title: "Gotch", Isbn: "AA", Director: &Director{firstName: "Raj", lastName: "Tevar"}})
	movies = append(movies, Movie{ID: "2", title: "Game of thrones", Isbn: "AB", Director: &Director{firstName: "Digaaz", lastName: "Tevar"}})
	movies = append(movies, Movie{ID: "3", title: "Man vs wild", Isbn: "ABC", Director: &Director{firstName: "Vipul", lastName: "Tevar"}})
	movies = append(movies, Movie{ID: "4", title: "Man vs wild", Isbn: "ABC", Director: &Director{firstName: "Vipul", lastName: "Tevar"}})

	for index, item := range movies {
		if item.ID == "2" {
			movies = append(movies[:index], movies...)
		}
	}

	var fields []int

	fields = append(fields, 1, 2, 3, 4, 5, 6)

	fmt.Println(fields)

	index := 2
	fmt.Println(fields[:index])   //1,2
	fmt.Println(fields[:index+1]) //1,2,3
	//fields = append(fields[:index], fields...) // 1,2,1,2,3,4,5,6
	//fmt.Println(fields)
	fields = append(fields[:index], fields[index+1:]...) // 1,2,4,5,6
	fmt.Println(fields)

}
