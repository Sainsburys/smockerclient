package mock

type ContextOption func(context *Context) *Context

func WithCallLimit(times int) ContextOption {
	return func(context *Context) *Context {
		if context == nil {
			context = &Context{}
		}
		context.Times = times
		return context
	}
}
