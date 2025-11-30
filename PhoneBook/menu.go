package main

import (
	"bufio"
	"fmt"
	"strings"
)

func DisplayMenu() {
	line := strings.Repeat("=", 50)
	fmt.Println("\n" + line)
	fmt.Println("		PHONEBOOK APPLICATION")
	fmt.Println(line)
	fmt.Println("\n  1.  Add New Contact")
	fmt.Println("  2.  List All Contacts")
	fmt.Println("  3.  View Contact Details")
	fmt.Println("  4.  Update Contact")
	fmt.Println("  5.  Delete Contact")
	fmt.Println("  6.  Search Contacts")
	fmt.Println("  7.  Toggle Favorite")
	fmt.Println("  8.  List Favorite Contacts")
	fmt.Println("  9.  Show Statistics")
	fmt.Println("  0.  Exit")
	fmt.Println("\n" + line)
}


func GetInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// ShowStatistics shows contact stats
func ShowStatistics(total, fav, reg int) {
	line := strings.Repeat("=", 80)
	fmt.Println("\n" + line)
	fmt.Println("                            STATISTICS")
	fmt.Println(line)
	fmt.Printf("\n  Total Contacts:    %d\n", total)
	fmt.Printf("  Favorite Contacts: %d\n", fav)
	fmt.Printf("  Regular Contacts:  %d\n", reg)
	fmt.Println("\n" + line)
}

func ShowWelcome() {
	fmt.Println("\n🎉 Welcome to Phonebook Application!")
}

// ShowGoodbye shows exit message
func ShowGoodbye(file string) {
	fmt.Println("\n👋 Thank you for using Phonebook Application!")
	fmt.Printf("All data has been saved to %s\n\n", file)
}
