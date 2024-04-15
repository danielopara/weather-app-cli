# Weather App

This is a simple weather application that retrieves current weather information for a specified city.

## How to Run

### Using Command Line

To run the application from the command line, follow these steps:

1. Ensure you have Go installed on your system. You can download and install it from [here](https://golang.org/dl/).
2. Clone this repository to your local machine.
3. Navigate to the directory where the `main.go` file is located.
4. Create a file named `apiConfig` in the same directory and add your OpenWeatherMap API key to it. The API key should be in JSON format with the key named `OpenWeatherMapApiKey`. For example:

```json
{
    "OpenWeatherMapApiKey": "YOUR_API_KEY_HERE"
}
```

5. Run the following command to build and execute the application:

```sh
go run main.go
```

6. Follow the on-screen instructions to enter the city name when prompted.

### Using Endpoint

Alternatively, you can run the application as a web service with an HTTP endpoint. To do this, follow these steps:

1. Ensure you have Go installed on your system. You can download and install it from [here](https://golang.org/dl/).
2. Clone this repository to your local machine.
3. Navigate to the directory where the `main.go` file is located.
4. Create a file named `apiConfig` in the same directory and add your OpenWeatherMap API key to it as described above.
5. Uncomment the code for using the endpoint in the `main()` function.
6. Run the following command to build and execute the application:

```terminal
go run main.go
```

7. The application will start a web server running on port `1000`. You can now make HTTP requests to `http://localhost:1000/weather/CITY_NAME` to retrieve weather information for the specified city.
