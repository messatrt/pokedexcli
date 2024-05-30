package pokiapi

type LocationArea struct {
	Count int `json:"count"`
	//in go a url is best represented using a pointer
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
