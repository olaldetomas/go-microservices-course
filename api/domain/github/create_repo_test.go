package github

import ()

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := &CreateRepoRequest{
		Name:        "go-introduction",
		Description: "A golang introduction repository",
		Homepage:    "http://github.com",
		Private:     true,
		HasIssues:   false,
		HasProjects: true,
		HasWiki:     false,
	}

	// MArshal takes an input interface and attempts to create a valid json string
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	fmt.Println(string(bytes))
}
