package config

import (
	"strings"
)

const EnvVar = "ProENV"

type Environment struct {
	str, major, minor string
}

// str is of the form:  [major-]minor
func NewEnv(str string) Environment {
	env := Environment{str: str}
	if i := strings.IndexByte(str, '.'); i >= 0 {
		env.major, env.minor = str[:i], str[i+1:]
	} else {
		env.minor = str
	}
	return env
}

func (e *Environment) String() string {
	return e.str
}

func (e *Environment) Major() string {
	return e.major
}

func (e *Environment) MajorIs(v string) bool {
	return e.major == v
}

func (e *Environment) Minor() string {
	return e.minor
}

func (e *Environment) MinorIs(v string) bool {
	return e.minor == v
}

func (e *Environment) IsDev() bool {
	return e.minor == `dev`
}

func (e *Environment) IsTest() bool {
	return e.minor == `test`
}

func (e *Environment) IsCI() bool {
	return e.minor == `ci`
}

func (e *Environment) IsQA() bool {
	return e.minor == `qa`
}

func (e *Environment) IsPreview() bool {
	return e.minor == `preview`
}

func (e *Environment) IsProduction() bool {
	return e.minor == `production`
}

func (e *Environment) ConfigDir() string {
	name := "config"
	if e.major != "" {
		name += "_" + e.major
	}
	return name
}

func (e *Environment) Vars() []string {
	vars := []string{EnvVar + `=` + e.str}
	if e.major != "" {
		vars = append(vars, EnvVar+`_Major=`+e.major)
	}
	if e.minor != "" {
		vars = append(vars, EnvVar+`_Minor=`+e.minor)
	}
	if e.major != "" {
		vars = append(vars, `ProConfigDir=`+e.ConfigDir())
	}
	return vars
}
