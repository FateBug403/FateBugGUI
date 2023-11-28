package api

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"time"
)

type infoCollectionApi struct {
	
}

func (receiver *infoCollectionApi) Start(ctx context.Context) string {
	go func() {
		time.Sleep(10*time.Second)
		log.Println("开始运行线程")
		runtime.EventsOn(ctx,"stop", func(optionalData ...interface{}) {
			log.Println("监听到停止，退出程序")
		})
		log.Println("停止监听")
	}()
	//runtime.LogInfo(ctx,"trojantest")
	//taskQueue, err := global.GVA_RmqC.OpenQueue("infoCollection")
	//err = taskQueue.PublishBytes([]byte("测试一下下"))
	//if err != nil {
	//	log.Println(err)
	//}
	log.Println("退出")
	return "test"
}

var InfoCollectionApi = new(infoCollectionApi)