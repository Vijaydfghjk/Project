package main

import "fmt"

// User defines the data model for our in-memory store
type User struct {
	ID   int
	Name string
}

// DataStore is an in-memory data store
type DataStore struct {
	users map[int]User
}

func NewDataStore() *DataStore {

	return &DataStore{

		users: make(map[int]User),
	}
}

func (ds *DataStore) Create(user User) {

	ds.users[user.ID] = user
}

func (ds *DataStore) Read(id int) (User, bool) {

	user, exists := ds.users[id] // this is map ds.users[id]  user returning the value.

	return user, exists
}

func (ds *DataStore) GetAllusers() []User {

	allusers := make([]User, 0, len(ds.users))

	for _, user := range ds.users {

		allusers = append(allusers, user)
	}

	return allusers
}

func (ds *DataStore) Update(id int, newname string) {

	if newuser, exists := ds.users[id]; exists {

		newuser.Name = newname
		ds.users[id] = newuser

	}

}

func (ds *DataStore) Delete(id int) {

	delete(ds.users, id)
}

func main() {

	getstore := NewDataStore()

	store1 := User{ID: 20, Name: "Ram"}
	store2 := User{ID: 30, Name: "Mano"}
	store3 := User{ID: 31, Name: "Thamsan"}

	getstore.Create(store3)
	getstore.Create(store1)
	getstore.Create(store2)

	reading, check := getstore.Read(20)

	fmt.Println("values:", reading, check)

	alsuers := getstore.GetAllusers()

	fmt.Println("Geting the all values:", alsuers)

	getstore.Update(20, "Raj")

	alsuers2 := getstore.GetAllusers()

	fmt.Println("afterUpdate :", alsuers2)

	getstore.Delete(31)
	alsuer3 := getstore.GetAllusers()
	fmt.Println("After delete :", alsuer3)

}
