package sizer

import (
	"errors"
	"strings"
)

type Unit struct {
	multiplicator string
	unit          string
}

const (
	bit   = "b"
	byt   = "B"
	octet = "o"
)

const (
	none  = ""
	kilo  = "k"
	mega  = "M"
	giga  = "G"
	tera  = "T"
	peta  = "P"
	exa   = "E"
	zetta = "Z"
	yotta = "Y"
)

var Bi = Unit{none, bit}
var Kb = Unit{kilo, bit}
var Mb = Unit{mega, bit}
var Gb = Unit{giga, bit}
var Oc = Unit{none, octet}
var Ko = Unit{kilo, octet}
var Mo = Unit{mega, octet}
var Go = Unit{giga, octet}
var By = Unit{none, byt}
var KB = Unit{kilo, byt}
var MB = Unit{mega, byt}
var GB = Unit{giga, byt}

var multiplicators map[string]string
var factors map[string]uint
var units map[string]string

func init() {

	factors = map[string]uint{
		none: 0,
		kilo: 1,
		mega: 2,
		giga: 3,
		tera: 4,
		peta: 5,
		exa: 6,
		zetta: 7,
		yotta: 8 }

	capitalize := func(s string) string {
		if len(s) == 0 {
			return s
		}
		return strings.ToUpper(string(s[0])) + s[1:]
	}

	units = make(map[string]string)
	initUnit := func(unitname string, fullname string, cst string) {
		units[unitname] = cst
		units[fullname] = cst
		units[capitalize(fullname)] = cst
		units[fullname + "s"] = cst
		units[capitalize(fullname + "s")] = cst
	}
	initUnit("b", "bit", bit)
	initUnit("B", "byte", byt)
	initUnit("o", "octet", octet)

	multiplicators = make(map[string]string)
	initMultiplicator := func(multiplicatorname string, fullname string, cst string) {
		multiplicators[multiplicatorname] = cst
		multiplicators[fullname] = cst
		multiplicators[capitalize(fullname)] = cst
	}
	initMultiplicator("", "", none)
	initMultiplicator("k", "kilo", kilo)
	initMultiplicator("M", "mega", mega)
	initMultiplicator("G", "giga", giga)
	initMultiplicator("T", "tera", tera)
	initMultiplicator("P", "peta", peta)
	initMultiplicator("E", "exa", exa)
	initMultiplicator("Z", "zetta", zetta)
	initMultiplicator("Y", "yotta", yotta)	
}

func (u Unit) String() string {
	return u.multiplicator + u.unit
}

func ParseUnit(b []byte) (Unit, error) {

	var u Unit
	var s = strings.Replace(string(b), " ", "", -1)

	for i := 0; i < len(s); i++ {
		mul, ok := multiplicators[s[:i]]
		if ok {
			un, ok := units[s[i:]]
			if ok {
				u.multiplicator = mul
				u.unit = un
				return u, nil
			}
		}
	}
	return u, errors.New("Unable to parse the unit used: " + string(b))
}
