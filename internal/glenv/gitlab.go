package glenv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type projectJson struct {
	Id   int    `json:"id"`
	Path string `json:"path_with_namespace"`
}

type variableJson struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type groupJson struct {
	Id       int    `json:"id"`
	FullPath string `json:"full_path"`
}

// This function returns you all env. variables for given project and
// also env. variables of all groups the project belong to.
//
// The function is using personal token for auth. with GitLab.
//
// What if I have same variable in group and in project  but with different
// values?
//
// If the env. variable is present in some group and also it's redefined
// with different value in project, the project's value will be used.
//
func GetAllProjectVariables(token string, host string, project string) (map[string]string, error) {
	result := make(map[string]string)
	projectURL := createProjectURL(host, project)

	// get all groups and vars for each project's group
	groups, err := getGroups(token, projectURL)
	if err != nil {
		return result, err
	}

	for _, grp := range groups {
		groupURL := createGroupURL(host, grp.FullPath)
		groupVars, err := getVariables(token, groupURL+"/variables")
		if err != nil {
			// skip the group we cannot read variables for
			errMsg := fmt.Sprintf("%v\n", err)
			os.Stderr.WriteString(errMsg)
		} else {
			addToMap(groupVars, result)
		}
	}

	// get project vars. Why we get project vars after group vars
	// is not accidentally. The project vars have higher priority
	// and will override the group vars.
	projectVars, err := getVariables(token, projectURL+"/variables")
	if err != nil {
		return result, err
	}
	addToMap(projectVars, result)

	// populate CI_PROJECT_* variables
	projectInfo, err := getProjectInfo(token, projectURL)
	if err != nil {
		return result, err
	}

	result["CI_PROJECT_ID"] = strconv.Itoa(projectInfo.Id)
	result["CI_PROJECT_PATH"] = projectInfo.Path

	return result, nil
}

func addToMap(src []variableJson, dest map[string]string) {
	for _, v := range src {
		dest[v.Key] = v.Value
	}
}

// creates `https://{host}/api/v4/groups/{group}` URL, where
// slashes in group name are replaced by `%2F`
func createGroupURL(host string, group string) string {
	escGroup := strings.ReplaceAll(group, "/", "%2F")
	url := fmt.Sprintf("https://%s/api/v4/groups/%s", host, escGroup)
	return url
}

// creates valid `https://{host}/api/v4/projects/{project}` URL
func createProjectURL(host string, project string) string {
	escProject := strings.ReplaceAll(project, "/", "%2F")
	url := fmt.Sprintf("https://%s/api/v4/projects/%s", host, escProject)
	return url
}

// get project's all groups from GitLab
func getGroups(token string, projectURL string) ([]groupJson, error) {

	// prepare request
	req, err := http.NewRequest("GET", projectURL+"/groups", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	// send it
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// parse response
	respBody, _ := ioutil.ReadAll(resp.Body)
	var groups []groupJson = make([]groupJson, 20)
	err = json.Unmarshal(respBody, &groups)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func getProjectInfo(token string, projectURL string) (projectJson, error) {

	var projectInfo projectJson

	// prepare request
	req, err := http.NewRequest("GET", projectURL, nil)
	if err != nil {
		return projectInfo, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	// send it
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return projectInfo, err
	}

	if resp.StatusCode != 200 {
		return projectJson{}, fmt.Errorf("Error in %s (HTTP %d)", projectURL, resp.StatusCode)
	}

	// parse response
	respBody, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respBody, &projectInfo)
	if err != nil {
		return projectInfo, err
	}

	return projectInfo, nil
}

// get the project/group variables from GitLab
func getVariables(token string, variablesURL string) ([]variableJson, error) {

	// prepare request
	req, err := http.NewRequest("GET", variablesURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	// send it
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 403 {
		return nil, fmt.Errorf("You don't have access to %s (HTTP %d)", variablesURL, resp.StatusCode)
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error in %s (HTTP %d)", variablesURL, resp.StatusCode)
	}

	// parse response
	respBody, _ := ioutil.ReadAll(resp.Body)
	var variables []variableJson = make([]variableJson, 20)
	err = json.Unmarshal(respBody, &variables)
	if err != nil {
		return nil, err
	}

	return variables, nil
}
