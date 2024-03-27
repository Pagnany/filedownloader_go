package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	time_start := time.Now()

	var wg sync.WaitGroup
	file, _ := os.Open("urls/urls_2.csv")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		link := scanner.Text()
		wg.Add(1)
		go func(n string) {
			defer wg.Done()
			DownloadPicture(n)
		}(link)
	}

	wg.Wait()

	elapsed := time.Now().Sub(time_start)
	println("Total time in s: ")
	println(elapsed / 1000000000)
}

func DownloadPicture(link string) {
	split_str := strings.Split(link, "/")
	name := split_str[len(split_str)-1]

	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}

	content, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		log.Fatal(err2)
	}

	res.Body.Close()

	err3 := os.WriteFile("files/"+name, content, 0644)
	if err3 != nil {
		log.Fatal(err3)
	}
}
