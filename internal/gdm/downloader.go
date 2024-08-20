package gdm

type downloadClient struct {
	fileName string
	url string
	totalSections int
	totalSize int
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