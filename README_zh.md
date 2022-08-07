# Robi

不少组织都有一些数据处理的需求，大家一般比较习惯用 Excel 来处理文件，然后通过整理成邮件发送给不同的人。

作为 Go 开发者，[RichFrame](https://github.com/richoffice/richframe/blob/main/README_zh.md) 可以帮我们做一些事情。Go 语言的强类型特性，虽然会让一些严谨的应用更加健壮，但对于一些比较随意的小工具来说，则有一点过于繁琐了。所以我尝试结合 [goja](https://github.com/dop251/goja)，实现了一个基于 javascript 的执行平台。这就是 [robi](https://github.com/richoffice/robi)。

目前，通过 robi 我们可以编写一些 JavaScript 脚本，实现 Excel 数据分析和发邮件的功能，未来还会增加许多办公自动化的功能。

## 快速入门

robi 是 Go 语言编写，通过源代码编译需要有 Go 语言环境。

下载 [robi](https://github.com/richoffice/robi) 代码，项目主目录执行：


```
go install .\cmd\robi\
```

确保构建出的 robi 已经在 path 里，执行：

```
robi .\testfiles\robidemo\simple.js .\testfiles\robidemo\src\sample_file.xlsx .\testfiles\robidemo\src\sample_file_out.xlsx
```

可以看到经过简单的数据分析，即可得到结果 Excel 文件。其中数据分析的功能是通过 RichFrame 实现的，大家可以到 [RichFrame](https://github.com/richoffice/richframe/blob/main/README_zh.md) 看一下使用方法。

## 发送邮件功能

还可以通过 Robi 发送邮件按，在 js 所在目录，新建一个 ./conf/conf.json 文件：

```
{
    "mail":{
        "smtp": "smtphz.qiye.163.com",
        "port": "465",
        "user": "****",
        "password": "********"
    }
}
```

以上是 163 企业邮箱的配置，然后在 js 中可以增加：

```
robi.Mailer.Send("", ["xxxx@163.com","xxx@126.com"],["xxx@gmail.com","xxx@139.com"], "测试邮件", "<b>测试</b>", ["atachment.xlsx"])
```

即可实现邮件发送功能。





