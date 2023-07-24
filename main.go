package main

import (
	"net/http"

	"github.com/dickidarmawansaputra/belajar-go-database-migration/app"
	"github.com/dickidarmawansaputra/belajar-go-database-migration/controller"
	"github.com/dickidarmawansaputra/belajar-go-database-migration/helper"
	"github.com/dickidarmawansaputra/belajar-go-database-migration/middleware"
	"github.com/dickidarmawansaputra/belajar-go-database-migration/repository"
	"github.com/dickidarmawansaputra/belajar-go-database-migration/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	// https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md#unversionedd
	// contoh command buat file migration: migrate create -ext sql -dir database/migrations create_category_table
	// command untuk jalankan migration: migrate -database "mysql://dickids:rahasia@tcp(localhost:3306)/belajar-go-database-migration" -path database/migrations up
	// command untuk jalankan rollback: migrate -database "mysql://dickids:rahasia@tcp(localhost:3306)/belajar-go-database-migration" -path database/migrations down
	// command untuk jalankan migration ke versi tertentu: migrate -database "mysql://dickids:rahasia@tcp(localhost:3306)/belajar-go-database-migration" -path database/migrations up 2
	// command untuk jalankan rollback ke versi tertentu: migrate -database "mysql://dickids:rahasia@tcp(localhost:3306)/belajar-go-database-migration" -path database/migrations down 1
	db := app.InitDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.Router(categoryController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
