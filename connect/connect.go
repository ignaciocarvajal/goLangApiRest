package connect

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"

	"github.com/ignaciocarvajal/goLangApiRest/structure"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var connection *gorm.DB

const engine_sql string = "mysql"

// en variables de entorno
const username string = "root"
const password string = "root"
const database string = "userdummy"

func InitializeDatabase() {
	connection = ConnectORM(CreateStringConnection())
	log.Println("La coneccion de base de datos fue exitosa")

}

func CloseConnection() {
	connection.Close()
	log.Println("La coneccion de base de datos fue cerrada")
}

func ConnectORM(connectStrig string) *gorm.DB {

	db, err := gorm.Open(engine_sql, connectStrig)
	if err != nil {
		log.Println(err)
		return nil
	}
	return db
}

func GetUser(id string) structure.User {
	user := structure.User{}
	err := connection.Find(&user, id)
	if err != nil {
		log.Println(err)
	}
	return user
}

func CreateStringConnection() string {
	return username + ":" + password + "@/" + database + "?charset=utf8&parseTime=True&loc=Local"
}
