package main

import (
	"github.com/charmingbiswas/golang-download-manager/internal/gdm"
)

func main() {
	gdm.InitFlags()
	gdm.NewDownloadClient()
}