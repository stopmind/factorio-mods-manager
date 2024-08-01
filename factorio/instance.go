package factorio

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Instance struct {
	Path    string
	Version Version
}

func NewInstance(path string) (*Instance, error) {
	raw, err := os.ReadFile(filepath.Join(path, "data/base/info.json"))

	if err != nil {
		return nil, err
	}

	var info struct {
		Version string `json:"version"`
	}

	if err := json.Unmarshal(raw, &info); err != nil {
		return nil, err
	}

	i := &Instance{
		Path: path,
	}

	if i.Version, err = VersionFromString(info.Version); err != nil {
		return nil, err
	}

	return i, nil
}
