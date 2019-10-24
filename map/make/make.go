package main

import "fmt"

type Vertex struct {
	Lat, long float64
}

var m map[string]Vertex

func main() {

	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		Lat:  40.68433,
		long: -74.39967,
	}

	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Bell Labs"].Lat)

	// literals
	var g = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(g)

	var t = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(t)
}
