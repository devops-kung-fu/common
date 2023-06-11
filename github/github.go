// Package github provides functions to interact with github
package github

import (
	"encoding/json"
	"fmt"

	"github.com/kirinlabs/HttpRequest"
)

type Response struct {
	URL       string `json:"url"`
	AssetsURL string `json:"assets_url"`
	UploadURL string `json:"upload_url"`
	HTMLURL   string `json:"html_url"`
	TagName   string `json:"tag_name"`
}

// LatestReleaseTag returns the latest released version tag from github, assumes version is a tag
func LatestReleaseTag(owner, repo string) (version string, err error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

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
