package taskset

import (
	"mdtk/lib"
	"mdtk/taskset/lang"
	"mdtk/taskset/group"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_TaskData(t *testing.T) {
	t.Run("getAttributesFromDesc", func(t *testing.T) {
		type E struct {
			attr []string
			desc string
		}

		tests := lib.TestCases[string, E] {
			{Name: "ok", TestArg: "[attr1 attr2 attr3] description", 
			Expected: E{attr: []string{"attr1", "attr2", "attr3"}, desc: "description"}},
			{Name: "rest is trimed space", TestArg: "[attr1 attr2 attr3]   description    ", 
			Expected: E{attr: []string{"attr1", "attr2", "attr3"}, desc: "description"}},
			{Name: "many spaces", TestArg: "[attr1  attr2   attr3 ] description", 
			Expected: E{attr: []string{"attr1", "attr2", "attr3"}, desc: "description"}},
			{Name: "アルファベット以外もちゃんと取り出せる", TestArg: "[attr1 属性２ attr3] description", 
			Expected: E{attr: []string{"attr1", "属性２", "attr3"}, desc: "description"}},
			{Name: "no attrs", TestArg: "description", 
			Expected: E{attr: []string{}, desc: "description"}},
			{Name: "no attrs but has bracket", TestArg: "[] description", 
			Expected: E{attr: []string{}, desc: "description"}},
			{Name: "no description", TestArg: "[attr1 attr2 attr3] ", 
			Expected: E{attr: []string{"attr1", "attr2", "attr3"}, desc: ""}},
			{Name: "If it is not the front, it remains so.", TestArg: "* [attr1 attr2 attr3] description", 
			Expected: E{attr: []string{}, desc: "* [attr1 attr2 attr3] description"}},
			{Name: "attr:num ok", TestArg: "[attr:1 attr:2 attr:3] ", 
			Expected: E{attr: []string{"attr:1", "attr:2", "attr:3"}, desc: ""}},
		}

		tests.Run(t, func(i int) {
			tt := tests.Get(i)
			td := TaskData{Description: []string{tt.TestArg}}
			attrs, desc := td.getAttributesFromDesc()

			assert.Equal(t, tt.Expected.attr, attrs)
			assert.Equal(t, tt.Expected.desc, desc)
		})
	})

	t.Run("GetAttrsAndSet", func(t *testing.T) {
		type A struct {
			desc string
			group string
			lang string
		}
		type E struct {
			attr []string
			desc string
		}
		tests := lib.TestCases[A, E] {
			{Name: "ok", 
			TestArg: A{desc: "[attr1 attr2 attr3] description"}, 
			Expected: E{attr: []string{"attr1", "attr2", "attr3"}, desc: "description"}},
			{Name: "rest is trimed space", 
			TestArg: A{desc: "[attr1 attr2 attr3]   description    "},
			Expected: E{attr: []string{"attr1", "attr2", "attr3"}, desc: "description"}},
			{Name: "no attrs", 
			TestArg: A{desc: "description"},
			Expected: E{attr: []string{}, desc: "description"}},
			{Name: "no attrs but has bracket", 
			TestArg: A{desc: "[] description"},
			Expected: E{attr: []string{}, desc: "description"}},
			{Name: "no description", 
			TestArg: A{desc: "[attr1 attr2 attr3] "},
			Expected: E{attr: []string{"attr1", "attr2", "attr3"}, desc: ""}},
			{Name: "If it is not the front, it remains so.", 
			TestArg: A{desc: "* [attr1 attr2 attr3] description"},
			Expected: E{attr: []string{}, desc: "* [attr1 attr2 attr3] description"}},
			{Name: "private group has attr 'hidden'", 
			TestArg: A{desc: "[attr1 attr2 attr3] description", group: "_group"}, 
			Expected: E{attr: []string{"attr1", "attr2", "attr3", "hidden"}, desc: "description"}},
			{Name: "no-shell lang has attr 'hidden'", 
			TestArg: A{desc: "[attr1 attr2 attr3] description", group: "_group", lang: "go"}, 
			Expected: E{attr: []string{"attr1", "attr2", "attr3", "hidden"}, desc: "description"}},
		}

		tests.Run(t, func(i int) {
			tt := tests.Get(i)
			if tt.TestArg.group == "" { tt.TestArg.group = "group" }
			if tt.TestArg.lang == "" { tt.TestArg.lang = lang.ShellLangs }

			td := TaskData{
				Description: []string{tt.TestArg.desc}, 
				Group: group.Group(tt.TestArg.group), 
				Lang: lang.Lang(tt.TestArg.lang),
			}
			td.GetAttrsAndSet()

			assert.Equal(t, tt.Expected.attr, td.Attributes)
			assert.Equal(t, tt.Expected.desc, td.Description[0])
		})
	})

	t.Run("HasAttr", func(t *testing.T) {
		type A struct {
			attr []string
			want string
		}
		tests := lib.TestCases[A, bool] {
			{Name: "ok", 
			TestArg: A{attr: []string{"attr1", "attr2", "attr3"}, want: "attr2"}, Expected: true},
			{Name: "bad", 
			TestArg: A{attr: []string{"attr1", "attr2", "attr3"}, want: "yahoooooooo"}, Expected: false},
		}

		tests.Run(t, func(i int) {
			tt := tests.Get(i)
			b := TaskData{Attributes: tt.TestArg.attr}.HasAttr(tt.TestArg.want)
			assert.Equal(t, tt.Expected, b)
		})
	})

	t.Run("HasAttrThatPrefixIs", func(t *testing.T) {
		type A struct {
			attr []string
			want string
		}
		type E struct {
			attr string
			has bool
		}
		tests := lib.TestCases[A, E] {
			{Name: "ok", 
			TestArg: A{attr: []string{"attr1", "attr:2", "attr3"}, want: "attr:"}, 
			Expected: E{attr: "attr:2", has: true}},
			{Name: "bad", 
			TestArg: A{attr: []string{"attr1", "attr:2", "attr3"}, want: "yahoooooooo:"}, 
			Expected: E{has: false}},
		}

		tests.Run(t, func(i int) {
			tt := tests.Get(i)
			s, b := TaskData{Attributes: tt.TestArg.attr}.HasAttrThatPrefixIs(tt.TestArg.want)

			if assert.Equal(t, tt.Expected.has, b) {
				assert.Equal(t, tt.Expected.attr, s)
			}
		})
	})

	t.Run("GetPriority", func(t *testing.T) {
		tests := lib.TestCases[[]string, int] {
			{Name: "No priority attr", TestArg: []string{}, Expected: 0},
			{Name: "Has positive priority attr", TestArg: []string{"priority:5"}, Expected: 5},
			{Name: "Has negative priority attr", TestArg: []string{"priority:-5"}, Expected: -5},
			{Name: "Overflow", TestArg: []string{"priority:10"}, Expected: 0},
			{Name: "Underflow", TestArg: []string{"priority:-10"}, Expected: 0},
			{Name: "Bad priority, not number", TestArg: []string{"priority:abc"}, Expected: 0},
			{Name: "Bad priority, empty", TestArg: []string{"priority:"}, Expected: 0},
		}

		tests.Run(t, func(i int) {
			tt := tests.Get(i)
			ii := TaskData{Attributes: tt.TestArg}.GetPriority()
			assert.Equal(t, tt.Expected, ii)
		})
	})
}
