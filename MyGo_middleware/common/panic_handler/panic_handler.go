package panic_handler

//func PanicHandler(ctx *context.Context) {
//	if panicErr := recover(); panicErr != nil {
//		logger.Builder(ctx).
//			Field("errorMsg", panicErr.(error).Error()).
//			Field("stackLog", string(debug.Stack())).
//			Build().ToCriticalLog()
//		return
//	}
//}
