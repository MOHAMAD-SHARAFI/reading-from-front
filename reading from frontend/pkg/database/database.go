package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type impl struct {
	db *gorm.DB
}

type Database interface {
	DB() *gorm.DB
	Close() error
}

func Connect(dbPath string, debug bool) (Database, error) {
	var (
		err        error
		db         *gorm.DB
		gormLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:        time.Second, // Slow SQL threshold
				LogLevel:             logger.Info, // Log level
				ParameterizedQueries: false,       // Don't include params in the SQL log
				Colorful:             true,        // Disable color
			},
		)
	)

	if debug {
		db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{Logger: gormLogger})
	} else {
		db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	}
	if err != nil {
		return nil, err
	}

	return &impl{db: db}, nil
}

func (m *impl) DB() *gorm.DB {
	return m.db
}

func (m *impl) Close() error {
	conn, err := m.db.DB()
	if err != nil {
		return err
	}

	return conn.Close()
}
