package radar

import (
	"sync"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-radar-data")

type Output struct {
	msgType      string  `md:"msgType"`
	timestamp    int64   `md:"timestamp"`
	icaoHexCode  string  `md:"icaoHexCode"`
	callsign     string  `md:"callsign"`
	altitude     int32   `md:"altitude"`
	latitude     float32 `md:"latitude"`
	longitude    float32 `md:"longitude"`
	onGround     int32   `md:"onGround"`
	groundSpeed  float32 `md:"groundSpeed"`
	track        float32 `md:"track"`
	verticalRate int32   `md:"verticalRate"`
}

type Input struct {
	payload []byte `md:"payload,required"`
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

	input := &Input{}
	err = context.GetInputObject(input)
	if err != nil {
		return true, err
	}

	payload := input.payload

	log.Debugf("Input: %s", payload)

	output := &Output{
		msgType:      "1221",
		timestamp:    0,
		icaoHexCode:  "1212",
		callsign:     "",
		altitude:     0,
		latitude:     0,
		longitude:    0,
		onGround:     0,
		groundSpeed:  0,
		track:        0,
		verticalRate: 0,
	}
	err = context.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
