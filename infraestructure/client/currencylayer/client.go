package currencylayer

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
)

type response struct {
	Success bool        `json:"success"`
	Quotes  interface{} `json:"quotes"`
	Error   struct {
		Code int    `json:"code"`
		Info string `json:"info"`
	} `json:"error"`
}

const (
	client        = "http://apilayer.net/api"
	defaultSource = "USD"
)

type ClientHTTP struct {
	apiKey string
}

func New(apiKey string) ClientHTTP {
	return ClientHTTP{apiKey: apiKey}
}

func (c ClientHTTP) Convert(from, to string, amount float64) (float64, error) {
	url := fmt.Sprintf("%s/live?access_key=%s&currencies=%s,%s&source=%s&format=1", client, c.apiKey, strings.ToUpper(from), strings.ToUpper(to), defaultSource)
	return adapter(url, from, to, amount)
}

// adapter this conversion may not be exact, it was done in this way so as not to pay unnecessarily for the api converter, since it is only a demo
func adapter(url, from, to string, amount float64) (float64, error) {
	request, err := doRequest(http.MethodGet, url)
	if err != nil {
		return 0, fmt.Errorf("adapter.doRequest: %v", err)
	}

	body, err := json.Marshal(request.Quotes)
	if err != nil {
		return 0, fmt.Errorf("adapter.marshal: %v", err)
	}

	currencies := make(map[string]float64, 0)
	err = json.Unmarshal(body, &currencies)
	if err != nil {
		return 0, fmt.Errorf("adapter.unmarshal: %v", err)
	}

	return adapterConverter(currencies, from, to, amount), nil
}

func adapterConverter(currencies map[string]float64, from, to string, amount float64) float64 {
	exchangeRateFrom := currencies[fmt.Sprintf("%s%s", defaultSource, from)]
	amount /= exchangeRateFrom

	exchangeRateTo := currencies[fmt.Sprintf("%s%s", defaultSource, to)]

	return toFixedFloat(amount*exchangeRateTo, 2)
}

func doRequest(method, url string) (response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return response{}, fmt.Errorf("NewRequest(): %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return response{}, fmt.Errorf("DefaultClient: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return response{}, fmt.Errorf("ocurred when request url: %s, status code: %d", url, res.StatusCode)
	}

	responseData := response{}
	if err := deserialize(res.Body, &responseData); err != nil {
		body, errReadAll := ioutil.ReadAll(res.Body)
		if errReadAll != nil {
			return response{}, fmt.Errorf("reading Body deserialization: %v", errReadAll)
		}

		return response{}, fmt.Errorf("deserialization: %v, body was: %s", err, string(body))
	}

	if !responseData.Success {
		return response{}, fmt.Errorf("code: %d, info: %s", responseData.Error.Code, responseData.Error.Info)
	}

	return responseData, nil
}

func deserialize(src io.ReadCloser, dst interface{}) error {
	if err := json.NewDecoder(src).Decode(dst); err != nil {
		return err
	}
	return nil
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixedFloat(num float64, numberDecimals int) float64 {
	output := math.Pow(10, float64(numberDecimals))
	return float64(round(num*output)) / output
}
