package geocodexyz

// ReverseLookupRequest is the request sent to the API requesting the json data
type ReverseLookupRequest struct {
	Locate       string `json:"locate"`
	Format       string `json:"format"`
	LanguageCode string `json:"language"`
	Region       string `json:"region"`
}

// toMap is needed for the requests. It seems that the parameter names actually change depending on
// POST or GET, so we have asome duplicated fields
func (input *ReverseLookupRequest) toMap() map[string]string {
	data := map[string]string{
		"locate": input.Locate,
		"geoit":  input.Format,
		"lang":   input.LanguageCode,
		"region": input.Region,
	}
	// there are 5 options
	if input.Format == "" {
		data["geoit"] = "json"
	}
	if input.Format == "json" {
		data["json"] = "1"
	}
	if input.Format == "" {
		data["lang"] = "en"
	}

	return data
}
