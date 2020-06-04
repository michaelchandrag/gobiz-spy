package geocoding

import (
	"os"
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type (
	GeocodingResponse struct {
		Results 		[]Results 		`json:"results"`
	}

	Results struct {
		AddressComponents 		[]AddressComponents 		`json:"address_components"`
		FormattedAddress 		string 						`json:"formatted_address"`
	}

	AddressComponents struct {
		LongName 		string 			`json:"long_name"`
		ShortName 		string 			`json:"short_name"`
		Types 			[]string 		`json:"types"`
	}
)

func ConvertLatLngToAddress (lat string, lng string) (result GeocodingResponse, err error) {
	googleApiKey := os.Getenv("GOOGLE_API_KEY")
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%s,%s&key=%s", lat, lng, googleApiKey)
		fmt.Println(url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	defer res.Body.Close()

	// printResponse(*res)
	json.NewDecoder(res.Body).Decode(&result)
	return result, nil
}

func printResponse (res http.Response) {
	fmt.Println(res.Status)
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bodyBytes))
}