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
	db.Model(&ProductStock{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")
	db.Model(&ProductStockLog{}).AddForeignKey("product_stock_id", "product_stocks(id)", "CASCADE", "CASCADE")

	// manual migrate
	manualMigrate(db)

	// add index
	db.Model(&Product{}).AddIndex("idx_product_name", "name")
	db.Model(&Sales{}).AddIndex("idx_sales_price", "price")
	db.Model(&ProductColor{}).AddIndex("idx_product_color_color", "color")
	db.Model(&ProductSize{}).AddIndex("idx_product_size_size", "size")
	db.Model(&ProductStockLog{}).AddIndex("idx_product_stock_log_stock", "stock")
	db.Model(&ProductStock{}).AddIndex("idx_product_stock_stock", "stock")
	db.Model(&ProductStock{}).AddIndex("idx_product_stock_sku", "sku")

	// add unique
	db.Model(&User{}).AddUniqueIndex("idx_username_", "username")
	db.Model(&Product{}).AddUniqueIndex("idx_product_name_deleted_", "name", "is_deleted")
	db.Model(&Product{}).AddUniqueIndex("idx_product_uri_deleted_", "uri", "is_deleted")
	db.Model(&ProductStock{}).AddUniqueIndex("idx_product_stock_sku_", "sku")
	db.Model(&ProductSize{}).AddUniqueIndex("idx_product_size_sku_", "sku")
	db.Model(&ProductColor{}).AddUniqueIndex("idx_product_color_pid_sizeid_color_", "product_id", "product_size_id", "color")
	db.Model(&ProductColor{}).AddUniqueIndex("idx_product_color_sku", "sku")

	fmt.Println("=== Auto Migration has beed processed ===")
}

func manualMigrate(db *gorm.DB) {
	/* if !db.HasTable(&ProductSize{}) {
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
	} */

	var count uint

	// Scan
	type Result struct {
		ID int
	}

	// Use scanner/valuer
	type UserValuer struct {
		gorm.Model
		Result
	}

	// insert user system if not exists
	db.Model(&User{}).Count(&count)
	if count == 0 {
		// Gorm Insert
		user := User{Username: "system", Fullname: "system"}
		db.Create(&user)

		// Update user system id
		db.Exec("UPDATE users SET id=? WHERE fullname=?", 0, "system")
	}
}
