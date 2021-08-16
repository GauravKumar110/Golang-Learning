package method

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	models "Golang-Learning/mini_project/models"

	"io/ioutil"
)

type post struct {
	Id          int32  `json:"id"`
	AuthorId    int32  `json:"author_id"`
	Title       string `json:"title"`
	MetaTitle   string `json:"meta_title"`
	Slug        string `json:"slug"`
	Summary     string `json:"summary"`
	Published   bool   `json:"published"`
	PublishedAt string `json:"published_at"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

type postData struct {
	Comments []post `json:"post"`
}

type response struct {
	Status  bool
	Message string
}

func Create(w http.ResponseWriter, r *http.Request) {

	DbObj, err := models.DatabaseConnection()

	if err != nil {
		panic("Database error")
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	var post post
	currentTime := time.Now()
	json.Unmarshal(body, &post)

	query := fmt.Sprintf("Insert into posts (author_id,title,meta_title,slug,summary,published,published_at,content,created_at,updated_at) value ('%d','%s','%s','%s','%s',%v,'%s','%s','%s','%s')", post.AuthorId, post.Title, post.MetaTitle, post.Slug, post.Summary, post.Published, currentTime.Format("2006-01-02 15:04:05"), post.Content, currentTime.Format("2006-01-02 15:04:05"), currentTime.Format("2006-01-02 15:04:05"))

	results, err := DbObj.Query(query)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(results)
	}

	res := response{true, "Record Insert"}
	json.NewEncoder(w).Encode(res)
}

func GetRecords(w http.ResponseWriter, r *http.Request) {
	DbObj, err := models.DatabaseConnection()

	if err != nil {
		panic("Database error")
	}

	results, err := DbObj.Query("select id, author_id, title, meta_title, slug, summary, published,published_at,content,created_at,updated_at from posts order by id desc")

	if err != nil {
		panic(err.Error())
	}

	posts := []post{}

	for results.Next() {

		var post post

		err := results.Scan(&post.Id, &post.AuthorId, &post.Title, &post.MetaTitle, &post.Slug, &post.Summary, &post.Published, &post.PublishedAt, &post.Content, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}

		posts = append(posts, post)
	}

	var postData postData
	postData.Comments = posts

	if err != nil {
		fmt.Println("Eror while Marshaling")
	}

	json.NewEncoder(w).Encode(postData)
}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {

	DbObj, err := models.DatabaseConnection()

	if err != nil {
		panic("Database error")
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	var post post
	currentTime := time.Now()
	json.Unmarshal(body, &post)

	query := fmt.Sprintf("Update posts set deleted_at = '%s' where id = %d ", currentTime.Format("2006-01-02 15:04:05"), post.Id)

	results, err := DbObj.Query(query)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(results)
	}

	res := response{true, "Record Deleted"}
	json.NewEncoder(w).Encode(res)

}

func UpdateRecord(w http.ResponseWriter, r *http.Request) {

	DbObj, err := models.DatabaseConnection()

	if err != nil {
		panic("Database error")
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	var post post
	currentTime := time.Now()
	json.Unmarshal(body, &post)

	fmt.Println(post)

	if post.Id == 0 {
		panic("Id can not be null")
	}

	if post.Content == "" {
		panic("Content field can not be null")
	}

	query := fmt.Sprintf("Update posts set content = '%s', updated_at = '%s' where id = %d ", post.Content, currentTime.Format("2006-01-02 15:04:05"), post.Id)
	results, err := DbObj.Query(query)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(results)
	}

	res := response{true, "Record Update"}
	json.NewEncoder(w).Encode(res)
}
