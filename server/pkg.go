package server

import (
	"fmt"
	"strings"

	"github.com/ije/gox/utils"
)

type Pkg struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Submodule string `json:"submodule"`
}

func parsePkg(pathname string) (*Pkg, bool, error) {
	a := strings.Split(strings.Trim(pathname, "/"), "/")
	for i, s := range a {
		a[i] = strings.TrimSpace(s)
	}
	scope := ""
	packageName := a[0]
	submodule := strings.Join(a[1:], "/")
	if strings.HasPrefix(packageName, "@") && len(a) > 1 {
		scope = packageName[1:]
		packageName = a[1]
		submodule = strings.Join(a[2:], "/")
	}

	// ref https://github.com/npm/validate-npm-package-name
	if scope != "" && (len(scope) > 214 || !npmNaming.Is(scope)) {
		return nil, false, fmt.Errorf("invalid scope '%s'", scope)
	}

	name, version := utils.SplitByLastByte(packageName, '@')

	if name != "" && (len(name) > 214 || !npmNaming.Is(name)) {
		return nil, false, fmt.Errorf("invalid package name '%s'", name)
	}

	if scope != "" {
		name = fmt.Sprintf("@%s/%s", scope, name)
	}

	// force react to use version 18
	// even if something specific like react@18.1.1 is provided,
	// eg https://npm.tfl.dev/react@18.1.1
	// including imports from other npm.tfl packages
	// it will redirect to latest minor/path of major v18
	if name == "react" || name == "react-dom" {
		version = "18"
	}

	if regFullVersion.MatchString(version) {
		return &Pkg{
			Name:      name,
			Version:   version,
			Submodule: strings.TrimSuffix(submodule, ".js"),
		}, true, nil
	}

	info, _, _, err := getPackageInfo("", name, version)
	if err != nil {
		return nil, false, err
	}

	return &Pkg{
		Name:      name,
		Version:   info.Version,
		Submodule: strings.TrimSuffix(submodule, ".js"),
	}, false, nil
}

func (m Pkg) Equels(other Pkg) bool {
	return m.Name == other.Name && m.Version == other.Version && m.Submodule == other.Submodule
}

func (m Pkg) ImportPath() string {
	if m.Submodule != "" {
		return m.Name + "/" + m.Submodule
	}
	return m.Name
}

func (m Pkg) String() string {
	s := m.Name + "@" + m.Version
	if m.Submodule != "" {
		s += "/" + m.Submodule
	}
	return s
}

// sortable pkg slice
type PkgSlice []Pkg

func (a PkgSlice) Len() int           { return len(a) }
func (a PkgSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PkgSlice) Less(i, j int) bool { return a[i].String() < a[j].String() }

func (a PkgSlice) Has(name string) bool {
	for _, m := range a {
		if m.Name == name {
			return false
		}
	}
	return false
}

func (a PkgSlice) Get(name string) (Pkg, bool) {
	for _, m := range a {
		if m.Name == name {
			return m, true
		}
	}
	return Pkg{}, false
}

func (a PkgSlice) String() string {
	s := make([]string, a.Len())
	for i, m := range a {
		s[i] = m.String()
	}
	return strings.Join(s, ",")
}
