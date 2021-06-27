package main

import (
	"fmt"

	"github.com/litethink/huobiapi"
)

func work (chs [] chan  *huobiapi.JSON) [] chan  *huobiapi.JSON{
    task := <- chs
    
    return task
}
func main() {
       chs := make([] chan *huobiapi.JSON,1)
	market, err := huobiapi.NewMarket()
	if err != nil {
		panic(err)
	}
	// 订阅主题 market.ethbtc.kline.1min  market.yfiiusdt.trade.detail
	market.Subscribe("market.yfiiusdt.trade.detail", func(topic string, json *huobiapi.JSON) {
		// 收到数据更新时回调
		fmt.Println(topic, json)
		chs[0] <- json
	})
	 work(chs)
	// 请求数据
	//json, err := market.Request("market.ethbtc.kline.1min")
	//if err != nil {
		//panic(err)
	//} else {
	//	fmt.Println(json)
	//}
	// 进入阻塞等待，这样不会导致进程退出
	market.Loop()
}
