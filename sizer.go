package sizer

type Size struct {
	unit Unit
	value Value
}

func abs(a int) uint {
	var s = a
	if a < 0 {
		s = -a
	}
	return uint(s)
}

func (s Size) String() string {
	return s.value.String() + s.unit.String()
}

func ParseStringSize(str string) (Size, error) {
	return ParseSize([]byte(str))
}

func ParseSize(b []byte) (Size, error) {

	var s Size
	var separator int
	var err error
	
	for i := 0; i < len(b); i++ {
		if !(('0' <= b[i] && b[i] <= '9') || b[i] == ',' || b[i] == '.' || b[i] == '-') {
			separator = i
			break
		}
	}
	
	s.value, err = ParseValue(b[:separator])
	if err != nil {
		return s, err
	}
	
	s.unit, err = ParseUnit(b[separator:])
	if err != nil {
		return s, err
	}
	
	return s, nil
}

func (s Size) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

func (s *Size) UnmarshalJSON(b []byte) (error) {
	var err error
	if len(b) >= 2 && b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1:len(b)-1]
	}
	*s, err = ParseSize(b)
	return err
}

func (s Size) ConvertTo(u Unit) Size {

	if u != s.unit {

		// unit
		if u.unit == bit && (s.unit.unit == byt || s.unit.unit == octet) {
			s.value = s.value.toBit()
		} else if (u.unit == byt || u.unit == octet) && s.unit.unit == bit {
			s.value = s.value.toByte()
		}

		// multiplicator
		fromFactor := int(factors[s.unit.multiplicator])
		toFactor := int(factors[u.multiplicator])
		//resultingFactor := math.Pow(1024, math.Abs(float64(fromFactor - toFactor)))
		resultingFactor := float64(uint(1) << (10*abs(fromFactor - toFactor)))
		if fromFactor - toFactor < 0 {
			s.value = s.value.divide(resultingFactor)
		} else {
			s.value = s.value.multiply(resultingFactor)
		}

		s.unit = u
	}
	return s
}

func (s Size) Value() Value {
	return s.value
}

func (s Size) Unit() Unit {
	return s.unit
}

