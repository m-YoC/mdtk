package code

import (
)

func (code Code) RemoveEmbedDescComment() Code {
	return code.RemoveEmbedComment("desc")
}

func (code Code) GetEmbedDescText() []string {
	return code.GetEmbedCommentText("desc")
}
