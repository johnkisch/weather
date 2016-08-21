// This code fetches weather conditions from the weather undergroud API
// to be displayed in the i3 status bar.

// You can edit this code!
// Click here and start typing.

package main

import "fmt"
import "os"
//import "encoding/json"
import "net/http"
import "io/ioutil"

// function for error handling
func do_err(){
    fmt.Println("Some error happened")
    os.Exit(1)
}

func main() {
    // make HTTP GET request
    resp, err := http.Get("http://api.wunderground.com/api/a771a3d22176fe4f/conditions/q/CA/San_Francisco.json")
    // if there is an error call our error handling function
    if err != nil { do_err() }
    // close our HTTP connection
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {do_err() }
    fmt.Printf("%s", body)

}
