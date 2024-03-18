package godemo

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func FindLine(htmlBody []byte) {
	parseBody := bytes.NewBuffer(htmlBody)
	doc, err := html.Parse(parseBody)

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v \n", err)
		os.Exit(1)
	}

	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}

}
