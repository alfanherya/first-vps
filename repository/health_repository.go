package repository

import (
	"context"
	"first-app/model/response"
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HealthRepository struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

func NewHealthRepository(log *logrus.Logger) *HealthRepository {
	return &HealthRepository{
		Log: log,
	}
}

func (r *HealthRepository) checkDB(tx *gorm.DB, health *response.CheckDBResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// ping database
	db, err := tx.DB()
	if err != nil {
		health.Status = "down"
		health.Message = fmt.Sprintf("db down: %v", err)
		r.Log.Fatalf(fmt.Sprintf("db down: %v", err))
		return err
	}

	err = db.PingContext(ctx)
	if err != nil {
		health.Status = "down"
		health.Message = fmt.Sprintf("db down: %v", err)
		r.Log.Fatalf(fmt.Sprintf("db down: %v", err))
		return err
	}

	// Database is up, add more statistics
	health.Status = "Up"
	health.Message = "It's healthy"

	// Get database stats (like open connection, in use, idle, etc)
	dbStats := db.Stats()
	health.OpenConnections = strconv.Itoa(dbStats.OpenConnections)
	health.InUse = strconv.Itoa(dbStats.InUse)
	health.Idle = strconv.Itoa(dbStats.Idle)
	health.WaitCount = strconv.FormatInt(dbStats.WaitCount, 10)
	health.WaitDuration = dbStats.WaitDuration.String()
	health.MaxIdleClosed = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	health.MaxLifetimeClosed = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		health.Message = "The database is experiencing heavy load."
	}
	if dbStats.WaitCount > 1000 {
		health.Message = "The database has a high number of wait events, indicating potential bottlenecks."
	}
	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		health.Message = "Many idle connections are being closed, consider revising the connection pool settings."
	}
	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		health.Message = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return nil
}
