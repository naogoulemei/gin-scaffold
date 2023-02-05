package lib

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSetGinTraceContext(t *testing.T) {
	trace := &TraceContext{
		Trace: Trace{
			TraceId: "sssss",
		},
		CSpanId: "test trace",
	}

	ginCtx := &gin.Context{}
	_ = SetGinTraceContext(ginCtx, trace)
	trace, ok := GetTraceContext(ginCtx)
	if !ok {
		t.Log("fail")
		return
	}
	t.Log(*trace)

	ctx := context.Background()
	ctx = SetTraceContext(ctx, trace)
	trace, ok = GetTraceContext(ctx)
	if !ok {
		t.Log("fail")
		return
	}
	t.Log(*trace)
}
