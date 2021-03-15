# czkit tools

## App

使用方式：

    czkit app -i 
    czkit app --init

将创建一个模板工程

如果当前目录下没有"cmd"目录，会先创建一个cmd目录，如果当前目录下有"cmd"目录，则在"cmd"目录
下创建一个"test"目录工程:

    test
    ├── Makefile
    ├── config.go
    └── main.go

该工程引用了"github.com/spf13/cobra"，来做flag解析。其中config.go为Toml格式的配置
文件读取。