package plugs

import (
	"testing"
)

func TestPlugForAlbaniaPlugType(t *testing.T) {

	albania := Display("Albania")
	if albania.Plug_type != "C / F" {
		t.Errorf("Albanian plug type incorrect")
	}

}

func TestPlugForAlbaniaVoltz(t *testing.T) {
	albania := Display("Albania")
	if albania.Volts != "230 V" {
		t.Errorf("Albanian voltz is incorrect")
	}

}
func TestPlugForAlbaniaHertz(t *testing.T) {
	albania := Display("Albania")
	if albania.Hertz != "50 Hz" {
		t.Errorf("Albanian hertz is incorrect")
	}
}

func TestPlugForConverter(t *testing.T) {
	if DoINeedAConverter("UK", "USA") == "No" {
		t.Errorf("Converter Method is incorrect for UK and USA")
	}
}

func TestPlugForConverterNoConverterNeeded(t *testing.T) {
	if DoINeedAConverter("Serbia", "Montenegro") == "Yes" {
		t.Errorf("Converter Method is incorrect for Serbia and Montenegro")
	}
}
