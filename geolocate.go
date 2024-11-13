package main

import (
	"fmt"
        "github.com/go-resty/resty/v2"
        "github.com/tidwall/gjson"
)

func geolocate(macAddressJson string, apiKey string)(lat string, lng string){

        request := fmt.Sprintf("https://www.googleapis.com/geolocation/v1/geolocate?key=%s", apiKey)

        rclient := resty.New()
        resp,apierr := rclient.R().
                SetHeader("Content-Type", "application/json").
                SetBody(macAddressJson).
                Post(request)

        if apierr != nil {
                fmt.Println("API Call Bombed")
                fmt.Println(apierr)
        }
        location := string(resp.Body())

        latJson := gjson.Get(location, "location.lat")
	lat = latJson.String()

        lngJson := gjson.Get(location, "location.lng")
	lng = lngJson.String()

	return
}
