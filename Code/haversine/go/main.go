package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"time"
)

type Pair struct {
	X0 float64 `json:"x0"`
	Y0 float64 `json:"y0"`
	X1 float64 `json:"x1"`
	Y1 float64 `json:"y1"`
}

type JSONInput struct {
	Pairs []Pair `json:"pairs"`
}

func HaversineOfDegrees(X0, Y0, X1, Y1, R float64) float64 {
	dY := math.Pi / 180 * (Y1 - Y0)
	dX := math.Pi / 180 * (X1 - X0)
	Y0 = math.Pi / 180 * Y0
	Y1 = math.Pi / 180 * Y1

	RootTerm := math.Pow(math.Sin(dY/2), 2) + math.Cos(Y0)*math.Cos(Y1)*math.Pow(math.Sin(dX/2), 2)
	Result := 2 * R * math.Asin(math.Sqrt(RootTerm))

	return Result
}

// Input is slower than python because python's json module is probably written in C and golang json is written in pure golang
func main() {
	EarthRadiuskm := 6371.0
	Sum := 0.0
	Count := 0

	JSONFile, err := ioutil.ReadFile("points.json")
	if err != nil {
		panic(err)
	}

	StartTime := time.Now()

	var data JSONInput
	err = json.Unmarshal(JSONFile, &data)
	if err != nil {
		panic(err)
	}

	MidTime := time.Now()

	for _, Pair := range data.Pairs {
		Sum += HaversineOfDegrees(Pair.X0, Pair.Y0, Pair.X1, Pair.Y1, EarthRadiuskm)
		Count += 1
	}

	Average := Sum / float64(Count)
	EndTime := time.Now()

	fmt.Printf("Result: %f\n", Average)
	fmt.Printf("Input = %s seconds\n", MidTime.Sub(StartTime))
	fmt.Printf("Math = %s seconds\n", EndTime.Sub(MidTime))
	fmt.Printf("Total = %s seconds\n", EndTime.Sub(StartTime))
	fmt.Printf("Throughput = %f haversines/second\n", float64(Count)/EndTime.Sub(StartTime).Seconds())
}
