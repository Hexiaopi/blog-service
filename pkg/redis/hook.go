package redis

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	tracerLog "github.com/opentracing/opentracing-go/log"
)

type RedisTracingHook struct {
	tracer opentracing.Tracer
}

var _ redis.Hook = RedisTracingHook{}

// NewHook creates a new go-redis hook instance and that will collect spans using the provided tracer.
func NewHook(tracer opentracing.Tracer) redis.Hook {
	return &RedisTracingHook{
		tracer: tracer,
	}
}

func (hook RedisTracingHook) createSpan(ctx context.Context, operationName string) (opentracing.Span, context.Context) {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		childSpan := hook.tracer.StartSpan(operationName, opentracing.ChildOf(span.Context()))
		return childSpan, opentracing.ContextWithSpan(ctx, childSpan)
	}

	return opentracing.StartSpanFromContextWithTracer(ctx, hook.tracer, operationName)
}

func (hook RedisTracingHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	span, ctx := hook.createSpan(ctx, "redis")
	span.SetTag("db.type", "redis")
	span.LogFields(tracerLog.String("cmd", cmd.String()))
	return ctx, nil
}

func (hook RedisTracingHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	span := opentracing.SpanFromContext(ctx)
	defer span.Finish()

	if err := cmd.Err(); err != nil {
		recordError(ctx, "db.error", span, err)
	}
	return nil
}

func (hook RedisTracingHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	span, ctx := hook.createSpan(ctx, "pipeline")
	span.SetTag("db.type", "redis")
	span.SetTag("db.redis.num_cmd", len(cmds))
	return ctx, nil
}

func (hook RedisTracingHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	span := opentracing.SpanFromContext(ctx)
	defer span.Finish()

	for i, cmd := range cmds {
		if err := cmd.Err(); err != nil {
			recordError(ctx, "db.error"+strconv.Itoa(i), span, err)
		}
	}
	return nil
}

func recordError(ctx context.Context, errorTag string, span opentracing.Span, err error) {
	if err != redis.Nil {
		span.SetTag(string(ext.Error), true)
		span.SetTag(errorTag, err.Error())
	}
}

func WithTrace(tracer opentracing.Tracer) Option {
	return func(conf *Config) {
		conf.Tracer = tracer
	}
}
