package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (a *Addon) CheckAddonVersionMatch(ver string) bool {
	return a.LatestFiles[2].DisplayName == ver
}

func (a *Addon) LoadAddonInfo() error {

	resp, err := http.Get("https://addons-ecs.forgesvc.net/api/v2/addon/279257")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &a); err != nil {
		return err
	}
	return nil
}

type Addon struct {
	Attachments     []Attachment    `json:"attachments"`
	Authors         []Author        `json:"authors"`
	Categories      []Categorie     `json:"categories"`
	CategorySection CategorySection `json:"categorySection"`
	DateCreated     string          `json:"dateCreated"`
	DateModified    string          `json:"dateModified"`
	DateReleased    string          `json:"dateReleased"`
	DefaultFileID   int             `json:"defaultFileId"`
	//DownloadCount          int                     `json:"downloadCount"`
	GameID                 int                     `json:"gameId"`
	GameName               string                  `json:"gameName"`
	GamePopularityRank     int                     `json:"gamePopularityRank"`
	GameSlug               string                  `json:"gameSlug"`
	GameVersionLatestFiles []GameVersionLatestFile `json:"gameVersionLatestFiles"`
	ID                     int                     `json:"id"`
	IsAvailable            bool                    `json:"isAvailable"`
	IsExperiemental        bool                    `json:"isExperiemental"`
	IsFeatured             bool                    `json:"isFeatured"`
	LatestFiles            []LatestFile            `json:"latestFiles"`
	Name                   string                  `json:"name"`
	PopularityScore        float32                 `json:"popularityScore"`
	PortalName             string                  `json:"portalName"`
	PrimaryCategoryID      int                     `json:"primaryCategoryId"`
	PrimaryLanguage        string                  `json:"primaryLanguage"`
	Slug                   string                  `json:"slug"`
	Status                 int                     `json:"status"`
	Summary                string                  `json:"summary"`
	WebsiteUrl             string                  `json:"websiteUrl"`
}

type Attachment struct {
	ID           int    `json:"id"`
	ProjectID    int    `json:"projectId"`
	Description  string `json:"description"`
	IsDefault    bool   `json:"isDefault"`
	Status       int    `json:"status"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	Title        string `json:"title"`
	Url          string `json:"url"`
}

type Author struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	ProjectID         int    `json:"projectId"`
	ProjectTitleID    int    `json:"projectTitleId"`
	ProjectTitleTitle string `json:"projectTitleTitle"`
	TwitchID          int    `json:"twitchId"`
	Url               string `json:"url"`
	UserID            int    `json:"userId"`
}

type Categorie struct {
	AvatarID   int    `json:"avatarId"`
	AvatarUrl  string `json:"avatarUrl"`
	CategoryID int    `json:"categoryId"`
	GameID     int    `json:"gameId"`
	Name       string `json:"name"`
	ParentID   int    `json:"parentId"`
	ProjectID  int    `json:"projectId"`
	RootID     int    `json:"rootId"`
	Url        string `json:"url"`
}

type CategorySection struct {
	ExtraIncludePattern     string `json:"extraIncludePattern"`
	GameCategoryID          int    `json:"gameCategoryId"`
	GameID                  int    `json:"gameId"`
	ID                      int    `json:"id"`
	InitialInclusionPattern string `json:"initialInclusionPattern"`
	Name                    string `json:"name"`
	PackageType             int    `json:"packageType"`
	Path                    string `json:"path"`
}

type GameVersionLatestFile struct {
	FileType          int    `json:"fileType"`
	GameVersion       string `json:"gameVersion"`
	GameVersionFlavor string `json:"gameVersionFlavor"`
	ProjectFileID     int    `json:"projectFileId"`
	ProjectFileName   string `json:"projectFileName"`
}

type LatestFile struct {
	AlternateFileID            int    `json:"alternateFileId"`
	CategorySectionPackageType int    `json:"categorySectionPackageType"`
	DisplayName                string `json:"displayName"`
	DownloadUrl                string `json:"downloadUrl"`
	FileDate                   string `json:"fileDate"`
	FileLength                 int    `json:"fileLength"`
	FileName                   string `json:"fileName"`
	FileStatus                 int    `json:"fileStatus"`
	FileTypeID                 int    `json:"fileTypeId"`
	GameID                     int    `json:"gameId"`
	GameVersionDateReleased    string `json:"gameVersionDateReleased"`
	GameVersionFlavor          string `json:"gameVersionFlavor"`
	GameVersionID              int    `json:"gameVersionId"`
	GameVersionMappingID       int    `json:"gameVersionMappingId"`
	HasInstallScript           bool   `json:"hasInstallScript"`
	ID                         int    `json:"id"`
	IsAlternate                bool   `json:"isAlternate"`
	IsAvailable                bool   `json:"isAvailable"`
	IsCompatibleWithClient     bool   `json:"isCompatibleWithClient"`
	IsServerPack               bool   `json:"isServerPack"`
	PackageFingerprint         int    `json:"fileType"`
	PackageFingerprintID       int    `json:"packageFingerprintId"`
	ProjectId                  int    `json:"projectId"`
	ProjectStatus              int    `json:"projectStatus"`
	ReleaseType                int    `json:"releaseType"`
	RenderCacheId              int    `json:"renderCacheId"`
	RestrictProjectFileAccess  int    `json:"restrictProjectFileAccess"`
}
