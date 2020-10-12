package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/blugelabs/bluge"
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatal("must specify src path")
	}

	if flag.NArg() < 2 {
		log.Fatal("must specify dest path")
	}

	cfg := bluge.DefaultConfig(flag.Arg(1))
	idx, err := bluge.OpenOfflineWriter(cfg, 100, 10)
	if err != nil {
		log.Fatalf("error opening index writer: %v", err)
	}

	err = walkDirectoryForIndexing(flag.Arg(0), idx)
	if err != nil {
		log.Fatal(err)
	}

	err = idx.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func walkDirectoryForIndexing(path string, idx *bluge.OfflineWriter) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil && filepath.Ext(path) == ".json" {
			pageDoc, err2 := readParseMapPage(path)
			if err2 != nil {
				return err2
			}
			return idx.Insert(pageDoc)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

type Page struct {
	Title string `json:"title"`
	Date string `json:"date"`
	Type string `json:"type"`
	PermaLink string `json:"permalink"`
	Content string `json:"content"`
}

func readParseMapPage(path string) (*bluge.Document, error) {
	pageBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var page Page
	err = json.Unmarshal(pageBytes, &page)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling json '%s': %v", path, err)
	}

	doc := bluge.NewDocument(page.PermaLink).
		AddField(bluge.NewTextField("title", page.Title).StoreValue()).
		AddField(bluge.NewKeywordField("type", page.Type).StoreValue().Aggregatable()).
		AddField(bluge.NewTextField("content", html.UnescapeString(page.Content)).StoreValue().HighlightMatches()).
		AddField(bluge.NewCompositeFieldExcluding("_all", []string{"_id"}))

	pageDate, err := time.Parse(time.RFC3339, page.Date)
	if err == nil && !pageDate.IsZero() {
		doc.AddField(bluge.NewDateTimeField("updated", pageDate))
	} else {
		fmt.Printf("date: %v err: %v\n", pageDate, err)
	}

	return doc, nil
}


