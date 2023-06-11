// Package github provides functions to interact with github
package github

import (
	"encoding/json"
	"fmt"

	"github.com/kirinlabs/HttpRequest"
)

const githubAPI = "https://api.github.com"

// Response is the response from github
type Response struct {
	TagName string `json:"tag_name"`
}

// LatestReleaseTag returns the latest released version tag from github, assumes version is a tag
func LatestReleaseTag(owner, repo string) (version string, err error) {
	url := fmt.Sprintf("%s/repos/%s/%s/releases/latest", githubAPI, owner, repo)

	req := HttpRequest.NewRequest()

	resp, _ := req.JSON().Get(url)
	defer func() {
		_ = resp.Close()
	}()
	body, _ := resp.Body()
	if resp.StatusCode() == 200 {
		var response Response
		err = json.Unmarshal(body, &response)
		if err == nil {
			return response.TagName, nil
		}
	}
	err = fmt.Errorf("error retrieving latest release for %s/%s (%v)", owner, repo, resp.Response().Status)

	return
}
