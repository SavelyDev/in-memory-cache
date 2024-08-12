# in-memory-cache

Implementation of an in-memory cache with TTL support for each record, Go package

## Example

Get:
`go get -u -v  github.com/SavelyDev/in-memory-cache`

Code:

```
package main

import (
	"fmt"
	"github.com/SavelyDev/in-memory-cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId, _ := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId, _ = cache.Get("userId")

	fmt.Println(userId)
}
```

Output:

```
42
<nil>
```
