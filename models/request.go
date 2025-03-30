package models

type RequestBody struct {
	Slug   string `json:"slug"`
	Render int    `json:"render"`
	UUID   string `json:"uuid,omitempty"`
	Skip   int    `json:"skip,omitempty"`
}
