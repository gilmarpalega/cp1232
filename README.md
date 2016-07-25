A simple conversion from Utf-8 to charset windows 1232, mainly for gofpdf (https://github.com/jung-kurt/gofpdf) which doesn't support uft-8 till now

## Installation
go get github.com/gilmarpalega/cp1232

```go
package main

import (
	"github.com/gilmarpalega/cp1232"
	"fmt"
)

func main() {
	text := FromUtf8("“Não existe almoço grátis”")
	fmt.Println(text)
}
```
