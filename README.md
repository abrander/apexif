# apexif

[![PkgGoDev](https://pkg.go.dev/badge/github.com/abrander/apexif)](https://pkg.go.dev/github.com/abrander/apexif)

This is a simple Go package for reading EXIF data from various file
formats. It's build for speed and simplicity. The package does not
have any dependencies.

All data must be loaded in memory (or mmapped) before calling `Identify()`.

### Supported file formats

- [x] CR2
- [x] CRW
- [x] HEIC
- [x] JPEG
- [x] PNG
- [x] TIFF
- [x] WebP

#### Supported container types

- [x] ISOBMFF (MPEG-4 Part 12)
- [x] RIFF
- [x] TIFF

These are not file formats and only interesting for developers of
this package.

### Usage example

```go
package main

import (
	"fmt"
	"os"

	"github.com/abrander/apexif"
)

func main() {
	data, err := os.ReadFile("example.jpeg")
	if err != nil {
		panic(err)
	}

	f, err := apexif.Identify(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Yay. We got a %s image.\n", f.Name())

	e, err := f.Exif()
	if err != nil {
		panic(err)
	}

	make, _ := e.Make()
	fmt.Printf("The image was captured by a %s camera.\n", make)
}
```

### License

This package is licensed under the MIT license. See LICENSE for details.
