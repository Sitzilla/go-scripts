package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

type Page struct {
    Title string
    Body []byte
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)

    // if error return nil
    if err != nil {
        return nil, err
    }

    return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
    // test to create and write to file then read from it
    p1 := &Page{Title: "TestPage", Body: []byte("This is sample data for the wiki page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
    
    // test to write code to localhost:8080
    // serve the code by typing:
    // go build wiki.go
    // ./wiki
    http.HandleFunc("/view/", viewHandler)
    http.ListenAndServe(":8080", nil)
}


