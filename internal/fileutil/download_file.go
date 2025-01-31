package fileutil

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadFile downlads the file from the url and saves it to the
// filepath
func DownloadFile(url string, filePath string) (err error) {
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
