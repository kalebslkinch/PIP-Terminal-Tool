package main

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"

)


func getPopularPodcasts() {
	handler("https://podcast-index-platform.vercel.app/api/trpc/podIndex.homePage?batch=1&input=%7B%220%22%3A%7B%22json%22%3Anull%2C%22meta%22%3A%7B%22values%22%3A%5B%22undefined%22%5D%7D%7D%7D")
}

func getHomepageData() {
	handler("https://podcast-index-platform.vercel.app/api/cachedAPI/homePage")
}

func handler(url string) {

req, err := http.NewRequest("GET",  url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	fmt.Println("Press any key to return to menu")
	fmt.Scanln()
	main()
	
	
}


func menu() tview.Primitive {
	list := tview.NewList().
		AddItem("Get Popular Podcasts", "", '1', nil).
		AddItem("Get Trending Podcasts", "", '2', nil).
		AddItem("Quit", "", 'q', func() {
			tview.NewApplication().Stop()
		})
	return list
	
}


func main() {

	terminal := tview.NewApplication()
	terminal.SetRoot(menu(), true)
	terminal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			terminal.Stop()
		}
		if event.Key() == tcell.KeyRune {
			if event.Rune() == '1' {
			terminal.Stop()
			getPopularPodcasts()

		}
			if event.Rune() == '2' {
			terminal.Stop()
			getHomepageData()

			}
			
		}
		

		return event
	})



	if err := terminal.Run(); err != nil {
		panic(err)
	}
}