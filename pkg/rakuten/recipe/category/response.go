package category

type LargeCategory struct {
	Id   int32  `json:"categoryId,string"`
	Name string `json:"categoryName"`
	Url  string `json:"categoryUrl"`
}

type MediumCategory struct {
	Id       int32  `json:"categoryId"`
	Name     string `json:"categoryName"`
	Url      string `json:"categoryUrl"`
	ParentId int    `json:"parentCategoryId,string"`
}

type SmallCategory struct {
	Id       int32  `json:"categoryId"`
	Name     string `json:"categoryName"`
	Url      string `json:"categoryUrl"`
	ParentId int    `json:"parentCategoryId,string"`
}

type Data struct {
	Result struct {
		Large  []LargeCategory  `json:"large"`
		Medium []MediumCategory `json:"medium"`
		Small  []SmallCategory  `json:"small"`
	} `json:"result"`
}
