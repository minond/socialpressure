package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/jinzhu/now"
)

const (
	GetTaskURL  = "https://beta.todoist.com/API/v8/tasks/"
	GetTasksURL = "https://beta.todoist.com/API/v8/tasks"
)

type Todoist struct {
	Auth
}

type TodoistTask struct {
	ID           int    `json:"id"`
	ProjectID    int    `json:"project_id"`
	Content      string `json:"content"`
	Completed    bool   `json:"completed"`
	Order        int    `json:"order"`
	Indent       int    `json:"indent"`
	Priority     int    `json:"priority"`
	URL          string `json:"url"`
	CommentCount int    `json:"comment_count"`
	Due          struct {
		Recurring bool        `json:"recurring"`
		String    string      `json:"string"`
		Date      TodoistDate `json:"date"`
		Datetime  time.Time   `json:"datetime"`
		Timezone  string      `json:"timezone"`
	} `json:"due"`
}

type TodoistDate struct {
	time.Time
}

type TodoistQuery struct {
	TaskID  string `json:"task_id"`
	Message string `json:"message"`
	OkToday bool
}

func (td *TodoistDate) UnmarshalJSON(bytes []byte) (err error) {
	// Expected layout: Mon Jan 2 15:04:05 -0700 MST 2006
	// Todoist layout: 2016-09-01
	loc := time.Now().Location()
	td.Time, err = time.ParseInLocation("2006-01-02", strings.Trim(string(bytes), "\""), loc)
	return err
}

func (client Todoist) Do(req *http.Request) (*http.Response, error) {
	httpClient := http.Client{}
	return httpClient.Do(req)
}

func (client Todoist) Request(method, url string) *http.Request {
	req, _ := http.NewRequest(method, url, nil)

	query := req.URL.Query()
	query.Add("token", client.Auth.Token)

	req.URL.RawQuery = query.Encode()
	return req
}

func (client Todoist) GetTask(id string) (TodoistTask, error) {
	var task TodoistTask
	res, err := client.Do(client.Request("GET", GetTaskURL+id))

	err = unmarshal(res, &task, err)
	return task, err
}

func (client Todoist) GetTasks() ([]TodoistTask, error) {
	var tasks []TodoistTask
	res, err := client.Do(client.Request("GET", GetTasksURL))

	err = unmarshal(res, &tasks, err)
	return tasks, err
}

func (client Todoist) Query(query TodoistQuery) (Query, error) {
	task, err := client.GetTask(query.TaskID)

	if err != nil {
		return Query{}, err
	}

	return Query{
		Message: query.Message,
		Ok:      task.Due.Date.After(now.EndOfDay()),
	}, nil
}
