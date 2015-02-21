package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Image struct {
	Url string `json:"url"`
}

type BingResponse struct {
	Images []Image `json:"images"`
}

func retreive(baseUrl string, numOfImages int, offset int) (string, error) {
	url := fmt.Sprintf("%s/HPImageArchive.aspx?format=js&idx=%d&n=%d&mkt=en-US", baseUrl, offset, numOfImages)
	resp, err := http.Get(url)
	if nil != err {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return "", err
	}

	return string(body), nil
}

func RetreiveImagesUrls(baseUrl string, numOfImages int) ([]string, error) {
	cntr := 0
	offset := 0
	result := make([]string, numOfImages)

	for {
		str, _ := retreive(baseUrl, 5, 5*offset)

		x := new(BingResponse)
		if err := json.NewDecoder(strings.NewReader(str)).Decode(x); err != nil {
			return []string{}, err
		}

		if 0 == len(x.Images) {
			return []string{}, errors.New(fmt.Sprintf("Wrong images count was received %d of %d", cntr, numOfImages))
		}

		for _, el := range x.Images {
			result[cntr] = baseUrl + el.Url
			cntr++

			if cntr == numOfImages {
				return result, nil
			}
		}
		offset++
	}
}
