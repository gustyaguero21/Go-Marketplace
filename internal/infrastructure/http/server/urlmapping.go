package server

import (
	"database/sql"
	"go-marketplace/cmd/config"
	"log"
	"net/http"

	"go-marketplace/internal/infrastructure/database"
	mysqluser "go-marketplace/internal/infrastructure/database/user"

	"github.com/gin-gonic/gin"
)

func UrlMapping(r *gin.Engine) {

	db, err := database.DatabaseConn()
	if err != nil {
		log.Fatal(err)
	}

	initUserDb(db)

	api := r.Group("/api/go-marketplace")
	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func initUserDb(db *sql.DB) {
	newMysqlUser := mysqluser.NewMysqlUsers(db)
	newMysqlUser.CreateUsersTable(config.UsersTableQuery)
}
