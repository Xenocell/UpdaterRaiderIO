//go:generate goversioninfo -icon=resource/Icon_FullColor.ico -manifest=resource/goversioninfo.exe.manifest
package main

import (
	"UpdateRaiderIO/models"
	"fmt"
	"log"
)

func main() {
	var game *models.Game = &models.Game{}
	var addon *models.Addon = &models.Addon{}

	if err := game.SearchGame(); err != nil {
		log.Println("Game not found!")
		pause()
		return
	}
	log.Printf("Game found - %s", game.GetPath())

	sa, err := game.SearchAddon()
	if err != nil {
		log.Println("RaiderIO not found in the game!")
	}
	if sa {
		log.Printf("RaiderIO found in the game - %s", game.GetCurrentVersionLocalAddon())
	}

	log.Println("Downloading the latest information about RaiderIO!")
	if err := addon.LoadAddonInfo(); err != nil {
		log.Println("Failed to download information about the latest version of RaiderIO!")
		pause()
		return
	}
	log.Printf("Information about the latest version of RaiderIO uploaded successfully - %s", addon.LatestFiles[2].DisplayName)
	if sa {
		if status := addon.CheckAddonVersionMatch(game.GetCurrentVersionLocalAddon()); status {
			log.Println("Your RaiderIO does not require an update!")
			pause()
			return
		}
	}
	log.Println("Starting the download.")
	if err := game.DownloadAddon(addon); err != nil {
		log.Println("Failed to load RaiderIO!")
		pause()
		return
	}
	log.Println("Starting to unzip RadarIO")
	if err := game.InstallAddon(addon); err != nil {
		log.Println("Failed to unzip RaiderIO!")
		pause()
		return
	}
	log.Println("RaiderIO update completed successfully!")

	pause()
}

func pause() {
	fmt.Println("\nPress enter to close the console.")
	fmt.Scanln()
}
