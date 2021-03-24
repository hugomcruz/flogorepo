package timestamp

import (
	"time"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnTimestamp{})
}

type fnTimestampMillis struct {
}

func (fnTimestampMillis) Name() string {
	return "currentTimestampMillis"
}

func (fnTimestampMillis) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false

}

func (fnTimestampMillis) Eval(params ...interface{}) (interface{}, error) {

	now := time.Now()
	timestamp := now.UnixNano() / 1000000

	return timestamp, nil
}
