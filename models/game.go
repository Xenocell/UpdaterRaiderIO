package models

import (
	"UpdateRaiderIO/helper"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"
	"golang.org/x/sys/windows/registry"
)

type Game struct {
	Path                     string
	CurrentVersionLocalAddon string
}

func New() *Game {
	return &Game{}
}

func (g *Game) SearchGame() error {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Wow6432Node\Blizzard Entertainment\World of Warcraft`, registry.QUERY_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	path, _, err := key.GetStringValue("InstallPath")
	if err != nil {
		return err
	}
	if len(path) == 0 {
		return err
	}
	g.Path = path
	return nil
}

func (g *Game) SearchAddon() (bool, error) {
	data, err := ioutil.ReadFile(g.GetAddonPath() + "\\RaiderIO\\RaiderIO.toc")
	if err != nil {
		return false, err
	}

	for _, line := range strings.Split(string(data), "\n") {
		if strings.Contains(line, "Version") {
			firstSymbol := strings.Index(line, "(")
			lastSymbol := strings.Index(line, ")")
			g.CurrentVersionLocalAddon = line[firstSymbol+1 : lastSymbol]
			break
		}
	}

	return true, nil
}

func (g *Game) DownloadAddon(addon *Addon) error {
	resp, err := http.Get(addon.LatestFiles[2].DownloadUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.OpenFile(g.GetAddonPath()+"\\"+addon.LatestFiles[2].FileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"Downloading",
	)
	io.Copy(io.MultiWriter(file, bar), resp.Body)
	return nil
}

func (g *Game) InstallAddon(addon *Addon) error {
	if err := helper.Unzip(g.GetAddonPath()+"\\"+addon.LatestFiles[2].FileName, g.GetAddonPath()); err != nil {
		return err
	}
	if err := os.Remove(g.GetAddonPath() + "\\" + addon.LatestFiles[2].FileName); err != nil {
		log.Println("Failed to delete temporary files!")
	}
	return nil
}

func (g *Game) GetCurrentVersionLocalAddon() string {
	return g.CurrentVersionLocalAddon
}

func (g *Game) GetAddonPath() string {
	return g.Path + "Interface\\AddOns"
}

func (g *Game) GetPath() string {
	return g.Path
}
