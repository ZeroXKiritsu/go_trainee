package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
)

type Config struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DBName       string
	SSLMode      string
	MaxReconn    int
	MinReconn    int
	PingNoEvents int
}

const ordersCustomTable = "orderssave"

type Postgres struct {
	DB          *pgx.Conn
	listner     *pgx.Conn
	config      Config
	listnerStop chan int
}

func NewPostgres(cfg Config) (*Postgres, error) {

	urlPostgres := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host,
		cfg.Port, cfg.DBName, cfg.SSLMode)

	db, err := pgx.Connect(context.Background(), urlPostgres)

	if err != nil {
		return nil, fmt.Errorf("failed to connect postgres %s", err)
	}

	err = db.Ping(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to ping postgres %s", err)
	}

	createListner := func(config Config) {
		return 
	}

	listner := createListner(cfg)

	return &Postgres{DB: db, listner: listner, config: cfg, listnerStop: make(chan int)}, nil
}
