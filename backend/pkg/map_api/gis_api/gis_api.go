package gisapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	GeoType = "geo"
	BizType = "biz"
)

type Config struct {
	Key    string
	Radius int
	Sort   string
}

type gisApi struct {
	cfg Config
}

func New(cfg Config) *gisApi {
	return &gisApi{
		cfg: cfg,
	}
}

func (a *gisApi) GetCompany(categories map[string]bool, lat, lon float64) (map[string][]string, error) {
	categoryCompanies := make(map[string][]string)
	for category, _ := range categories {
		url := fmt.Sprintf(`https://catalog.api.2gis.com/3.0/items?q=%s&point=%f%%2C%f&radius=%d&sort_point=%f%%2C%f&sort=%s&key=%s`,
			category, lat, lon, a.cfg.Radius, lat, lon, a.cfg.Sort, a.cfg.Key)
		response, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		jsonBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		defer func() {
			if err := response.Body.Close(); err != nil {
				logrus.Error(err)
			}
		}()
		var сompanyList CompanyList
		if err := json.Unmarshal(jsonBytes, &сompanyList); err != nil {
			return nil, err
		}
		categoryCompanies[category] = a.comapanyListToCompany(&сompanyList)
	}
	return categoryCompanies, nil
}

func (a *gisApi) comapanyListToCompany(сompanyList *CompanyList) []string {
	companies := make([]string, len(сompanyList.Result.Items))
	for i, item := range сompanyList.Result.Items {
		companies[i] = item.Name
	}
	return companies
}
