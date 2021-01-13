package edtf

import (
	"testing"
)

var tests = map[string]map[string]string{
	"MY": map[string]string{
		"04/1972":   "1972-04",
		"03/-1980":  "-1980-03",
		"c. 3/1984": "1984-~03",
	},
	"MDY": map[string]string{
		"6/30/2010":     "2010-06-30",
		"c. 02/29/2020": "2020-02-29~",
	},
	"EARLY": map[string]string{
		"early 1970s":    "1970-01/1970-04",
		"c. early 1950s": "1950-~01/1950-~04",
		"Early 1960s":    "1960-01/1960-04",
		"Early -0200s":   "-0200-01/-0200-04",
	},
	"MID": map[string]string{
		"mid 1970s":    "1970-05/1970-08",
		"c. mid 1950s": "1950-~05/1950-~08",
		"Mid 1960s":    "1960-05/1960-08",
	},
	"LATE": map[string]string{
		"late 1970s":    "1970-09/1970-12",
		"c. late 1950s": "1950-~09/1950-~12",
		"Late 1960s":    "1960-09/1960-12",
	},
	"DECADE": map[string]string{
		"1930s":   "193X",
		"c 1980s": "~198X-01-01/~198X-12-31",
	},
	"RANGE": map[string]string{
		"1970 - 1980":    "1970/1980",
		"1980-1990":      "1980/1990",
		"c. 1994 -2010":  "~1994/~2010",
		"c. 2018- 2020":  "~2018/~2020",
		"c. -0100- 2020": "~-0100/~2020",
	},
	"YYYY": map[string]string{
		"1900":     "1900",
		"c. 1843":  "1843~",
		"c. -0200": "-0200~",
	},
	"MDYHUMAN": map[string]string{
		"Mar 03 1960":  "1960-03-03",
		"Jul 4, 1979":  "1979-07-04",
		"Jan 2, -1982": "-1982-01-02",
	},
}

func TestToEDTFDate(t *testing.T) {

	for _, feature_tests := range tests {

		for input, expected := range feature_tests {

			t.Logf("ToEDTFDate '%s'", input)

			edtf_d, err := ToEDTFDate(input)

			if err != nil {
				t.Logf("Failed to create EDTF date from '%s' (%s), %v", input, expected, err)
				t.Fail()
				continue
			}

			if expected != edtf_d.EDTF {
				t.Logf("Unexpected EDTF string. Expected '%s' but got '%s'", expected, edtf_d.EDTF)
				t.Fail()
				continue
			}
		}

	}

}

func TestToEDTFString(t *testing.T) {

	for _, feature_tests := range tests {

		for input, expected := range feature_tests {

			t.Logf("ToEDTFString '%s'", input)

			edtf_str, err := ToEDTFString(input)

			if err != nil {
				t.Logf("Failed to parse '%s', %v", input, err)
				t.Fail()
				continue
			}

			if edtf_str != expected {
				t.Logf("Invalid result from '%s'. Expected '%s' but got '%s'.", input, expected, edtf_str)
				t.Fail()
				continue
			}
		}

	}

}

func TestEDTFStringFromMY(t *testing.T) {

	for input, expected := range tests["MY"] {

		t.Logf("EDTFStringFromMY '%s'", input)

		edtf_str, err := EDTFStringFromMY(input)

		if err != nil {
			t.Logf("Failed to parse '%s', %v", input, err)
			t.Fail()
			continue
		}

		if edtf_str != expected {
			t.Logf("Invalid result from '%s'. Expected '%s' but got '%s'.", input, expected, edtf_str)
			t.Fail()
			continue
		}
	}

}

func TestEDTFStringFromMDY(t *testing.T) {

	for input, expected := range tests["MDY"] {

		t.Logf("EDTFStringFromMDY '%s'", input)

		edtf_str, err := EDTFStringFromMDY(input)

		if err != nil {
			t.Logf("Failed to parse '%s', %v", input, err)
			t.Fail()
			continue
		}

		if edtf_str != expected {
			t.Logf("Invalid result from '%s'. Expected '%s' but got '%s'.", input, expected, edtf_str)
			t.Fail()
			continue
		}
	}

}

