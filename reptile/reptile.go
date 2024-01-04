package reptile

import (
	"HackerReptile/templateType"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/thep0y/go-logger/log"
	"time"
)

func CreatActionQuery(steps []templateType.Step) {
	for _, stepsVal := range steps {
		fmt.Println(stepsVal.Action, stepsVal.Args)
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
