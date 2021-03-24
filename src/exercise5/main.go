package main

import (
	"encoding/xml"
	parser "exercise4"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type siteMap struct {
	XMLName      xml.Name `xml:"urlset"`
	XMLNamespace string   `xml:"xmlns",attr`
	Urls         []siteMapURL
}

type siteMapURL struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}
type siteMapLinks map[string]struct{}

var emptyStruct struct{}

func main() {
	reqURL := flag.String("url", "http://www.example.com", "The URL to site map")
	flag.Parse()

	siteMap := make(siteMapLinks)
	siteMap["/"] = emptyStruct

	getSiteMap(*reqURL, &siteMap)

	fmt.Println(siteMap)
}

func getResponse(url string) *http.Response {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())

		os.Exit(1)
	}

	return response
}

func getSiteMap(reqURL string, siteMap *siteMapLinks) {
	var newLinks []string
	response := getResponse(reqURL)
	defer response.Body.Close()
	links, err := parser.Parse(response.Body)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	baseURL := &url.URL{
		Scheme: response.Request.URL.Scheme,
		Host:   response.Request.URL.Host,
	}

	for _, link := range links {
		_, alreadyAdded := (*siteMap)[link.Href]

		if isLocalLink(link.Href, baseURL.String()) && !alreadyAdded {
			(*siteMap)[link.Href] = emptyStruct
			newLinks = append(newLinks, link.Href)
		}
	}

	for _, newLink := range newLinks {
		getSiteMap(newLink, siteMap)
	}
}

func isLocalLink(link string, baseURL string) bool {
	return strings.HasPrefix(link, "/") ||
		(strings.Contains(link, baseURL) &&
			(strings.HasPrefix(link, "www.") || strings.HasPrefix(link, "http://") || strings.HasPrefix(link, "https://")))
}

func convertHrefs(links []parser.Link, baseUrl url.URL) []string {
	var hrefs []string

	for _, link := range links {
		switch {
		case strings.HasPrefix(link.Href, "/"):
			hrefs = append(hrefs, baseUrl.String()+link.Href)
		}
	}

	return hrefs
}
