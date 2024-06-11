package code

import (
	"strings"
    "regexp"
	"mdtk/grtask"
	"mdtk/args"
)

var embed_comment_rex_map = map[string]*regexp.Regexp{}

func init() {
	embed_comment_rex_map["config"] = regexp.MustCompile(getEmbedCommentRegexStr("config"))
	embed_comment_rex_map["embed"] = regexp.MustCompile(getEmbedCommentRegexStr("embed"))
	embed_comment_rex_map["task"] = regexp.MustCompile(getEmbedCommentRegexStr("task"))
	embed_comment_rex_map["func"] = regexp.MustCompile(getEmbedCommentRegexStr("func"))
	embed_comment_rex_map["desc"] = regexp.MustCompile(getEmbedCommentRegexStr("desc"))
	embed_comment_rex_map["args"] = regexp.MustCompile(getEmbedCommentRegexStr("args"))
}

func getEmbedCommentRegexStr(key string) string {
	// #key> comment...
	const comment_str = "[^ \t\n](?:.*[^ \t\n])?"
	var key_str = "(?:#|//)" + key + ">"
	return "(?m)^[ \t]*" + key_str + "[ \t]+(" + comment_str + ")[ \t]*$"
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


func (code Code) GetEmbedCommentText(key string) []string {
	res := []string{}
	embeds := code.GetEmbedComment(key)

	if len(embeds) == 0 {
		return res
	}

	for _, embed := range embeds {
		res = append(res, embed[1])
	}

	return res
}

func (code Code) RemoveEmbedComment(key string) Code {
	embeds := code.GetEmbedComment(key)

	if len(embeds) == 0 {
		return code
	}

	res := string(code)
	for _, embed := range embeds {
		res = strings.Replace(res, embed[0] + "\n", "", 1)
	}

	return Code(res)
}


