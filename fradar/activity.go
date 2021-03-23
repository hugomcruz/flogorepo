package fradar

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/project-flogo/core/data/coerce"
)

// log is the default package logger
//var log = logger.GetLogger("tibco-activity-radar")

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

	payload := context.GetInput("payload")
	b, err := coerce.ToBytes(payload)

	//Decompress the payload message
	r, _ := gzip.NewReader(bytes.NewReader(b))
	result, _ := ioutil.ReadAll(r)

	data := string(result)

	// Split the data string into lines
	dataLines := strings.Split(data, "\n")

	var outputArray = []Output{}

	for _, s := range dataLines {
		//Split the lines in the comma
		planeRecord := strings.Split(s, ",")

		timestamp, _ := strconv.ParseInt(planeRecord[1], 10, 64)

		if planeRecord[0] == "1" {

			output := &Output{
				msgType:      planeRecord[0],
				timestamp:    timestamp,
				icaoHexCode:  planeRecord[2],
				callsign:     planeRecord[3],
				altitude:     0,
				latitude:     0,
				longitude:    0,
				onGround:     0,
				groundSpeed:  0,
				track:        0,
				verticalRate: 0,
			}

			outputArray = append(outputArray, *output)

		}

	}

	if err != nil {
		return true, err
	}

	//log.Debugf("Input: %s", data)

	context.SetOutput("radardata", outputArray)

	return true, nil
}
