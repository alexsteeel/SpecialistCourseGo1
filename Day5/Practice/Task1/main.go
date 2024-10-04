package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, pagesChan chan Page) {
	// Получить URL:
	response, _ := http.Get(url)

	defer response.Body.Close()

	// Получить тело ответа
	body, _ := ioutil.ReadAll(response.Body)
	pagesChan <- Page{url, len(body)}
}

func main() {
	pagesChan := make(chan Page)

	pages := []string{
		"https://github.com",
		"https://yandex.ru",
		"https://google.com",
		"https://yahoo.com",
		"https://habr.com",
	}

	for _, page := range pages {
		go responseSize(page, pagesChan)
		resPage := <-pagesChan
		fmt.Println(resPage)
	}
}
