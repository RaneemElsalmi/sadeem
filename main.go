package main

import (
    "errors"
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "test/controllers"

    "github.com/go-michi/michi"
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    "github.com/jmoiron/sqlx"
    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

func main() {
    fmt.Println("Starting migration...")

    err := godotenv.Load("./a/.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    migrationsPath := GetRootpath("db/migrations")
    fmt.Println("Migration path: ", migrationsPath)

    mig, err := migrate.New(
        "file://"+migrationsPath,
        os.Getenv("DATABASE_URL"),
    )
    if err != nil {
        log.Fatal(err)
    }

    if err := mig.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
        log.Fatal(err.Error())
    } else if errors.Is(err, migrate.ErrNoChange) {
        log.Printf("No migrations to apply: %s", err.Error())
    }

    controllers.SetDB(db)

    r := michi.NewRouter()
    r.Route(
    	"/",
    	func(sub *michi.Router) {
        r.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
        sub.HandleFunc("GET users", controllers.IndexUserHandler)
        sub.HandleFunc("GET users/{id}", controllers.ShowUserHandler)
		sub.HandleFunc("PUT users/{id}", controllers.UpdateUserHandler)
		// sub.HandleFunc("DELETE users/{id}", controllers.DeleteUserHandler)
        sub.HandleFunc("POST signup", controllers.SignUpHandler)
        sub.HandleFunc("GET vendorsss", controllers.IndexvendorHandler)
        sub.HandleFunc("GET vendors/{id}", controllers.ShowvendorHandler)
    },
    )

    fmt.Println("Server running on port 8080")
    http.ListenAndServe(":8000", r)
}

func GetRootpath(dir string) string {
    wd, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    return filepath.ToSlash(filepath.Join(wd, dir))
}
