// Package destinationfinder provides functionality for
// determining where a flight with a given callsign is
// destined.
package destinationfinder

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type HolidayExtrasDestinationFinder struct {
}

/**
 * Retrieves the lat long of the destination (as a simple string, we're not interested in doing
 * any real processing with this, just using it as an index. Uses the flightaware website as
 * a datasource, and parses some js embedded in the page. As such this is potentially
 * brittle, but the function defintion should stand, even if we were to plugin a different
 * data source.
 **/
func (destination_finder HolidayExtrasDestinationFinder) GetDestinationFromCallsign(callsign string) (lat_long string, err error) {
	if callsign == "" {
		return "", errors.New("Not going to get latlong from an empty callsign")
	}
	flight_url := "http://www.holidayextras.co.uk/flight/" + callsign

	resp, err := http.Get(flight_url)
	if err != nil {
		return "", errors.New("Failed to retrieve flight details from " + flight_url)
	}
	defer resp.Body.Close()

	holidayextras_html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("Failed to retrieve flight details from " + flight_url)
	}

	return destination_finder.ExtractDestinationFromHTML(holidayextras_html)
}

func (destination_finder *HolidayExtrasDestinationFinder) ExtractDestinationFromHTML(html []byte) (lat_long string, err error) {
	if strings.Index(string(html), "arrival_latlng") == -1 {
		return "", errors.New("Failed to arrival_latlng in html ")
	}

	tmp_strings := strings.Split(string(html), "arrival_latlng = '")
	lat_long = strings.Split(tmp_strings[1], "'")[0]
	return lat_long, nil
}
