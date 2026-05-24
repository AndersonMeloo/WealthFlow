package models

import "time"

type Customer struct {
	ID              uint      `gorm:"primaryKey"`
	ClienteNome     string    `gorm:"column:cliente_nome;not null"`
	ClienteEmail    string    `gorm:"column:cliente_email;uniqueIndex;not null"`
	TipoSolicitacao string    `gorm:"column:tipo_solicitacao;not null"`
	ValorPatrimonio float64   `gorm:"column:valor_patrimonio;not null"`
	Status          string    `gorm:"column:status;not null;default:Aguardando Análise"`
	Prioridade      string    `gorm:"column:prioridade"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

func (Customer) TableName() string {
	return "clientes"
}
