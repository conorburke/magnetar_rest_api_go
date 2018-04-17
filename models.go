package main

type user struct {
	ID           int
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
	ID        int
	Title     string
	ToolType  string
	Price     float64
	ToolOwner int
}
