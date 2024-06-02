package parse

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_FlagData(t *testing.T) {
	t.Run("MatchName", func(t *testing.T) {
		fd := FlagData{Name: "--flag", Alias: []string{"-f", "--hoge"}}

		assert.True(t, fd.MatchName("--flag"))
		assert.True(t, fd.MatchName("--hoge"))
		assert.False(t, fd.MatchName("--huga"))
	})

	t.Run("HasValueInt", func(t *testing.T) {
		fd1 := FlagData{Name: "--flag", Alias: []string{"-f", "--hoge"}, HasValue: true, Value: "314"}
		fd2 := FlagData{Name: "--flag", Alias: []string{"-f", "--hoge"}, HasValue: true, Value: "hoge"}

		assert.True(t, fd1.HasValueInt())
		assert.False(t, fd2.HasValueInt())
	})

	t.Run("HasValueUint", func(t *testing.T) {
		fd1 := FlagData{Name: "--flag", Alias: []string{"-f", "--hoge"}, HasValue: true, Value: "314"}
		fd2 := FlagData{Name: "--flag", Alias: []string{"-f", "--hoge"}, HasValue: true, Value: "-314"}
		fd3 := FlagData{Name: "--flag", Alias: []string{"-f", "--hoge"}, HasValue: true, Value: "hoge"}

		assert.True(t, fd1.HasValueUint())
		assert.False(t, fd2.HasValueUint())
		assert.False(t, fd3.HasValueUint())
	})

	t.Run("ValueInt", func(t *testing.T) {
		fd1 := FlagData{Name: "--flag", Alias: []string{"-f", "--hoge"}, HasValue: true, Value: "314", DefaultValue: "27"}
		fd2 := FlagData{Name: "--flag", Alias: []string{"-f", "--hoge"}, HasValue: true, Value: "hoge", DefaultValue: "27"}

		assert.Equal(t, int64(314), fd1.ValueInt())
		assert.Equal(t, int64(27), fd2.ValueInt())
	})

	t.Run("ValueUint", func(t *testing.T) {
		fd1 := FlagData{Name: "--flag", Alias: []string{"-f", "--hoge"}, HasValue: true, Value: "314", DefaultValue: "27"}
		fd2 := FlagData{Name: "--flag", Alias: []string{"-f", "--hoge"}, HasValue: true, Value: "-314", DefaultValue: "27"}

		assert.Equal(t, uint64(314), fd1.ValueUint())
		assert.Equal(t, uint64(27), fd2.ValueUint())
	})
}

func Test_Flag(t *testing.T) {
	t.Run("Set", func(t *testing.T) {
		f := Flag{}
		f.Set("--flag", []string{"-f"}).SetHasValue("default")

		assert.Equal(t, 1, len(f))
		assert.Equal(t, "--flag", f[0].Name)
		assert.Equal(t, "-f", f[0].Alias[0])
		assert.True(t, f[0].HasValue)
		assert.Equal(t, "default", f[0].Value)
		assert.Equal(t, "", f[0].Description)
	})

	t.Run("GetIndex", func(t *testing.T) {
		f := Flag{}
		f.Set("--flag1", []string{"-f1"}).SetHasValue("default1")
		f.Set("--flag2", []string{"-f2"})
		f.Set("--flag3", []string{"-f3"}).SetHasValue("default3")

		assert.Equal(t, 0, f.GetIndex("-f1"))
		assert.Equal(t, 1, f.GetIndex("--flag2"))
		assert.Equal(t, 2, f.GetIndex("-f3"))
		assert.Equal(t, -1, f.GetIndex("-f4"))
	})

	t.Run("GetData", func(t *testing.T) {
		f := Flag{}
		f.Set("--flag1", []string{"-f1"}).SetHasValue("default1")
		f.Set("--flag2", []string{"-f2"})
		f.Set("--flag3", []string{"-f3"}).SetHasValue("default3")

		assert.True(t, f.GetData("-f1").HasValue)
		assert.Equal(t, "--flag1", f.GetData("-f1").Name)
		assert.False(t, f.GetData("-f2").HasValue)
		assert.Equal(t, "--flag3", f.GetData("--flag3").Name)
		assert.Equal(t, "", f.GetData("-f5").Name)
	})
}
