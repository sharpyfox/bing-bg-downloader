package http

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRetreiveSimple(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"images\":[{\"url\":\"/az/hprichbg/rb/BigHornSheep_EN-US6016496175_1920x1080.jpg\"}]}")
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	urls, err := RetreiveImagesUrls(server.URL, 1)

	assert.Nil(t, err)
	assert.Equal(t, []string{server.URL + "/az/hprichbg/rb/BigHornSheep_EN-US6016496175_1920x1080.jpg"}, urls)
}

func TestImgsNumChecking(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"images\":[]}")
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	_, err := RetreiveImagesUrls(server.URL, 2)
	assert.NotNil(t, err)
}

func TestImgsRecalls(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"images\":[{\"url\":\"/az/hprichbg/rb/BigHornSheep_EN-US6016496175_1920x1080.jpg\"},{\"url\":\"/az/hprichbg/rb/BigHornSheep_EN-US6016496175_1920x1080.jpg\"},{\"url\":\"/az/hprichbg/rb/BigHornSheep_EN-US6016496175_1920x1080.jpg\"},{\"url\":\"/az/hprichbg/rb/BigHornSheep_EN-US6016496175_1920x1080.jpg\"},{\"url\":\"/az/hprichbg/rb/BigHornSheep_EN-US6016496175_1920x1080.jpg\"}]}")
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	urls, err := RetreiveImagesUrls(server.URL, 10)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(urls))
}
