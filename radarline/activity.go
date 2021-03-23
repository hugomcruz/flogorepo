package radarline

import (
	"sync"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/hugomcruz/flogorepo/fradar"
)

// log is the default package logger
//var log = logger.GetLogger("tibco-activity-radar")

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

	linedata := context.GetInput("linedata").(fradar.Output)

	if err != nil {
		return true, err
	}

	//log.Debugf("Input: %s", data)

	context.SetOutput("msgType", linedata.msgType)
	context.SetOutput("callsign", linedata.callsign)
	context.SetOutput("icaohex", linedata.icaoHexCode)

	return true, nil
}
