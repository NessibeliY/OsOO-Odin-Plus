package requestapi

import (
	"encoding/json"
	"errors"
	"exchange-rates-client/nyeltay/internal/models"
	"io/ioutil"
	"net/http"
)

type RequestAPI struct{}

func New() *RequestAPI {
	return &RequestAPI{}
}

func (r *RequestAPI) Api() error {
	currenciesAPI := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"
	err := r.unmarshalJSON(currenciesAPI, &models.Currencies)
	if err != nil {
		return err
	}

	return nil
}

func (r *RequestAPI) unmarshalJSON(s string, v interface{}) error {
	response, err := http.Get(s)
	if err != nil {
		return errors.New("Error getting API: " + err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New("API request fails with status: " + response.Status)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.New("Error reading API: " + err.Error())
	}

	err = json.Unmarshal(responseData, v)
	if err != nil {
		return errors.New("Error unmarshaling JSON: " + err.Error())
	}

	return nil
}
