package main

import (
	"flag"
	"fmt"
	"github.com/HuguesBt/search/pkg/search"
	"log"
)

var (
	action string
	id     int
	title  string
	text   string
	tags   string
	query  string
)

func initFlags() {
	flag.StringVar(&action, "action", "", "Action to realize; add/get/search")
	flag.StringVar(&title, "title", "", "Title to add")
	flag.StringVar(&text, "text", "", "Text to add")
	flag.StringVar(&tags, "tags", "", "Tags to add")
	flag.IntVar(&id, "id", 0, "Id for get")
	flag.StringVar(&query, "query", "", "Query to search")
	flag.Parse()
}

func initDb() {
	err := search.InitDB()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initFlags()
	initDb()
	defer search.DbObj.GetDb().Close()

	switch action {
	case "add":
		if title == "" {
			log.Fatal("Title is required")
		} else if text == "" {
			log.Fatal("Text is required")
		} else if doc, err := search.DbObj.AddDocument(search.Document{
			Title: title,
			Text:  text,
			Tags:  tags,
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Added document", doc.Id)
		}
		break
	case "update":
		if tags == "" {
			log.Fatal("tags is required")
		} else if id == 0 {
			log.Fatal("Id is required")
		} else if doc, err := search.DbObj.UpdateDocumentTag(search.Document{
			Id:   int64(id),
			Tags: tags,
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Added document", doc.Id)
		}
		break
	case "get":
		if id == 0 {
			log.Fatal("Id is required")
		} else if doc, err := search.DbObj.GetDocument(int64(id)); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Get document")
			fmt.Println(doc)
		}
		break
	case "search":
		if query == "" {
			log.Fatal("Query is required")
		} else if results, err := search.DbObj.SearchDocuments(query); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Résultats de la recherche pour '%s':\n", query)
			for _, doc := range results {
				fmt.Printf("Document ID: %d, Title: %s, Text: %s, Tags: %s\n", doc.Id, doc.Title, doc.Text, doc.Tags)
			}
		}
		break
	default:
		fmt.Println("Invalid action")
	}
}
