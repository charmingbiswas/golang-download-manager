package gdm

import (
	"fmt"
	"net/http"
	"strconv"
)

type downloadClient struct {
	fileName string
	url string
	totalSections int
	totalSize int64
	fileType string
}


func NewDownloadClient() *downloadClient {
	return &downloadClient{
		url: URL,
		fileName: Filename,
		totalSections: 8, // hardcoding this value, just to be safe. This denotes the number of requests the app makes to the server to download the resource. Too many requests can make the server block our IP
		totalSize: 0, // will be updated later via API
		fileType: "", // will be updated later via API
	}
}

func (d *downloadClient) StartDownload() error {
	// fetch metadata
	err := d.fetchMetaData()
	if err != nil {
		return err
	}
	fmt.Println(d)
	return nil
}

func (d *downloadClient) fetchMetaData() error {
	req, err := http.NewRequest("HEAD", d.url, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	
	if res.StatusCode > 299 {
		return fmt.Errorf("invalid response from server")
	}
	
	fileSize, err := strconv.Atoi(res.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

	contentType := res.Header.Get("Content-Type")
	fileType, ok := MimeTypes[contentType]
	if !ok {
		return fmt.Errorf("file type currently not supported")
	}
	
	d.fileType = fileType
	d.totalSize = int64(fileSize)
	
	return nil
}