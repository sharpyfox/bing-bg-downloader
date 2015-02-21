package utils

import (
	"errors"
	"fmt"
	"strings"
)

func PreProcessUrl(url string, resolution string) (string, error) {
	defResolution := "1920x1080"
	if !strings.Contains(url, defResolution) {
		return "", errors.New(fmt.Sprintf("Seems that wrong url was provided: %s", url))
	} else {
		return strings.Replace(url, defResolution, resolution, 1), nil
	}
}
