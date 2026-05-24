package repository

import (
	"context"

	"github.com/AndersonMeloo/WealthFlow/internal/models"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Criar(ctx context.Context, customer *models.Customer) error
	BuscarPorEmail(ctx context.Context, email string) (*models.Customer, error)
	Atualizar(ctx context.Context, customer *models.Customer) error
}

type GormCustomerRepository struct {
	db *gorm.DB
}

// NovoRepositorioCliente cria e retorna um repositório de clientes baseado em GORM.
func NovoRepositorioCliente(db *gorm.DB) CustomerRepository {
	return &GormCustomerRepository{db: db}
}

// Criar persiste um novo cliente no banco de dados.
func (r *GormCustomerRepository) Criar(ctx context.Context, customer *models.Customer) error {
	return r.db.WithContext(ctx).Create(customer).Error
}

// BuscarPorEmail retorna o cliente correspondente ao email informado.
// Retorna (nil, error) caso ocorra um erro durante a consulta.
func (r *GormCustomerRepository) BuscarPorEmail(ctx context.Context, email string) (*models.Customer, error) {
	var customer models.Customer
	if err := r.db.WithContext(ctx).Where("cliente_email = ?", email).First(&customer).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

// Atualizar salva alterações de um cliente existente no banco de dados.
func (r *GormCustomerRepository) Atualizar(ctx context.Context, customer *models.Customer) error {
	return r.db.WithContext(ctx).Save(customer).Error
}
