package DB

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	config "lab3/config"
	prod "lab3/internal/products"
)

var Connect *gorm.DB
var conf = config.GetConfig()

func init() {
	Connect = ConnectToDb()
}
var products *[]prod.Product

func ConnectToDb() *gorm.DB {

	host := "host=" + conf.HostPostgre
	port := " port=" + conf.PortPostgre
	user := " user=" + conf.UserNamePostgre
	password := " password=" + conf.PasswordPostgre
	dbname := " dbname=" + conf.DBname
	sslmode := " sslmode=" + conf.Sslmode
	dsn := host + user + password + dbname + port + sslmode
	con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Подключение к базе невозможно.")
		return  nil
	}
	return con
}

func InsertProducts(prod []prod.Product){
	err := Connect.Model(prod).Table("Products").Create(prod) 
		if err.Error != nil {
			fmt.Println("Вставка невозможна.")
		}
	}

func SelectAllProducts(){
	err := Connect.Table("Products").Find(&products)
	fmt.Println(products)
	if err.Error != nil {
		fmt.Println(err.Error)
	}
}

func UserSelect(){
	var typeName string
	var param string
	fmt.Println("Какую выборку хотите осуществить?\nВведите (1) - по типу\n(2) - по рейтингу\n(3) - по цене")
	fmt.Scanf("%s\n", &param)
	switch param {
		case "1":
			fmt.Println("Введите тип продукта(wire/wireless)")
			param = "type = '"
			fmt.Scanf("%s\n", &typeName)
			param = param + typeName + "'"
		case "2":
			fmt.Println("Введите рейтинг")
			fmt.Scanf("%s\n", &typeName)
			fmt.Println("(1) - больше\n(2) - меньше")
			fmt.Scanf("%s\n", &param)
			switch param {
				case "1":
					param = "rate > " + typeName
				case "2":
					param = "rate < " + typeName
			}
		case "3":
			fmt.Println("Введите цену")
			fmt.Scanf("%s\n", &typeName)
			fmt.Println("(1) - больше\n(2) - меньше")
			fmt.Scanf("%s\n", &param)
			switch param {
				case "1":
					param = "price > " + typeName
				case "2":
					param = "price < " + typeName
			}
	}
	Connect.Table("Products").Where(param).Find(&products)
	fmt.Println(products)
}

func UpdDel(){

	tx := Connect.Begin()
	errUpd := tx.Table("Products").Where("price = ?",1428).Update("name", "Vgpframed")
	if errUpd.Error != nil {
		fmt.Println("Невозможно обновить запись.")
	} else {
		fmt.Println("Запись обновлена.")
	}
	errDel := tx.Table("Products").Where("price = ?",1570).Delete("")
	if errDel.Error != nil {
		fmt.Println("Невозможно удалить запись.")
	} else {
		fmt.Println("Запись удалена.")
	}
	//tx.Rollback()
	tx.Commit()
}