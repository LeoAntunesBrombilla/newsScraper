package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gocolly/colly/v2"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

func fetchGloboNews() {
	c := colly.NewCollector()

	c.OnHTML(".feed-post-body", func(e *colly.HTMLElement) {
		title := e.ChildText(".feed-post-link")
		fmt.Println("News Title:", title)
	})

	err := c.Visit("https://g1.globo.com/")
	if err != nil {
		log.Fatal(err)
	}
}

func fetchTerraNews() {
	c := colly.NewCollector()

	counter := 1
	icon := "🌎"

	c.OnHTML("a.card-news__text--title", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.Text)

		cyan := color.New(color.FgCyan).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()

		fmt.Printf("%s %s %s %s\n", cyan(fmt.Sprintf("[%d]", counter)), yellow(icon), "News Title:", title)

		counter++
	})

	err := c.Visit("https://www.terra.com.br/noticias/")

	if err != nil {
		log.Fatal(err)
	}

}

func fetchArticleText(url string, htmlComponent string, htmlElement string) {
	c := colly.NewCollector()

	c.OnHTML(htmlComponent, func(e *colly.HTMLElement) {
		e.ForEach(htmlElement, func(_ int, el *colly.HTMLElement) {
			fmt.Println(el.Text)
		})
	})

	err := c.Visit(url)

	if err != nil {
		log.Fatal("Voce precisa de uma assinatura para ler esta notícia.")
	}
}

func fetchNews(htmlElement, url, articleComponent, articleElement string) {
	var continueFetching = true
	titles := make([]string, 0, 10)
	links := make(map[string]string)

	c := colly.NewCollector()

	c.OnHTML(htmlElement, func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.Text)
		link := e.Attr("href")

		titles = append(titles, title)
		links[title] = link
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	start := 0
	end := 10

	for continueFetching {
		displayTitles := make([]string, 0, 10)
		if end > len(titles) {
			end = len(titles)
		}

		displayTitles = append(displayTitles, titles[start:end]...)

		if end < len(titles) && len(titles) > 10 {
			displayTitles = append(displayTitles, "Carregar mais notícias...")
		}

		prompt := promptui.Select{
			Label: "Selecione uma notícia",
			Items: displayTitles,
			Size:  min(11, len(displayTitles)),
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if result == "Carregar mais notícias..." {
			fmt.Println("Carregando mais notícias...")
			start += 10
			end += 10
		} else {
			fmt.Println("Acessando notícia...")
			url := links[result]

			fetchArticleText(url, articleComponent, articleElement)
			continueFetching = false
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var rootCmd = &cobra.Command{Use: "newscli"}

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		newsSource := []string{"Globo", "UOL", "Terra"}

		prompt := promptui.Select{
			Label: "Select News Source",
			Items: newsSource,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", result)

		switch result {
		case "Globo":
			fmt.Println("Fetching news from Globo...")
			fetchNews(".container", "https://globo.com/", "div.post ", "p")
		case "UOL":
			fmt.Println("Fetching news from UOL...")
			fetchNews("a.relatedList__related", "https://www.uol.com.br/", "div.c-news__body", "p")
		case "Terra":
			fmt.Println("Fetching news from Terra...")
			fetchNews("a.card-news__text--title", "https://www.terra.com.br/noticias/", "div.article__content--body", "p.text")
		default:
			fmt.Println("Invalid choice")
		}
	}

	rootCmd.Execute()
}
