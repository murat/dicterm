package api

// Meta is the metadata for the API.
type Meta struct {
	ID        string   `json:"id"`
	UUID      string   `json:"uuid"`
	Sort      string   `json:"sort"`
	Src       string   `json:"src"`
	Section   string   `json:"section"`
	Stems     []string `json:"stems"`
	Offensive bool     `json:"offensive"`
}

// Sound ...
type Sound struct {
	Audio string `json:"audio"`
	Ref   string `json:"ref"`
	Stat  string `json:"stat"`
}

// HeadwordInfo is the headword info.
type HeadwordInfo struct {
	Headword       string `json:"hw"`
	Pronunciations []struct {
		MerriemWebster string `json:"mw"`
		Sound          Sound  `json:"sound,omitempty"`
		AudioURL       string `json:"audio_url,omitempty"`
	}
}

// Collegiate is the Collegiate Dictionary API response.
type Collegiate struct {
	Meta            Meta                `json:"meta"`
	Headword        HeadwordInfo        `json:"hwi"`
	FunctionalLabel string              `json:"fl"`
	Inflections     []map[string]string `json:"ins"`
	Date            string              `json:"date"`
	Etymologies     [][]string          `json:"et"`
	Shortdef        []string            `json:"shortdef"`
}
