goddg
========

Golang wrapper for DuckDuckGo Instant Answer API

Installation
---
```
go get -u github.com/danielwetan/goddg
```

Usage
---
```
package main 

import (
    "fmt"
    "github.com/danielwetan/goddg"
)

func main() {
    res, _ := goddg.Query("facebook")
    fmt.Println(res)
}
```

```
go run main.go
```

Author
---
- [Daniel Saputra](https://www.linkedin.com/in/danielwetan/)

License
---
This project is licensed under the MIT License - see the [LICENSE](https://github.com/danielwetan/goddg/blob/master/LICENSE) file for details
