package reptile

import (
	"HackerReptile/templateType"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/spf13/viper"
	"github.com/thep0y/go-logger/log"
	"strconv"
	"time"
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

	for _, stepsVal := range steps {
		switch stepsVal.Action {
		case "navigate":
			actions = append(actions, chromedp.Navigate(stepsVal.Args.URL))
		case "click":
			actions = append(actions, chromedp.Click(stepsVal.Args.Xpath, chromedp.BySearch))
		case "text":
			actions = append(actions, chromedp.SendKeys(stepsVal.Args.Xpath, stepsVal.Args.Value, chromedp.BySearch))
		case "sleep":
			times, err := strconv.ParseInt(stepsVal.Args.Value, 10, 64)
			if err != nil {
				log.Error(err)
			}
			actions = append(actions, chromedp.Sleep(time.Duration(times)*time.Second))
		case "waitload":
			actions = append(actions, chromedp.WaitVisible(stepsVal.Args.Xpath, chromedp.BySearch))
		case "extract":
			var tmp string
			actions = append(actions, chromedp.AttributeValue(stepsVal.Args.Xpath, stepsVal.Args.Attribute, &tmp, nil, chromedp.BySearch))
			variables[stepsVal.Name] = &tmp
		case "keyboard":
			actions = append(actions, chromedp.KeyEvent(GetKbKey(stepsVal.Args.Keys)))
		}
	}
	actions = append(actions, chromedp.Sleep(10*time.Second))
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
