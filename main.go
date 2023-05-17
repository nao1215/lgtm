package main

import (
	"bytes"
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

	f, err := images.ReadFile(filepath.Join("images", imageFiles[rand.Intn(len(imageFiles))].Name()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}

	if err := clipboard.Write(bytes.NewReader(f)); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}
	fmt.Printf("copy LGTM image to the clipboard. Execute Ctrl-C")
}
