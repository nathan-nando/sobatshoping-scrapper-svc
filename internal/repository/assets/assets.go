package assets

import "github.com/nathan-nando/sobatshoping-scrapper-svc/domain"

type AssetRepository struct {
	UserAgents []domain.UserAgent
	UrlScraps  []domain.UrlScrap
}
