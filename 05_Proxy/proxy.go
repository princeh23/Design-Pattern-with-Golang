package main

import "fmt"

type payment interface {
	pay()
}

type WXPayment struct {
}

func (p WXPayment) pay() {
	fmt.Println("WX pay successful")
}

type AliPayment struct {
}

func (p AliPayment) pay() {
	fmt.Println("Ali pay successful")
}

type paymentProxy struct {
	realPaymentWay payment
}

func (p paymentProxy) pay() {
	fmt.Println("检查格式...")
	fmt.Println("记录日志...")
	fmt.Println("准备支付...")
	p.realPaymentWay.pay()
	fmt.Println("支付完成")
}
func main() {
	proxy := paymentProxy{}
	proxy.realPaymentWay = WXPayment{}
	proxy.pay()
	proxy.realPaymentWay = AliPayment{}
	proxy.pay()
}
