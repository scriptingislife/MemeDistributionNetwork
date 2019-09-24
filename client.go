package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
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
	ip := "<ip>:<port>" // CHANGE ME
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(32)
	filename := fmt.Sprintf("%d.jpg", num)
	fileURL := fmt.Sprintf("http://%s/memes/%s", ip, filename)
	fmt.Println(fileURL)
	filepath := filepath.Join(os.TempDir(), filename)

	err := DownloadFileFromURL(fileURL, filepath)
	if err != nil {
		panic(err)
	}

	OpenFile(filepath)

}
