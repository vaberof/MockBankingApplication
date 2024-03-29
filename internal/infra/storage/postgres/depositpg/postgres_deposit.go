package depositpg

import "time"

type PostgresDeposit struct {
	Id              uint
	SenderId        uint
	SenderUsername  string
	SenderAccountId uint
	PayeeId         uint
	PayeeUsername   string
	PayeeAccountId  uint
	Amount          uint
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
