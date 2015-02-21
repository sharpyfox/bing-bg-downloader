package http

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type FilesDownloader struct {
	cl *http.Client
	wg *sync.WaitGroup
}

func (f *FilesDownloader) extractFileName(u string) (string, error) {
	fileURL, err := url.Parse(u)

	if err != nil {
		return "", err
	}

	segments := strings.Split(fileURL.Path, "/")
	return segments[len(segments)-1], nil
}

func (f *FilesDownloader) downloadFile(basePath string, url string) {
	fileName, err := f.extractFileName(url)

	if nil != err {
		log.Panicln(err)
	}

	log.Println("downloading " + fileName)

	fullPath := filepath.Join(basePath, fileName)
	if _, err = os.Stat(fullPath); os.IsNotExist(err) {
		out, err := os.Create(fullPath)

		if nil != err {
			log.Panicln(err)
		}

		defer out.Close()
		resp, err := f.cl.Get(url)

		if nil != err {
			log.Panicln(err)
		}
		defer resp.Body.Close()
		_, err = io.Copy(out, resp.Body)

		if nil != err {
			log.Panicln(err)
		}
	} else {
		log.Println(fileName + " is already downloaded previously")
	}
	f.wg.Done()
}

func CreateDownloader() FilesDownloader {
	cl := &http.Client{}
	wg := &sync.WaitGroup{}

	return FilesDownloader{cl, wg}
}

func (f *FilesDownloader) Download(basePath string, urls []string) {
	for _, u := range urls {
		f.wg.Add(1)
		go f.downloadFile(basePath, u)
	}

	f.wg.Wait()
}
