package main

import (
    "fmt"
    "github.com/gocolly/colly"
)


type item struct {
    Name string `json:"name"`
    Price string `json:"price"`
    ImgUrl string `json:"imgurl"`
}

func main() {
    c := colly.NewCollector(
        colly.AllowedDomains("j2store.net"),
        )

    c.OnHTML("", func(h *colly.HTMLElement) {
        fmt.Println(h.Text)
    })

    c.Visit()
}
