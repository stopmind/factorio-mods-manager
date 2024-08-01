package factorio

import (
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major, Minor, Patch int
}

func (v Version) GreeterOrEqualTo(other Version) bool {
	if v.Major != other.Major {
		return v.Major > other.Major
	}
	if v.Minor != other.Minor {
		return v.Minor > other.Minor
	}
	if v.Patch != other.Patch {
		return v.Patch > other.Patch
	}

	return true
}

func (v Version) Equals(other Version) bool {
	return v.Major == other.Major &&
		v.Minor == other.Minor &&
		v.Patch == other.Patch
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func VersionFromString(raw string) (Version, error) {
	parts := strings.Split(raw, ".")

	v := Version{Patch: 0}

	var err error
	if v.Major, err = strconv.Atoi(parts[0]); err != nil {
		return Version{}, err
	}
	if v.Minor, err = strconv.Atoi(parts[1]); err != nil {
		return Version{}, err
	}

	if len(parts) == 3 {
		if v.Patch, err = strconv.Atoi(parts[2]); err != nil {
			return Version{}, err
		}
	}

	return v, nil
}

func ZeroVersion() Version {
	return Version{
		Major: 0,
		Minor: 0,
		Patch: 0,
	}
}
