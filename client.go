package geocodexyz

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-resty/resty/v2"
)

var httpClient *resty.Client

// LocationLookupResult is a successful lookup for a location passed in, such as "123 Main Street, Salem, NH"
type LocationLookupResult struct {
	Standard struct {
		Stnumber string `json:"stnumber"`
		Addresst string `json:"addresst"`
		Postal   struct {
		} `json:"postal"`
		Region string `json:"region"`
		Zip    struct {
		} `json:"zip"`
		Prov        string `json:"prov"`
		City        string `json:"city"`
		Countryname string `json:"countryname"`
		Confidence  string `json:"confidence"`
	} `json:"standard"`
	Longt string `json:"longt"`
	Alt   struct {
	} `json:"alt"`
	Elevation struct {
	} `json:"elevation"`
	Latt string `json:"latt"`
}

// APIError holds information about the rate limit or location miss errors
type APIError struct {
	Description string `json:"description"`
	Code        string `json:"code"`
	Requests    int    `json:"requests"`
}

func (input *APIError) Error() string {
	return input.Code
}

// makeCall actually calls the target API. Unfortunately, they return error codes as HTTP 200, with the shapes of the
// json returns changing depending on the error, so this method gets pretty messy with some of the paths. A great example
// is that some errors place the message in the "description" field and sometimes it's in the "message" field, so we need to do a bunch of checks
func makeCall(params map[string]string) ([]byte, *APIError, error) {
	if httpClient == nil {
		httpClient = resty.New()
	}

	b := []byte{}

	request := httpClient.R()
	if _, ok := params["auth"]; !ok {
		params["auth"] = config.APIKey
	}
	if config.SendMethod == http.MethodGet {
		request.SetQueryParams(params)
	} else if config.SendMethod == http.MethodPost {
		request.SetFormData(params)
	} else {
		// unsupported
		return b, nil, errors.New("method must be GET or POST")
	}
	response, err := request.Execute(config.SendMethod, "https://geocode.xyz")
	if err != nil {
		return b, nil, err
	}
	code := response.StatusCode()
	b = response.Body()
	if code != http.StatusOK && code != http.StatusForbidden {
		// unknown error
		return b, nil, fmt.Errorf("unknown status code: %d", code)
	}

	// since they return errors with an HTTP 200, it's not as simple to decode
	// so first we will put it in a map[string]interface{} to determine if there
	// is an error field
	resultMap := map[string]interface{}{}
	err = json.Unmarshal(b, &resultMap)
	if err != nil {
		return b, nil, err
	}

	// it's an error, and see the opening comment
	if _, ok := resultMap["error"]; ok {
		// it is an error, so we can return the error
		errorMap, ok := resultMap["error"].(map[string]interface{})
		if !ok {
			fmt.Printf("\t\t%+v\n", resultMap["error"])
			return b, nil, errors.New("could not convert the error")
		}
		apiError := &APIError{
			Code: errorMap["code"].(string),
		}
		description := "an error occurred"
		if foundDescription, ok := errorMap["description"].(string); ok {
			description = foundDescription
		} else if foundMessage, ok := errorMap["message"].(string); ok {
			description = foundMessage
		}
		apiError.Description = description

		if errorMap["code"] == "006" {
			// it has a requests field stored as a string
			apiError.Requests, _ = strconv.Atoi(errorMap["requests"].(string))
		}
		return b, apiError, apiError
	}

	return b, nil, nil
}
