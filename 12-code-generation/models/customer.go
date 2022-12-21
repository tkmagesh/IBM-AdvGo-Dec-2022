package models

/* //go:generate go run ../col-gen.go -P models -N Customer */
//go:generate col-gen-app -P models -N Customer
//go:generate go fmt Customers.go
type Customer struct {
	Id      int
	Name    string
	Address string
}
