package main

import "fmt"

func sumInt(list []int) int {
	var result int
	for _, no := range list {
		result += no
	}
	return result
}

func sumFloat(list []float32) float32 {
	var result float32
	for _, no := range list {
		result += no
	}
	return result
}

// type constraint
type Numbers interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// func sum[T int | float32 | int32 | int64](list []T) T {
func sum[T Numbers](list []T) T {
	var result T
	for _, no := range list {
		result += no
	}
	return result
}

func filter[T any](list []T, predicate func(T) bool) []T {
	result := []T{}
	for _, item := range list {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func main() {
	ints := []int{3, 1, 4, 2, 5, 8, 6, 7, 9}
	// fmt.Println(sumInt(ints))
	fmt.Println(sum(ints))
	fmt.Println("Even ints :", filter(ints, func(x int) bool {
		return x%2 == 0
	}))

	floats := []float32{3, 1, 4, 2, 5, 8, 6, 7, 9}
	// fmt.Println(sumFloat(floats))
	fmt.Println(sum(floats))
	fmt.Println("More than 5 floats :", filter(floats, func(x float32) bool {
		return x > 5
	}))

	products := []Product{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	stationaryProducts := filter(products, func(p Product) bool {
		return p.Category == "Stationary"
	})
	for _, p := range stationaryProducts {
		fmt.Println(p)
	}

}
