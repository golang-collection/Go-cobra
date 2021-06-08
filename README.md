# Go-cobra
[Cobra](https://github.com/spf13/cobra)既是一个用于创建强大的现代CLI应用程序的库，也是一个用于生成应用程序和命令文件的程序。使用该仓库可以很方便的创建命令行应用，本项目是Cobra的一个简单教程。

## 使用方式
### 引入cobra
```go
import "github.com/spf13/cobra"
```
### cobra项目结构
```
appName/
  cmd/
    root.go
  main.go
```
其中root.go类似于构建之后的项目根目录，所以一般用于介绍项目。
```go
var rootCmd = &cobra.Command{
	Use:   "myCobra",
	Short: "MyCobra is a test case",
	Long:  `Have fun`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```
按照上面的代码项目启动后回输出如下所示内容
![](https://tva1.sinaimg.cn/large/008i3skNly1grauc1tdbqj31a20ne0va.jpg)
可以按照需要自己添加其他命令，如version.go与add.go。

还可以配合flag组合使用，更多见[cobra.dev](https://cobra.dev/)

## 参考文献
- [Go常见库](https://segmentfault.com/a/1190000023382214)