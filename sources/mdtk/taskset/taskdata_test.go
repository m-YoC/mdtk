package taskset

import (
	"mdtk/taskset/group"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_TaskData(t *testing.T) {
	t.Run("LangIsContainedIn", func(t *testing.T) {
		tests := []struct {
			name string
			lang string
			list []string
			expected bool
		} {
			{"ok", "huga", []string{"hoge", "huga", "piyo"}, true},
			{"bad", "fizz", []string{"hoge", "huga", "piyo"}, false},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				td := TaskData{Lang: tt.lang}

				assert.Equal(t, tt.expected, td.LangIsContainedIn(tt.list))
			})
		}
	})


	t.Run("getAttributesFromDesc", func(t *testing.T) {
		tests := []struct {
			name string
			accepted string
			expected []string
			restdesc string
		} {
			{"ok", "[attr1 attr2 attr3] description", []string{"attr1", "attr2", "attr3"}, "description"},
			{"rest is trimed space", "[attr1 attr2 attr3]   description    ", []string{"attr1", "attr2", "attr3"}, "description"},
			{"no attrs", "description", []string{}, "description"},
			{"no attrs but has bracket", "[] description", []string{}, "description"},
			{"no description", "[attr1 attr2 attr3] ", []string{"attr1", "attr2", "attr3"}, ""},
			{"If it is not the front, it remains so.", "* [attr1 attr2 attr3] description", []string{}, "* [attr1 attr2 attr3] description"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				td := TaskData{Description: []string{tt.accepted}}
				attrs, desc := td.getAttributesFromDesc()

				assert.Equal(t, tt.expected, attrs)
				assert.Equal(t, tt.restdesc, desc)
			})
		}
	})

	t.Run("GetAttrsAndSet", func(t *testing.T) {
		tests := []struct {
			name string
			accepted string
			expected []string
			restdesc string
			group string
			lang string
		}{
			{"ok", "[attr1 attr2 attr3] description", []string{"attr1", "attr2", "attr3"}, "description", "group", ShellLangs},
			{"rest is trimed space", "[attr1 attr2 attr3]   description    ", []string{"attr1", "attr2", "attr3"}, "description", "group", ShellLangs},
			{"no attrs", "description", []string{}, "description", "group", ShellLangs},
			{"no attrs but has bracket", "[] description", []string{}, "description", "group", ShellLangs},
			{"no description", "[attr1 attr2 attr3] ", []string{"attr1", "attr2", "attr3"}, "", "group", ShellLangs},
			{"If it is not the front, it remains so.", "* [attr1 attr2 attr3] description", []string{}, "* [attr1 attr2 attr3] description", "group", ShellLangs},
			{"private group has attr 'hidden'", "[attr1 attr2 attr3] description", []string{"attr1", "attr2", "attr3", "hidden"}, "description", "_group", ShellLangs},
			{"no-shell lang has attr 'hidden'", "[attr1 attr2 attr3] description", []string{"attr1", "attr2", "attr3", "hidden"}, "description", "group", "go"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				td := TaskData{Description: []string{tt.accepted}, Group: group.Group(tt.group), Lang: tt.lang}
				td.GetAttrsAndSet()

				assert.Equal(t, tt.expected, td.Attributes)
				assert.Equal(t, tt.restdesc, td.Description[0])
			})
		}
	})

	t.Run("HasAttr", func(t *testing.T) {
		tests := []struct {
			name string
			data []string
			accepted string
			expected bool
		}{
			{"ok", []string{"attr1", "attr2", "attr3"}, "attr2", true},
			{"bad", []string{"attr1", "attr2", "attr3"}, "yahoooooooo", false},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				td := TaskData{Attributes: tt.data}

				assert.Equal(t, tt.expected, td.HasAttr(tt.accepted))
			})
		}
	})
}
