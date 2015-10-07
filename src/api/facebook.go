package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
)

var apiKey string = "CAACEdEose0cBAH3YUUv1rUHZAT8OaOgraVWtuDzV2NSfxiAxL1C6R8gvtZBhwtThgwYZAOo8gI8KsjVqxIElDFCO3PwodKZBKqdcmmLHwMcZCZCwdg2NbJSAe9QyF0pp0k9Bxg5gVIXsmkZBYwuZAnSywNsGbidxvZBMUVVaFV13q4OXOPvVmniJplhFYkjDiD1f3auM12OuPQsrkuDaWkHGUFdXBMAhZAa3MZD"

func callApi() {
    resp, err := http.Get("https://graph.facebook.com/v2.2/me/feed?access_token=apiKey")

    if err != nil {
        fmt.Printf("Error occured... exiting program.")
        os.Exit(1)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    fmt.Printf("%s\n", string(body))
}

func main() {
    callApi()
}