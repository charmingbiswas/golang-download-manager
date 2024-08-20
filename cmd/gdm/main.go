package main

import (
	"github.com/charmingbiswas/golang-download-manager/internal/gdm"
)

func main() {
	gdm.InitFlags()

	if(gdm.URL != "") {
		panic("Please provide a valid url for download")
	}
}