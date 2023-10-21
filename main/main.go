package main

import (
	"fmt"
	"GoContactApp/contacts"
	"GoContactApp/fileio"
)

func main() {
	// Read contacts from the file (if any)
	err := fileio.ReadContacts()
	if err != nil {
		fmt.Println("Error reading contacts:", err)
	}

	for {
		fmt.Println("\nContact Management System")
		fmt.Println("1. Add a Contact")
		fmt.Println("2. Delete a Contact")
		fmt.Println("3. List All Contacts")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter contact details:")
			var name, phone, email string
			fmt.Print("Name: ")
			fmt.Scanln(&name)
			fmt.Print("Phone: ")
			fmt.Scanln(&phone)
			fmt.Print("Email: ")
			fmt.Scanln(&email)

			contacts.AddContact(name, phone, email)
			fileio.WriteContacts()
			fmt.Println("Contact added successfully!")

		case 2:
			fmt.Print("Enter the name to delete: ")
			var deleteName string
			fmt.Scanln(&deleteName)

			contacts.DeleteContact(deleteName)
			fileio.WriteContacts()
			fmt.Println("Contact deleted successfully!")

		case 3:
			fmt.Println("List of Contacts:")
			for name, contact := range contacts.ContactsMap {
				fmt.Printf("Name: %s, Phone: %s, Email: %s\n", name, contact.Phone, contact.Email)
			}

		case 4:
			fmt.Println("Exiting the Contact Management System.")
			return

		default:
			fmt.Println("Invalid choice. Please choose a valid option.")
		}
	}
}
