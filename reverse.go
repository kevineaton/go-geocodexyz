package geocodexyz

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// ReverseLookupRequest is the request sent to the API requesting the json data
type ReverseLookupRequest struct {
	LanguageCode string  `json:"language"`
	Region       string  `json:"region"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
}

// Note: these have been autogenerated and not been validated yet so may (will) change.

// ReverseLocationLookupAPIResult is the raw return for a location look up. Not all fields are always returned. Note the remote API converts everything
// to strings, including elevation, etc. So it's challenging to convert them correctly here. You may want to check their values and convert them
// by hand. Often, you will use the simpler version below.
type ReverseLocationLookupAPIResult struct {
	Statename string `json:"statename"`
	Distance  string `json:"distance"`
	Elevation string `json:"elevation"`
	Osmtags   struct {
		Wikipedia  string `json:"wikipedia"`
		Wikidata   string `json:"wikidata"`
		Place      string `json:"place"`
		Name       string `json:"name"`
		LandArea   string `json:"land_area"`
		AdminLevel string `json:"admin_level"`
		Boundary   string `json:"boundary"`
		Type       string `json:"type"`
	} `json:"osmtags"`
	State     string `json:"state"`
	Latt      string `json:"latt"`
	City      string `json:"city"`
	Prov      string `json:"prov"`
	Geocode   string `json:"geocode"`
	Geonumber string `json:"geonumber"`
	Country   string `json:"country"`
	Stnumber  string `json:"stnumber,omitempty"`
	Staddress string `json:"staddress"`
	Inlatt    string `json:"inlatt"`
	Alt       struct {
		Loc []struct {
			Staddress string `json:"staddress"`
			Stnumber  string `json:"stnumber"`
			Postal    string `json:"postal"`
			Latt      string `json:"latt"`
			City      string `json:"city"`
			Prov      string `json:"prov"`
			Longt     string `json:"longt"`
			Class     struct {
			} `json:"class"`
			Dist string `json:"dist"`
		} `json:"loc"`
	} `json:"alt"`
	Timezone         string `json:"timezone"`
	Region           string `json:"region"`
	Postal           string `json:"postal"`
	Longt            string `json:"longt"`
	RemainingCredits struct {
	} `json:"remaining_credits"`
	Confidence string `json:"confidence"`
	Inlongt    string `json:"inlongt"`
	Class      struct {
	} `json:"class"`
	Adminareas struct {
		Admin6 struct {
			Wikipedia        string `json:"wikipedia"`
			Wikidata         string `json:"wikidata"`
			Population       string `json:"population"`
			Name             string `json:"name"`
			SourcePopulation string `json:"source_population"`
			NistStateFips    string `json:"nist_state_fips"`
			AdminLevel       string `json:"admin_level"`
			NistFipsCode     string `json:"nist_fips_code"`
			Level            string `json:"level"`
			Boundary         string `json:"boundary"`
			Type             string `json:"type"`
			BorderType       string `json:"border_type"`
		} `json:"admin6"`
		Admin8 struct {
			Wikipedia  string `json:"wikipedia"`
			Wikidata   string `json:"wikidata"`
			Place      string `json:"place"`
			Name       string `json:"name"`
			LandArea   string `json:"land_area"`
			AdminLevel string `json:"admin_level"`
			Level      string `json:"level"`
			Boundary   string `json:"boundary"`
			Type       string `json:"type"`
		} `json:"admin8"`
	} `json:"adminareas"`
	Altgeocode string `json:"altgeocode"`
}

// ReverseLocationResult is the main result used by the look up and simplifies some of the values
type ReverseLocationResult struct {
	Lat        float64                         `json:"lat"`
	Lng        float64                         `json:"lng"`
	Street     string                          `json:"street"`
	City       string                          `json:"city"`
	State      string                          `json:"state"`
	Postal     string                          `json:"postal"`
	Country    string                          `json:"country"`
	Confidence float64                         `json:"confidence"`
	RawRequest *ReverseLocationLookupAPIResult `json:"raw"`
}

// toMap is needed for the requests. It seems that the parameter names actually change depending on
// POST or GET, so we have asome duplicated fields
func (input *ReverseLookupRequest) toMap() map[string]string {
	data := map[string]string{
		"locate": fmt.Sprintf("%v,%v", input.Lat, input.Lng),
		"geoit":  "json",
		"lang":   input.LanguageCode,
		"region": input.Region,
	}
	if input.LanguageCode == "" {
		data["lang"] = "en"
	}

	return data
}

// ReverseLookup performs a reverse lookup based upon the provided parameters. At a minimum, you should provide
// the Lat and Lng fields
func ReverseLookup(input *ReverseLookupRequest) (*ReverseLocationResult, *APIError, error) {
	result := &ReverseLocationResult{}
	resultBytes, apiError, err := makeCall(input.toMap())
	if apiError != nil || err != nil {
		return result, apiError, err
	}

	// we got the call back, so let's destructure
	apiResult := &ReverseLocationLookupAPIResult{}
	err = json.Unmarshal(resultBytes, apiResult)
	if err != nil {
		err = errors.New("could not find a suitable location")
		apiErr := &APIError{
			Code:        "008",
			Description: "could not find a suitable location or parse the response",
		}
		return result, apiErr, err
	}

	// now convert
	result.Lat = input.Lat
	result.Lng = input.Lng
	result.Street = fmt.Sprintf("%v %s", apiResult.Stnumber, apiResult.Staddress)
	result.City = apiResult.City
	result.State = apiResult.State
	result.Country = apiResult.Country
	result.Postal = apiResult.Postal
	result.Confidence, _ = strconv.ParseFloat(apiResult.Confidence, 64)
	result.RawRequest = apiResult

	return result, nil, nil
}
