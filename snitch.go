package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/table"
	"github.com/schollz/wifiscan"
	"strings"

)

var apiKey = "SET ONE!"

var lattitude string
var longitude string
var plus_code string
var locality string
var administrative_area_level_3 string
var administrative_area_level_2 string
var administrative_area_level_1 string
var country string



func main() {
	//
	// Below is all example code from https://github.com/osquery/osquery-go
	//

	socket := flag.String("socket", "", "Path to osquery socket file")
	flag.Parse()
	if *socket == "" {
		log.Fatalf(`Usage: %s --socket SOCKET_PATH`, os.Args[0])
	}

	server, err := osquery.NewExtensionManagerServer("geosnitch", *socket)
	if err != nil {
		log.Fatalf("Error creating extension: %s\n", err)
	}

	server.RegisterPlugin(table.NewPlugin("geosnitch", GeosnitchColumns(), GeosnitchGenerate))
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}



func GeosnitchColumns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("lattitude"),
		table.TextColumn("longitude"),
		table.TextColumn("plus_code"),
		table.TextColumn("locality"),
		table.TextColumn("administrative_area_level_3"),
		table.TextColumn("administrative_area_level_2"),
		table.TextColumn("administrative_area_level_1"),
		table.TextColumn("country"),
	}
}

func GeosnitchGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	// This is the trigger function that fires when anything is selected from the 'geosnitch' table
	//

	var MACS []string
	var macAddress string

	wifis,_ := wifiscan.Scan()

	// FIXME: This is a lousy way to handle JSON
	//

	for _, w := range wifis {
		macAddress = fmt.Sprintf("{ \"macAddress\":\"%s\"} ", w.SSID)
		MACS = append(MACS, macAddress)
	}

	macAddresses:= strings.Join(MACS, ", ")
	macAddresses = strings.Trim(macAddresses, ",")

	macAddressJson := fmt.Sprintf("{\n\t\"wifiAccessPoints\": [\n\t\t%s\n\t]\n}\n", macAddresses)

	// Get a Lat/Long pair from Google Maps
	//

	lattitude, longitude = geolocate(macAddressJson, apiKey)

	// Convert Lat/Long to Google Maps data
	//

	plus_code,locality,administrative_area_level_3,administrative_area_level_2,administrative_area_level_1 := geocode(lattitude, longitude, apiKey)

	return []map[string]string{
		{
			"lattitude":			lattitude,
			"longitude":			longitude,
			"plus_code":			plus_code,
			"locality":			locality,
			"administrative_area_level_3":	administrative_area_level_3,
			"administrative_area_level_2":	administrative_area_level_2,
			"administrative_area_level_1":	administrative_area_level_1,
			"country": country,
		},
	}, nil
}

