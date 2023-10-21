package fileio

import (
	"GoContactApp/contacts"
	"encoding/json"
	"os"
)

const filename = "contacts.json"

func ReadContacts() error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &contacts.ContactsMap)
	if err != nil {
		return err
	}

	return nil
}

func WriteContacts() error {
	data, err := json.Marshal(contacts.ContactsMap) // Change variable name to 'data'
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, os.ModePerm) // Use the existing 'err' variable
	if err != nil {
		return err
	}

	return nil
}

