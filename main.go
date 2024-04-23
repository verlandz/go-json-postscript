package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/verlandz/go-json-postscript/postscripts"
)

const (
	INPUT_FILE_PATH    = "files/input.json"
	OUTPUT_FILE_PATH   = "files/output.json"
	OUTPUT_FILE_INDENT = "    "
	OUTPUT_FILE_PERM   = 0644
)

func main() {
	// create watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// watch input file
	err = watcher.Add(INPUT_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}

	// listening for events
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					generateOutput()
					log.Printf("(src:%v) Success!\n", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				panic(err)
			}
		}
	}()

	// blocks
	for {
	}
}

func generateOutput() {
	// read input
	input, err := os.ReadFile(INPUT_FILE_PATH)
	if err != nil {
		panic(err)
	}

	// unmarshal data
	data := map[string]interface{}{}
	if err := json.Unmarshal(input, &data); err != nil {
		panic(err)
	}

	// postscripts
	postscripts.PostScripts(data)

	// write output
	output, err := json.MarshalIndent(data, "", OUTPUT_FILE_INDENT)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(OUTPUT_FILE_PATH, output, OUTPUT_FILE_PERM); err != nil {
		panic(err)
	}
}
