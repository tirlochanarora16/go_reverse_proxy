package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/tirlochanarora16/go_reverse_proxy/internal/lb"
)

func StartConfigFileWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Failed to create file watcher:", err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					log.Println("Starting file watcher...", ok)
					return
				}

				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("[CONFIG] Change detected in config file:", event.Name)

					ParseConfigFile()

				}

				if event.Has(fsnotify.Write) {
					log.Println("[WATCHER] Written to config.yml file", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}

				log.Println("[WATCHER ERROR]:", err)
			}
		}
	}()

	err = watcher.Add(fmt.Sprintf("./%s", lb.ConfigFileName))

	if err != nil {
		log.Fatal(err)
	}
}
