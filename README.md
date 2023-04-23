gowhois
===

gowhois is whois command implemented by Golang

## Installation
```sh
go get -u github.com/famasoon/gowhois
```

## Usage
```sh
gowhois example.com
```

## Importing
```go
import (
    "github.com/famasoon/gowhois/whois"
)
```

## Example
### Get whois on example.com
```go
package main

import (
    "fmt"

    "github.com/famasoon/gowhois/whois"
)

func main() {
    result, _ := whois.Whois("example.com")
    fmt.Println(result)
}
```

## Thanks
- When I implemented it, inspired [whois-go](https://github.com/likexian/whois-go) (Developed by [Li Kexian](https://www.likexian.com/en-US/)).
- [Whois servers list](https://github.com/cheenanet/whois-servers-list/blob/master/whois.min.json) provided by [Cheena](https://twitter.com/cheenanet)
