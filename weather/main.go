package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Location struct {
		Name           string `json:"name"`
		Country        string `json:"country"`
		LocalTimeEpoch int    `json:"localtime_epoch"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		TempC            float32 `json:"temp_c"`
		Condition        struct {
			Text string `json:"text"`
		} `json:"condition"`
		FeelsLikeC float32 `json:"feelslike_c"`
	} `json:"current"`
}

func main() {
	var api, place string
	if len(os.Args) > 1 {
		place = os.Args[1]
	} else {
		place = "narayanghat"
	}

	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Failed to read envs")
	} else {
		api = envs["WEATHER_KEY"]
	}

	path := "http://api.weatherapi.com/v1/current.json?key=" + api + "&q=" + place + "&aqi=no"
	res, err := http.Get(path)
	if err != nil {
		log.Fatal("Error while calling the api")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Cannot Read the body")
	}
	// fmt.Println("\n JSON: \n", string(body), "\n")

	var info WeatherResponse
	error := json.Unmarshal(body, &info)
	if error != nil {
		log.Fatal("Cannot parse JSON!")
	}
	// fmt.Println("Struct: ", info, "\n")

	date := time.Unix(int64(info.Current.LastUpdatedEpoch), 0)
	fmt.Printf("%s, %s [%s]\n=> %s (%.0fC) Feels Like %.0fC",
		info.Location.Name,
		info.Location.Country,
		date.Format(time.ANSIC),
		info.Current.Condition.Text,
		info.Current.TempC,
		info.Current.FeelsLikeC,
	)
}
