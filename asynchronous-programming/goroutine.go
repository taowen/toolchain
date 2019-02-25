package main

import (
    "fmt"
    "math/rand"
)

func searchKeyword(kw string) chan int {
    fmt.Printf("search %s\n", kw);

    pages := make(chan string)
    urls := make(chan string)
    rooms := make(chan int)

    pages <- fmt.Sprintf("http://seed.page/%s", kw)

    go fetchPages(pages, urls)
    go fetchUrls(urls, rooms)

    return rooms
}

func fetchPages(pages chan string, urls chan string) {
    curPage := 0
    maxPageCount := 10

    for ;; {
        page := <-pages
        fmt.Println(page)

        for i := 0; i < 10; i++{
            urls <- fmt.Sprintf("http://sub.link/%d", i)
        }

        for j := 0; j < 2 && curPage < maxPageCount; j++ {
            pages <- fmt.Sprintf("http://%d.page/", j)
            curPage += 1
        }
    }
}

func fetchUrls(urls chan string, rooms chan int) {
    for ;; {
        url := <-urls
        fmt.Println(url)

        rooms <- rand.Int()
    }
}

func main() {
    keyword := "keyword"

    rooms := searchKeyword(keyword)
    for {
        room := <-rooms
        fmt.Println(room)
    }
}
