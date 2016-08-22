// This code fetches weather conditions from the weather undergroud API

package main

import "github.com/Jeffail/gabs"
//import "flag"
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
    conditions string
    temp string
}

func main() {

    // initialize argument variables
    var city, state string

    // assign arguments to variables
    city  = os.Args[1]
    state = os.Args[2]

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

        conditions.conditions = jsonParsed.Path("current_observation.weather").Data().(string)
        conditions.temp = jsonParsed.Path("current_observation.temperature_string").Data().(string)
        fmt.Printf("%s %s", conditions.temp, conditions.conditions)

    }
}
