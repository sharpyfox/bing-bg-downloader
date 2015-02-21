package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlPreprocessing(t *testing.T) {
	url := "http://www.bing.com//az/hprichbg/rb/BigHornSheep_EN-US6016496175_1920x1080.jpg"
	expected := "http://www.bing.com//az/hprichbg/rb/BigHornSheep_EN-US6016496175_1024x768.jpg"

	p, err := PreProcessUrl(url, "1024x768")
	assert.Nil(t, err)
	assert.Equal(t, expected, p)
}

func TestUrlFormatChecking(t *testing.T) {
	url := "http://www.bing.com//az/hprichbg/rb/BigHornSheep_EN-US6016496175_1024x768.jpg"

	_, err := PreProcessUrl(url, "1024x768")
	assert.NotNil(t, err)
}
