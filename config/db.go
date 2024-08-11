package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// var DB *sql.DB

// // init DB initialize the database connection
// func InitDb() {
// 	// Load environment variables
// 	InitEnv()

// 	// Get database connection string from environtment variable
// 	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		GetEnv("DB_HOST", "localhost"),
// 		GetEnv("DB_PORT", "5432"),
// 		GetEnv("DB_USER", "postgres"),
// 		GetEnv("DB_PASSWORD", ""),
// 		GetEnv("DB_NAME", "postgres"))

// 	var err error
// 	DB, err = sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatalf("Error opening database: %q", err)
// 	}

// 	fmt.Println("Successfully connected to the database!")

// }
//

func NewDatabase(viper *viper.Viper, log *logrus.Logger) *gorm.DB {
	user := viper.GetString("database.postgres.username")
	password := viper.GetString("database.postgres.password")
	host := viper.GetString("database.postgres.host")
	port := viper.GetString("database.postgres.port")
	database := viper.GetString("database.postgres.name")
	idleConnection := viper.GetInt("database.postgres.pool.idle")
	maxConnection := viper.GetInt("database.postgres.pool.max")
	maxLifeTimeConnection := viper.GetInt("database.postgres.pool.lifetime")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, user, password, database)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxOpenConns(maxConnection)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	// if err := db.AutoMigrate(&entity.User{}, &entity.Session{}, &entity.Website{}, &entity.WebsiteEvent{}, &entity.EventData{}, &entity.SessionData{}, &entity.Team{}, &entity.TeamUser{}, &entity.Report{}); err != nil {
	// 	log.Fatalf("failed to migrate database: %v", err)
	// }

	return db
}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}
