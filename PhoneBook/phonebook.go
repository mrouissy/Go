package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Phonebook manages contacts
type Phonebook struct {
	Contacts []Contact
	NextID   int
	filename string
}

// New creates a new phonebook
func New(file string) *Phonebook {
	pb := &Phonebook{
		Contacts: []Contact{},
		NextID:   1,
		filename: file,
	}
	pb.load()
	return pb
}

func (pb *Phonebook) AddContact(name, phone, email, addr string) {
	c := Contact{
		ID:         pb.NextID,
		Name:       name,
		Phone:      phone,
		Email:      email,
		Address:    addr,
		IsFavorite: false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	pb.Contacts = append(pb.Contacts, c)
	pb.NextID++
	pb.save()
	fmt.Printf("\n✓ Contact '%s' added!\n", name)
}

func (pb *Phonebook) ListContacts() {
	if len(pb.Contacts) == 0 {
		fmt.Println("\nNo contacts found.")
		return
	}

	line := strings.Repeat("=", 80)
	fmt.Println("\n" + line)
	fmt.Println("           ALL CONTACTS")
	fmt.Println(line)

	for _, c := range pb.Contacts {
		c.Display()
	}
}


func (pb *Phonebook) ViewContact(id int) {
	c, ok := pb.find(id)
	if !ok {
		fmt.Println("\n✗ Contact not found!")
		return
	}

	line := strings.Repeat("=", 80)
	fmt.Println("\n" + line)
	fmt.Println("           CONTACT DETAILS")
	fmt.Println(line)
	c.Display()
}

func (pb *Phonebook) UpdateContact(id int, name, phone, email, addr string) {
	for i, c := range pb.Contacts {
		if c.ID == id {
			if name != "" {
				pb.Contacts[i].Name = name
			}
			if phone != "" {
				pb.Contacts[i].Phone = phone
			}
			if email != "" {
				pb.Contacts[i].Email = email
			}
			if addr != "" {
				pb.Contacts[i].Address = addr
			}
			pb.Contacts[i].UpdatedAt = time.Now()
			pb.save()
			fmt.Println("\n✓ Contact updated!")
			return
		}
	}
	fmt.Println("\n✗ Contact not found!")
}


func (pb *Phonebook) DeleteContact(id int) {
	for i, c := range pb.Contacts {
		if c.ID == id {
			pb.Contacts = append(pb.Contacts[:i], pb.Contacts[i+1:]...)
			pb.save()
			fmt.Println("\n✓ Contact deleted!")
			return
		}
	}
	fmt.Println("\n✗ Contact not found!")
}

// SearchContacts finds contacts
func (pb *Phonebook) SearchContacts(query string) {
	var results []Contact

	for _, c := range pb.Contacts {
		if c.MatchesQuery(query) {
			results = append(results, c)
		}
	}

	if len(results) == 0 {
		fmt.Println("\n✗ No contacts found.")
		return
	}

	line := strings.Repeat("=", 80)
	fmt.Println("\n" + line)
	fmt.Printf("                      SEARCH RESULTS (%d found)\n", len(results))
	fmt.Println(line)

	for _, c := range results {
		c.Display()
	}
}

// ToggleFavorite toggles favorite status
func (pb *Phonebook) ToggleFavorite(id int) {
	for i, c := range pb.Contacts {
		if c.ID == id {
			pb.Contacts[i].IsFavorite = !pb.Contacts[i].IsFavorite
			pb.Contacts[i].UpdatedAt = time.Now()
			pb.save()
			if pb.Contacts[i].IsFavorite {
				fmt.Println("\n✓ Added to favorites!")
			} else {
				fmt.Println("\n✓ Removed from favorites!")
			}
			return
		}
	}
	fmt.Println("\n✗ Contact not found!")
}

// ListFavorites shows favorite contacts
func (pb *Phonebook) ListFavorites() {
	var favs []Contact
	for _, c := range pb.Contacts {
		if c.IsFavorite {
			favs = append(favs, c)
		}
	}

	if len(favs) == 0 {
		fmt.Println("\nNo favorites found.")
		return
	}

	line := strings.Repeat("=", 80)
	fmt.Println("\n" + line)
	fmt.Printf("                      FAVORITE CONTACTS (%d)\n", len(favs))
	fmt.Println(line)

	for _, c := range favs {
		c.Display()
	}
}

// GetStatistics returns stats
func (pb *Phonebook) GetStatistics() (total, fav, reg int) {
	total = len(pb.Contacts)
	for _, c := range pb.Contacts {
		if c.IsFavorite {
			fav++
		}
	}
	reg = total - fav
	return
}

// save saves to JSON file
func (pb *Phonebook) save() {
	data, err := json.MarshalIndent(pb, "", "  ")
	if err != nil {
		fmt.Println("Error saving:", err)
		return
	}

	err = os.WriteFile(pb.filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing:", err)
	}
}

// load loads from JSON file
func (pb *Phonebook) load() {
	data, err := os.ReadFile(pb.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Error reading:", err)
		return
	}

	err = json.Unmarshal(data, pb)
	if err != nil {
		fmt.Println("Error parsing:", err)
	}
}

// find finds contact by ID
func (pb *Phonebook) find(id int) (Contact, bool) {
	for _, c := range pb.Contacts {
		if c.ID == id {
			return c, true
		}
	}
	return Contact{}, false
}
