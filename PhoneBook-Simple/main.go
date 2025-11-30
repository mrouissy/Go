package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Contact struct {
	ID         int
	Name       string
	Phone      string
	Email      string
	Address    string
	IsFavorite bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Data struct {
	Contacts []Contact
	NextID   int
}

var contacts []Contact
var nextID = 1
var filename = "contacts.json"

func main() {
	reader := bufio.NewReader(os.Stdin)
	loadData()

	clearScreen()
	showWelcome()

	for {
		showMenu()
		fmt.Print("\nChoice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			clearScreen()
			addContact(reader)
		case "2":
			clearScreen()
			listContacts()
		case "3":
			clearScreen()
			viewContact(reader)
		case "4":
			clearScreen()
			updateContact(reader)
		case "5":
			clearScreen()
			deleteContact(reader)
		case "6":
			clearScreen()
			searchContacts(reader)
		case "7":
			clearScreen()
			toggleFavorite(reader)
		case "8":
			clearScreen()
			listFavorites()
		case "9":
			clearScreen()
			showStats()
		case "0":
			clearScreen()
			fmt.Println("\n👋 Goodbye!")
			fmt.Printf("Data saved to %s\n\n", filename)
			return
		default:
			fmt.Println("\n✗ Invalid choice!")
		}
	}
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func showWelcome() {
	fmt.Println("\n🎉 Welcome to Phonebook!")
	fmt.Printf("📁 File: %s\n", filename)
}

func showMenu() {
	line := strings.Repeat("=", 60)
	fmt.Println("\n" + line)
	fmt.Println("                    PHONEBOOK")
	fmt.Println(line)
	fmt.Println("\n  1. Add Contact")
	fmt.Println("  2. List All Contacts")
	fmt.Println("  3. View Contact")
	fmt.Println("  4. Update Contact")
	fmt.Println("  5. Delete Contact")
	fmt.Println("  6. Search Contacts")
	fmt.Println("  7. Toggle Favorite")
	fmt.Println("  8. List Favorites")
	fmt.Println("  9. Statistics")
	fmt.Println("  0. Exit")
	fmt.Println("\n" + line)
}

func addContact(reader *bufio.Reader) {
	fmt.Println("\n--- Add Contact ---")
	
	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	
	fmt.Print("Phone: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)
	
	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	
	fmt.Print("Address: ")
	addr, _ := reader.ReadString('\n')
	addr = strings.TrimSpace(addr)

	if name == "" || phone == "" {
		fmt.Println("\n✗ Name and Phone required!")
		return
	}

	c := Contact{
		ID:         nextID,
		Name:       name,
		Phone:      phone,
		Email:      email,
		Address:    addr,
		IsFavorite: false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	contacts = append(contacts, c)
	nextID++
	saveData()
	fmt.Printf("\n✓ Contact '%s' added!\n", name)
}

func listContacts() {
	if len(contacts) == 0 {
		fmt.Println("\nNo contacts found.")
		return
	}

	line := strings.Repeat("=", 60)
	fmt.Println("\n" + line)
	fmt.Println("                  ALL CONTACTS")
	fmt.Println(line)

	for _, c := range contacts {
		printContact(c)
	}
}

func viewContact(reader *bufio.Reader) {
	fmt.Print("\nContact ID: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("\n✗ Invalid ID!")
		return
	}

	for _, c := range contacts {
		if c.ID == id {
			line := strings.Repeat("=", 60)
			fmt.Println("\n" + line)
			fmt.Println("              CONTACT DETAILS")
			fmt.Println(line)
			printContact(c)
			return
		}
	}
	fmt.Println("\n✗ Contact not found!")
}

func updateContact(reader *bufio.Reader) {
	fmt.Println("\n--- Update Contact ---")
	fmt.Print("Contact ID: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("\n✗ Invalid ID!")
		return
	}

	found := false
	for i := range contacts {
		if contacts[i].ID == id {
			found = true
			fmt.Println("(Leave blank to keep current)")
			
			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			if name != "" {
				contacts[i].Name = name
			}
			
			fmt.Print("Phone: ")
			phone, _ := reader.ReadString('\n')
			phone = strings.TrimSpace(phone)
			if phone != "" {
				contacts[i].Phone = phone
			}
			
			fmt.Print("Email: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)
			if email != "" {
				contacts[i].Email = email
			}
			
			fmt.Print("Address: ")
			addr, _ := reader.ReadString('\n')
			addr = strings.TrimSpace(addr)
			if addr != "" {
				contacts[i].Address = addr
			}
			
			contacts[i].UpdatedAt = time.Now()
			saveData()
			fmt.Println("\n✓ Contact updated!")
			break
		}
	}
	
	if !found {
		fmt.Println("\n✗ Contact not found!")
	}
}

func deleteContact(reader *bufio.Reader) {
	fmt.Print("\nContact ID to delete: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("\n✗ Invalid ID!")
		return
	}

	fmt.Print("Are you sure? (yes/no): ")
	confirm, _ := reader.ReadString('\n')
	confirm = strings.ToLower(strings.TrimSpace(confirm))
	
	if confirm != "yes" && confirm != "y" {
		fmt.Println("\n✗ Cancelled.")
		return
	}

	for i, c := range contacts {
		if c.ID == id {
			contacts = append(contacts[:i], contacts[i+1:]...)
			saveData()
			fmt.Println("\n✓ Contact deleted!")
			return
		}
	}
	fmt.Println("\n✗ Contact not found!")
}

func searchContacts(reader *bufio.Reader) {
	fmt.Print("\nSearch (name/phone/email): ")
	query, _ := reader.ReadString('\n')
	query = strings.TrimSpace(query)
	
	if query == "" {
		fmt.Println("\n✗ Query cannot be empty!")
		return
	}

	q := strings.ToLower(query)
	var results []Contact
	
	for _, c := range contacts {
		if strings.Contains(strings.ToLower(c.Name), q) ||
			strings.Contains(strings.ToLower(c.Phone), q) ||
			strings.Contains(strings.ToLower(c.Email), q) {
			results = append(results, c)
		}
	}

	if len(results) == 0 {
		fmt.Println("\n✗ No contacts found.")
		return
	}

	line := strings.Repeat("=", 60)
	fmt.Println("\n" + line)
	fmt.Printf("            SEARCH RESULTS (%d found)\n", len(results))
	fmt.Println(line)

	for _, c := range results {
		printContact(c)
	}
}

func toggleFavorite(reader *bufio.Reader) {
	fmt.Print("\nContact ID: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("\n✗ Invalid ID!")
		return
	}

	for i := range contacts {
		if contacts[i].ID == id {
			contacts[i].IsFavorite = !contacts[i].IsFavorite
			contacts[i].UpdatedAt = time.Now()
			saveData()
			if contacts[i].IsFavorite {
				fmt.Println("\n✓ Added to favorites!")
			} else {
				fmt.Println("\n✓ Removed from favorites!")
			}
			return
		}
	}
	fmt.Println("\n✗ Contact not found!")
}

func listFavorites() {
	var favs []Contact
	for _, c := range contacts {
		if c.IsFavorite {
			favs = append(favs, c)
		}
	}

	if len(favs) == 0 {
		fmt.Println("\nNo favorites found.")
		return
	}

	line := strings.Repeat("=", 60)
	fmt.Println("\n" + line)
	fmt.Printf("            FAVORITES (%d)\n", len(favs))
	fmt.Println(line)

	for _, c := range favs {
		printContact(c)
	}
}

func showStats() {
	total := len(contacts)
	fav := 0
	for _, c := range contacts {
		if c.IsFavorite {
			fav++
		}
	}
	reg := total - fav

	line := strings.Repeat("=", 60)
	fmt.Println("\n" + line)
	fmt.Println("                STATISTICS")
	fmt.Println(line)
	fmt.Printf("\n  Total:     %d\n", total)
	fmt.Printf("  Favorites: %d\n", fav)
	fmt.Printf("  Regular:   %d\n", reg)
	fmt.Println("\n" + line)
}

func printContact(c Contact) {
	star := ""
	if c.IsFavorite {
		star = " ⭐"
	}

	fmt.Printf("\n  ID: %d%s\n", c.ID, star)
	fmt.Printf("  Name:    %s\n", c.Name)
	fmt.Printf("  Phone:   %s\n", c.Phone)
	fmt.Printf("  Email:   %s\n", c.Email)
	fmt.Printf("  Address: %s\n", c.Address)
	fmt.Printf("  Created: %s\n", c.CreatedAt.Format("2006-01-02 15:04"))
	fmt.Println(strings.Repeat("-", 60))
}

func saveData() {
	data := Data{
		Contacts: contacts,
		NextID:   nextID,
	}
	
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error saving:", err)
		return
	}

	err = os.WriteFile(filename, bytes, 0644)
	if err != nil {
		fmt.Println("Error writing:", err)
	}
}

func loadData() {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Error reading:", err)
		return
	}

	var data Data
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("Error parsing:", err)
		return
	}

	contacts = data.Contacts
	nextID = data.NextID
}
