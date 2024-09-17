package gdm

import (
	"flag"
	"fmt"
	"log"
)

var (
	URL string
	Filename string
	MimeTypes map[string]string
)

func initFlags() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Please use the following flags\n")
		flag.PrintDefaults()
	}
	flag.StringVar(&URL, "url", "", "url to the resource you want to download")
	flag.StringVar(&URL, "u", "", "url to the resource you want to download")
	flag.StringVar(&Filename, "filename", "video", "filename to to save as")
	flag.StringVar(&Filename, "f", "video", "filename to to save as")
	flag.Parse()

	// check if url to the download resource is provided or not
	if URL == "" {
		log.Fatal("Please provide a valid url for download")
	}
}

func initMimeTypes() {
	MimeTypes = map[string]string{
		"audio/aac": ".aac",
		"audio/wav": ".wav",
		"image/avif": ".avif",
		"image/gif": ".gif",
		"image/jpeg": ".jpeg",
		"image/png": ".png",
		"video/x-msvideo": ".avi",
		"video/mp4": ".mp4",
		"video/mpeg": ".mpeg",
		"video/mp2t": ".ts",
		"application/vnd.ms-powerpoint": ".ppt",
		"application/vnd.ms-excel": ".xls",
		"application/zip": ".zip",
	}
}


func InitApp() {
	initFlags()
	initMimeTypes()
}