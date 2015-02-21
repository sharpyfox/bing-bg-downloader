package main

import (
	"flag"
	"fmt"
	"github.com/sharpyfox/bing-bg-downloader/http"
	"github.com/sharpyfox/bing-bg-downloader/path"
	"github.com/sharpyfox/bing-bg-downloader/utils"
	"log"
	"os"
)

func main() {
	bpPtr := flag.String("bp", "~/Pictures/wallpapers", "base path to save images")
	numOfImages := flag.Int("n", 5, "how many images to retreive")
	resPtr := flag.String("r", "1920x1080", "resolution to use")
	showVersion := flag.Bool("version", false, "print version string")
	flag.Parse()

	if *showVersion {
		fmt.Println(utils.Version("bing-bg-downloader"))
		return
	}

	pathToSvPic, err := path.PreProcessPath(*bpPtr)

	if nil != err {
		log.Panicln(err)
	}

	os.MkdirAll(pathToSvPic, 0755)

	urls, err := http.RetreiveImagesUrls("http://bing.com", *numOfImages)
	if nil != err {
		log.Panicln(err)
	}

	for i, u := range urls {
		urls[i], err = utils.PreProcessUrl(u, *resPtr)
		if nil != err {
			log.Panicln(err)
		}
	}

	d := http.CreateDownloader()
	d.Download(pathToSvPic, urls)
}
