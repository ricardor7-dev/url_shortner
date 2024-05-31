package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

type Config struct {
	PostgresHost	string
	PostgresPort	string
    PostgresUser     string
    PostgresPassword string
    PostgresDB       string
	PostgresSSMLMode bool
}

func LoadConfig() Config {
	sslmode, _ := strconv.ParseBool(os.Getenv("POSTGRES_SSMLMODE"))
    return Config{
		PostgresHost: 		os.Getenv("POSTGRES_HOST"),
		PostgresPort: 		os.Getenv("POSTGRES_PORT"),
        PostgresUser:     os.Getenv("POSTGRES_USER"),
        PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
        PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresSSMLMode: sslmode,
    }
}

func InitDB(config Config) *pgx.Conn{
    var err error
	//postgresql://teste_owner:S4x3OpVWMvth@ep-polished-sky-a29p88fr.eu-central-1.aws.neon.tech/teste?sslmode=require
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
	config.PostgresUser, config.PostgresPassword, config.PostgresHost, config.PostgresPort, config.PostgresDB)
	
	if config.PostgresPort ==""{
		dsn = fmt.Sprintf("postgres://%s:%s@%s/%s",
        config.PostgresUser, config.PostgresPassword, config.PostgresHost, config.PostgresDB)
	}
	if config.PostgresSSMLMode{
		dsn +="sslmode=require"
	}
    db, err = pgx.Connect(context.Background(), dsn)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
	return db
}
