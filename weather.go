// This code fetches weather conditions from the weather undergroud API
// to be displayed in the i3 status bar.


package main

import "github.com/Jeffail/gabs"
import "fmt"
import "os"
import "net/http"
import "io/ioutil"

// function for error handling
func do_err(){
    fmt.Println("Some error happened")
    os.Exit(1)
}

type Weather struct {
    weather string
    temp_f float64
    temp_c float64
}

func main() {
    // select a place from which to query weather
    city  := os.Args[1]
    state := os.Args[2]

    // make HTTP GET request
    resp, err := http.Get("http://api.wunderground.com/api/a771a3d22176fe4f/conditions/q/"+ state +"/"+ city + ".json")

    // if there is an error call our error handling function
    if err != nil { do_err() }

    // close our HTTP connection
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {do_err() }
    //fmt.Printf("%s", body)

    // parse our JSON from wunderground API
    jsonParsed, err := gabs.ParseJSON(body)

    invalid := jsonParsed.ExistsP("response.error.type")
    if invalid{
        do_err()
    } else{
        var conditions Weather

        conditions.weather = jsonParsed.Path("current_observation.weather").Data().(string)
        conditions.temp_f = jsonParsed.Path("current_observation.temp_f").Data().(float64)
        conditions.temp_c = jsonParsed.Path("current_observation.temp_c").Data().(float64)

        fmt.Println(conditions)
    }


}
