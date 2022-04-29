package api

// Response is the response from the dictionary API.
type Response []struct {
	Meta     Meta     `json:"meta"`
	FL       string   `json:"fl"`
	Shortdef []string `json:"shortdef"`
}

// Meta API response
type Meta struct {
	ID    string   `json:"id"`
	Stems []string `json:"stems"`
}
