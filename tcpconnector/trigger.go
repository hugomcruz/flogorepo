package tcpconnector

import (
	"bufio"
	"bytes"
	"context"
	"net"

	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/carlescere/scheduler"
)

// Create a new logger
var log = logger.GetLogger("trigger-tcpconnector")

// MyTriggerFactory My Trigger factory
type MyTriggerFactory struct {
	metadata *trigger.Metadata
}

// NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &MyTriggerFactory{metadata: md}
}

// New Creates a new trigger instance for a given id
func (t *MyTriggerFactory) New(config *trigger.Config) trigger.Trigger {
	return &MyTrigger{metadata: t.metadata, config: config}
}

// MyTrigger is a stub for your Trigger implementation
type MyTrigger struct {
	metadata *trigger.Metadata
	config   *trigger.Config
	timers   []*scheduler.Job
	handlers []*trigger.Handler
	conn     net.Conn
}

// Initialize implements trigger.Init.Initialize
func (t *MyTrigger) Initialize(ctx trigger.InitContext) error {
	t.handlers = ctx.GetHandlers()
	return nil
}

// Metadata implements trigger.Trigger.Metadata
func (t *MyTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Start implements trigger.Trigger.Start
func (t *MyTrigger) Start() error {
	// start the trigger

	log.Info("Starting TCP connection to listen for incoming messages")
	port := t.config.GetSetting("port")
	hostname := t.config.GetSetting("hostname")

	conn, _ := net.Dial("tcp", hostname+":"+port)
	t.conn = conn

	//message, _ := bufio.NewReader(conn).ReadString('\n')
	scanner := bufio.NewScanner(conn)
	scanner.Split(ScanCRLF)

	for scanner.Scan() {
		payload := scanner.Text()
		log.Debugf("Received payload: %v", payload)

		trgData := make(map[string]interface{})
		trgData["payload"] = payload

		log.Debug("Processing handlers")
		for _, handler := range t.handlers {
			results, err := handler.Handle(context.Background(), trgData)
			if err != nil {
				log.Error("Error starting action: ", err.Error())
			}
			log.Debugf("Ran Handler: [%s]", handler)
			log.Debugf("Results: [%v]", results)
		}

	}

	return nil
}

// Stop implements trigger.Trigger.Start
func (t *MyTrigger) Stop() error {
	// stop the trigger
	conn := t.conn
	conn.Close()
	return nil
}

// Add on functions needed

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func ScanCRLF(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\r', '\n'}); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}
