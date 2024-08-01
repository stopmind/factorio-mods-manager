package factorio

import "strings"

type DependenceType int

const (
	Optional DependenceType = iota
	Required
	Conflict
)

func (t DependenceType) String() string {
	switch t {
	case Optional:
		return "Optional"
	case Required:
		return "Required"
	case Conflict:
		return "Conflict"
	}

	return "Unknown"
}

type Dependence struct {
	Type    DependenceType
	Version Version
	ModID   string
}

func DependenceFromString(raw string) (Dependence, error) {
	clearString := strings.ReplaceAll(raw, "~", "")
	clearString = strings.ReplaceAll(clearString, "(", "")
	clearString = strings.ReplaceAll(clearString, ")", "")

	d := Dependence{
		Type: Required,
	}

	if strings.HasPrefix(clearString, "?") {
		d.Type = Optional
	}

	if strings.HasPrefix(clearString, "!") {
		d.Type = Conflict
	}

	clearString = strings.ReplaceAll(clearString, "?", "")
	clearString = strings.ReplaceAll(clearString, "!", "")

	parts := strings.Split(clearString, ">=")
	if len(parts) == 2 {
		d.ModID = parts[0]
		var err error
		if d.Version, err = VersionFromString(strings.TrimSpace(parts[1])); err != nil {
			return Dependence{}, err
		}
	} else {
		d.ModID = parts[0]
		d.Version = ZeroVersion()
	}

	d.ModID = strings.TrimSpace(d.ModID)

	return d, nil
}
