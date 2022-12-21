package models

/* //go:generate go run ../col-gen.go -P models -N Product */
//go:generate col-gen-app -P models -N Product
//go:generate go fmt Products.go
type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}
