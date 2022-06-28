package logger

//
//import (
//	"MyGo_middleware/common/context"
//	"MyGo_middleware/logs"
//	"runtime"
//	"time"
//)
//
//const (
//	DateFormatDateTimePattern = "2006-01-02 15:04:05"
//)
//
//var eventLogger *logs.BeeLogger
//
////自己封装的业务日志对象
//type bizLogger struct {
//	logger *logs.BeeLogger
//	df     *defaultField
//	logStr string
//}
//
//type defaultField struct {
//	TraceId      string                 `json:"trace_id,omitempty"`
//	Scene        string                 `json:"scene,omitempty"`
//	Ability      string                 `json:"ability,omitempty"`
//	RedisKey     string                 `json:"redis_key,omitempty"`
//	LogType      string                 `json:"log_type,omitempty"`
//	Message      interface{}            `json:"message,omitempty"` //普通的message字段
//	ErrorReason  string                 `json:"error_reason,omitempty"`
//	ErrorMessage string                 `json:"error_message,omitempty"`
//	Fields       map[string]interface{} `json:"fields,omitempty"`
//	DateTime     string                 `json:"date_time"`
//	Method       string                 `json:"method,omitempty"`
//}
//
//func Builder(ctx *context.Context) (bl *bizLogger) {
//	return initField(ctx)
//}
//
////构造bizLogger对象
//func BuilderWithTraceId(traceId string) (bl *bizLogger) {
//	return initField(context.NewContext(traceId))
//}
//
////构造bizLogger对象
//func BuilderWithNotCtx() (bl *bizLogger) {
//	return initField(nil)
//}
//
////对field做一些默认的初始化
////公共参数
//func initField(ctx *context.Context) *bizLogger {
//	bl := &bizLogger{logger: eventLogger, df: &defaultField{
//		Fields: make(map[string]interface{}),
//	}}
//	if ctx != nil {
//		bl.df.TraceId = ctx.GetTraceId()
//		bl.df.Scene = ctx.Scene
//		bl.df.Ability = ctx.Ability
//		//bl.df.RedisKey = ctx.EngineBestHashKey
//	}
//	bl.df.DateTime = time.Now().Format(DateFormatDateTimePattern)
//	//定义执行的方法
//	if pc, _, _, ok := runtime.Caller(2); ok {
//		f := runtime.FuncForPC(pc)
//		bl.df.Method = f.Name()
//	}
//	return bl
//}
//
//func (l *bizLogger) Business() *bizLogger {
//	l.df.LogType = "business"
//	return l
//}
//
//func (l *bizLogger) Error(errReason string, errMessage string) *bizLogger {
//	l.df.ErrorReason = errReason
//	l.df.ErrorMessage = errMessage
//	return l
//}
//
////key value 形式的打印 打点日志 value可以为任何值
//func (l *bizLogger) Field(key string, value interface{}) *bizLogger {
//	l.df.Fields[key] = value
//	return l
//}
//
//func (l *bizLogger) Build() *bizLogger {
//	bytes := logBackWrap(l.df, l.df.TraceId)
//	l.logStr = string(bytes)
//	return l
//}
//
////将fieldMap转为json并输出Error级别的日志
//func (l *bizLogger) ToErrorLog() {
//	l.doControlLog(func() {
//		l.logger.Error(l.logStr)
//	})
//}
//
////将fieldMap转为json并输出Info级别的日志
//func (l *bizLogger) ToInfoLog() {
//	l.doControlLog(func() {
//		l.logger.Info(l.logStr)
//	})
//}
//
//func (l *bizLogger) doControlLog(f func()) {
//	f()
//	l.logger.Flush()
//	l.df = nil
//}
//
////将fieldMap转为json并输出Error级别的日志
//func (l *bizLogger) ToCriticalLog() {
//	l.doControlLog(func() {
//		l.logger.Critical(l.logStr)
//	})
//}
