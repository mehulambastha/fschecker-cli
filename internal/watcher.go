package internal

import (
	"fmt"
	_ "github.com/fsnotify/fsnotify"
	_ "log"
)

func WatchDir(dir string) {
	fmt.Println("In internal...")
}
