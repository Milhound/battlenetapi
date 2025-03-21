package gamedata

import "fmt"

const (
	ConnectedRealmsIndexEndpoint = "/data/wow/connected-realm/index"
	ConnectedRealmEndpoint       = "/data/wow/connected-realm/%d"
	ConnectedRealmSearchEndpoint = "/data/wow/search/connected-realm"
)

type Realm struct {
	ID     int       `json:"id"`
	Name   localized `json:"name"`
	Region struct {
		Name localized `json:"name"`
		ID   int       `json:"id"`
	} `json:"region"`
	Category localized `json:"category"`
	Locale   string    `json:"locale"`
	Type     struct {
		Name localized `json:"name"`
		Type string    `json:"type"`
	} `json:"type"`
	Slug string `json:"slug"`

	Timezone   string `json:"timezone"`
	Tournament bool   `json:"is_tournament"`
}

type RealmSearchResult struct {
	Key  href `json:"key"`
	Data struct {
		Realms []Realm `json:"realms"`
		ID     int     `json:"id"`
		Queue  bool    `json:"has_queue"`
		Status struct {
			Name localized `json:"name"`
			Type string    `json:"type"`
		} `json:"status"`
		Population struct {
			Name localized `json:"name"`
			Type string    `json:"type"`
		} `json:"population"`
	} `json:"data"`
}

type RealmStatusParams struct {
	Status   string
	Timezone string
	OrderBy  string
	Page     int
}

type URLFormatterImpl struct{}

func (u URLFormatterImpl) FormatURL(baseURL, endpoint, namespace, region string, options interface{}) string {
	o := options.(RealmStatusParams)
	requestURL := fmt.Sprintf("%s%s?namespace=%s-%s", baseURL, endpoint, namespace, region)
	requestURL = fmt.Sprintf("%s&status.type=%s&realms.timezone=%s&orderby=%s&_page=%d", requestURL, o.Status, o.Timezone, o.OrderBy, o.Page)
	return requestURL
}

type localized struct {
	IT string `json:"it_IT"`
	RU string `json:"ru_RU"`
	GB string `json:"en_GB"`
	TW string `json:"zh_TW"`
	KR string `json:"ko_KR"`
	US string `json:"en_US"`
	MX string `json:"es_MX"`
	BR string `json:"pt_BR"`
	ES string `json:"es_ES"`
	CN string `json:"zh_CN"`
	FR string `json:"fr_FR"`
	DE string `json:"de_DE"`
}

type ConnectedRealmSearchAPI struct {
	Page        int                 `json:"page"`
	PageSize    int                 `json:"pageSize"`
	MaxPageSize int                 `json:"maxPageSize"`
	PageCount   int                 `json:"pageCount"`
	Results     []RealmSearchResult `json:"results"`
}
