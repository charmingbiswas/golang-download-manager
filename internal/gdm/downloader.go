package gdm

type download struct {
	filename string
	url string
	totalSections int // number of requests to make to the server to downloads sections of the file. Generally will be 8, more than this, server may rate limit us
}

func(d *download) NewDownloader(url string, filename string) *download {
	return &download{url: url, filename: filename}
}

