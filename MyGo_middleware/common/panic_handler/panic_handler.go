package panic_handler

import (
	"MyGo_middleware/common/context"
	"MyGo_middleware/common/logger"
	"runtime/debug"
)

func PanicHandler(ctx *context.Context) {
	if panicErr := recover(); panicErr != nil {
		logger.Builder(ctx).
			Field("errorMsg", panicErr.(error).Error()).
			Field("stackLog", string(debug.Stack())).
			Build().ToCriticalLog()
		return
	}
}
