package main

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func getPopularPodcasts(terminal *tview.Application) {
	handler("https://podcast-index-platform.vercel.app/api/trpc/podIndex.homePage?batch=1&input=%7B%220%22%3A%7B%22json%22%3Anull%2C%22meta%22%3A%7B%22values%22%3A%5B%22undefined%22%5D%7D%7D%7D","API", terminal)
}

func getHomepageData(terminal *tview.Application) {
	handler("https://podcast-index-platform.vercel.app/api/cachedAPI/homePage","API", terminal)
}

func getTrendingPageData(terminal *tview.Application) {
	handler("https://podcast-index-platform.vercel.app/api/cachedAPI/trendingPage","API", terminal)
}

func getPodcastById(id string, terminal *tview.Application) {
	handler("https://podcast-index-platform.vercel.app/api/searchBy/searchById/?id="+ id, "SEARCH",terminal)
}

func getTrendingPodcasts(max string, terminal *tview.Application) {
	handler("https://podcast-index-platform.vercel.app/api/searchBy/searchByTrending?max=" + max ,"SEARCH",terminal)
}

func handler(url string, handlerType string, terminal *tview.Application) {
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

	if handlerType == "API" {		
		showAPIMenu(terminal)
	}

	if handlerType == "SEARCH" {
		showSearchMenu(terminal)
	}
	
}

func mainMenu() tview.Primitive {
	list := tview.NewList().
		AddItem("Check API Endpoints", "", '1', nil).
		AddItem("Search", "", '2', nil).
		AddItem("Menu Option 3", "", '3', nil).
		AddItem("Quit", "", '4', nil)
	return list
}

func apiMenu() tview.Primitive {
	list := tview.NewList().
		AddItem("View Popular Podcasts Data", "", '1', nil).
		AddItem("View Home Page Cached Data", "", '2', nil).
		AddItem("View Trending Page Cached Data", "", '3', nil).
		AddItem("Return to Menu", "", '4', nil)
	return list
	
}

func mainSearchMenu() tview.Primitive {
	list := tview.NewList().
		AddItem("Podcast", "", '1', nil).
		AddItem("Episode", "", '2', nil).
		AddItem("Return to Menu", "", '3', nil)
	return list
}

func searchPodcastMenu() tview.Primitive {
	list := tview.NewList().
		AddItem("Search by Feed ID", "", '1', nil).
		AddItem("Search Trending Podcasts", "", '2', nil).
		AddItem("Return to Menu", "", '3', nil)
	return list
}

func searchEpisodeMenu() tview.Primitive {
	list := tview.NewList().
		AddItem("Search by ID", "", '1', nil).
		AddItem("Search Trending Podcasts", "", '2', nil).
		AddItem("Return to Menu", "", '3', nil)
	return list
}

func showPodcastSearchMenu(terminal *tview.Application) {
	terminal.SetRoot(searchPodcastMenu(), true)
	terminal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
			case '1': 
				fmt.Println("Enter Feed ID")
				var id string
				fmt.Scanln(&id)
				getPodcastById(id,terminal)

			case '2':
				fmt.Println("Enter Max")
				var max string
				fmt.Scanln(&max)
				getTrendingPodcasts(max,terminal)

			case '3':
				terminal.Stop()
				main()
		}

		return event
	})
}

func showSearchMenu(terminal *tview.Application) {
	terminal.SetRoot(mainSearchMenu(), true)
	terminal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		
		if event.Key() == tcell.KeyRune {

		switch event.Rune() {
		case '1':
			showPodcastSearchMenu(terminal)
		case '2':
			searchEpisodeMenu()
		case '3':
			terminal.Stop()
				main()
		}
	}

		return event
	})
}
			
func showAPIMenu(terminal *tview.Application) {

	terminal.SetRoot(apiMenu(), true)
	terminal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	
		if event.Key() == tcell.KeyRune {
			
			if event.Rune() == '1' {
				getPopularPodcasts(terminal)
			}	
			if event.Rune() == '2' {
				getHomepageData(terminal)
			}
			if event.Rune() == '3' {
				getTrendingPageData(terminal)	
			}
			if event.Rune() == '4'  {
				terminal.Stop()
				main()
			}
		}
		
		return event
	})
}

func main() {

	terminal := tview.NewApplication()

	terminal.SetRoot(mainMenu(), true)
	terminal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			terminal.Stop()
		}
		if event.Key() == tcell.KeyRune {
			if event.Rune() == '1' {
				showAPIMenu(terminal)
		}
			if event.Rune() == '2' {
				showSearchMenu(terminal)
			}
			if event.Rune() == '3' {
				getTrendingPageData(terminal)  
			}
			
		}
		
		return event
	})

	if err := terminal.Run(); err != nil {
		panic(err)
	}
}