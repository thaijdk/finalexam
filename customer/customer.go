package customer

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thaijdk/finalexam/database"
)

type Customer struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func PostHandler(c *gin.Context) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("faltal", err.Error())

	}
	defer db.Close()
	cus := Customer{}
	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if (cus.Name != "") && (cus.Email != "") && (cus.Status != "") {
		query := `
	INSERT INTO customer (name, email, status) VALUES ($1, $2, $3) RETURNING id
	`
		var id int
		row := db.QueryRow(query, cus.Name, cus.Email, cus.Status)
		err = row.Scan(&id)
		if err != nil {
			log.Fatal("can't scan id", err.Error())
		}
		cus.ID = id
		c.JSON(http.StatusCreated, cus)
	} else {
		c.JSON(http.StatusInternalServerError, "Empty Input")
	}
}

func GetByIdHandler(c *gin.Context) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("faltal", err.Error())

	}
	defer db.Close()
	stmt, _ := db.Prepare("SELECT id, name, email, status FROM customer WHERE id=$1")
	id := c.Param("id")
	row := stmt.QueryRow(id)
	cus := Customer{}
	err2 := row.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
	if err2 != nil {
		log.Fatal("error", err.Error())
	}
	c.JSON(http.StatusOK, cus)
}

func GetHandler(c *gin.Context) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("faltal", err.Error())

	}
	defer db.Close()
	stmt, _ := db.Prepare("SELECT id, name, email, status FROM customer")
	customers := []Customer{}
	rows, _ := stmt.Query()
	for rows.Next() {
		cus := Customer{}
		err := rows.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
		if err != nil {
			log.Fatal(err.Error())
		}
		customers = append(customers, cus)
	}
	c.JSON(http.StatusOK, customers)
}

func UpdateHandler(c *gin.Context) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("faltal", err.Error())

	}
	defer db.Close()
	cus := Customer{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cus.ID = id
	stmt, err := db.Prepare("UPDATE customer SET name=$2, email=$3, status=$4 WHERE id= $1;")
	if err != nil {
		log.Fatal("prepare error", err.Error())
	}
	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if (cus.Name != "") && (cus.Email != "") && (cus.Status != "") {
		if _, err := stmt.Exec(cus.ID, cus.Name, cus.Email, cus.Status); err != nil {
			log.Fatal("exec error ", err.Error())
		}
	}
	c.JSON(http.StatusOK, cus)
}

func DeleteByIdHandler(c *gin.Context) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("faltal", err.Error())

	}
	defer db.Close()
	stmt, _ := db.Prepare("DELETE FROM customer WHERE id=$1")
	id := c.Param("id")
	if _, err := stmt.Exec(id); err != nil {
		log.Fatal("exec error ", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}
