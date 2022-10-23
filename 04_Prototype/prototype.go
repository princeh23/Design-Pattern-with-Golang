package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type keyWord struct {
	//想序列化字段必须大写，要不然私有，json包获取不到
	Word           string `json:"word"`
	LastUpdateTime time.Time
}

func (w *keyWord) clone() *keyWord {
	//深拷贝方法一：取w原本的地址
	//cloneW := *w
	//return &cloneW

	//深拷贝方法二：序列化反序列化
	var cloneKeyWord keyWord
	fmt.Printf("%+v\n", w)
	b, err := json.Marshal(w)

	if err != nil {
		log.Println(err)
		return nil
	}
	//json.Unmarshal必须为指针
	if err = json.Unmarshal(b, &cloneKeyWord); err != nil {
		log.Println(err)
		return nil
	}
	return &cloneKeyWord
}

func main() {
	protoKeyWord := keyWord{
		Word:           "test",
		LastUpdateTime: time.Now(),
	}
	fmt.Printf("%+v\n", protoKeyWord)
	cloneKeyWord := protoKeyWord.clone()
	fmt.Printf("%+v\n", cloneKeyWord)
}
