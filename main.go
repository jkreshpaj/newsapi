package newsapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

// TopHeadlines gets top news
func (n *NewsAPI) TopHeadlines(p RequestParams) (*APIResult, error) {
	result := APIResult{}

	query := GetParams(&p)

	req, err := http.NewRequest("GET", "https://newsapi.org/v2/top-headlines"+query, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+n.APIKey)

	resp, err := n.Client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(body), &result)

	return &result, nil
}

// Everything gets all news based on query
func (n *NewsAPI) Everything(p RequestParams) (*APIResult, error) {
	result := APIResult{}

	query := GetParams(&p)

	log.Println(query)

	req, err := http.NewRequest("GET", "https://newsapi.org/v2/everything"+query, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+n.APIKey)

	resp, err := n.Client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(body), &result)

	return &result, nil
}

// Sources gets the subset of news publishers
func (n *NewsAPI) Sources(p RequestParams) (*APIResultSources, error) {
	result := APIResultSources{}

	query := GetParams(&p)

	log.Println(query)

	req, err := http.NewRequest("GET", "https://newsapi.org/v2/sources"+query, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+n.APIKey)

	resp, err := n.Client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(body), &result)

	return &result, nil
}

//GetParams returns query string from RequestParams
func GetParams(p *RequestParams) string {
	var query string
	params := reflect.ValueOf(p).Elem()
	for i := 0; i < params.NumField(); i++ {
		key := strings.ToLower(string(params.Type().Field(i).Name[0])) + string(params.Type().Field(i).Name[1:])
		if params.Field(i).String() != "" && params.Field(i).String() != "<int Value>" {
			if query == "" {
				query += "?" + key + "=" + params.Field(i).String()
			} else {
				query += "&" + key + "=" + params.Field(i).String()
			}
		} else if params.Field(i).String() == "<int Value>" && params.Field(i).Int() != 0 {
			val := fmt.Sprintf("%d", params.Field(i).Int())
			if query == "" {
				query += "?" + key + "=" + val
			} else {
				query += "&" + key + "=" + val
			}
		}
	}
	return query
}
