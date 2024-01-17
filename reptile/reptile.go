package reptile

import (
	"HackerReptile/templateType"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/spf13/viper"
	"github.com/thep0y/go-logger/log"
	"os"
	"strings"
)

// 存储网页中获取的的变量
var variables = make(map[interface{}]interface{})
var pngs = make(map[interface{}]interface{})
var tmps = make(map[interface{}]interface{})

// CreatActionQuery
//
//	@Description: 创建chromedp任务列表
//	@param steps
func CreatActionQuery(steps []templateType.Step) {
	// 定义动作的存储列表
	var actions []chromedp.Action
	// 循环读取每一步动作，交给处理函数处理，并返回chromedp.aciton
	for _, stepsVal := range steps {
		actions = append(actions, ActionHandle(stepsVal)...)
	}

	// 下面的代码是测试使用
	//actions = append(actions, chromedp.Sleep(10*time.Second))
	RunAction(actions)
	fmt.Println(variables)
	for key, value := range variables {
		a := value
		b := key.(*string)
		if actualPointer, ok := a.(*string); ok {
			fmt.Println(*b, " is: ", *actualPointer)
		}
	}

	for key, value := range pngs {
		a := value
		b := key.(*string)
		path := *b
		lastIndex := strings.LastIndex(path, "/")
		_, err := os.Stat(path[:lastIndex])
		if err != nil {
			err := os.Mkdir(path[:lastIndex], 0755)
			if err != nil {
				log.Fatal(err)
			}
			log.Info("[+] 创建目录 " + path[:lastIndex])
		}
		if actualPointer, ok := a.(*[]byte); ok {
			if err := os.WriteFile(*b, *actualPointer, 0o644); err != nil {
				log.Fatal(err)
			}

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
