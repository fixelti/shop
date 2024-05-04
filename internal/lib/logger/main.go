package logger

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"shop/internal/common/models"
)

type Logger interface {
	Info(ctx context.Context, msg string)
	Debug(ctx context.Context, msg string)
	Error(ctx context.Context, err error)
	Warning(ctx context.Context, msg string)
}

type keyType int

const key = keyType(0)

type logger struct {
	next slog.Handler
}

func New(level slog.Level) Logger {
	handler := slog.Handler(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	handler = &logger{next: handler}
	slog.SetDefault(slog.New(handler))
	return &logger{}
}

func (l *logger) Enabled(ctx context.Context, rec slog.Level) bool {
	return l.next.Enabled(ctx, rec)
}

func (l *logger) Handle(ctx context.Context, rec slog.Record) error {
	if c, ok := ctx.Value(key).(models.LogCtx); ok {
		if c.UserID != 0 {
			rec.Add("userID", c.UserID)
		}
		if c.OP != "" {
			rec.Add("op", c.OP)
		}
	}
	return l.next.Handle(ctx, rec)
}

func (l *logger) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &logger{next: l.next.WithAttrs(attrs)}
}

func (l *logger) WithGroup(name string) slog.Handler {
	return &logger{next: l.next.WithGroup(name)}
}

func WithLogUserID(ctx context.Context, userID uint) context.Context {
	if c, ok := ctx.Value(key).(models.LogCtx); ok {
		c.UserID = userID
		return context.WithValue(ctx, key, c)
	}
	return context.WithValue(ctx, key, models.LogCtx{UserID: userID})
}

func WithOP(ctx context.Context, op string) context.Context {
	if c, ok := ctx.Value(key).(models.LogCtx); ok {
		c.OP = op
		return context.WithValue(ctx, key, c)
	}
	return context.WithValue(ctx, key, models.LogCtx{OP: op})
}

type errorWithLogCtx struct {
	next error
	ctx  models.LogCtx
}

func (e *errorWithLogCtx) Error() string {
	return e.next.Error()
}

func WrapError(ctx context.Context, err error) error {
	c := models.LogCtx{}
	if x, ok := ctx.Value(key).(models.LogCtx); ok {
		c = x
	}
	return &errorWithLogCtx{
		next: err,
		ctx:  c,
	}
}

func (l *logger) Info(ctx context.Context, msg string) {
	slog.InfoContext(ctx, msg)
}

func (l *logger) Error(ctx context.Context, err error) {
	var errorWithLogCtx *errorWithLogCtx
	if errors.As(err, &errorWithLogCtx) {
		ctx = context.WithValue(ctx, key, errorWithLogCtx.ctx)
		slog.ErrorContext(ctx, errorWithLogCtx.Error())
		return
	}

	slog.ErrorContext(ctx, err.Error())
}

func (l *logger) Warning(ctx context.Context, msg string) {
	slog.WarnContext(ctx, msg)
}

func (l *logger) Debug(ctx context.Context, msg string) {
	slog.DebugContext(ctx, msg)
}
