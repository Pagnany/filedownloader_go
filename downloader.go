package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
	"sync"
	"time"
)

func main() {
	time_start := time.Now()

	var wg sync.WaitGroup
	file, _ := os.Open("urls/urls_2.csv")
	defer file.Close()

	// Get all paths in slice
	links := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		links = append(links, scanner.Text())
	}

	// Remove duplicates from slice
	slices.Sort(links)
	links = slices.Compact(links)

	for _, link := range links {
		wg.Add(1)
		go func(n string) {
			defer wg.Done()
			DownloadPicture(n)
		}(link)
	}
	wg.Wait()

	// Get the time
	elapsed := time.Now().Sub(time_start)
	println("Total time in s: ")
	println(elapsed / 1_000_000_000)
}

func DownloadPicture(link string) {
	split_str := strings.Split(link, "/")
	name := split_str[len(split_str)-1]

	// Does File exist
	_, errpath := os.Stat("files/" + name)
	if errpath == nil {
		return
	}

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
