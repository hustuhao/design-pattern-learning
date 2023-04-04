

相关代码目录：
```
├── adaptor_pattern # 适配器模式代码目录
│ └── logger_adaptor # 日志适配器目录
│     ├── adapter # 日志适配器的具体实现
│     │ ├── logrus_adapter.go # logrus日志适配器实现
│     │ └── zap_adapter.go    # zap日志适配器实现
│     ├── logger # 日志接口和相关的基础组件
│     │ ├── logger.go
│     │ ├── loghook.go
│     │ └── stack.go
│     └── logger_adptor_test.go # 测试代码
```



适配（Adaptation）是指使不同的事物相互匹配、协调或适应彼此，以满足特定需求或实现特定目标的过程。
在计算机领域中，适配通常指将软件或硬件系统修改为能够与其他系统协同工作或兼容的过程。

生活中的适配器：通常是指能够将不同类型的电子设备接口转换成其他设备所需要的接口类型的电子配件。例如，手机充电器适配器可以将家用电源插头转换为USB接口，以便给手机充电；HDMI转VGA适配器可以将高清晰度数字信号转换为模拟信号以连接老式显示器；同时，还有一些适配器可以将不同国家或地区的电源标准转换成其他国家或地区所需的电源标准，以确保设备在不同地区使用时能够正常工作。
来软件开中中，适配器模式用来不兼容的接口转化为可兼容的接口，将一个类的接口转换成客户端所期望的另一个接口，从而使原本由于接口不匹配而无法一起工作的两个类能够协同工作。这种情况通常出现在需要重复使用现有代码但其接口与所需环境不匹配时，使用适配器模式可以避免修改现有代码并提高代码的复用性。

具体的例子：
- 一个例子是使用适配器将一个第三方库的接口转换成符合自己代码需求的接口。
- 另一个例子是在不改变现有系统的前提下，将某个数据源输出格式转换为另一种数据格式。

适配器模式在实际开发中有很多应用场景，常见的例子：
- 数据库适配器：不同的数据库具有不同的接口和查询语言，将它们转换为统一的接口以提供相同的查询语言和功能。
- 日志适配器：不同的日志库可能具有不同的接口和日志格式，将它们转换为统一的接口以提供相同的日志格式和记录功能。
- 文件系统适配器：不同的操作系统具有不同的文件系统接口，将它们转换为统一的接口以提供相同的文件读写和管理功能。
- 网络协议适配器：不同的网络协议具有不同的接口和通信方式，将它们转换为统一的接口以提供相同的通信方式和数据传输功能。

总之，适配器模式可以用于解决各种接口不兼容的问题，并提高代码的可维护性和复用性。


## 适配器模式实现自定义的 HTTP 客户端
以下是使用适配器模式将一个第三方http库的接口转换成符合自己代码需求的接口的示例代码，很好地说明了适配器模式的基本思想：将一个类的接口转换成另一个类所期望的接口，使得两个不兼容的类可以一起工作。

在这个示例中，我们定义了一个原始 http 库的接口 HttpLibrary 和一个自定义的 http 客户端接口 HttpClient，然后使用适配器模式将 HttpLibrary 适配为 HttpClient。

适配器 HttpLibAdapter 实现了 HttpClient 接口，并在其中包含了一个 HttpLibrary 对象，将原始 http 库的 Get 方法转换为自定义 http 客户端接口的 DoRequest 方法。

当需要使用自定义 http 客户端接口时，我们只需要实例化适配器并调用其 DoRequest 方法即可。

实例代码：
```go
// 原始http库的接口
type HttpLibrary interface {
    Do(req *http.Request) (*http.Response, error)
}

// 假设需要的 http 客户端接口为：
type HttpClient interface {
    DoRequest(method, url string, headers map[string]string, body io.Reader) (*HttpResponse, error)
}

// HttpResponse 是自定义的响应结构体
type HttpResponse struct {
    StatusCode int
    Body       []byte
}

// HttpLibAdapter 是适配器，将 HttpLibrary 转换为 HttpClient 接口
type HttpLibAdapter struct {
    lib HttpLibrary
}

func (a *HttpLibAdapter) DoRequest(method, url string, headers map[string]string, body io.Reader) (*HttpResponse, error) {
    req, err := http.NewRequest(method, url, body)
    if err != nil {
        return nil, err
    }
    for k, v := range headers {
        req.Header.Set(k, v)
    }
    resp, err := a.lib.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    b, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    return &HttpResponse{
        StatusCode: resp.StatusCode,
        Body:       b,
    }, nil
}
```

## 日志适配器

```go
// 统一的日志接口
type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

// logrus 日志适配器
type LogrusAdapter struct {
    logger *logrus.Logger
}

func (a *LogrusAdapter) Debug(msg string) {
    a.logger.Debug(msg)
}
// 省略了实现 Info,Warn,Error 方法

// zap 日志适配器
type ZapAdapter struct {
    logger *zap.Logger
}

func (a *ZapAdapter) Debug(msg string) {
a.logger.Debug(msg)
}
// 省略了实现 Info,Warn,Error 方法

// 使用日志适配器进行日志记录
func DoSomething(logger Logger) {
    // 记录日志
    logger.Debug("doing something")
    // ...
    return
}
```



## 对 SLF4J 的讨论
SLF4J 是一种为 Java 应用程序提供日志记录接口的框架，它提供了一组抽象的接口和类，可以让开发人员在代码中使用通用的、与特定实现无关的日志 API。

SLF4J 本身并不是一种适配器模式的实现，而是通过门面模式（Facade Pattern）来实现统一接口的封装。具体来说，SLF4J 将不同的日志库（如 Logback、Log4j2、JDK Logging 等）的实现封装在不同的适配器（Adapter）中，这些适配器实现了 SLF4J 所定义的接口，从而使得应用程序可以通过 SLF4J 的接口来调用底层日志库的实现。

因此，对于 SLF4J 来说，它不是一个传统意义上的适配器模式的实现，而更倾向于门面模式的实现。但从功能上看，它确实可以被认为是一种日志适配器，因为它将不同的日志库转换为相同的日志接口，达到了适配的效果。








# 参考

[1] The adapter pattern in Go:https://bitfieldconsulting.com/golang/adapter

[2] 设计模式之美:适配器模式：代理、适配器、桥接、装饰，这四个模式有何区别

[3] https://github.com/sirupsen/logrus