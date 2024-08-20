package main

import (
	"log"

	"github.com/charmingbiswas/golang-download-manager/internal/gdm"
)

func main() {
	gdm.InitApp()
	dc := gdm.NewDownloadClient()
	err := dc.StartDownload()
	if err != nil {
		log.Fatal(err)
	}
}