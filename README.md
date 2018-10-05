sizer
=====

Handle sizes in an easy way with Golang

* Author: Julien CHAUMONT (http://julienc.io)
* Version: 1.0
* Date: 2014-08-08
* Licence: MIT
* Url: https://github.com/julienc91/sizer

## Tests

This package includes a set of unit tests you can run with the following command:

    go test
    
## Usage

### Parse sizes from a string

```golang
package main

import (
    "fmt"
    "github.com/julienc91/sizer"
)

func main() {

    var stringSize = "54 kilo octets"
    size, _ := sizer.ParseStringSize(stringSize)
    fmt.Println(size)    
}
```
    
Result:

    54ko
    
When a size is parsed, it is automatically formatted this way:

    <numeric_value>[multiplicator]<base unit>

### Parse sizes from a JSON file
    
```
package main

import (
    "encoding/json"
    "fmt"
    "github.com/julienc91/sizer"
)

func main() {

    var jsonData = []byte(`[{"date": "2014-08-08", "size": "21ko"},
                            {"date": "2014-08-07", "size": "23ko"}]`)
    type sizes struct {
        Date string     "json:`date`"
        Size sizer.Size "json:`size`"
    }

    var s []sizes
    err := json.Unmarshal(jsonData, &s)
    if err != nil {
        panic(err)
    }

    fmt.Println(s)
}
```
    
Result:

    [{2014-08-08 21ko} {2014-08-07 23ko}]

### Convert sizes to another unit

```
package main

import (
    "fmt"
    "github.com/julienc91/sizer"
)

func main() {

    size, _ := sizer.ParseStringSize("54Mb")
    size = size.ConvertTo(sizer.Ko)
    fmt.Println(size)
}
```

Result:

    6912ko
    
Never forget the floating point precision.

## Compatible units

### Base unit

These are the available base units (without multiplicator factors):

* Bit: "b" or "bit"
* Byte: "B" or "byte"
* Octet: "o" or "octet"

Base units can be used in their fullname form with a plural form and/or a capital letter.

Conversion:

> 1 Octet = 1 Byte
> 1 Octet = 8 Bits

### Multiplicators

Add a multiplicator to a base unit to form another unit:

* Kilo: "k" or "kilo"
* Mega: "M" or "mega"
* Giga: "G" or "giga"
* Tera: "T" or "tera"
* Peta: "P" or "peta"
* Exa: "E" or "exa"
* Zetta: "Z" or "zetta"
* Yotta: "Y" or "yotta"

Multiplicators can be used in their fullname form with a plural form and/or a capital letter. A multplicator is not required.

Conversion:

> 1 Kilo  = 1024
> 1 Mega  = 1024k
> 1 Giga  = 1024M
> 1 Tera  = 1024G
> 1 Peta  = 1024T
> 1 Exa   = 1024P
> 1 Zetta = 1024E
> 1 Yotta = 1024Z

### Predefined units

A few units are already defined in the package and can be used directly:

    sizer.Bi // b
    sizer.Kb // kb
    sizer.Mb
    sizer.Gb
    sizer.Oc // o
    sizer.Ko // ko
    sizer.Mo
    sizer.Go
    sizer.By // B
    sizer.KB // kB
    sizer.MB
    sizer.GB

If the unit you want is not one of them, use the `ParseUnit` function to build your own:

    Po, err := sizer.ParseUnit([]byte("Peta Octet"))


## Size

A size is formed by a value and a unit. The value must be numeric, and the unit must be one of the compatible units listed above. The value must come before the unit.

Examples of parsable sizes:

    "54Go", "17ko", "3.2 Mo", "1.8B", "9,78965o", "8 Tera Octets", "-3 Poctets", "-17.2 EB", "2 exa bits"
    
Example of invalid sizes:

    "Go" // No value
    "54" // No unit
    "1Tera" // Invalid unit: no base unit
    "24,8,2Mb" // Invalid value: not numeric
    "Go54" // Value comes after unit
    "17 Ks" // Invalid unit: unknown base unit
    "" // No value and no unit
