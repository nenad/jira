package jiradata

import (
	"strings"
)

// Find will search the transitions for one that matches
// the given name.  It will return a valid trantion that matches
// or nil
func (t Transitions) Find(name string) *Transition {
	name = strings.ToLower(name)
	matches := []Transitions{}
	for _, trans := range t {
		if strings.Compare(strings.ToLower(trans.Name), name) == 0 {
			return trans
		}
		if strings.Contains(strings.ToLower(trans.Name), name) {
			matches = append(matches, trans)
		}
	}
	if len(matches) > 0 {
		return matches[0]
	}
	return nil
}
