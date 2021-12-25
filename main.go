package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	models "weather_api/models"

	"github.com/joho/godotenv"
)

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main() {
		log.Print(models.Name())
		
		getWeatherHandler := http.HandlerFunc(getWeather)
		http.Handle("/weather", getWeatherHandler)
		http.ListenAndServe(":80", nil)
}

func getWeather(w http.ResponseWriter, r *http.Request) {
		API_KEY := os.Getenv("OPEN_WEATHER_API_KEY")

		query := r.URL.Query()
		city := query.Get("city")
		lang := query.Get("lang")
		units := query.Get("units")
		
		resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s&lang=%s", city, API_KEY, units, lang))

		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Fprintln(w, resp.Status, string(body))
}