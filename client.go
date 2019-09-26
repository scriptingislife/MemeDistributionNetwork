package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"golang.org/x/net/html"
)

// DownloadFileFromURL ...
func DownloadFileFromURL(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// OpenFile ...
func OpenFile(filepath string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", filepath).Start()
	case "windows":
		err = exec.Command("explorer.exe", filepath).Start()
	case "darwin":
		err = exec.Command("open", filepath).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err

}

func main() {

	host := "<serv-ip>:8000" // CHANGE ME
	dir := "popups"          // "popups" or "memes"

	// Seed our random number generator
	rand.Seed(time.Now().UnixNano())

	resp, _ := http.Get(fmt.Sprintf("http://%s/%s/", host, dir))

	tokenizer := html.NewTokenizer(resp.Body)

	files := make([]string, 0)

	// https://drstearns.github.io/tutorials/tokenizing/
	for {
		//get the next token type
		tokenType := tokenizer.Next()

		//if it's an error token, we either reached
		//the end of the file, or the HTML was malformed
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				//end of the file, break out of the loop
				break
			}
			//otherwise, there was an error tokenizing,
			//which likely means the HTML was malformed.
			//since this is a simple command-line utility,
			//we can just use log.Fatalf() to report the error
			//and exit the process with a non-zero status code
			log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
		}

		//process the token according to the token type...
		if tokenType == html.StartTagToken {
			t := tokenizer.Token()
			if t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key == "href" {
						files = append(files, a.Val)
						break
					} // if href
				} // for attr
			} // if anchor
		}
	}

	resp.Body.Close()

	filename := files[rand.Intn(len(files))]
	fileURL := fmt.Sprintf("http://%s/%s/%s", host, dir, filename)
	fmt.Println(fileURL)
	filepath := filepath.Join(os.TempDir(), filename)

	err := DownloadFileFromURL(fileURL, filepath)
	if err != nil {
		panic(err)
	}

	OpenFile(filepath)

}
