package main

import (
	"fmt"
	"time"

	"github.com/Bcly-512/go-iec104/asdu"
	"github.com/Bcly-512/go-iec104/cs104"
)

type myClient struct{}

func main() {
	var err error

	option := cs104.NewOption()
	if err = option.AddRemoteServer("127.0.0.1:2404"); err != nil {
		panic(err)
	}

	mycli := &myClient{}

	client := cs104.NewClient(mycli, option)

	client.LogMode(true)

	client.SetOnConnectHandler(func(c *cs104.Client) {
		c.SendStartDt() // 发送startDt激活指令
	})
	err = client.Start()
	if err != nil {
		panic(fmt.Errorf("Failed to connect. error:%v\n", err))
	}

	for {
		time.Sleep(time.Second * 100)
	}

}
func (myClient) InterrogationHandler(asdu.Connect, *asdu.ASDU) error {
	return nil
}

func (myClient) CounterInterrogationHandler(asdu.Connect, *asdu.ASDU) error {
	return nil
}
func (myClient) ReadHandler(asdu.Connect, *asdu.ASDU) error {
	return nil
}

func (myClient) TestCommandHandler(asdu.Connect, *asdu.ASDU) error {
	return nil
}

func (myClient) ClockSyncHandler(asdu.Connect, *asdu.ASDU) error {
	return nil
}
func (myClient) ResetProcessHandler(asdu.Connect, *asdu.ASDU) error {
	return nil
}
func (myClient) DelayAcquisitionHandler(asdu.Connect, *asdu.ASDU) error {
	return nil
}
func (myClient) ASDUHandler(asdu.Connect, *asdu.ASDU) error {
	return nil
}
