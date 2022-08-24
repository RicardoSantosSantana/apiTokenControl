package apiTokenControl

type SalesTerms_valueStruct struct {
	Number int32  `json:"number"`
	Unit   string `json:"unit"`
}

type SalesTerms_values struct {
	Id     string                 `json:"id"`
	Name   string                 `json:"name"`
	Struct SalesTerms_valueStruct `json:"value_struct"`
}

type SalesTerms struct {
	Id           string                 `json:"id"`
	Name         string                 `json:"name"`
	Value_id     string                 `json:"value_id"`
	Value_name   string                 `json:"value_name"`
	Value_struct SalesTerms_valueStruct `json:"value_struct"`
	Values       []SalesTerms_values    `json:"values"`
}

type Pictures struct {
	Id         string `json:"id"`
	Url        string `json:"url"`
	Secure_url string `json:"secure_url"`
	Size       string `json:"size"`
	Max_size   string `json:"max_size"`
	Quality    string `json:"quality"`
}
type Description struct {
	Text         string `json:"text"`
	Plain_text   string `json:"plain_text"`
	Last_updated string `json:"last_updated"`
	Date_created string `json:"date_created"`
}

type ItemsIds struct {
	Results []string `json:"results"`
}

type Token struct {
	Access_token  string `json:"access_token"`
	Token_type    string `json:"token_type"`
	Expires_in    int    `json:"expires_in"`
	Scope         string `json:"scope"`
	User_id       int    `json:"user_id"`
	Refresh_token string `json:"refresh_token"`
}

type SInitialParams struct {
	Client_id     string `json:"client_id"`
	Client_secret string `json:"client_secret"`
	Code          string `json:"code"`
	Redirect_uri  string `json:"redirect_uri"`
}

type Errors struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Status  int    `json:"status"`
}

type Items struct {
	Id                 string       `json:"id"`
	Site_id            string       `json:"site_id"`
	Title              string       `json:"title"`
	Subtitle           string       `json:"subtitle"`
	Seller_id          int32        `json:"seller_id"`
	Category_id        string       `json:"category_id"`
	Official_store_id  string       `json:"official_store_id"`
	Price              float32      `json:"price"`
	Base_price         float32      `json:"base_price"`
	Original_price     float32      `json:"original_price"`
	Currency_id        string       `json:"currency_id"`
	Initial_quantity   int32        `json:"initial_quantity"`
	Available_quantity int32        `json:"available_quantity"`
	Sold_quantity      int32        `json:"sold_quantity"`
	Sale_terms         []SalesTerms `json:"sale_terms"`
	Buying_mode        string       `json:"buying_mode"`
	Listing_type_id    string       `json:"listing_type_id"`
	Start_time         string       `json:"start_time"`
	Stop_time          string       `json:"stop_time"`
	Condition          string       `json:"condition"`
	Permalink          string       `json:"permalink"`
	Thumbnail_id       string       `json:"thumbnail_id"`
	Thumbnail          string       `json:"thumbnail"`
	Secure_thumbnail   string       `json:"secure_thumbnail"`
	Status             string       `json:"status"`
	Warranty           string       `json:"warranty"`
	Catalog_product_id string       `json:"catalog_product_id"`
	Domain_id          string       `json:"domain_id"`
	Health             float32      `json:"health"`
	Pictures           []Pictures   `json:"pictures"`
	Description        Description  `json:"description"`
}
