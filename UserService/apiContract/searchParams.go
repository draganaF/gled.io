package apicontract

type SearchParams struct {
	Name     string
	LastName string
	Email    string
	Role     int

	FromNumberOfBoughtTickets  int
	FromNumberOfResevedTickets int
	FromNumberOfSoldTickets    int

	ToNumberOfBoughtTickets  int
	ToNumberOfResevedTickets int
	ToNumberOfSoldTickets    int
}