func TestEDTFStringFromEARLY(t *testing.T) {

	for input, expected := range tests["MDY"] {

		t.Logf("EDTFStringFromMDY '%s'", input)

		edtf_str, err := EDTFStringFromMDY(input)

		if err != nil {
			t.Logf("Failed to parse '%s', %v", input, err)
			t.Fail()
			continue
		}

		if edtf_str != expected {
			t.Logf("Invalid result from '%s'. Expected '%s' but got '%s'.", input, expected, edtf_str)
			t.Fail()
			continue
		}
	}

}

func TestEDTFStringFromMID(t *testing.T) {

	for input, expected := range tests["MID"] {

		t.Logf("EDTFStringFromMID '%s'", input)

		edtf_str, err := EDTFStringFromMID(input)

		if err != nil {
			t.Logf("Failed to parse '%s', %v", input, err)
			t.Fail()
			continue
		}

		if edtf_str != expected {
			t.Logf("Invalid result from '%s'. Expected '%s' but got '%s'.", input, expected, edtf_str)
			t.Fail()
			continue
		}
	}

}

func TestEDTFStringFromLATE(t *testing.T) {

	for input, expected := range tests["LATE"] {

		t.Logf("EDTFStringFromLATE '%s'", input)

		edtf_str, err := EDTFStringFromLATE(input)

		if err != nil {
			t.Logf("Failed to parse '%s', %v", input, err)
			t.Fail()
			continue
		}

		if edtf_str != expected {
			t.Logf("Invalid result from '%s'. Expected '%s' but got '%s'.", input, expected, edtf_str)
			t.Fail()
			continue
		}
	}

}

func TestEDTFStringFromDECADE(t *testing.T) {

	for input, expected := range tests["DECADE"] {

		t.Logf("EDTFStringFromDECADE '%s'", input)

		edtf_str, err := EDTFStringFromDECADE(input)

		if err != nil {
			t.Logf("Failed to parse '%s', %v", input, err)
			t.Fail()
			continue
		}

		if edtf_str != expected {
			t.Logf("Invalid result from '%s'. Expected '%s' but got '%s'.", input, expected, edtf_str)
			t.Fail()
			continue
		}
	}

}

func TestEDTFStringFromRANGE(t *testing.T) {

	for input, expected := range tests["RANGE"] {

		t.Logf("EDTFStringFromRANGE '%s'", input)

		edtf_str, err := EDTFStringFromRANGE(input)

		if err != nil {
			t.Logf("Failed to parse '%s', %v", input, err)
			t.Fail()
			continue
		}

		if edtf_str != expected {
			t.Logf("Invalid result from '%s'. Expected '%s' but got '%s'.", input, expected, edtf_str)
			t.Fail()
			continue
		}
	}

}

func TestEDTFStringFromYYYY(t *testing.T) {

	for input, expected := range tests["YYYY"] {

		t.Logf("EDTFStringFromYYYY '%s'", input)

		edtf_str, err := EDTFStringFromYYYY(input)

		if err != nil {
			t.Logf("Failed to parse '%s', %v", input, err)
			t.Fail()
			continue
		}

		if edtf_str != expected {
			t.Logf("Invalid result from '%s'. Expected '%s' but got '%s'.", input, expected, edtf_str)
			t.Fail()
			continue
		}
	}

}

func TestEDTFStringFromMDYHUMAN(t *testing.T) {

	for input, expected := range tests["MDYHUMAN"] {

		t.Logf("EDTFStringFromMDYHUMAN '%s'", input)

		edtf_str, err := EDTFStringFromMDYHUMAN(input)

		if err != nil {
			t.Logf("Failed to parse '%s', %v", input, err)
			t.Fail()
			continue
		}

		if edtf_str != expected {
			t.Logf("Invalid result from '%s'. Expected '%s' but got '%s'.", input, expected, edtf_str)
			t.Fail()
			continue
		}
	}

}
