# NewsAPI SDK for Go

Read the full API documentation on https://newsapi.org

### Installation

```sh
$ go get github.com/flakaal/newsapi
```

### Examples

#### Create new api instance

You need to get an api key from https://newsapi.org/register

```
api := newsapi.NewsAPI{APIKey: "<YOUR API KEY>"}
```

#### Request Parameters

```
params := newsapi.RequestParams{Country: "us"}
```

Creates a query string for the request

### Top Headlines

```
params := newsapi.RequestParams{Country: "us"}

res, err := api.TopHeadlines(params)
if err != nil {
	log.Print(err)
}
fmt.Println(res)
```

### Everything

```
params := newsapi.RequestParams{Q: "bitcoin"}

res, err := api.Everything(params)
if err != nil {
	log.Print(err)
}
```

### Sources

```
params := newsapi.RequestParams{Category: "science"}

res, err := api.Sources(params)
if err != nil {
	log.Print(err)
}
```
