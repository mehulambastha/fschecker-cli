package internal

import (
	_ "fmt"
	"github.com/fsnotify/fsnotify"
	"log"
)

func WatchDir(dir string) {
	watcher, err := fsnotify.NewWatcher()
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event: ", event)

				if event.Has(fsnotify.Write) {
					log.Println("A file was modified at: ", event.Name)
				} else if event.Has(fsnotify.Create) {
					log.Println("A new file was created: ", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("An error occured: ", err)
			}
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}
