package method

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "Golang-Learning/mini_project/models"

	"io/ioutil"

	"time"
)

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

func Insert(w http.ResponseWriter, r *http.Request) {

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

	query := fmt.Sprintf("Insert into comment (post_id,parent_id,published,content,created_at,updated_at) value (%d,%d,%v,'%s','%s','%s')", comment.PostId, comment.ParentId, comment.Published, comment.Content, currentTime.Format("2006-01-02 15:04:05"), currentTime.Format("2006-01-02 15:04:05"))

	results, err := DbObj.Query(query)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(results)
	}

	res := response{true, "Record Insert"}
	json.NewEncoder(w).Encode(res)
}

func GetResult(w http.ResponseWriter, r *http.Request) {
	DbObj, err := models.DatabaseConnection()

	if err != nil {
		panic("Database error")
	}

	results, err := DbObj.Query("select id, post_id, parent_id, published, content, created_at, updated_at from comment order by id desc")

	if err != nil {
		panic(err.Error())
	}

	comments := []comment{}

	for results.Next() {

		var comment comment

		err := results.Scan(&comment.Id, &comment.PostId, &comment.ParentId, &comment.Published, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}

		comments = append(comments, comment)
	}

	var commentData commentData
	commentData.Comments = comments

	if err != nil {
		fmt.Println("Eror while Marshaling")
	}

	json.NewEncoder(w).Encode(commentData)
}
