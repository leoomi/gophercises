// package main

// import (
// 	parser "exercise4/htmlLinkParser"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	path := os.Args[1]

// 	reader, err := os.Open(path)

// 	if err != nil {
// 		fmt.Printf(err.Error())
// 		return
// 	}

// 	links, err := parser.Parse(reader)

// 	if err != nil {
// 		fmt.Printf(err.Error())
// 		return
// 	}

// 	fmt.Printf("%s", links)
// }

package htmllinkparser

import (
	"os"
	"testing"
)

func TestEx1(t *testing.T) {
	reader, _ := os.Open("ex1.html")
	links, _ := Parse(reader)

	href := "/other-page"
	text := "A link to another page"

	assertLink(t, links, Link{href, text})
}

func TestEx2(t *testing.T) {
	reader, _ := os.Open("ex2.html")
	links, _ := Parse(reader)

	var expectedLinks = []Link{
		Link{"https://www.twitter.com/joncalhoun", "Check me out on twitter oh no"},
		Link{"https://github.com/gophercises", "Gophercises is on Github !"},
	}

	for _, l := range expectedLinks {
		assertLink(t, links, l)
	}
}

func TestEx3(t *testing.T) {
	reader, _ := os.Open("ex3.html")
	links, _ := Parse(reader)

	var expectedLinks = []Link{
		Link{"https://twitter.com/marcusolsson", "@marcusolsson"},
	}

	for _, l := range expectedLinks {
		assertLink(t, links, l)
	}
}

func TestEx4(t *testing.T) {
	reader, _ := os.Open("ex4.html")
	links, _ := Parse(reader)

	var expectedLinks = []Link{
		Link{"/dog-cat", "dog cat"},
	}

	for _, l := range expectedLinks {
		assertLink(t, links, l)
	}
}

func assertLink(t *testing.T, links []Link, expected Link) {
	l, ok := testGetLink(links, expected.Href)

	if !ok {
		t.Errorf("href \"%s\" not found", expected.Href)
	}

	if l.Text != expected.Text {
		t.Errorf("Expected \"%s\" to be text for the link but found: \"%s\"", expected.Text, l.Text)
	}
}

func testGetLink(links []Link, href string) (Link, bool) {
	for _, l := range links {
		if l.Href == href {
			return l, true
		}
	}

	return Link{}, false
}
