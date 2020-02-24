package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Image struct {
	title string
	body  io.ReadCloser
	err   error
}

func downloadAllImages(path, baseURL string, imageURLs []string) {
	images := make([]Image, 0)
	chImages, chFinished := make(chan Image), make(chan bool)
	for _, imageURL := range imageURLs {
		url := ""
		if strings.HasPrefix(imageURL, "http://") || strings.HasPrefix(imageURL, "https://") {
			url = imageURL
		} else {
			url = baseURL + imageURL
		}
		go func() {
			res, err := http.Get(url)
			if err != nil {
				return
			}
			parts := strings.Split(url, "/")
			img := Image{
				title: parts[len(parts)-1],
				body:  res.Body,
				err:   nil,
			}
			fmt.Println(img.title)
			chImages <- img
			chFinished <- true
		}()
	}
	for c := 0; c < len(imageURLs); {
		select {
		case img := <-chImages:
			if img.err != nil {
				panic(img.err)
			}
			images = append(images, img)
			defer img.body.Close()
		case <-chFinished:
			c++
		}
	}
	for _, img := range images {
		file, err := os.Create(path + "/" + img.title)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		_, err = io.Copy(file, img.body)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ./photos <URL>")
		return
	}
	baseURL := os.Args[1]
	res, err := http.Get(baseURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	path := strings.TrimPrefix(baseURL, "http://")
	path = strings.TrimPrefix(path, "https://")
	images := make([]string, 0)
	imgRegex := "<img.*?src=\"(.*?)\""
	r := regexp.MustCompile(imgRegex)
	matches := r.FindAllStringSubmatchIndex(string(body), -1)
	for i := len(matches) - 1; i >= 0; i-- {
		img := body[matches[i][2]:matches[i][3]]
		images = append(images, string(img))
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
	downloadAllImages(path, baseURL, images)
}
