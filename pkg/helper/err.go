package helper

import (
	"context"
	"log/slog"
	"os"
)

func Die(err error) {
	slog.LogAttrs(
		context.Background(),
		slog.LevelError,
		err.Error(),
	)
	os.Exit(1)
}

func DieMsg(err error, msg string) {
	slog.LogAttrs(
		context.Background(),
		slog.LevelError,
		msg,
		slog.String("error", err.Error()),
	)
	os.Exit(1)
}

func MaybeDie(err error, msg string) {
	if err != nil {
		DieMsg(err, msg)
	}
}

func MaybeDieErr(err error) {
	if err != nil {
		Die(err)
	}
}
