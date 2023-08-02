package main

import (
	"encoding/json"
)

type Storage struct {
	binariesSummary         BinariesSummary  `json::"binariesSummary, omitempty"`
	fileStoreSummary        FileStoreSummary `json:"fileStoreSummary, omitempty"`
	repositoriesSummaryList []Repositories   `json:"repositoriesSummaryList, omitempty"`
}

type BinariesSummary struct {
	binariesCount  string `json:"binariesCount"`
	binariesSize   string `json:"binariesSize, omitempty"`
	artifactsSize  string `json:"artifactsSize, omitempty"`
	optimization   string `json:"optimization, omitempty"`
	itemssCount    string `json:"itemsCount, omitempty"`
	artifactsCount string `json:"artifactsCount, omitempty"`
}

type FileStoreSummary struct {
	storageType      string `json:"storageType"`
	storageDirectory string `json:"storageDirectory, omitempty"`
	totalSpace       string `json:"totalSpace, omitempty"`
	usedSpace        string `json:"usedSpace, omitempty"`
	freeSpace        string `json:"freeSpace, omitempty"`
}

type Repositories struct {
	repoKey          string `json:"repoKey"`
	repoType         string `json:"repoType"`
	foldersCount     int    `json:"foldersCount, omitempty"`
	filesCount       int    `json:"filesCount, omitempty"`
	usedSpace        string `json:"usedSpace, omitempty"`
	usedSpaceInBytes int    `json:"usedSpaceInBytes, omitempty"`
	itemsCount       int    `json:"itemsCount, omitempty"`
	packageType      string `json:"packageType, omitempty"`
	projectKey       string `json:"projectKey, omitempty"`
	percentage       string `json:"percentage, omitempty"`
}

func New(storageJson string) Storage {
	var storage Storage
	json.Unmarshal([]byte(storageJson), &storage)
	return storage
}
