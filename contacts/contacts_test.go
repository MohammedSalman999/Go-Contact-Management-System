package contacts

import (
	"testing"
)

func TestAddContact(t *testing.T) {
	// Initialize a test contact
	testContact := Contact{
		Name:  "John Doe",
		Phone: "123-456-7890",
		Email: "john@example.com",
	}

	// Add the test contact to the contacts map
	AddContact(testContact.Name, testContact.Phone, testContact.Email)

	// Retrieve the added contact from the contacts map
	addedContact, found := ContactsMap[testContact.Name]

	// Check if the added contact is found in the contacts map
	if !found {
		t.Errorf("Expected contact was not added.")
	}

	// Verify if the added contact matches the test contact
	if addedContact != testContact {
		t.Errorf("Added contact does not match the expected contact.")
	}
}

func TestDeleteContact(t *testing.T) {
	// Initialize a test contact
	testContact := Contact{
		Name:  "Jane Smith",
		Phone: "987-654-3210",
		Email: "jane@example.com",
	}

	// Add the test contact to the contacts map
	ContactsMap[testContact.Name] = testContact

	// Delete the test contact from the contacts map
	DeleteContact(testContact.Name)

	// Verify if the contact was deleted from the contacts map
	_, found := ContactsMap[testContact.Name]

	// Check if the contact was successfully deleted
	if found {
		t.Errorf("Contact was not deleted as expected.")
	}
}
