package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const base_url = "https://api.github.com"

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func GetIssues(owner, repo string) ([]Issue, error) {
	url := strings.Join([]string{base_url, "repos", owner, repo, "issues"}, "/")
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("get issues failed: %s", resp.Status)
	}

	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return issues, nil
}

func GetIssue(owner, repo, issue_number string) (*Issue, error) {
	url := strings.Join([]string{base_url, "repos", owner, repo, "issues", issue_number}, "/")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get issue failed: %s", resp.Status)
	}

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &issue, nil
}

func CreateIssue(owner, repo, username, token string, fields map[string]string) error {

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(fields)
	if err != nil {
		return err
	}
	url := strings.Join([]string{base_url, "repos", owner, repo, "issues"}, "/")
	request, err := http.NewRequest(http.MethodPost, url, &buf)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(username, token)
	client := http.Client{}
	_, err = client.Do(request)
	return err
}

func UpdateIssue(owner, repo, number, username, token string, fields map[string]string) error {

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(fields)
	if err != nil {
		return err
	}
	url := strings.Join([]string{base_url, "repos", owner, repo, "issues", number}, "/")
	request, err := http.NewRequest(http.MethodPatch, url, &buf)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(username, token)
	client := &http.Client{}
	_, err = client.Do(request)
	return err
}

func main() {
	// todo: use preferred text editor
	issues, err := GetIssues("torbiak", "gopl")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	for _, issue := range issues {
		fmt.Println(issue)
	}

	issue, err := GetIssue("torbiak", "gopl", "1")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(*issue)
	payload := make(map[string]string)
	payload["title"] = "Test create issue for 4.11"
	payload["body"] = "RT"
	err = UpdateIssue("jukanntenn", "gopl-exercises", "1", "jukanntenn", "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", payload)
	if err != nil {
		fmt.Println(err)
	}

	err = CreateIssue("jukanntenn", "gopl-exercises", "jukanntenn", "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", payload)
	if err != nil {
		fmt.Println(err)
	}
}
