package gdm

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type downloadClient struct {
	fileName string
	url string
	totalSections int
	totalSize int64
	fileType string
	fileSizeForEachSection int64
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

	// start building number of sections to download the file in
	fileSizeForEachSection := d.totalSize/int64(d.totalSections)
	sections := make([][2]int64, d.totalSections)
	d.fileSizeForEachSection = fileSizeForEachSection
	// for initialize starting byte and ending byte for each section
	for i := range sections {
		// starting bytes of each section
		if i == 0 {
			// first section, starting byte
			sections[i][0] = 0
		} else {
			// other sections, starting byte
			sections[i][0] = sections[i - 1][1] + 1
		}

		// ending bytes of each section
		if i < d.totalSections - 1 {
			// every section, other than last section
			sections[i][1] = sections[i][0] + (fileSizeForEachSection)
		} else {
			// ending byte of last section
			sections[i][1] = d.totalSize
		}
	}

	var wg sync.WaitGroup
	for i, s := range sections {
		wg.Add(1)
		go func(i int, s [2]int64) {
			defer wg.Done()
			err = d.downloadSection(i, s)
			if err != nil {
				panic(err)
			}
		}(i, s)
	}
	wg.Wait()
	err = d.mergeFiles(sections)
	if err != nil {
		return err
	}
	return nil
}

func (d *downloadClient) fetchMetaData() error {
	req, err := http.NewRequest("HEAD", d.url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "gdm")
	req.Header.Set("Accept-Language", "en-GB,en;q=0.9,en-US;q=0.8,pt;q=0.7")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	
	if res.StatusCode > 299 {
		fmt.Println(res.StatusCode)
		return fmt.Errorf("invalid response from server")
	}
	fmt.Println(res)
	fileSize, err := strconv.Atoi(res.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

	contentType := res.Header.Get("Content-Type")
	fileType, ok := MimeTypes[contentType]
	if !ok {
		fileType = ".mp4" // hardcoded for now, later extract type from url if possible
		// return fmt.Errorf("file type currently not supported: %s", contentType)
	}
	
	d.fileType = fileType
	d.totalSize = int64(fileSize)
	
	return nil
}

func (d *downloadClient) downloadSection(i int, s[2]int64) error {
	req, err := http.NewRequest("GET", d.url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", s[0], s[1]))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("section-%d.tmp", i))
	if err != nil {
		return err
	}
	defer file.Close()
	buffer := make([]byte, 1024 * 1024) // buffer size
	fmt.Printf("Downloading bytes for section-%d\n", i)
	for {
		bytesRead, err := res.Body.Read(buffer)
		if bytesRead > 0 {
			_, err = file.Write(buffer[:bytesRead])
			if err != nil {
				return err
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
	}


	return nil
}

func (d *downloadClient) mergeFiles(sections [][2]int64) error {
	file, err := os.OpenFile(fmt.Sprintf("%s%s", d.fileName, d.fileType), os.O_CREATE | os.O_WRONLY |os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	for i := range sections {
		tempFileName := fmt.Sprintf("section-%d.tmp", i)
		bytesRead, err := os.ReadFile(tempFileName)
		if err != nil {
			return err
		}
		bytesWritten, err := file.Write(bytesRead)
		if err != nil {
			return err
		}
		err = os.Remove(tempFileName)
		if err != nil {
			return err
		}

		fmt.Printf("%d bytes merged\n", bytesWritten)
	}
	fmt.Printf("%s%s ready for use\n", d.fileName, d.fileType)
	return nil
}
