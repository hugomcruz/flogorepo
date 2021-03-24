package fradar

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
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

	var outputArray []map[string]string

	for _, s := range dataLines {

		//Split the lines in the comma
		planeRecord := strings.Split(s, ",")

		if planeRecord[0] == "1" {

			//timestamp, _ := strconv.ParseInt(planeRecord[1], 10, 64)
			localmap := map[string]string{
				"msgtype":      planeRecord[0],
				"timestamp":    planeRecord[1],
				"icaohexcode":  planeRecord[2],
				"callsign":     planeRecord[3],
				"altitude":     "",
				"latitude":     "",
				"longitude":    "",
				"onGround":     "",
				"groundSpeed":  "",
				"track":        "",
				"verticalRate": "",
			}
			outputArray = append(outputArray, localmap)
		} else if planeRecord[0] == "2" {

			//timestamp, _ := strconv.ParseInt(planeRecord[1], 10, 64)
			localmap := map[string]string{
				"msgtype":      planeRecord[0],
				"timestamp":    planeRecord[1],
				"icaohexcode":  planeRecord[2],
				"callsign":     "",
				"altitude":     planeRecord[3],
				"latitude":     planeRecord[4],
				"longitude":    planeRecord[5],
				"onGround":     planeRecord[6],
				"groundSpeed":  "",
				"track":        "",
				"verticalRate": "",
			}
			outputArray = append(outputArray, localmap)
		} else if planeRecord[0] == "3" {

			//timestamp, _ := strconv.ParseInt(planeRecord[1], 10, 64)
			localmap := map[string]string{
				"msgtype":      planeRecord[0],
				"timestamp":    planeRecord[1],
				"icaohexcode":  planeRecord[2],
				"callsign":     "",
				"altitude":     planeRecord[3],
				"latitude":     planeRecord[4],
				"longitude":    planeRecord[5],
				"onGround":     planeRecord[6],
				"groundSpeed":  "",
				"track":        "",
				"verticalRate": "",
			}
			outputArray = append(outputArray, localmap)
		} else if planeRecord[0] == "4" {

			//timestamp, _ := strconv.ParseInt(planeRecord[1], 10, 64)
			localmap := map[string]string{
				"msgtype":      planeRecord[0],
				"timestamp":    planeRecord[1],
				"icaohexcode":  planeRecord[2],
				"callsign":     "",
				"altitude":     "",
				"latitude":     "",
				"longitude":    "",
				"onGround":     "",
				"groundSpeed":  planeRecord[3],
				"track":        planeRecord[4],
				"verticalRate": planeRecord[5],
			}
			outputArray = append(outputArray, localmap)
		} else if planeRecord[0] == "5" {

			//timestamp, _ := strconv.ParseInt(planeRecord[1], 10, 64)
			localmap := map[string]string{
				"msgtype":      planeRecord[0],
				"timestamp":    planeRecord[1],
				"icaohexcode":  planeRecord[2],
				"callsign":     "",
				"altitude":     planeRecord[3],
				"latitude":     "",
				"longitude":    "",
				"onGround":     planeRecord[4],
				"groundSpeed":  "",
				"track":        "",
				"verticalRate": "",
			}
			outputArray = append(outputArray, localmap)
		} else if planeRecord[0] == "6" {

			//timestamp, _ := strconv.ParseInt(planeRecord[1], 10, 64)
			localmap := map[string]string{
				"msgtype":      planeRecord[0],
				"timestamp":    planeRecord[1],
				"icaohexcode":  planeRecord[2],
				"callsign":     "",
				"altitude":     planeRecord[3],
				"latitude":     "",
				"longitude":    "",
				"onGround":     "",
				"groundSpeed":  "",
				"track":        "",
				"verticalRate": "",
				"squak":        planeRecord[4],
			}
			outputArray = append(outputArray, localmap)
		}
	}

	if err != nil {
		return true, err
	}

	//log.Debugf("Input: %s", data)

	context.SetOutput("radardata", outputArray)

	return true, nil
}
