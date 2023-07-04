package utils

import (
	"github.com/antchfx/htmlquery"
	"regexp"
	"strings"
)

func ExcerptHTML(htmlCode []byte) (result []string, err error) {
	doc, err := htmlquery.Parse(strings.NewReader(string(htmlCode)))
	if err != nil {
		return
	}

	xpath := "//*/@*[contains(., '/content/')]"
	nodes, err := htmlquery.QueryAll(doc, xpath)
	if err != nil {
		return
	}

	if len(nodes) == 0 {
		return
	}

	for _, node := range nodes {
		intelText := htmlquery.InnerText(node)

		re := regexp.MustCompile(`/content/(\w+)`)
		matches := re.FindAllStringSubmatch(intelText, -1)

		for _, match := range matches {
			result = append(result, match[1])
		}
	}
	return
}
