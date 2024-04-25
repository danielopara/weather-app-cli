package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)


type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type weather struct {
	Name string `json:"name"`
	Main struct{
		Celsius float64 `json:"temp"`
	}`json:"main"`
}

func loadApiConfig (fileName string) (apiConfigData, error){
	bytes, err := ioutil.ReadFile(fileName)

	if err != nil{
		return apiConfigData{}, err
	}

	var configData apiConfigData
	err = json.Unmarshal(bytes, &configData)
	if err != nil{
		return apiConfigData{}, err
	}

	return configData, nil
}

func welcome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome\n"))
}

func query(city string) (weather, error){
	apiConfig, err := loadApiConfig("C:/Users/Dell/Documents/goLand-work/weather-app/apiConfig")
	if err != nil {
		return weather{}, err
	}

	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)
	if err != nil {
		return weather{}, err
	}
	defer response.Body.Close()

	var data weather
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return weather{}, err
	}

	data.Main.Celsius -= 273.15
	return data, nil
}

/** Using endpoint */
// func main() {
// 	http.HandleFunc("/welcome", welcome)

// 	http.HandleFunc("/weather/", 
// 	func(w http.ResponseWriter, r *http.Request) {
// 		city := strings.SplitN(r.URL.Path, "/", 3)[2]
// 		data, err := query(city)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(data)
// 	})
// 	http.ListenAndServe(":1000", nil)
// }



/** Asking for the user input on cli
 or Using flag on the cli
*/

func main() {
	RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use: "weather-app",
	Short : "Checks the weather of a city",
}

var weatherCmd = &cobra.Command{
	Use: "checks",
	Short : "Checks weather",
	Run: func(c *cobra.Command, args []string) {
		var city string
	flag.StringVar(&city, "city", "", "City for weather information")
	flag.Parse()

	if city == "" {
		fmt.Print("Enter city name: ")
		_, err := fmt.Scanln(&city)
		if err != nil {
			fmt.Println("Error reading city name:", err)
			os.Exit(1)
		}
	}

	data, err := query(city)
	if err != nil {
		fmt.Println("Error querying weather information:", err)
		os.Exit(1)
	}

	fmt.Printf("Weather information for %s:\n", strings.Title(city))
	fmt.Printf("Temperature: %.2f°C\n", data.Main.Celsius)
	},
}

func init(){
	RootCmd.AddCommand(weatherCmd)
}