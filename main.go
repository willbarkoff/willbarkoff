package main

import (
	"fmt"
	"math/rand"
	"os"
	"text/template"
	"time"

	"github.com/mmcdole/gofeed"
)

type Post struct {
	Title string
	URL   string
	Date  string
}

type TemplateData struct {
	Posts       []Post
	Greeting    string
	GeneratedAt string
}

var greetings = []string{"Hello", "Howdy", "Hi", "G'day", "Hey", "Hiya"}

var t = template.Must(template.ParseFiles("readme-template.md"))

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	greetingNum := rand.Intn(len(greetings))

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://willbarkoff.dev/feed.xml")
	items := feed.Items

	posts := []Post{}

	for i, item := range items {
		if i > 3 {
			break
		}

		p := Post{}
		p.Title = item.Title
		p.URL = item.Link
		p.Date = item.PublishedParsed.Format("January 2, 2006")
		posts = append(posts, p)
	}

	_ = os.Remove("readme.md")
	f, err := os.Create("readme.md")
	defer f.Close()

	if err != nil {
		panic(err)
	}
	t.Execute(f, TemplateData{Posts: posts, Greeting: greetings[greetingNum], GeneratedAt: time.Now().Format("January 2, 2006 at 3:04PM")})

	fmt.Println(feed.Title)
}
