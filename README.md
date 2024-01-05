# 开发中！！！！！！！

# HackerReptile

一款基于nuclei yaml模版的爬虫，可根据用户编写的yaml模版实现爬虫功能。

旨在帮助渗透测试工作人员在日常渗透测试、攻防演练过程中快速爬取需要的信息

## 技术栈

采用yaml模版编写爬虫的规则

使用chromedp库，所有爬虫都是基于无头爬虫进行

## 使用方法

```
-y  --yaml      选择使用的yaml模版
    --debg      开启debug模式
```

## 模版编写

### 模版头部

与nuclei类似

首先为id

然后为一些信息。例如名称，作者，备注

```yaml
id: dvwa-headless-automatic-login
info:
  name: DVWA Headless Automatic Login
  author: pdteam
  Description: high
```

### 动作

```
navigate                # 访问
    Args: 
        URL             # 访问的地址
    
text                    # 输入内容
    Args:
        by: xpath       # 进行匹配的规则，暂时只支持xpath
        xpath:          # xpath
        value:          # 输入的值
      
keyboard                # 模拟键盘按键
    Args:
        keys:           # 模拟的按键

sleep                   # 等待
    Args:
        value:          # 等待的时间
        
waitload                # 等待加载
    Args:
        by: xpath       # 进行匹配的规则，暂时只支持xpath
        xpath:          # 等待加载的元素路径
        
extract                 # 获取内容
    Args:
        by: xpath        # 进行匹配的规则，暂时只支持xpath
        xpath:           # xpath
        target:          # 获取的内容，如属性或text内容
        attribute:       # 具体的属性内容如href
        name:            # 存放的参数名称
```

### 例子

```yaml
id: dvwa-headless-automatic-login
info:
  name: DVWA Headless Automatic Login
  author: pdteam
  Description: high
headless:
  - steps:
      - args:
          url: https://www.baidu.com
        action: navigate
      - action: waitload
      - args:
          by: xpath
          xpath: /html/body/div[2]/div[1]/div[5]/div/div/form/span[1]/input
          value: haochen1204
        action: text
      - action: keyboard
        args:
          keys: '\r\n'
      - action: waitload
      - args:
          by: xpath
          xpath: /html/body/div[3]/div[3]/div[1]/div[3]/div[1]/div/div[1]/h3/a
          target: attribute
          attribute: href
        name: info
        action: extract
      - args:
          by: xpath
          xpath: /html/body/div[3]/div[3]/div[1]/div[3]/div[4]/div/div[1]/h3/a
          target: attribute
          attribute: href
        name: info2
        action: extract
```