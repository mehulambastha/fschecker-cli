package internal

import (
	"fmt"
	"log"
	"slices"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

func IsMediaFile(fileName string) bool {
	splitted := strings.Split(fileName, ".")
	extension := strings.TrimSpace(strings.ToLower(splitted[len(splitted)-1]))

	mediaExtensions := []string{
		"jpg", "jpeg", "png", "gif", "bmp", "tiff", "webp", // Image formats
		"mp3", "wav", "flac", "aac", "ogg", "m4a", // Audio formats
		"mp4", "avi", "mkv", "mov", "wmv", "webm", "flv", // Video formats
		"pdf", "epub", "mobi", "azw3", "doc", "docx", // Document formats
		"xls", "xlsx", "ppt", "pptx", "odt", "ods", "odp", // Office formats
		"txt", "md", "rtf", "log", // Text formats
		"zip", "rar", "7z", "tar", "gz", "bz2", "xz", // Compressed formats
		"iso", "img", "bin", // Disk image formats
		"svg", "psd", "ai", "eps", "pdf", "xcf", // Design/vector formats
	}

	if slices.Contains(mediaExtensions, extension) {
		return true
	}
	return false
}

func WatchDir(dir string) {
	watcher, err := fsnotify.NewWatcher()
	defer watcher.Close()

	debouncingTimerState := make(map[string]time.Time)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					debouncingTimerState[event.Name] = time.Now()

					go func(fileName string) {
						time.Sleep(150 * time.Millisecond)
						if time.Since(debouncingTimerState[fileName]) >= 150*time.Millisecond {
							addAndCommit(fileName)
						}
					}(event.Name)

				} else if event.Has(fsnotify.Create) {

					// Checking if this is a media file, and that if it needs to be added as soon as it is created.
					if IsMediaFile(event.Name) {
						addAndCommit(event.Name)
					} else {
						continue
					}
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

func addAndCommit(fileName string) {
	fmt.Println("adding and commtting..", fileName)
}
