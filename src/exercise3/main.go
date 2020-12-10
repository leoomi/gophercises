package main

import (
	"encoding/json"
	"exercise3/handlers"
	"exercise3/models"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	filePath := flag.String("file", "gopher.json", "The CSV file to be read. The default is problems.csv")
	flag.Parse()

	file, err := ioutil.ReadFile(*filePath)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var data map[string]models.StoryArc

	if err = json.Unmarshal(file, &data); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	intro, ok := data["intro"]
	if !ok {
		fmt.Println("The json does not contain an intro story")
		os.Exit(1)
	}

	fmt.Println(intro)
	http.HandleFunc("/", handlers.StoryHandler(data))
	http.ListenAndServe(":8080", nil)
}
