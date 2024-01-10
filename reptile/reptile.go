package reptile

import (
	"HackerReptile/templateType"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/spf13/viper"
	"github.com/thep0y/go-logger/log"
)

// 存储网页中获取的的变量
var variables = make(map[string]interface{})

// CreatActionQuery
//
//	@Description: 创建chromedp任务列表
//	@param steps
func CreatActionQuery(steps []templateType.Step) {
	// 定义动作的存储列表
	var actions []chromedp.Action
	// 循环读取每一步动作，交给处理函数处理，并返回chromedp.aciton
	for _, stepsVal := range steps {
		actions = append(actions, ActionHandle(stepsVal))
	}
	// 下面的代码是测试使用
	//actions = append(actions, chromedp.Sleep(10*time.Second))
	RunAction(actions)
	fmt.Println(variables)
	for key, value := range variables {
		a := value
		if actualPointer, ok := a.(*string); ok {
			fmt.Println(key, " is: ", *actualPointer)
		}
	}
}

// RunAction
//
//	@Description: 开启chromdp任务
//	@param actions
func RunAction(actions []chromedp.Action) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	if viper.GetBool("debug") {
		// 创建一个带有选项的新context，用于关闭无头模式
		ctx, cancel = chromedp.NewExecAllocator(ctx,
			append(chromedp.DefaultExecAllocatorOptions[:],
				chromedp.Flag("headless", false), // 关闭无头模式
			)...,
		)
		defer cancel()
	}

	// 创建一个新的chromedp context
	ctx, cancel = chromedp.NewContext(ctx)

	// 处理任务
	tasks := chromedp.Tasks(actions)
	// 执行任务
	err := chromedp.Run(ctx, tasks)
	if err != nil {
		log.Fatal(err)
	}
}
