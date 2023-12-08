package internal

import (
	"bytes"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/labbs/castle/config"
	"github.com/labbs/castle/modules/repository/domain"
	"github.com/pkg/errors"
)

var publicKeys *ssh.PublicKeys = nil
var err error

func CloneRepository(repo domain.Repository, test bool) error {
	if repo.Type == "git" {
		err = CloneGitRepository(repo, test)
		if err != nil {
			return err
		}
	}
	return nil
}

func CloneGitRepository(repo domain.Repository, test bool) error {
	path := fmt.Sprintf("%s/repos/%s", config.AppConfig.TemporaryFolder, repo.Id)
	buf := new(bytes.Buffer)

	gitCloneOptions := &git.CloneOptions{
		URL:      repo.Url,
		Progress: buf,
	}

	if repo.SSHKey != "" {
		publicKeys, err = ssh.NewPublicKeysFromFile("git", repo.SSHKey, repo.SSHKeyPassphrase)
		if err != nil {
			return err
		}

		gitCloneOptions.Auth = publicKeys
	}

	_, err := git.PlainClone(path, false, gitCloneOptions)

	if err != nil {
		os.RemoveAll(path)
		return errors.New("failed to clone repository, error: " + err.Error() + ", output: " + buf.String())
	}

	if test {
		os.RemoveAll(path)
	}

	return nil
}
