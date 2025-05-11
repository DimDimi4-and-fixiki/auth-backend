package numverify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/DimDimi4-and-fixiki/auth-back/internal/model"
)

const (
	baseURL = "http://apilayer.net/api/validate"
)

type client struct {
	apiKey string
}

type PhoneValidationResponse struct {
	Valid               bool   `json:"valid"`
	Number              string `json:"number"`
	LocalFormat         string `json:"local_format"`
	InternationalFormat string `json:"international_format"`
	CountryPrefix       string `json:"country_prefix"`
	CountryCode         string `json:"country_code"`
	CountryName         string `json:"country_name"`
	Location            string `json:"location"`
	Carrier             string `json:"carrier"`
	LineType            string `json:"line_type"`
}

func (c *client) ValidatePhoneNumber(phoneNumber string) (*model.PhoneValidationResult, error) {
	params := url.Values{}
	params.Add("access_key", c.apiKey)
	params.Add("number", phoneNumber)

	// Формируем полный URL запроса
	requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Выполняем HTTP GET-запрос
	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("unable to call numverify API: %v", err)
	}
	defer resp.Body.Close()

	// Декодируем JSON-ответ
	var result PhoneValidationResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("unable to decode numverify resp: %v", err)
	}

	return validationResultToModel(&result), nil
}

func validationResultToModel(result *PhoneValidationResponse) *model.PhoneValidationResult {
	return &model.PhoneValidationResult{
		IsValid:       result.Valid,
		CountryCode:   result.CountryCode,
		CountryPrefix: result.CountryPrefix,
		Location:      result.Location,
		Carrier:       result.Carrier,
	}
}
