package main

type user struct {
	ID           int
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	PhoneNumber  string
	Address1     string
	Address2     string
	City         string
	Region       string
	Zipcode      string
}

type tool struct {
	ID        int
	Title     string
	Category  string
	Price     float64
	ToolOwner int
	StartDate float64
	EndDate   float64
}

type toolOwner struct {
	ID        	 int
	Title     	 string
	Category  	 string
	Price     	 float64
	ToolOwner 	 int
	StartDate    float64
	EndDate      float64
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	PhoneNumber  string
	Address1     string
	Address2     string
	City         string
	Region       string
	Zipcode      int
}
