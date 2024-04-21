package logger

import (
	"context"
	"log/slog"
	"os"
	"shop/internal/common/models"
)

type Logger interface {
	Info(ctx context.Context, msg string)
	Debug(ctx context.Context, msg string)
	Error(ctx context.Context, msg string)
	Warning(ctx context.Context, msg string)
}

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
	if c, ok := ctx.Value(0).(models.LogCtx); ok {
		if c.UserID != 0 {
			rec.Add("userID", c.UserID)
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
	if c, ok := ctx.Value(0).(models.LogCtx); ok {
		c.UserID = userID
		return context.WithValue(ctx, 0, c)
	}
	return context.WithValue(ctx, 0, models.LogCtx{UserID: userID})
}

func (l *logger) Info(ctx context.Context, msg string) {
	slog.InfoContext(ctx, msg)
}

func (l *logger) Error(ctx context.Context, msg string) {
	slog.ErrorContext(ctx, msg)
}

func (l *logger) Warning(ctx context.Context, msg string) {
	slog.WarnContext(ctx, msg)
}

func (l *logger) Debug(ctx context.Context, msg string) {
	slog.DebugContext(ctx, msg)
}
