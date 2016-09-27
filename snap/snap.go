package snap

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/fsouza/go-dockerclient"
)

const SNAP_IMAGE = "capture"

type Snap struct {
	url       string
	dir       string
	dockerURI string
}

func NewSnap(url string, dockerURI string) *Snap {
	return &Snap{
		url:       url,
		dockerURI: dockerURI,
	}
}

func (s *Snap) Snap() io.Reader {
	return nil
}

func (s *Snap) Clear() {
	if s.dir != "" {
		os.RemoveAll(s.dir)
	}
}

func (s *Snap) takeSnap() (err error) {
	dir, err := ioutil.TempDir("", "snap")
	if err != nil {
		return
	}
	s.dir = dir

	err = s.runSnapDocker()
	return
}

func (s *Snap) snapFilename() string {
	return fmt.Sprintf("%s/snap.png", s.dir)
}

func (s *Snap) runSnapDocker() (err error) {
	client, err := docker.NewClient(s.dockerURI)
	if err != nil {
		return
	}

	container, err := client.CreateContainer(docker.CreateContainerOptions{
		Name: "",
		Config: &docker.Config{
			Image:        SNAP_IMAGE,
			Cmd:          []string{fmt.Sprintf("capture %s %s", s.url, "/tmp/snap/snap.png")},
			OpenStdin:    true,
			StdinOnce:    true,
			AttachStdin:  true,
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
		},
	})

	if err != nil {
		return
	}

	// remove on finish
	defer func() {
		client.RemoveContainer(docker.RemoveContainerOptions{
			ID:    container.ID,
			Force: true,
		})
	}()

	hostConfig := &docker.HostConfig{
		Binds: []string{
			fmt.Sprintf("%s:/tmp/snap", s.dir),
		},
	}

	err = client.StartContainer(container.ID, hostConfig)
	if err != nil {
		return
	}
	_, err = client.WaitContainer(container.ID)
	return
}
