package global

type siteJSON struct {
	ID       string
	Language *uint16
}

type domainJSON struct {
	ID       string
	Host     string
	Language *uint16
}

type storeObject struct {
	Name   string
	Fields []*storeObjectField
}

type languageJSON struct {
	ID                     uint16
	EnglishName            string
	AlternativeEnglishName *string
}

type storeObjectField struct {
	Name string
	Type string
}

type globalJSON struct {
	Domain  *domainJSON
	Site    *siteJSON
	Domains []*domainJSON

	Languages    []*languageJSON
	UserRights   map[uint16]string
	StoreObjects []*storeObject
}
