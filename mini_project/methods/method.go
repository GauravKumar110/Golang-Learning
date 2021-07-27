package method

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	models "Golang-Learning/mini_project/models"

	"github.com/gorilla/mux"
)

type users struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type comments struct {
	Id        int32  `json:"id"`
	PostId    int32  `json:"PostId"`
	ParentId  int32  `json:"ParentId"`
	Published bool   `json:"Published"`
	Content   string `json:"Content"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

type response struct {
	Status  bool
	Message string
}

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/blog", recordInsert).Methods("POST")
	myRouter.HandleFunc("/blog", recordSelect).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func recordInsert(w http.ResponseWriter, r *http.Request) {

	DbObj, err := models.DatabaseConnection()

	if err != nil {
		panic("Database error")
	}

	results, err := DbObj.Query("Insert into comment (post_id,parent_id,published,content,created_at,updated_at) value (1,1,0,'test','2021-07-26','2021-07-26')")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(results)
	}

	res := response{true, "Record Insert"}

	/*data, err := json.Marshal(res)

	if err != nil {
		fmt.Println("Eror while Marshaling")
	}*/

	//fmt.Println(string(data))
	json.NewEncoder(w).Encode(res)
}

func recordSelect(w http.ResponseWriter, r *http.Request) {

	DbObj, err := models.DatabaseConnection()

	if err != nil {
		panic("Database error")
	}

	results, err := DbObj.Query("select id, post_id, parent_id, published, content, created_at, updated_at from comment order by id desc")

	if err != nil {
		panic(err.Error())
	}

	commentData := []comments{}

	for results.Next() {

		comment := comments{}

		err := results.Scan(&comment.Id, &comment.PostId, &comment.ParentId, &comment.Published, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}

		commentData = append(commentData, comment)

		w.WriteHeader(http.StatusOK)

	}
	data, err := json.Marshal(commentData)

	if err != nil {
		fmt.Println("Eror while Marshaling")
	}

	fmt.Println(string(data))
	json.NewEncoder(w).Encode(commentData)
}

/*
func homepage1(w http.ResponseWriter, r *http.Request) {

	DbObj, err := models.DatabaseConnection()

	if err != nil {
		panic("Database error")
	}

	fmt.Println("Endpoint")
	//	results, err := DbObj.Query("Select id, email, phone from members order by id desc limit 3,3")

	results, err := DbObj.Query("Insert into comment (post_id,parent_id,published,content,created_at,updated_at) value (1,1,0,'test','2021-07-26','2021-07-26')")

	if err != nil {
		panic(err.Error())
	}

	userData := []users{}

	for results.Next() {
		fmt.Println(results.Scan())
		user := users{}

		err := results.Scan(&user.Id, &user.Email, &user.Phone)
		if err != nil {
			panic(err.Error())
		}

		userData = append(userData, user)

		w.WriteHeader(http.StatusOK)

		//log.Print(user.Id)
		//log.Printf(user.Email)
		//log.Printf(user.Phone)
	}
	data, err := json.Marshal(userData)

	if err != nil {
		fmt.Println("Eror while Marshaling")
	}

	fmt.Println(string(data))
	json.NewEncoder(w).Encode(userData)
}*/
