package database

import (
	"fmt"

	"github.com/fnxr21/invoice-system/internal/model"
	"github.com/fnxr21/invoice-system/pkg/mysql"
)

func RunMigration() {
	var err error
	// main migration
	err = mysql.DB.AutoMigrate(
		&model.Invoice{},
		&model.Customer{},
		&model.Item{},
		&model.InvoiceItem{},
	)

	if err != nil {
		fmt.Println(err)
		panic("DB Migration Failed ")
	}

	fmt.Println("All Migration Success")
}
