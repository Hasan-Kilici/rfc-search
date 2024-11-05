package utils

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/cache/diskcache"
	"github.com/geziyor/geziyor/cache"
	"github.com/geziyor/geziyor/client"
	"github.com/PuerkitoBio/goquery"
)

func Search(title string) []RFC {
	var rfcs []RFC

	baseURL := "https://www.rfc-editor.org/search/rfc_search_detail.php"
	queryParams := url.Values{}
	queryParams.Add("title", title)
	queryParams.Add("pubstatus[]", "Any")
	queryParams.Add("pub_date_type", "any")
	fullURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{fullURL},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			rfcs = parseRFCData(g, r)
		},
		Cache:       diskcache.New(".cache"),
		CachePolicy: cache.RFC2616,
	}).Start()

	return rfcs
}

func parseRFCData(g *geziyor.Geziyor, r *client.Response) []RFC {
	var rfcs []RFC

	r.HTMLDoc.Find("table.gridtable tr").Each(func(index int, row *goquery.Selection) {
		if index == 0 {
			return 
		}

		var rfc RFC
		rfc.Number = parseRFCNumber(row.Find("td").Eq(0).Text())
		rfc.Files = parseFiles(row.Find("td").Eq(1).Text())
		rfc.Title = strings.TrimSpace(row.Find("td").Eq(2).Text())
		rfc.Authors = parseAuthors(row.Find("td").Eq(3).Text())
		rfc.Date = strings.TrimSpace(row.Find("td").Eq(4).Text())
		rfc.MoreInfo = strings.TrimSpace(row.Find("td").Eq(5).Text())
		rfc.Status = strings.TrimSpace(row.Find("td").Eq(6).Text())

		rfcs = append(rfcs, rfc)
	})

	return rfcs
}

func parseRFCNumber(numberText string) int {
	parts := strings.Split(numberText, "RFC")
	if len(parts) > 1 {
		trimmed := strings.TrimSpace(parts[1])
		if number, err := strconv.Atoi(trimmed); err == nil {
			return number
		}
	}
	return 0
}

func parseFiles(filesText string) []string {
	files := strings.Split(filesText, ",")
	for i := range files {
		files[i] = strings.TrimSpace(strings.ToLower(files[i]))
	}
	return files
}

func parseAuthors(authorsText string) []string {
	authors := strings.Split(authorsText, ",")
	for i := range authors {
		authors[i] = strings.TrimSpace(authors[i])
	}
	return authors
}
