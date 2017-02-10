package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func GetResult(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	fmt.Println(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("return status code is :%s\n", resp.Status)
	}

	var results IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &results, nil
}
