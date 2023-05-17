package main

import (
	"embed"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/skanehira/clipboard-image/v2"
)

//go:embed images/*
var images embed.FS

func main() {
	rand.Seed(time.Now().UnixNano())

	imageFiles, err := images.ReadDir("images")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}

	f, err := os.Open(filepath.Join("images", imageFiles[rand.Intn(len(imageFiles))].Name()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}
	defer f.Close()

	if err := clipboard.Write(f); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}
	fmt.Printf("copy LGTM image to the clipboard. Execute Ctrl-C")
}
