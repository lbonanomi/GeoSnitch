package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)



func geocode(lattitude string, longitude string, apiKey string)(plus_code string, locality string, administrative_area_level_2 string, administrative_area_level_1 string, country string){

	request := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%s,%s&key=%s", lattitude, longitude, apiKey)

	client := resty.New()
	resp,apierr := client.R().
		SetHeader("Content-Type", "application/json").
		Post(request)

	if apierr != nil {
		fmt.Println("API Call Bombed")
		fmt.Println(apierr)
	}
	location := string(resp.Body())

	values := gjson.Get(location, "results.#.address_components.#.short_name").Array()
	labels := gjson.Get(location, "results.#.address_components.#.types.0").Array()

	for c,_ := range(values) {
		clam := labels[c].Array()

		for cc,_ := range(clam) {
			if (clam[cc].Str == "plus_code") {
				plus_code = values[c].Array()[0].Str
				locality = values[c].Array()[1].Str
				administrative_area_level_2 = values[c].Array()[2].Str
				administrative_area_level_1 = values[c].Array()[3].Str
				country = values[c].Array()[4].Str
			}
		}
	}

	return
}
