package tag_soup

import (
  il "github.com/brysgo/gofigure/input_log"
)

type Tag string

type TagGroup il.Input

type TagSoup struct {
  il.InputLog
}

func New() *TagSoup {
  return &TagSoup{}
}

func (self *TagSoup) AddTagGroup(tagGroup TagGroup) *TagSoup {
  self.InputLog = self.InputLog.AddInput(il.Input(tagGroup))
  return self
}

func (self *TagSoup) NearestTag(tag Tag) {
}
