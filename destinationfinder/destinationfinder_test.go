package destinationfinder

import "testing"
import "io/ioutil"
import "fmt"

func TestHolidayExtrasExtraction(t *testing.T) {
	var df HolidayExtrasDestinationFinder
	data, err := ioutil.ReadFile("./testdata/holidayextras.html")
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to open holidayextras.html")
		return
	}
	latlong, err := df.ExtractDestinationFromHTML(data)
	if err != nil {
		fmt.Println("ExtractDestinationFromHTML errored with")
		fmt.Println(err)
		t.Fail()
	} else if latlong != "36.847621,10.21709" {
		fmt.Println("Failed to extract correct lat-long: " + latlong)
		t.Fail()
	}
}

func TestPlaneFinderExtraction(t *testing.T) {
	var df PlaneFinderDestinationFinder
	data, err := ioutil.ReadFile("./testdata/planefinder.html")
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to open planefinder.html")
		return
	}
	latlong, err := df.ExtractDestinationFromHTML(data)
	if err != nil {
		fmt.Println("ExtractDestinationFromHTML errored with")
		fmt.Println(err)
		t.Fail()
	} else if latlong != "63.985,-22.6056" {
		fmt.Println("Failed to extract correct lat-long: " + latlong)
		t.Fail()
	}
}
