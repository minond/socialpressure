package provider

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	GET_TASKS = "https://beta.todoist.com/API/v8/tasks"
)

type Todoist struct {
	Auth
}

type TodoistTask struct {
	Id           int    `json:"id"`
	ProjectId    int    `json:"project_id"`
	Content      string `json:"content"`
	Completed    bool   `json:"completed"`
	Order        int    `json:"order"`
	Indent       int    `json:"indent"`
	Priority     int    `json:"priority"`
	Url          string `json:"url"`
	CommentCount int    `json:"comment_count"`
	Due          struct {
		Recurring bool   `json:"recurring"`
		String    string `json:"string"`
		Date      string `json:"date"`
		Datetime  string `json:"datetime"`
		Timezone  string `json:"timezone"`
	} `json:"due"`
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

func (client Todoist) GetTasks() ([]TodoistTask, error) {
	var tasks []TodoistTask
	res, err := client.Do(client.Request("GET", GET_TASKS))

	if err != nil {
		return tasks, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return tasks, err
	}

	err = json.Unmarshal(body, &tasks)
	return tasks, err
}
