```go
package main

import (
	"github.com/ayushkr12/logger"
)

func main() {
	logger.LogInfo("This is an info message")
	logger.LogError("Something went wrong!")
	logger.LogSuccess("Operation completed successfully")
	logger.LogWarning("This is a warning")
	logger.LogDebug("Debugging details here")
}
```
