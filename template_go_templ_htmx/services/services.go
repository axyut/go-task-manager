package services

import (
	"template/db"

	"golang.org/x/exp/slog"
)

func GlobalService(log *slog.Logger, cs *db.DBTable) string {
	return "Global Service Usage"
}
