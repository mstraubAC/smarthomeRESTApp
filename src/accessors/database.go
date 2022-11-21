package accessors

import (
	"context"
	"mstraubAC/smarthome-restService/configuration"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type DatabaseAccessor struct {
	// public
	Config *configuration.Config
	Logger *zap.Logger

	// private
	db *pgxpool.Pool
}

func (h *DatabaseAccessor) GetSqlConnection() (*pgxpool.Pool, error) {
	if h.db == nil {
		h.Logger.Info("No PostgreSQL connection available. Creating new connection pool")
		dbConn, err := pgxpool.New(context.Background(), h.Config.DBUrl)
		if err != nil {
			h.Logger.Error("PostgresSQL connection could not be established: " + err.Error())
			return h.db, err
		}
		h.db = dbConn
	}

	return h.db, nil
}
