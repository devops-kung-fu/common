package github

import (
	"encoding/json"
	"fmt"

	"github.com/kirinlabs/HttpRequest"
)

type Response struct {
	URL             string `json:"url"`
	AssetsURL       string `json:"assets_url"`
	UploadURL       string `json:"upload_url"`
	HTMLURL         string `json:"html_url"`
	ID              int64  `json:"id"`
	NodeID          string `json:"node_id"`
	TagName         string `json:"tag_name"`
	TargetCommitish string `json:"target_commitish"`
	Name            string `json:"name"`
	Draft           bool   `json:"draft"`
	Prerelease      bool   `json:"prerelease"`
	CreatedAt       string `json:"created_at"`
	PublishedAt     string `json:"published_at"`
	TarballURL      string `json:"tarball_url"`
	ZipballURL      string `json:"zipball_url"`
}

// LatestVersion returns the latest released version tag from github, assumes version is a tag
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
