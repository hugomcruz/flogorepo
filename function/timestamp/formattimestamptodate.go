package timestamp

import (
	"time"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnTsToDate{})
}

type fnTsToDate struct {
}

func (fnTsToDate) Name() string {
	return "converttsdatestr"
}

func (fnTsToDate) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeInt64, data.TypeString}, false

}

func (fnTsToDate) Eval(params ...interface{}) (interface{}, error) {
	timestamp, err := coerce.ToInt64(params[0])
	if err != nil {
		return nil, err
	}

	dateLayout, err := coerce.ToString(params[1])
	if err != nil {
		return nil, err
	}
	tm := time.Unix(timestamp, 0)

	dateTime := tm.Format(dateLayout)

	return dateTime, nil
}
