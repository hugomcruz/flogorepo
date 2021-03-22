package fradar

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("tibco-activity-radar")

type Output struct {
	msgType      string
	timestamp    int64
	icaoHexCode  string
	callsign     string
	altitude     int32
	latitude     float32
	longitude    float32
	onGround     int32
	groundSpeed  float32
	track        float32
	verticalRate int32
}

type Input struct {
	payload []byte
}

const (
	ivCounterName = "counterName"
	ivIncrement   = "increment"
	ivReset       = "reset"

	ovValue = "value"
)

// CounterActivity is a Counter Activity implementation
type CounterActivity struct {
	sync.Mutex
	metadata *activity.Metadata
	counters map[string]int
}

// NewActivity creates a new CounterActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &CounterActivity{metadata: metadata, counters: make(map[string]int)}
}

// Metadata implements activity.Activity.Metadata
func (a *CounterActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *CounterActivity) Eval(context activity.Context) (done bool, err error) {

	payload := context.GetInput("payload").([]byte)

	//Decompress the payload message
	r, _ := gzip.NewReader(bytes.NewReader(payload))
	result, _ := ioutil.ReadAll(r)

	data := string(result)

	fmt.Print(data)

	if err != nil {
		return true, err
	}

	//log.Debugf("Input: %s", data)

	context.SetOutput("msgType", "test")

	return true, nil
}
