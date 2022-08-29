package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/opentracing/opentracing-go"
)

func TestTrace(t *testing.T) {
	ctx := context.Background()
	span, ctx := opentracing.StartSpanFromContext(ctx, "test_operation")
	defer span.Finish()
	inner(ctx)
}

func inner(ctx context.Context) {
	fmt.Println("dd")
	_ = ctx
}
