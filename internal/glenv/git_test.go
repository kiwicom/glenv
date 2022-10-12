package glenv

import "testing"

func Test_parseRemoteURL(t *testing.T) {
	host, project := parseRemoteURL("ssh://git@somegitlab.com:mygroup/myrepo.git")
	if host != "somegitlab.com" {
		t.FailNow()
	}

	if project != "mygroup/myrepo.git" {
		t.FailNow()
	}

	host, project = parseRemoteURL("git@somegitlab.com:mygroup/myrepo.git")
	if host != "somegitlab.com" {
		t.FailNow()
	}

	if project != "mygroup/myrepo.git" {
		t.FailNow()
	}
}
