package main

import (
	"errors"
	"github.com/afex/hystrix-go/hystrix"
	"log"
	"math/rand"
	"time"
)

func init() {
	hystrix.ConfigureCommand("seckill", hystrix.CommandConfig{
		Timeout: 1, // cmd的超时时间，一旦超时则返回失败
		MaxConcurrentRequests: 5, // 最大并发请求数
		RequestVolumeThreshold: 3, // 熔断探测前的调用次数
		SleepWindow: 1000, // 熔断发生后的等待恢复时间
		ErrorPercentThreshold:10, // 失败占比
	})
}


var Gcount, Gerror int

func TestHystrix() error {
	query := func() error {
		var err error
		r := rand.Float64()
		Gcount++
		if r < 1 {
			err = errors.New("bad luck")
			Gerror++
			return err
		} else {
			time.Sleep(20*time.Microsecond)
		}

		return nil
	}

	var err error
	err = hystrix.Do("seckill", func() error {
		err = query()
		return err
	}, nil)

	return err
}

func main() {

	for i :=0 ; i <100; i++ {
		err := TestHystrix()
		if err != nil  {
			log.Printf("testHystrix error:%v", err)
		}
	}

	log.Printf("Gcount:%d Gerror:%d", Gcount, Gerror)


}
