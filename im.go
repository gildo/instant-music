package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter songname/lyrics/artist or other\n>")

	input, _ := consoleReader.ReadString('\n')
	search := url.QueryEscape(input)
	fmt.Println("Searching...")

	doc, err := goquery.NewDocument("https://www.youtube.com/results?search_query=" + search)
	if err != nil {
		log.Fatal(err)
	}

	var videos = []map[string]string{}

	doc.Find(".section-list .yt-lockup-content").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".yt-lockup-title a").Text()
		videoURL, _ := s.Find(".yt-lockup-title a").Attr("href")
		videos = append(videos, map[string]string{title: videoURL})
	})

	for i, v := range videos {
		for k := range v {
			fmt.Printf("%d => %s\n", i, k)
		}
	}

	fmt.Println("Pick one: ")

	var choose int
	fmt.Scanf("%d", &choose)

	selected := videos[choose]

	var videoLink string

	for _, v := range selected {
		videoLink = v
	}

	commands := []string{
		"--extract-audio",
		"--audio-format", "mp3",
		"--audio-quality", "0",
		"http://www.youtube.com" + videoLink}

	cmd := exec.Command("youtube-dl", commands...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}
