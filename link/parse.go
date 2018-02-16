package link

import (
	"io"
	"golang.org/x/net/html"
	"strings"
)

//Link represent a link (<a href="....">) in a HTML tag document

type Link struct {
	Href string
	Text string
}

//function Parse take a HTML document and return
// a slise of links parsed from it

func Parse(r io.Reader) ([]Link, error)  {

	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	nodes := linkNodes(doc)
	var links []Link

	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) Link  {
	var result Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			result.Href = attr.Val
			break
		}
	}

	result.Text = text(n)
	return result
}

func text(n *html.Node) string  {
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	var result string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result += text(c)

	}

	return strings.Join(strings.Fields(result), " ")
}

func linkNodes(n *html.Node) []*html.Node  {
	if n.Type == html.ElementNode && n.Data == "a"{
		return []*html.Node{n}
	}

	var res []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res = append(res, linkNodes(c)...)
	}
	return res
}
