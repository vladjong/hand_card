package mapapi

type IMapApi interface {
	GetCompany(categories map[string]bool, lat, lon float64) (map[string][]string, error)
}
