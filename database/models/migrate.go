package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	// auto migrate
	db.AutoMigrate(
		&User{},
		&Post{},
		&Product{},
		&Sales{},
		&ProductSize{},
		&ProductColor{},
		&ProductStock{},
		&ProductStockLog{},
	)

	// set up foreign keys
	db.Model(&Post{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Product{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Sales{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")
	db.Model(&ProductSize{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")
	db.Model(&ProductColor{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")
	db.Model(&ProductColor{}).AddForeignKey("product_size_id", "product_sizes(id)", "CASCADE", "CASCADE")

	// manual migrate
	// manualMigrate(db)

	// add index
	db.Model(&Product{}).AddIndex("idx_product_name", "name")

	// add unique
	db.Model(&Product{}).AddUniqueIndex("idx_product_name", "name")

	fmt.Println("=== Auto Migration has beed processed ===")
}

func manualMigrate(db *gorm.DB) {
	if !db.HasTable(&ProductSize{}) {
		db.Exec(`CREATE TABLE product_size (
			id int(10) unsigned NOT NULL AUTO_INCREMENT,
			created_at timestamp NULL DEFAULT NULL,
			updated_at timestamp NULL DEFAULT NULL,
			deleted_at timestamp NULL DEFAULT NULL,
			product_id int(10) unsigned DEFAULT NULL,
			size varchar(16) DEFAULT NULL,
			created_by int(10) unsigned DEFAULT NULL,
			updated_by int(10) unsigned DEFAULT NULL,
			deleted_by int(10) unsigned DEFAULT NULL,
			PRIMARY KEY (id),
			KEY idx_product_size_deleted_at (deleted_at),
			KEY product_size_product_id_products_id_foreign (product_id),
			CONSTRAINT product_size_product_id_products_id_foreign FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE ON UPDATE CASCADE
		  ) ENGINE=InnoDB DEFAULT CHARSET=latin1`)
	}
}
