package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	Name string
	ID   int
}

var users = []User{

	{ID: 1, Name: "张三"},

	{ID: 2, Name: "李四"},

	{ID: 3, Name: "王五"},
}

func main() {
	//http.HandleFunc("/users", handlerUsersJson)
	//http.ListenAndServe(":8080", nil)

	svr := gin.Default()
	svr.GET("/users", listUsers)
	svr.GET("/users/:id", getUser)
	svr.Run(":8080")
}

func listUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	found := false
	var u User
	for _, user := range users {
		if strings.EqualFold(id, strconv.Itoa(user.ID)) {
			u = user
			found = true
			break
		}
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
	} else {
		c.JSON(http.StatusOK, u)
	}
}

func handlerUsersJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		users, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "{\"message\": \""+err.Error()+"\"}")
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(users))
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "{\"message\": \"not found\"}")
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":

		w.WriteHeader(http.StatusOK)

		fmt.Fprintln(w, "ID:1,Name:张三")

		fmt.Fprintln(w, "ID:2,Name:李四")

		fmt.Fprintln(w, "ID:3,Name:王五")

	default:

		w.WriteHeader(http.StatusNotFound)

		fmt.Fprintln(w, "not found")

	}
}
