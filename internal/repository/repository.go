package repository

import (
	"bphn/artikel-hukum/pkg/log"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

type Repository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewDB(conf *viper.Viper, l *log.Logger) *gorm.DB {

	dbUser := conf.GetString("data.mysql.username")
	dbPassword := conf.GetString("data.mysql.password")
	dbHost := conf.GetString("data.mysql.host")
	dbPort := conf.GetString("data.mysql.port")
	dbName := conf.GetString("data.mysql.database")

	logger := zapgorm2.New(l.Logger)
	logger.SetAsDefault()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger})

	if err != nil {
		panic(err)
	}

	return db
}

func NewRepository(db *gorm.DB, logger *log.Logger) *Repository {
	return &Repository{db: db, logger: logger}
}
