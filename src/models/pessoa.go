package models

import (
	"github.com/guregu/null/zero"
	_ "gorm.io/driver/postgres"
)

type Pessoa struct {
	ID          uint     `json:"id" gorm:"primary key; autoIncrement" `
	Nome        *string  `json:"nome" gorm:"not null"  `
	CPF         zero.Int `json:"cpf" gorm:"not null;unique" `
	Nascimento  *string  `json:"nascimento"`
	Telefone    *uint64  `json:"telefone"`
	Email       *string  `json:"email" gorm:"unique" `
	Rua         *string  `json:"rua" gorm:"not null" `
	Numero      uint64   `json:"numero" gorm:"not null" `
	Bairro      *string  `json:"bairro" gorm:"not null" `
	Complemento *string  `json:"complemento"`
	Cidade      *string  `json:"cidade" gorm:"not null" `
	UF          *string  `json:"uf" gorm:"not null"  `
	CEP         zero.Int `json:"cep" gorm:"not null"  `
}
