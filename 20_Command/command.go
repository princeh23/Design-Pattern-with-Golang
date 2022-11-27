package main

import (
	"fmt"
	"time"
)

type ICommand interface {
	Execute() error
}

type StartCommand struct{}

func NewStartCommand() *StartCommand {
	return &StartCommand{}
}

func (c *StartCommand) Execute() error {
	fmt.Println("start")
	return nil
}

type ArchiveCommand struct{}

func NewArchiveCommand() *ArchiveCommand {
	return &ArchiveCommand{}
}

func (c *ArchiveCommand) Execute() error {
	fmt.Println("archive")
	return nil
}

func main() {
	//总体：正在实现一个后端
	//使用一个 goroutine 不断接收来自客户端请求的命令，并且将它放置到一个队列当中
	//然后我们在另外一个 goroutine 中来执行它

	eventChan := make(chan string)
	go func() {
		events := []string{"start", "archive", "start", "archive", "start", "start"}
		for _, e := range events {
			eventChan <- e
		}
	}()
	defer close(eventChan)

	// 使用命令队列缓存命令
	commands := make(chan ICommand, 1000)
	defer close(commands)

	go func() {
		for {
			// 从请求或者其他地方获取相关事件参数
			event, ok := <-eventChan
			if !ok {
				return
			}

			var command ICommand
			switch event {
			case "start":
				command = NewStartCommand()
			case "archive":
				command = NewArchiveCommand()
			}

			// 将命令入队
			commands <- command
		}
	}()

	for {
		select {
		case c := <-commands:
			c.Execute()
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1s")

		}
	}
}
