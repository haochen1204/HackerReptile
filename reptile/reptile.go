package reptile

import (
	"HackerReptile/templateType"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"github.com/thep0y/go-logger/log"
	"time"
)

func CreatActionQuery(steps []templateType.Step) {
	// 定义动作的存储列表
	var actions []chromedp.Action
	// 存储网页中获取的的变量
	variables := make(map[string]interface{})
	for _, stepsVal := range steps {
		switch stepsVal.Action {
		case "navigate":
			actions = append(actions, chromedp.Navigate(stepsVal.Args.URL))
		case "click":
			actions = append(actions, chromedp.Click(stepsVal.Args.Xpath, chromedp.BySearch))
		case "text":
			actions = append(actions, chromedp.SendKeys(stepsVal.Args.Xpath, stepsVal.Args.Value, chromedp.BySearch))
		case "waitload":
			actions = append(actions, chromedp.Sleep(1*time.Second))
		case "extract":
			var tmp string
			actions = append(actions, chromedp.AttributeValue(stepsVal.Args.Xpath, stepsVal.Args.Attribute, &tmp, nil, chromedp.BySearch))
			fmt.Println(tmp)
			variables[stepsVal.Name] = tmp
		case "keyboard":
			actions = append(actions, chromedp.KeyEvent(kb.Enter))
		}
	}
	actions = append(actions, chromedp.Sleep(10*time.Second))
	RunAction(actions)
	fmt.Println(variables)
}

func RunAction(actions []chromedp.Action) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 创建一个带有选项的新context，用于关闭无头模式
	ctx, cancel = chromedp.NewExecAllocator(ctx,
		append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false), // 关闭无头模式
		)...,
	)
	defer cancel()

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

func Test(url string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var data string

	// 创建一个带有选项的新context，用于关闭无头模式
	ctx, cancel = chromedp.NewExecAllocator(ctx,
		append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false), // 关闭无头模式
		)...,
	)
	defer cancel()

	// 创建一个新的chromedp context
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML("html", &data, chromedp.ByQuery),
		chromedp.Sleep(10*time.Second),
	); err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
