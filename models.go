package main

type user struct {
	Id           int
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	PhoneNumber  string
	Age          int
	Address1     string
	Address2     string
	City         string
	Region       string
	Zipcode      int
}

type tool struct {
	id        int
	title     string
	toolType  string
	price     float64
	toolOwner string
}
