package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"
	"xkcd-go/comic"
)

const comicTemplate = `
===================================================
Title: {{.Title}}
URL: {{.Num | getURL}}
Image: {{.Img}}
Alt: {{.Alt}}

{{.Transcript}}
===================================================
`

var id int

func init() {
	defer flag.Parse()
	flag.IntVar(&id, "id", 0, "Id of the comic to be read")
}

func renderOut(c *comic.Comic) {
	out := template.Must(template.New("comic").
		Funcs(template.FuncMap{"getURL": comic.GetURL}).
		Parse(comicTemplate))
	out.Execute(os.Stdout, c)
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	if isFlagPassed("id") {
		if currComic, err := comic.GetComic(id); err != nil {
			log.Fatal(err)
		} else {
			renderOut(currComic)
		}
	} else {
		input := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("=id> ")
			cmd, err := input.ReadString('\n')
			if err != nil {
				log.Fatalf("IO error: %v\n", err)
			}
			cmd = cmd[:len(cmd)-1]
			if cmd == "" {
				continue
			}
			if cmd == "q" || cmd == "exit" || cmd == "quit" || cmd == ":q" {
				os.Exit(0)
			}
			id, err := strconv.Atoi(cmd)
			if err != nil {
				fmt.Println("=err:> Id can only be a integer")
				continue
			}
			if currComic, err := comic.GetComic(id); err != nil {
				fmt.Printf("=err (calling GetComic):> %v\n", err)
			} else {
				renderOut(currComic)
			}
		}
	}
}
