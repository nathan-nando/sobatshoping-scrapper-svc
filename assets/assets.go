package assets

import (
	"encoding/json"
	"github.com/nathan-nando/sobatshoping-scrapper-svc/domain"
	"github.com/nathan-nando/sobatshoping-scrapper-svc/internal/utils/constants"
	"log"
	"os"
)

type Asset struct {
	UserAgent []domain.UserAgent
}

func NewAsset() *Asset {
	asset := Asset{}

	userAgentJson, err := os.ReadFile(constants.JsonUserAgents)

	if err != nil {
		log.Fatal(constants.JsonUserAgents + "not found")
	}

	err = json.Unmarshal([]byte(userAgentJson), &asset.UserAgent)

	return &asset
}
