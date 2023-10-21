package contacts

type Contact struct {
    Name  string
    Phone string
    Email string
}

var ContactsMap = make(map[string]Contact)

func AddContact(name, phone, email string) {
    contact := Contact{name, phone, email}
    ContactsMap[name] = contact
}

func DeleteContact(name string) {
    delete(ContactsMap, name)
}