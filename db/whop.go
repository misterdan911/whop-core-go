package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Whop *gorm.DB

func ConnectWhopDb() {

	// connect to database
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")

	dsn := username + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	Whop = database
}




/*
type Person struct {
	gorm.Model
	name   string
	age    int
	job    string
	salary int
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
*/


	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "root:@tcp(127.0.0.1:3306)/whop_core?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:@tcp(127.0.0.1:3306)/whop_core"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// Migrate the schema
	/*
	app.DB.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
	*/

	/*
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	*/
