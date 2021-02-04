package main 

import (

	"fmt"
)

type Book struct {

	name string
	author string 
	pages int
	
}

func (book Book) print_details() {
	fmt.Printf("Book %s was written by %s.", book.name, book.author) 
	fmt.Printf("\nIt contains %d pages.\n", book.pages) 
	
}

func main() {

	book1 := Book{"Life Of Truth", "Maheswar Reddy", 1729}
	book1.print_details()
	book1.name = "Kloudone"
	book1.pages = 2122



book1.print_details()
}