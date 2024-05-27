package code

import (
    "regexp"
	"mdtk/grtask"
	"mdtk/args"
)

var embed_comment_rex_map = map[string]*regexp.Regexp{}

func init() {
	embed_comment_rex_map["config"] = regexp.MustCompile(getEmbedCommentRegexStr("config"))
	embed_comment_rex_map["embed"] = regexp.MustCompile(getEmbedCommentRegexStr("embed"))
	embed_comment_rex_map["task"] = regexp.MustCompile(getEmbedCommentRegexStr("task"))
	embed_comment_rex_map["args"] = regexp.MustCompile(getEmbedCommentRegexStr("args"))
}

func getEmbedCommentRegexStr(key string) string {
	// #key> comment...
	const comment_str = "[^ \t\n](?:.*[^ \t\n])?"
	return "(?m)^[ \t]*#" + key + ">[ \t]+(" + comment_str + ")[ \t]*$"
}

type TaskDataSetInterface interface {
	GetTask(grtask.GroupTask, args.Args, bool, bool, int) (Code, error)
}


type Code string

func (code Code) GetEmbedComment(key string) [][]string {
	if _, ok := embed_comment_rex_map[key]; ok {
		res := embed_comment_rex_map[key].FindAllStringSubmatch(string(code), -1)
		return res
	}

	reg := getEmbedCommentRegexStr(key)
	rex := regexp.MustCompile(reg)
	res := rex.FindAllStringSubmatch(string(code), -1)
	return res
}


