package sizer

import (
	"encoding/json"
	"testing"
)

type testValue struct {
	value          string
	expectedError  bool
	expectedOutput string
	conversions    []testConversion
}
type testConversion struct {
	unit           Unit
	expectedOutput string
}
var testValues []testValue

func init() {
	testValues = append(testValues,
		// Good values
		testValue{"54Go", false, "54Go", []testConversion{{Mo, "55296Mo"}}},
		testValue{"17ko", false, "17ko", []testConversion{{Ko, "17ko"}}},
		testValue{"3.2 Mo", false, "3.2Mo", []testConversion{{Go, "0.003125Go"}}},
		testValue{"1.8B", false, "1.8B", []testConversion{{Oc, "1.8o"}}},
		testValue{"9,78965o", false, "9.78965o", []testConversion{{Bi, "78.3172b"}}},
		testValue{"8 Tera Octets", false, "8To", []testConversion{{Mo, "8388608Mo"}}},
		testValue{"-3 Poctets", false, "-3Po", []testConversion{{Unit{peta, bit}, "-24Pb"}}},
		testValue{"-17.2 EB", false, "-17.2EB", []testConversion{{Unit{exa, bit}, "-137.6Eb"}}},
		testValue{"2 exa bits", false, "2Eb", []testConversion{{Unit{peta, byt}, "256PB"}}},
		// Bad values
		testValue{"Go", true, "", nil},
		testValue{"54", true, "", nil},
		testValue{"1Tera", true, "", nil},
		testValue{"24,8,2Mb", true, "", nil},
		testValue{"Go54", true, "", nil},
		testValue{"", true, "", nil})
}


func Test_Sizer(t *testing.T) {

	for _, v := range testValues {
		s, err := ParseStringSize(v.value)
		if (err == nil) == v.expectedError {
			t.Error("Parse", v, err, s)
		} else if !v.expectedError {
			if s.String() != v.expectedOutput {
				t.Error("Format", v, s)
			}
		}
	}

}

func Test_SizerConvert(t *testing.T) {

	for _, v := range testValues {
		if !v.expectedError {
			s, _ := ParseStringSize(v.value)
			for _, tv := range v.conversions {
				s = s.ConvertTo(tv.unit)
				if s.String() != tv.expectedOutput {
					t.Error(s.String(), tv.expectedOutput)
				}
			}
		}
	}
}

func Test_SizerJson(t *testing.T) {

	for _, v := range testValues {
		if !v.expectedError {
			var data = []byte(`{"size":"` + v.value + `"}`)
			type unmarshalSize struct {
				Size Size `json:"size"`
			}
			var u unmarshalSize
			err := json.Unmarshal(data, &u)
			if err != nil {
				t.Error(err)
			}
			if u.Size.String() != v.expectedOutput {
				t.Error(u, v)
			}

			data, err = json.Marshal(u)
			if err != nil {
				t.Error(err, u)
			}
			if string(data) != `{"size":"` + v.expectedOutput + `"}` {
				t.Error(string(data))
			}
		}
	}
}
