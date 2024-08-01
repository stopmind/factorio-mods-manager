package factorio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Release struct {
	Version         Version
	FactorioVersion Version
	Dependencies    []Dependence
	BaseVersion     Version
}

type rawAPIRelease struct {
	Version string `json:"version"`
	Info    struct {
		FactorioVersion string   `json:"factorio_version"`
		Dependencies    []string `json:"dependencies"`
	} `json:"info_json"`
}

func releaseFromRawAPI(raw *rawAPIRelease) (*Release, error) {
	r := &Release{
		Dependencies: make([]Dependence, 0, len(raw.Info.Dependencies)),
	}

	var err error
	if r.Version, err = VersionFromString(raw.Version); err != nil {
		return nil, err
	}
	if r.FactorioVersion, err = VersionFromString(raw.Info.FactorioVersion); err != nil {
		return nil, err
	}

	for _, rawDependence := range raw.Info.Dependencies {
		var dependence Dependence
		if dependence, err = DependenceFromString(rawDependence); err != nil {
			return nil, err
		}

		if dependence.ModID == "base" {
			r.BaseVersion = dependence.Version
			continue
		}

		r.Dependencies = append(r.Dependencies, dependence)
	}

	return r, nil
}

type Mod struct {
	ID          string
	Description string
	Summary     string
	Title       string
	Author      string

	Releases []*Release
}

func GetMod(id string) (*Mod, error) {
	resp, err := http.Get(fmt.Sprintf("https://mods.factorio.com/api/mods/%v/full", id))

	if err != nil {
		return nil, err
	}

	var raw []byte
	if raw, err = io.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	var rawData struct {
		Description string           `json:"description"`
		Summary     string           `json:"summary"`
		Title       string           `json:"title"`
		Author      string           `json:"author"`
		Releases    []*rawAPIRelease `json:"releases"`
	}
	if err = json.Unmarshal(raw, &rawData); err != nil {
		return nil, err
	}

	mod := &Mod{
		ID:          id,
		Description: rawData.Description,
		Summary:     rawData.Summary,
		Title:       rawData.Title,
		Author:      rawData.Author,
		Releases:    make([]*Release, len(rawData.Releases)),
	}

	for i, rawRelease := range rawData.Releases {
		mod.Releases[i], err = releaseFromRawAPI(rawRelease)

		if err != nil {
			return nil, err
		}
	}

	return mod, nil
}
