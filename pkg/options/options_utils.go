package options

import (
	"github.com/antihax/optional"
	"strconv"
)

func StringToInt32(value string) optional.Int32 {
	d, _ := strconv.ParseInt(value, 10, 32)
	if d > 0 {
		return optional.NewInt32(int32(d))
	} else {
		return optional.EmptyInt32()
	}
}

func StringToOptional(value string) optional.String {
	if value == "" {
		return optional.EmptyString()
	} else {
		return optional.NewString(value)
	}
}
