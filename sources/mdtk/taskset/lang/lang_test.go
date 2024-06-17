package lang

import (
	"mdtk/config"
	"mdtk/lib"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Lang(t *testing.T) {
	t.Run("Set", func(t *testing.T) {
		tests := lib.TestCases[string, string] {
			{Name: "If get empty, set ShellLangs", TestArg: "", Expected: ShellLangs},
			{Name: "If get sh (contained in LangAlias), set ShellLangs", TestArg: "sh", Expected: ShellLangs},
			{Name: "If get pwsh (contained in LangAliasPwSh), set PwShLangs", TestArg: "pwsh", Expected: PwShLangs},
			{Name: "If get other lang, set as it is", TestArg: "dockerfile", Expected: "dockerfile"},
		}

		for _, tt := range tests {
			t.Run(tt.Name, func(t *testing.T) {
				var lang Lang
				lang.Set(tt.TestArg)
				assert.Equal(t, tt.Expected, string(lang))
			})
		}
	})

	t.Run("LangX", func(t *testing.T) {
		tests := lib.TestCases[string, int] {
			{Name: "Lang sh, get LANG_SHELL"  , TestArg: "sh"        , Expected: LANG_SHELL},
			{Name: "Lang bash, get LANG_SHELL", TestArg: "bash"      , Expected: LANG_SHELL},
			{Name: "Lang sh, get LANG_PWSH"   , TestArg: "pwsh"      , Expected: LANG_PWSH},
			{Name: "Lang sh, get LANG_SUB"    , TestArg: "dockerfile", Expected: LANG_SUB},
		}

		for _, tt := range tests {
			t.Run(tt.Name, func(t *testing.T) {
				var lang Lang
				lang.Set(tt.TestArg)
				assert.Equal(t, tt.Expected, lang.LangX().Iam())
			})
		}
	})

	t.Run("IsSub (Is it a appropriate complement set?)", func(t *testing.T) {
		tests := config.GetMergedLangAlias()

		for _, tt := range tests {
			t.Run(tt, func(t *testing.T) {
				var lang Lang
				lang.Set(tt)
				assert.Equal(t, lang.IsSub(), lang.LangX().Iam() == LANG_SUB)
			})
		}
	})
}

func Test_splitFirstAndOther(t *testing.T) {
	type S = []string
	type SS struct {
		f string
		o S
	}
	tests := lib.TestCases[S, SS] {
		{Name: "size 0", TestArg: S{}, Expected: SS{f: "echo", o: S{"Bad Exec Command"}}},
		{Name: "size 1", TestArg: S{"bash"}, Expected: SS{f: "bash", o: S{}}},
		{Name: "size 2", TestArg: S{"bash", "-c"}, Expected: SS{f: "bash", o: S{"-c"}}},
		{Name: "size 3", TestArg: S{"bash", "-x", "-c"}, Expected: SS{f: "bash", o: S{"-x", "-c"}}},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			actual_f, actual_o := splitFirstAndOther(tt.TestArg)
			assert.Equal(t, tt.Expected.f, actual_f)
			assert.Equal(t, tt.Expected.o, actual_o)
		})
	}
}


func Test_removeOpC(t *testing.T) {
	type S = []string
	tests := lib.TestCases[S, S] {
		{Name: "Only has -c", TestArg: S{"-c"}, Expected: S{}},
		{Name: "Has -cx", TestArg: S{"-cx"}, Expected: S{"-x"}},
		{Name: "Has -a -cx", TestArg: S{"-a", "-cx"}, Expected: S{"-a", "-x"}},
		{Name: "Has -cx -a", TestArg: S{"-cx", "-a"}, Expected: S{"-cx", "-a"}},
		{Name: "Only has -x", TestArg: S{"-x"}, Expected: S{"-x"}},
		{Name: "Has -Command", TestArg: S{"-Command"}, Expected: S{}},
		{Name: "Has -Command", TestArg: S{"-a", "-Command"}, Expected: S{"-a"}},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			actual := removeOpC(tt.TestArg)
			assert.Equal(t, tt.Expected, actual)
		})
	}
}

