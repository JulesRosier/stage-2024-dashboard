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

func DieMsg(msg string, err error) {
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
		DieMsg(msg, err)
	}
}

func MaybeDieErr(err error) {
	if err != nil {
		Die(err)
	}
}
