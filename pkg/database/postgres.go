package database

import (
	"fmt"
	"github.com/secmohammed/golang-kafka-grpc-poc/config"
	log "github.com/siruspen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
)

type Repository interface {
	Get() *gorm.DB
}
type databaseConnection struct {
	DB *gorm.DB
}

func NewDatabaseConnection(config config.Repository) Repository {
	user, err := config.GetString("app.db.username")
	if err != nil {
		log.Fatalf("database connection failed: %s", err)
	}
	password, err := config.GetString("app.db.password")
	if err != nil {
		log.Fatalf("database connection failed: %s", err)
	}
	database, err := config.GetString("app.db.database")
	if err != nil {
		log.Fatalf("database connection failed: %s", err)
	}
	host, err := config.GetString("app.db.host")
	if err != nil {
		log.Fatalf("database connection failed: %s", err)
	}
	port, err := config.GetInt("app.db.port")
	if err != nil {
		log.Fatalf("database connection failed: %s", err)
	}
	var enableLogging logger.Interface
	enableDBLogging, err := config.GetBool("app.db.log")
	if err != nil {
		log.Fatalf("database connection failed: %s", err)
	}
	if enableDBLogging {
		enableLogging = logger.Default
	}
	dsn := url.URL{
		User:     url.UserPassword(user, password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", host, port),
		Path:     database,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
		Logger: enableLogging,
	})
	if err != nil {
		log.Fatalf("database connection failed: %s", err)
	}
	sync, err := config.GetBool("app.db.sync")
	if err != nil {
		log.Fatalf("database connection failed: %s", err)

	}
	if sync {
		if err := synchronize(db); err != nil {
			log.Fatalf("database connection failed: %s", err)

		}
	}
	return &databaseConnection{DB: db}
}
func (d *databaseConnection) Get() *gorm.DB {
	return d.DB
}
