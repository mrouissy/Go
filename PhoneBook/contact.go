package main

import (
	"fmt"
	"strings"
	"time"
)

// Contact represents a contact in the phonebook
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

// Display prints the contact information
func (c Contact) Display() {
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
	fmt.Println(strings.Repeat("-", 80))
}

// MatchesQuery checks if the contact matches search
func (c Contact) MatchesQuery(query string) bool {
	q := strings.ToLower(query)
	n := strings.ToLower(c.Name)
	p := strings.ToLower(c.Phone)
	e := strings.ToLower(c.Email)
	
	return strings.Contains(n, q) || strings.Contains(p, q) || strings.Contains(e, q)
}
