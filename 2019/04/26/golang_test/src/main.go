package main

import (
	"archive/zip"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	tempDir := filepath.Join("./ttt", "temp")
	fmt.Println(tempDir)

	err := os.Rename("./test", "./test2")
	if err != nil {
		fmt.Println(err)
	}
	r, err := zip.OpenReader("./test.zip")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range r.File {
		rc, err := file.Open()
		if err != nil {
			//return err
			fmt.Println(err)
		}
		defer rc.Close()
		fmt.Println(file.FileInfo)
	}
}
