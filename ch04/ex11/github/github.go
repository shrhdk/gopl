package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const user = "shrhdk"
const repo = "go-training-course"
const baseURL = "https://api.github.com/"
const issueURL = baseURL + "repos/" + user + "/" + repo + "/issues"

func token(user, accessToken string) string {
	return base64.StdEncoding.EncodeToString([]byte(url.QueryEscape(user) + ":" + url.QueryEscape(accessToken)))
}

func fetch(method, url, body string) (int, string, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	if err != nil {
		return 0, "", err
	}

	req.Header.Set("Authorization", "Basic "+token(user, ""))

	res, err := client.Do(req)

	if err != nil {
		res.Body.Close()
		return 0, "", err
	}

	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		res.Body.Close()
		return 0, "", err
	}

	res.Body.Close()

	return res.StatusCode, string(resBody), nil
}

func CreateIssue(title, body string) error {
	param := createIssueParam{
		Title: title,
		Body:  body,
	}

	bytes, err := json.Marshal(param)

	if err != nil {
		return err
	}

	reqBody := string(bytes)

	statusCode, resBody, err := fetch("POST", issueURL, reqBody)

	if err != nil {
		return err
	}

	fmt.Println(resBody)

	if statusCode != http.StatusOK {
		return fmt.Errorf("create request failed: %d", statusCode)
	}

	return nil
}

func ReadIssue(num int) (Issue, error) {
	statusCode, resBody, err := fetch("GET", issueURL+fmt.Sprintf("/%d", num), "")

	if err != nil {
		return Issue{}, err
	}

	if statusCode != http.StatusOK {
		return Issue{}, fmt.Errorf("%v", resBody)
	}

	var issue Issue
	if err := json.Unmarshal([]byte(resBody), &issue); err != nil {
		return Issue{}, err
	}

	return issue, nil
}

func UpdateIssue(num int, title, body string) error {
	param := createIssueParam{
		Title: title,
		Body:  body,
	}

	bytes, err := json.Marshal(param)

	if err != nil {
		return err
	}

	reqBody := string(bytes)

	statusCode, resBody, err := fetch("PATCH", issueURL+fmt.Sprintf("/%d", num), reqBody)

	if err != nil {
		return err
	}

	fmt.Println(resBody)

	if statusCode != http.StatusOK {
		return fmt.Errorf("updated request failed: %d", statusCode)
	}

	return nil
}

func CloseIssue(num int) error {
	param := createIssueParam{
		State: "closed",
	}

	bytes, err := json.Marshal(param)

	if err != nil {
		return err
	}

	reqBody := string(bytes)

	statusCode, resBody, err := fetch("PATCH", issueURL+fmt.Sprintf("/%d", num), reqBody)

	if err != nil {
		return err
	}

	fmt.Println(resBody)

	if statusCode != http.StatusOK {
		return fmt.Errorf("close request failed: %d", statusCode)
	}

	return nil
}

type createIssueParam struct {
	Title     string `json:"title,omitempty"`
	Body      string `json:"body,omitempty"`
	Assignee  string `json:"assignee,omitempty"`
	Milestone int    `json:"milestone,omitempty"`
	State     string `json:"state,omitempty"`
}

type createIssueResponse struct {
	Id int `json:"id"`
}

type Issue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
