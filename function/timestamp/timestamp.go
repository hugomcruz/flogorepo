package timestamp

import (
	"fmt"
	"time"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnTimestamp{})
}

type fnTimestamp struct {
}

func (fnTimestamp) Name() string {
	return "formatdatetotimestamp"
}

func (fnTimestamp) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false

}

func (fnTimestamp) Eval(params ...interface{}) (interface{}, error) {
	dateString, err := coerce.ToString(params[0])
	if err != nil {
		return nil, err
	}

	dateLayout, err := coerce.ToString(params[1])
	if err != nil {
		return nil, err
	}

	timestamp := dateStringToTimestamp(dateString, dateLayout)

	return timestamp, nil
}

func dateStringToTimestamp(dateString string, dateLayout string) int64 {

	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")

	layout := dateLayout
	//time.LoadLocation("Asia/Ho_Chi_Minh_City")
	t, err := time.ParseInLocation(layout, dateString, loc)

	if err != nil {
		fmt.Println(err)
	}

	timestamp := t.Unix()

	return timestamp

}
