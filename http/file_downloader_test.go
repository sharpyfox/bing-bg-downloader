package http

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func Test(t *testing.T) {
	fPath := path.Join(os.TempDir(), "BigHornSheep_EN-US6016496175_1920x1080.jpg")
	os.Remove(fPath)

	_, err := os.Stat(fPath)
	assert.True(t, os.IsNotExist(err))
	defer os.Remove(fPath)

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "test")
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	d := CreateDownloader()
	d.Download(os.TempDir(), []string{server.URL + "/az/hprichbg/rb/BigHornSheep_EN-US6016496175_1920x1080.jpg"})

	_, err = os.Stat(fPath)
	assert.False(t, os.IsNotExist(err))
}
