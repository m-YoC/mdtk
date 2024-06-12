package group

import (
	"fmt"
	"regexp"
	"mdtk/base"
)

var validate_groupname_rex = regexp.MustCompile("^" + base.NameReg + "$")

type Group string

func (g Group) IsPrivate() bool {
	return len(g) >= 2 && g[0:1] == "_"
}

func (g Group) Validate() error {
	if !validate_groupname_rex.MatchString(string(g)) {
		return fmt.Errorf("Validation error: group name. => %s\n", g)
	}
	return nil
}

func (g Group) ValidatePublic() error {
	if err := g.Validate(); err != nil {
		return err
	}

	if g.IsPrivate() {
		return fmt.Errorf("Private group (beginning with an underscore) cannot be executed directly.\n")
	}
	return nil
}

func (g Group) ValidateEmptyIsSafe() error {
	if !validate_groupname_rex.MatchString(string(g)) && g != "" {
		return fmt.Errorf("Validation error: group name. => %s\n", g)
	}
	return nil
}

func (g Group) ValidatePublicEmptyIsSafe() error {
	if err := g.ValidateEmptyIsSafe(); err != nil {
		return err
	}

	if g.IsPrivate() {
		return fmt.Errorf("Private group (beginning with an underscore) cannot be executed directly.\n")
	}
	return nil
}


func (g Group) Match(group Group) bool {
	// If group is empty, return true
	return group == "" || group == g
}
