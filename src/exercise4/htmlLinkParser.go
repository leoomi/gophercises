package htmllinkparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(reader io.Reader) ([]Link, error) {
	node, err := html.Parse(reader)

	if err != nil {
		return nil, err
	}

	links := dfsNode(node)
	return links, nil
}

func dfsNode(node *html.Node) []Link {
	if node.Type == html.ElementNode && node.Data == "a" {
		return getLink(node)
	}

	links := []Link{}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		l := dfsNode(c)

		if len(l) > 0 {
			links = append(links, l...)
		}
	}

	return links
}

func getLink(node *html.Node) []Link {
	attrs := node.Attr

	for _, attr := range attrs {
		if attr.Key != "href" {
			continue
		}

		link := Link{
			Href: attr.Val,
			Text: dfsText(node),
		}

		return []Link{link}
	}

	return []Link{}
}

func dfsText(node *html.Node) string {
	if node.Type == html.TextNode {
		return strings.TrimSpace(node.Data)
	}

	var text string = ""
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		t := dfsText(c)

		if len(t) > 0 {
			text = strings.TrimSpace(text + " " + strings.Join(strings.Fields(t), " "))
		}
	}

	return text
}
