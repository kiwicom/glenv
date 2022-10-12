package glenv

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/kiwicom/glenv/internal/glenv/log"
	"strings"
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

	log.Debug("The repository's origin host:%s, project: %s", host, project)
	return host, project, nil
}

func getRemoteURL(path string) (string, error) {
	log.Debug("get 'origin' URL for repo: %s", path)
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
//
// The parsing itself is naive and straightforward without
// any reg.exp. I realised the reg.exp was confusing for me
// and few more bug make it "complicated"
func parseRemoteURL(URL string) (string, string) {
	log.Debug("Parsing origin URL:%s", URL)

	// normalize URL - remove schema and user
	normalizedURL := URL
	schemeIndex := strings.Index(URL, "://")
	if schemeIndex > 0 {
		normalizedURL = URL[schemeIndex+len("://"):]
	}

	userIndex := strings.Index(normalizedURL, "@")
	if userIndex > 0 {
		normalizedURL = normalizedURL[userIndex+len("@"):]
	}

	// divide URL to host and project
	hostIndex := strings.Index(normalizedURL, ":")
	if hostIndex < 0 {
		hostIndex = strings.Index(normalizedURL, "/")
		if hostIndex < 0 {
			return "", ""
		}
	}

	host := normalizedURL[0:hostIndex]
	project := normalizedURL[hostIndex+1:]

	return host, project
}
