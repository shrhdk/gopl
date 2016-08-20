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

const baseURL = "https://api.github.com"

type Issue struct {
	Title string `json:"title"`
	URL   string `json:"html_url"`
}

func (i Issue) String() string {
	return fmt.Sprintf("%s %s", i.Title, i.URL)
}

type Milestone struct {
	Title string `json:"title"`
	URL   string `json:"html_url"`
}

func (m Milestone) String() string {
	return fmt.Sprintf("%s %s", m.Title, m.URL)
}

type Collaborator struct {
	Login string `json:"login"`
	URL   string `json:"html_url"`
}

func (c Collaborator) String() string {
	return fmt.Sprintf("%s %s", c.Login, c.URL)
}

type Repository struct {
	Owner         string
	Name          string
	Issues        []Issue
	Milestones    []Milestone
	Collaborators []Collaborator
}

func (r Repository) String() string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "%s/%s\n", r.Owner, r.Name)

	fmt.Fprintln(&buf, "Issues:")
	for _, issue := range r.Issues {
		fmt.Fprintf(&buf, "\t%v\n", issue)
	}

	fmt.Fprintln(&buf, "Milestones:")
	for _, milestone := range r.Milestones {
		fmt.Fprintf(&buf, "\t%v\n", milestone)
	}

	fmt.Fprintln(&buf, "Collaborators:")
	for _, collaborator := range r.Collaborators {
		fmt.Fprintf(&buf, "\t%v\n", collaborator)
	}

	return buf.String()
}

func FetchRepository(owner, name, user, accessToken string) (*Repository, error) {
	issues, err := fetchIssues(owner, name, user, accessToken)
	if err != nil {
		return nil, err
	}

	milestones, err := fetchMilestones(owner, name, user, accessToken)
	if err != nil {
		return nil, err
	}

	collaborators, err := fetchCollaborators(owner, name, user, accessToken)
	if err != nil {
		return nil, err
	}

	return &Repository{owner, name, issues, milestones, collaborators}, nil
}

func fetchIssues(owner, name, user, accessToken string) ([]Issue, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/issues", baseURL, owner, name)

	_, body, err := fetch(url, user, accessToken)
	if err != nil {
		return nil, err
	}

	var issues []Issue
	if err := json.Unmarshal([]byte(body), &issues); err != nil {
		return nil, err
	}

	return issues, nil
}

func fetchMilestones(owner, name, user, accessToken string) ([]Milestone, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/milestones", baseURL, owner, name)

	_, body, err := fetch(url, user, accessToken)
	if err != nil {
		return nil, err
	}

	var milestones []Milestone
	if err := json.Unmarshal([]byte(body), &milestones); err != nil {
		return nil, err
	}

	return milestones, nil
}

func fetchCollaborators(owner, name, user, accessToken string) ([]Collaborator, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/collaborators", baseURL, owner, name)

	_, body, err := fetch(url, user, accessToken)
	if err != nil {
		return nil, err
	}

	var collaborators []Collaborator
	if err := json.Unmarshal([]byte(body), &collaborators); err != nil {
		return nil, err
	}

	return collaborators, nil
}

func token(user, accessToken string) string {
	return base64.StdEncoding.EncodeToString([]byte(url.QueryEscape(user) + ":" + url.QueryEscape(accessToken)))
}

func fetch(url, user, accessToken string) (int, string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return 0, "", err
	}

	req.Header.Set("Authorization", "Basic "+token(user, accessToken))

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
