package processradarpayload

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	ASetting string `md:"aSetting,required"`
}

type Input struct {
	payload []byte `md:"payload,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	bytesVal, _ := coerce.ToBytes(values["payload"])
	r.payload = bytesVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"payload": r.payload,
	}
}

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

func (o *Output) FromMap(values map[string]interface{}) error {

	strVal, _ := coerce.ToString(values["msgType"])
	o.msgType = strVal

	intVal, _ := coerce.ToInt64(values["timestamp"])
	o.timestamp = intVal

	strVal, _ = coerce.ToString(values["icaoHexCode"])
	o.icaoHexCode = strVal

	strVal, _ = coerce.ToString(values["callsign"])
	o.callsign = strVal

	int32Val, _ := coerce.ToInt32(values["altitude"])
	o.altitude = int32Val

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"msgType":      o.msgType,
		"timestamp":    o.timestamp,
		"icaoHexCode":  o.icaoHexCode,
		"callsign":     o.callsign,
		"altitude":     o.altitude,
		"latitude":     o.latitude,
		"longitude":    o.longitude,
		"onGround":     o.onGround,
		"groundSpeed":  o.groundSpeed,
		"track":        o.track,
		"verticalRate": o.verticalRate,
	}
}
