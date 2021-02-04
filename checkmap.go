// Go program to illustrate how to 
// check the key is available or not 

package main 

import "fmt"

// Main function 
func main() { 

	// Creating and initializing a map 
	m_a_p := map[int]string{ 
		90: "Dog", 
		91: "Cat", 
		92: "Cow", 
		93: "Bird", 
		94: "Rabbit", 
	} 

	fmt.Println("Original map: ", m_a_p) 

	// Checking the key is available 
	// or not in the m_a_p map 
	pet_name, ok := m_a_p[90] 
	fmt.Println("\nKey present or not:", ok) 
	fmt.Println("Value:", pet_name) 

	// Using blank identifier 
	_, ok1 := m_a_p[92] 
	fmt.Println("\nKey present or not:", ok1) 
} 
