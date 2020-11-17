package main

import (
	"fmt"
	"io/ioutil"
	"time"

	geojson "github.com/paulmach/go.geojson"
)

func main() {
	fmt.Println("JSON LOAD TEST")
	startTime := time.Now()
	file, _ := ioutil.ReadFile("c:\\data\\test\\sejong.geojson")

	fc1, _ := geojson.UnmarshalFeatureCollection(file)

	//fc2 := geojson.NewFeatureCollection()
	//err = json.Unmarshal(file, fc2)

	features := fc1.Features

	for _, v := range features {
		_ = v.BoundingBox

	}
	//endTime :=
	fmt.Println(time.Now().Sub(startTime))

}
