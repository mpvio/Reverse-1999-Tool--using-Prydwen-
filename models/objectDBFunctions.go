package models

import (
	"golangR99/constants"
)

type ObjectDB interface {
	GetType() string
	GetName() string
}

// characterDB functions
func (c CharacterDB) GetType() string {
	return constants.CHARACTER
}

func (c CharacterDB) GetName() string {
	return c.Name
}

// psychubeDB functions
func (p PsychubeDB) GetType() string {
	return constants.PSYCHUBE
}

func (p PsychubeDB) GetName() string {
	return p.Name
}
