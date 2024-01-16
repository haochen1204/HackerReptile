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
## 计划任务

适配nuclei模板：

- [x] 增加屏幕截图操作
- [ ] 增加执行javascript操作
- [ ] 增加右键点击操作
- [ ] 增加select选择器对html输入执行选择
- [ ] 增加上传文件操作
- [ ] 增加获取src属性操作
- [ ] 增加设置请求方式操作
- [ ] 增加设置请求头操作
- [ ] 增加删除请求头操作
- [ ] 增加设置包体操作

增加功能：

- [ ] 增加debug选项，开启debug模式时会出现网页，并延迟操作
- [ ] 增加从文件中读取数据，配合yaml表格进行写入功能，如批量读取excel中的账号密码进行登录
- [ ] 增加输出到文件的操作，暂定为输入excel格式
- [ ] 增加循环，可以根据在html中获取的数据或用户导入的数据重复某一操作不断获取数据
- [ ] 增加多线程操作

思考ing:

- 增加nuclei模板中的匹配器(matchers)
- 增加nuclei中的提取器(extractors)
- 增加nuclei中的变量器(variables)

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
  Description: 说明....
```

### 动作

```
navigate                # 访问
    Args: 
        URL             # 访问的地址
    
text                    # 输入内容
    Args:
        by: xpath       # 进行匹配的规则
        xpath:          # xpath
        value:          # 输入的值
    
client                  # 点击元素
    Args:
        by:             # 进行匹配的规则
        xpath:          # xpath
  
keyboard                # 模拟键盘按键
    Args:
        keys:           # 模拟的按键

sleep                   # 等待
    Args:
        value:          # 等待的时间
        
waitload                # 等待加载
    Args:
        by: xpath       # 进行匹配的规则
        xpath:          # 等待加载的元素路径
        
extract                 # 获取内容
    Args:
        by: xpath        # 进行匹配的规则
        xpath:           # xpath
        target:          # 获取的内容，如属性或text内容
        attribute:       # 具体的属性内容如href
        name:            # 存放的参数名称
screenshot               # 截图
    Args:
        to:              # 指定存放截图的文件
```

### 支持获取元素的方法

```
by: xpath
xpath:

by: id
id:

by: jspath
jspath: 

by: nodeid
nodeid:

by: query
query:

by: queryall
queryall:
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
      - action: sleep
        args:
          value: "1"
      - args:
          by: xpath
          xpath: /html/body/div[2]/div[1]/div[5]/div/div/form/span[1]/input
          value: haochen1204
        action: text
      - action: keyboard
        args:
          keys: "\r"
      - action: sleep
        args:
          value: "1"
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
      - args:
          by: jspath
          jspath: document.querySelector("#\\33  > div > div:nth-child(1) > h3 > a")
          target: attribute
          attribute: href
        name: info3
        action: extract

```