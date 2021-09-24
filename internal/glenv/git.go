package glenv

import (
	"fmt"
	"regexp"

	"github.com/go-git/go-git/v5"
)

// This function open the Git repository on given path and returns you the
// origin's host and project name.
//
// The function counts you cloned the repository SSH way (git@server:project).
// If repo is cloned HTTP/HTTPS way, the function ends and returns you empty
// strings and error
//
func GetHostAndProject(path string) (string, string, error) {
	remoteURL, err := getRemoteURL(path)
	if err != nil {
		return "", "", fmt.Errorf("cannot open repository %s", path)
	}

	host, project := parseRemoteURL(remoteURL)
	if host == "" || project == "" {
		return "", "", fmt.Errorf("cannot parse the URL %s", remoteURL)
	}

	return host, project, nil
}

func getRemoteURL(path string) (string, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return "", err
	}

	// get the origin URL
	origin, err := repo.Remote("origin")
	if err != nil {
		return "", err
	}
	URL := origin.Config().URLs[0]

	return URL, nil
}

// parse git@{host}:{project}.git and returns you
// host and project.
func parseRemoteURL(URL string) (string, string) {
	re := regexp.MustCompile(".*@(.*):(.*).git")
	res := re.FindStringSubmatch(URL)
	if len(res) < 3 {
		return "", ""
	}
	return res[1], res[2]
}
