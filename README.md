## Install

```
go get -u "github.com/ayushkr12/logger"
```

## Usage

```go
package main

import (
	log "github.com/ayushkr12/logger"
)

func main() {
	log.DisableWarn = false
	log.DisableInfo = false
	log.DisableDebug = false
	log.DisableSuccess = false
	log.EnableTimeStamp = false

	log.Info("This is an info message")
	log.Infof("This is an info message with format: %s", "formatted")
	log.Success("This is a success message")
	log.Successf("This is a success message with format: %s", "formatted")
	log.Warn("This is a warning message")
	log.Warnf("This is a warning message with format: %s", "formatted")
	log.Debug("This is a debug message")
	log.Debugf("This is a debug message with format: %s", "formatted")
	log.Error("This is an error message")
	log.Errorf("This is an error message with format: %s", "formatted")
}
```
