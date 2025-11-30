package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	file := "phonebook.json"
	pb := New(file)

	ClearScreen()
	ShowWelcome()

	for {
		DisplayMenu()
		choice := GetInput(reader, "\nEnter your choice: ")

		switch choice {
		case "1":
			ClearScreen()
			fmt.Println("\n--- Add New Contact ---")
			name := GetInput(reader, "Name: ")
			phone := GetInput(reader, "Phone: ")
			email := GetInput(reader, "Email: ")
			addr := GetInput(reader, "Address: ")

			if !ValidateContact(name, phone) {
				fmt.Println("\n✗ Name and Phone are required!")
			} else {
				pb.AddContact(name, phone, email, addr)
			}

		case "2":
			ClearScreen()
			pb.ListContacts()

		case "3":
			ClearScreen()
			idStr := GetInput(reader, "\nEnter Contact ID: ")
			id, err := ParseID(idStr)
			if err != nil {
				fmt.Println("\n✗ Invalid ID!")
			} else {
				pb.ViewContact(id)
			}

		case "4":
			ClearScreen()
			fmt.Println("\n--- Update Contact ---")
			idStr := GetInput(reader, "Enter Contact ID: ")
			id, err := ParseID(idStr)
			if err != nil {
				fmt.Println("\n✗ Invalid ID!")
				continue
			}

			fmt.Println("(Leave blank to keep current)")
			name := GetInput(reader, "Name: ")
			phone := GetInput(reader, "Phone: ")
			email := GetInput(reader, "Email: ")
			addr := GetInput(reader, "Address: ")

			pb.UpdateContact(id, name, phone, email, addr)

		case "5":
			ClearScreen()
			idStr := GetInput(reader, "\nEnter Contact ID to delete: ")
			id, err := ParseID(idStr)
			if err != nil {
				fmt.Println("\n✗ Invalid ID!")
			} else {
				confirm := GetInput(reader, "Are you sure? (yes/no): ")
				if ConfirmAction(confirm) {
					pb.DeleteContact(id)
				} else {
					fmt.Println("\n✗ Cancelled.")
				}
			}

		case "6":
			ClearScreen()
			query := GetInput(reader, "\nSearch (name/phone/email): ")
			if IsEmpty(query) {
				fmt.Println("\n✗ Query cannot be empty!")
			} else {
				pb.SearchContacts(query)
			}

		case "7":
			ClearScreen()
			idStr := GetInput(reader, "\nEnter Contact ID: ")
			id, err := ParseID(idStr)
			if err != nil {
				fmt.Println("\n✗ Invalid ID!")
			} else {
				pb.ToggleFavorite(id)
			}

		case "8":
			ClearScreen()
			pb.ListFavorites()

		case "9":
			ClearScreen()
			total, fav, reg := pb.GetStatistics()
			ShowStatistics(total, fav, reg)

		case "0":
			ClearScreen()
			ShowGoodbye(file)
			return

		default:
			fmt.Println("\n✗ Invalid choice!")
		}
	}
}
