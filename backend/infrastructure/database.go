package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDatabase interface {
	Connect() error
	Migrate(models []interface{}) error
	GetDB() *gorm.DB
}

type Database struct {
	DB *gorm.DB
	URL string
}

func NewDatabase(config Config) IDatabase {
	return &Database{
		URL: config.Database.URL,
		DB: nil,
	}
}

func (d *Database) Connect() error {
	db, err := gorm.Open(postgres.Open(d.URL), &gorm.Config{})
	if err != nil {
		return  err
	}

	d.DB = db
	return nil
}

func (d *Database) Migrate(models []interface{}) error {
	err := d.DB.AutoMigrate(models...)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) GetDB() *gorm.DB {
	return d.DB
}
