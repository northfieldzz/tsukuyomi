package ranking

type Recipe struct {
	Id             int      `json:"recipeId"`
	Title          string   `json:"recipeTitle"`
	Description    string   `json:"recipeDescription"`
	Url            string   `json:"recipeUrl"`
	Cost           string   `json:"recipeCost"`
	Indication     string   `json:"recipeIndication"`
	Material       []string `json:"recipeMaterial"`
	PublishDay     string   `json:"recipePublishDay"`
	NickName       string   `json:"nickname"`
	Rank           int      `json:"rank,string"`
	PickUp         int      `json:"pickup"`
	Shop           int      `json:"shop"`
	FoodImageUrl   string   `json:"foodImageUrl"`
	SmallImageUrl  string   `json:"smallImageUrl"`
	MediumImageUrl string   `json:"mediumImageUrl"`
}

type Data struct {
	Results []Recipe `json:"result"`
}
