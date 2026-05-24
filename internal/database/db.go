package database

import (
	"github.com/AndersonMeloo/WealthFlow/internal/models"
	glebaresqlite "github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Aqui vai AbrirPostgres, que é a função para abrir a conexão com o banco de dados PostgreSQL usando GORM
func AbrirPostgres(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Silent),
		TranslateError: true,
	})
}

//  Aqui vai AbrirSQLite conecta no banco SQLite puro Go usando nos testes
func AbrirSQLite(dsn string) (*gorm.DB, error) {
	return gorm.Open(glebaresqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		TranslateError: true,
	})
}

// MigraAutomaticamente aplica as migrations das entidades do sistema
func MigrarAutomaticamente(db *gorm.DB) error {
	return db.AutoMigrate(&models.Customer{}, &models.ProcessedEvent{})
}
