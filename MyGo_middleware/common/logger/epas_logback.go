package logger

import (
	"encoding/json"
	"os"
	"runtime"
	"strconv"
	"time"
)

type LogBack struct {
	Date         int64       `json:"date,omitempty"`
	Exception    string      `json:"exception"`
	TraceId      string      `json:"traceId,omitempty"`
	Level        string      `json:"level,omitempty"`
	Logger       string      `json:"logger,omitempty"`
	Ip           string      `json:"ip,omitempty"`
	TraceIgnored int         `json:"traceIgnored"`
	MethodName   string      `json:"methodname,omitempty"`
	ClassName    string      `json:"classname,omitempty"`
	Message      interface{} `json:"message,omitempty"`
	Platform     string      `json:"platform,omitempty"`
	MachineName  string      `json:"machineName,omitempty"`
	Application  string      `json:"application,omitempty"`
	Relative     string      `json:"relative,omitempty"`
}

func logBackWrap(message interface{}, traceId string) []byte {
	pc, _, _, ok := runtime.Caller(2)
	machineName, _ := os.Hostname()
	msg, _ := json.Marshal(message)
	lb := LogBack{
		Date:      time.Now().UnixNano() / 1e6,
		Exception: "",
		TraceId:   traceId,
		Level:     "INFO",
		Logger:    "event_logger",
		//Ip:           config.Addr,
		TraceIgnored: 0,
		Message:      string(msg),
		//Platform:     config.System.EpasPlatform,
		MachineName: machineName,
		//Application:  config.System.EpasApplication,
		Relative: strconv.FormatInt(time.Now().Unix(), 10),
	}
	if ok {
		lb.ClassName = "event_logger"
		lb.MethodName = runtime.FuncForPC(pc).Name()
	}
	info, _ := json.Marshal(lb)
	return info
}
