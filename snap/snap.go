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

func (s *Snap) Snap() (err error) {
	err = s.takeSnap()
	if err != nil {
		return
	}
	return
}

func (s *Snap) SnapFilename() string {
	return fmt.Sprintf("%s/snap.png", s.dir)
}

func (s *Snap) SnapReader() (reader io.ReadCloser, size int64, err error) {
	f, err := os.Open(s.SnapFilename())
	if err != nil {
		return
	}
	reader = f

	stat, err := f.Stat()
	if err != nil {
		return
	}

	size = stat.Size()
	return
}

func (s *Snap) Clear() {
	if s.dir != "" {
		os.RemoveAll(s.dir)
	}
}

func (s *Snap) takeSnap() (err error) {
	dir, err := ioutil.TempDir("/tmp", "snap")
	if err != nil {
		return
	}
	s.dir = dir

	err = s.runSnapDocker()
	return
}

func (s *Snap) runSnapDocker() (err error) {
	client, err := docker.NewClient(s.dockerURI)
	if err != nil {
		return
	}

	share := fmt.Sprintf("%s:/tmp/snap", s.dir)
	hostConfig := &docker.HostConfig{
		Binds: []string{
			share,
		},
	}

	container, err := client.CreateContainer(docker.CreateContainerOptions{
		Name: "",
		Config: &docker.Config{
			Image:        SNAP_IMAGE,
			Cmd:          []string{"capture", s.url, "/tmp/snap/snap.png"},
			OpenStdin:    true,
			StdinOnce:    true,
			AttachStdin:  true,
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
		},
		HostConfig: hostConfig,
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

	err = client.StartContainer(container.ID, nil)
	if err != nil {
		return
	}
	_, err = client.WaitContainer(container.ID)
	return
}
