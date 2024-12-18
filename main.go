package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	productHandler "practice6/handler/product"
	productService "practice6/service/product"
	variantService "practice6/service/variant"
	productStore "practice6/store/product"
	variantStore "practice6/store/variant"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

func connectToDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	dbuser := os.Getenv("DB_USERNAME")
	dbpass := os.Getenv("DB_PASS")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpass, dbhost, dbport, dbname)

	ds, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	return ds
}

func main() {

	db := connectToDB()

	pstr := productStore.New(db)
	vstr := variantStore.New(db)

	vsvc := variantService.New(vstr)
	psvc := productService.New(vsvc, pstr)

	handler := productHandler.New(psvc)

	router := mux.NewRouter()
	router.HandleFunc("/product", handler.Create).Methods(http.MethodPost)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println(err)
	}

}
