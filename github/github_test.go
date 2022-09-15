package github

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestLatestReleaseTag_Success(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.github.com/repos/devops-kung-fu/bomber/releases/latest",
		httpmock.NewBytesResponder(200, testResponse()))

	version, err := LatestReleaseTag("devops-kung-fu", "bomber")
	assert.NoError(t, err)
	assert.Equal(t, "v0.2.1", version)
	httpmock.GetTotalCallCount()
}

func TestLatestReleaseTag_Failure(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.github.com/repos/test/test/releases/latest",
		httpmock.NewBytesResponder(404, []byte("{}")))

	_, err := LatestReleaseTag("test", "test")
	assert.Error(t, err)
	assert.Equal(t, "error retrieving latest release for test/test (404)", err.Error())
	httpmock.GetTotalCallCount()
}

func testResponse() []byte {
	response := `
	{
		"url": "https://api.github.com/repos/devops-kung-fu/bomber/releases/76185923",
		"assets_url": "https://api.github.com/repos/devops-kung-fu/bomber/releases/76185923/assets",
		"upload_url": "https://uploads.github.com/repos/devops-kung-fu/bomber/releases/76185923/assets{?name,label}",
		"html_url": "https://github.com/devops-kung-fu/bomber/releases/tag/v0.2.1",
		"tag_name": "v0.2.1"
	}
	`
	return []byte(response)
}
