package mock

func NewMockCallLimit(t int) *Context {
	return &Context{Times: t}
}
