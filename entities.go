package newsapi

import "net/http"

//Service interface with all the api methods
type Service interface {
	TopHeadlines(p RequestParams) (*APIResult, error)
	Everything(p RequestParams) (*APIResult, error)
	Sources(p RequestParams) (*APIResultSources, error)
}

//RequestParams struct for api call params
type RequestParams struct {
	Q              string
	Sources        string
	Country        string
	Category       string
	Language       string
	Domains        string
	ExcludeDomains string
	SortBy         string
	From           string
	To             string
	PageSize       int
	Page           int
}

//APIResult response from the api for articles
type APIResult struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []Article
}

//Article type
type Article struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

//APIResultSources response from the api for sources
type APIResultSources struct {
	Status  string `json:"status"`
	Sources []Source
}

//Source type
type Source struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"en"`
	Country     string `json:"country"`
}

//NewsAPI creates new api instance
type NewsAPI struct {
	APIKey string
	Client http.Client
}
