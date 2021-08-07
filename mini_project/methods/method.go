package method

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	models "Golang-Learning/mini_project/models"

	"io/ioutil"

	"time"

	"github.com/gorilla/mux"
)

/*type users struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}*/

type comment struct {
	Id        int32  `json:"id"`
	PostId    int32  `json:"postId"`
	ParentId  int32  `json:"parentId"`
	Published bool   `json:"published"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type commentData struct {
	Comments []comment `json:"post"`
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

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	var comment comment
	currentTime := time.Now()

	byteCode := body
	json.Unmarshal(byteCode, &comment)

	//	fmt.Println(comment)

	//query := "Insert into comment (id,post_id,parent_id,published,content,created_at,updated_at) value (" + comment.Id + "," + comment.PostId + "," + comment.ParentId + "," + comment.Published + "," + comment.Content + "," + comment.CreatedAt + "," + comment.UpdatedAt + ")"

	query := fmt.Sprintf("Insert into comment (post_id,parent_id,published,content,created_at,updated_at) value (%d,%d,%v,'%s','%s','%s')", comment.PostId, comment.ParentId, comment.Published, comment.Content, currentTime.Format("2006-01-02 15:04:05"), currentTime.Format("2006-01-02 15:04:05"))

	//results, err := DbObj.Query("Insert into comment (post_id,parent_id,published,content,created_at,updated_at) value (1,1,0,'test','2021-07-26','2021-07-26')")

	results, err := DbObj.Query(query)

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

func recordSelect(w http.ResponseWriter, r http.Request) {

	DbObj, err := models.DatabaseConnection()

	if err != nil {
		panic("Database error")
	}

	results, err := DbObj.Query("select id, post_id, parent_id, published, content, created_at, updated_at from comment order by id desc")

	if err != nil {
		panic(err.Error())
	}

	comments := []comment{}
	//var commentData commentData

	for results.Next() {

		var comment comment

		err := results.Scan(&comment.Id, &comment.PostId, &comment.ParentId, &comment.Published, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}

		comments = append(comments, comment)

		//commentData.Comments = append(commentData.Comments, comment)

		//w.WriteHeader(http.StatusOK)

	}

	var commentData commentData
	commentData.Comments = comments
	//data, err := json.Marshal(comments)

	if err != nil {
		fmt.Println("Eror while Marshaling")
	}

	//fmt.Println(string(data))
	json.NewEncoder(w).Encode(commentData)

	//data['post'] = comments
	//return json{data}
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
