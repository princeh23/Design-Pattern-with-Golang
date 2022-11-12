package main

import "fmt"

type ICreateServer interface {
	CreateServer(cpu, mem float64) error
}

type AliyunClient struct{}

func (c *AliyunClient) CreateServer(cpu, mem int) error {
	fmt.Printf("aliyun client run success, cpu： %d, mem: %d\n", cpu, mem)
	return nil
}

type AliyunClientAdapter struct {
	Client AliyunClient
}

func (a *AliyunClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.CreateServer(int(cpu), int(mem))
	return nil
}

type AWSClient struct{}

func (c *AWSClient) RunInstance(cpu, mem float64) error {
	fmt.Printf("aws client run success, cpu： %f, mem: %f\n", cpu, mem)
	return nil
}

type AwsClientAdapter struct {
	Client AWSClient
}

func (a *AwsClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.RunInstance(cpu, mem)
	return nil
}

func main() {
	var a ICreateServer = &AliyunClientAdapter{
		Client: AliyunClient{},
	}

	a.CreateServer(1, 2)

	a = &AwsClientAdapter{
		Client: AWSClient{},
	}
	a.CreateServer(1, 2)
}
