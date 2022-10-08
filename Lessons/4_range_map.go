package range_maps

import (
	"fmt"

	"mapstructure"
)

func main() {
	maps()
}

// range
func ranges() {
	var arr = []string{"a", "b", "c"}
	for _, l := range arr {
		fmt.Println(l)
	}
}

// maps
type Point struct {
	X int `mapstructure:"X_point"`
	Y int `mapstructure:"Y_point"`
}

func maps() {
	pointMap := map[string]Point{}
	anotherPointMap := make(map[int]Point)

	pointMap["a"] = Point{1, 2}
	fmt.Println(pointMap)
	fmt.Println(pointMap["a"])

	anotherPointMap[1] = Point{3, 4}
	fmt.Println(anotherPointMap)
	fmt.Println(anotherPointMap[1])

	testmap := map[string]int{
		"X_point": 8,
		"Y_point": 9,
	}

	p1 := Point{}
	mapstructure.Decode(testmap, &p1)

	fmt.Println(p1)
}
