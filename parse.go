package reflectx

import (
	"fmt"
	"log"
	"os"
	"time"
)

var parseLogger = log.New(os.Stdout, "[Parse]", log.LstdFlags)

var (
	layout1 = "2006-01-02 15:04:05"
	layout2 = "2006-01-02T15:04:05"
)

// ParseTime parse Time from val
func ParseTime(val string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", val)
	if err != nil {
		t, err = time.Parse("2006-01-02T15:04:05", val)
		if err != nil {
			parseLogger.Println(fmt.Sprintf("[ParseTime]%v", err))
		}
	}
	return t
}

// ParseDuration parse Duration from val
func ParseDuration(val string) time.Duration {
	t, err := time.ParseDuration(val)
	if err != nil {
		parseLogger.Println(fmt.Sprintf("[ParseDuration]%v", err))
	}
	return t
}
