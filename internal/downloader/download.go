package downloader

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
)

func DownloadDataset(url, output string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		path := dest + "/" + f.Name

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
			continue 
		}

		outFile, err := os.Create(path)
		if err != nil {
			return err
		}

		rc, _ := f.Open()
		io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
	}
	return nil
}