package main

import "fmt"

type InputText struct {
	content string
}

func (in *InputText) Append(content string) {
	in.content += content
}

func (in *InputText) GetText() string {
	return in.content
}

func (in *InputText) Snapshot() *Snapshot {
	return &Snapshot{content: in.content}
}

func (in *InputText) Restore(s *Snapshot) {
	in.content = s.GetText()
}

type Snapshot struct {
	content string
}

func (s *Snapshot) GetText() string {
	return s.content
}

func main() {
	c := &InputText{}
	c.Append("ABC")
	fmt.Println(c.GetText())
	snap := c.Snapshot()
	c.Append("BCD")
	fmt.Println(c.GetText())
	c.Restore(snap)
	fmt.Println(c.GetText())
}
