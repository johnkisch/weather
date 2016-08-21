// This code fetches weather conditions from the weather undergroud API

package main

import "github.com/Jeffail/gabs"
import "flag"
import "fmt"
import "os"
import "net/http"
import "io/ioutil"


// function for error handling
func do_err(msg string){
    fmt.Printf("Some error happened: %s", msg)
    os.Exit(1)
}

// structure containing parsed weather data
type Weather struct {
    weather string
    temp_f float64
    temp_c float64
}

func main() {
    // parse flags
    scale := flag.String("scale", "f", "Temp in Celcius (C) or Fahrenheit (F)")

    flag.Parse()

    // initialize argument variables
    var city, state string

    // assign arguments to variables
    if flag.NArg() == 2 {
        city  = flag.Arg(0)
        state = flag.Arg(1)
    }

    // make HTTP GET request
    resp, err := http.Get("http://api.wunderground.com/api/a771a3d22176fe4f/conditions/q/"+ state +"/"+ city + ".json")

    // if there is an error call our error handling function
    if err != nil { do_err("http GET failed.") }

    // close our HTTP connection
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {do_err("ReadAll Failed.") }

    // parse our JSON from wunderground API
    jsonParsed, err := gabs.ParseJSON(body)

    // check for errors
    invalid := jsonParsed.ExistsP("response.error")
    ambig   := jsonParsed.ExistsP("response.results")

    if invalid {
        do_err(jsonParsed.Path("response.error.description").Data().(string))
    } else if ambig {
        do_err("Multiple results were returned for your query.")

    // if no errors exist, return result
    } else {

        var conditions Weather

        conditions.weather = jsonParsed.Path("current_observation.weather").Data().(string)
        conditions.temp_f = jsonParsed.Path("current_observation.temp_f").Data().(float64)
        conditions.temp_c = jsonParsed.Path("current_observation.temp_c").Data().(float64)
        fmt.Println(conditions.weather)

        if *scale == "f"{
            fmt.Printf("%.1f Degrees Fahrenheit", conditions.temp_f)
        } else if *scale == "c" {
            fmt.Printf("%.1f Degrees Celcius", conditions.temp_c)
        }
    }
}
