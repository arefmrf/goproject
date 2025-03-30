package models

type STokenResponse struct {
	Status bool               `json:"status"`
	Data   STokenResponseData `json:"data"`
}

type STokenResponseData struct {
	Token string `json:"token"`
}

type MinimalResponse struct {
	Status bool `json:"status"`
	Data   struct {
		Structure []Structure `json:"structure"`
		TopBanner interface{} `json:"top_banner"`
	} `json:"data"`
}

type Structure struct {
	ID                  string       `json:"id"`
	SectionType         string       `json:"section_type"`
	UUID                string       `json:"uuid"`
	Items               []Item       `json:"items"`
	Sort                []SortOption `json:"sort"`
	Filter              []Filter     `json:"filter"`
	Pagination          Pagination   `json:"pagination"`
	SuggestedCategories []string     `json:"suggested_categories"`
}

type Item struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	Price      Price    `json:"price"`
	Image      Image    `json:"image"`
	Colors     []Colors `json:"colors"`
	IsFake     bool     `json:"is_fake"`
	Href       string   `json:"href"`
	InCampaign Campaign `json:"in_campaign"`
	Badge      string   `json:"badge"`
	Rel        string   `json:"rel"`
	Target     string   `json:"target"`
	State      State    `json:"state"`
}
type Colors struct {
	Code  string `json:"code"`
	ID    string `json:"id"`
	Title string `json:"title"`
}
type Price struct {
	Price           int    `json:"price"`
	DiscountedPrice int    `json:"discounted_price"`
	Discount        int    `json:"discount"`
	StartAt         string `json:"start_at"`
	EndAt           string `json:"end_at"`
}

// Image represents the image of an item
type Image struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}

// Campaign represents information related to the campaign of an item
type Campaign struct {
	CampaignName string `json:"campaign_name"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}

// State represents the rating state of an item
type State struct {
	RateCount int `json:"rate_count"`
	Rate      int `json:"rate"`
}

// SortOption represents sorting options
type SortOption struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	IsSelected bool   `json:"is_selected"`
}

// Filter represents filter options
type Filter struct {
	FilterID    string       `json:"filter_id"`
	FilterType  string       `json:"filter_type"`
	Title       string       `json:"title"`
	IconName    string       `json:"icon_name,omitempty"` // omitempty to ignore if empty
	IsEmphasize bool         `json:"is_emphasize,omitempty"`
	Items       []FilterItem `json:"items,omitempty"`
}

// FilterItem represents items within a filter
type FilterItem struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Code       string `json:"code"`
	IsSelected bool   `json:"is_selected"`
}

// Pagination represents pagination information
type Pagination struct {
	CurrentPage int `json:"current_page"`
	Total       int `json:"total"`
	TotalPages  int `json:"total_pages"`
	Count       int `json:"count"`
	PerPage     int `json:"per_page"`
}
