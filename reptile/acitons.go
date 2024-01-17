package reptile

import (
	"HackerReptile/templateType"
	"github.com/chromedp/chromedp"
	"github.com/thep0y/go-logger/log"
	"strconv"
	"time"
)

// ActionHandle
// ************************************
//
//	@Description: 处理用户传入的每一步动作
//	@param s
//	@return chromedp.Action
//
// ************************************
func ActionHandle(s templateType.Step) []chromedp.Action {
	var action []chromedp.Action
	switch s.Action {
	// 处理不需要获取元素的动作
	case "navigate":
		action = append(action, NavigateAciton(s.Args.URL))
	case "keyboard":
		action = append(action, keyboardAction(s.Args.Keys))
	case "sleep":
		action = append(action, sleepAction(s.Args.Value))
	case "screenshot":
		var tmp []byte
		var tmpAddr string
		tmpAddr = s.Args.To + "/" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".png"
		pngs[&tmpAddr] = &tmp
		action = append(action, screenshotAction(&tmp))
	case "script":
		var tmp []byte
		tmps[&tmp] = &tmp
		action = append(action, scriptAction(s.Args.Code, &tmp))
	default:
		// 处理需要获取动作的元素
		t, v := HandleSelector(s.Args)
		switch s.Action {
		case "click":
			action = append(action, ClientAction(t, v))
		case "text":
			action = append(action, textAction(t, s.Args.Value, v))
		case "waitload":
			action = append(action, waitloadAction(t, v))
		case "extract":
			var tmp string
			variables[&s.Name] = &tmp
			action = append(action, extractAction(t, s.Args.Attribute, &tmp, v))
		}
	}
	return action
}

// HandleSelector
// ************************************
//
//	@Description: 处理用户选择获取元素的方式
//	@param s
//	@return func(s *chromedp.Selector)
//
// ************************************
func HandleSelector(s templateType.Args) (string, func(s *chromedp.Selector)) {
	switch s.By {
	case "xpath":
		return s.Xpath, chromedp.BySearch
	case "id":
		return s.Id, chromedp.ByID
	case "jspath":
		return s.Jspath, chromedp.ByJSPath
	case "nodeid":
		return s.Nodeid, chromedp.ByNodeID
	case "query":
		return s.Query, chromedp.ByQuery
	case "queryall":
		return s.Queryall, chromedp.ByQueryAll
	default:
		log.Error("[-] " + s.By + " 不是规定的表达式")
		return "", nil
	}
}

// NavigateAciton
// ************************************
//
//	@Description: 处理nnavigate 访问行为
//	@param url
//	@return chromedp.Action
//
// ************************************
func NavigateAciton(u string) chromedp.Action {
	return chromedp.Navigate(u)
}

// ClientAction
// ************************************
//
//	@Description: 处理点击动作
//	@param s
//	@param f
//	@return chromedp.Action
//
// ************************************
func ClientAction(s string, f func(s *chromedp.Selector)) chromedp.Action {
	return chromedp.Click(s, f)
}

// textAction
// ************************************
//
//	@Description: 处理输入动作
//	@param s
//	@param v
//	@param f
//	@return chromedp.Action
//
// ************************************
func textAction(s string, v string, f func(s *chromedp.Selector)) chromedp.Action {
	return chromedp.SendKeys(s, v, f)
}

// sleepAction
// ************************************
//
//	@Description: 处理等待动作
//	@param v
//	@return chromedp.Action
//
// ************************************
func sleepAction(v string) chromedp.Action {
	times, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Error(err)
	}
	return chromedp.Sleep(time.Duration(times) * time.Second)
}

// waitloadAction
// ************************************
//
//	@Description: 处理等待元素加载动作
//	@param s
//	@param f
//	@return chromedp.Action
//
// ************************************
func waitloadAction(s string, f func(s *chromedp.Selector)) chromedp.Action {
	return chromedp.WaitVisible(s, f)
}

// extractAction
// ************************************
//
//	@Description:处理获取值动作
//	@param s
//	@param a
//	@param v
//	@param f
//	@return chromedp.Action
//
// ************************************
func extractAction(s string, a string, v *string, f func(s *chromedp.Selector)) chromedp.Action {
	return chromedp.AttributeValue(s, a, v, nil, f)
}

// keyboardAction
// ************************************
//
//	@Description: 处理键盘按键动作
//	@param k
//	@return chromedp.Action
//
// ************************************
func keyboardAction(k string) chromedp.Action {
	return chromedp.KeyEvent(GetKbKey(k))
}

// screenshotAction
// ************************************
//
//	@Description: 处理屏幕截图动作
//	@param b
//	@return chromedp.Action
//
// ************************************
func screenshotAction(b *[]byte) chromedp.Action {
	return chromedp.FullScreenshot(b, 100)
}

// scriptAction
// ************************************
//
//	@Description: 处理执行js动作
//	@param c
//	@return chromedp.Action
//
// ************************************
func scriptAction(c string, t *[]byte) chromedp.Action {
	return chromedp.Evaluate(c, t)
}
