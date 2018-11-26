/*
type Cmd struct {  
    Path         string　　　//运行命令的路径，绝对路径或者相对路径  
    Args         []string　　 // 命令参数  
    Env          []string         //进程环境，如果环境为空，则使用当前进程的环境  
    Dir          string　　　//指定command的工作目录，如果dir为空，则comman在调用进程所在当前目录中运行  
    Stdin        io.Reader　　//标准输入，如果stdin是nil的话，进程从null device中读取（os.DevNull），stdin也可以时一个文件，否则的话则在运行过程中再开一个goroutine去  
　　　　　　　　　　　　　//读取标准输入  
    Stdout       io.Writer       //标准输出  
    Stderr       io.Writer　　//错误输出，如果这两个（Stdout和Stderr）为空的话，则command运行时将响应的文件描述符连接到os.DevNull  
    ExtraFiles   []*os.File 　　  
    SysProcAttr  *syscall.SysProcAttr  
    Process      *os.Process    //Process是底层进程，只启动一次  
    ProcessState *os.ProcessState　　//ProcessState包含一个退出进程的信息，当进程调用Wait或者Run时便会产生该信息．  
} 
*/

package main

import(
	"github.com/urfave/cli"
	"os"
	"fmt"
)

func main(){

	app := cli.NewApp()

	app.Commands = []cli.Command{
		commandA,
		commandB,
	}
	app.Before = func(c *cli.Context) error {
		return nil
	}
				
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error\n")
	}
	

}

var commandA = cli.Command{
	Name: 	"cmdA",
	ShortName: 	"a",
	Usage: 	"To print AAAA",
	Subcommands: []cli.Command{
		{
			Name:	"subA1",
			Usage:	"to print subA1",
			Action:	func_subA1,
		},
		{
			Name:	"subA2",
			Usage:	"to print subA2",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "StringA",
					Value: "valueA",
					Usage: "The usage of StringA",
				},
			},	
			Action:	func_subA2,
		},
	},
	
	Action: func( c *cli.Context) error {
		fmt.Println("This is TestA error becaus no sub command\n")
		return nil
	},
}

func func_subA1(c *cli.Context){
	fmt.Println("Call subA1 successfully!\n")
	fmt.Printf("subA1 input: %s\n", c.Args().Get(0))  //Is not println
}

func func_subA2(c *cli.Context){
	fmt.Println("Call subA1 successfully!\n")
	fmt.Printf("subA2 input: %s\n", c.String("StringA")) 
}

var commandB = cli.Command{
	Name: 	"cmdB",
	ShortName: 	"b",
	Usage: 	"The usage of cmdB",
	
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "StringB",
			Value: "valueB",
			Usage: "The usage of StringB",
		},
	},
	Action: func( c *cli.Context) error {
		fmt.Println("This is Test B error")
		fmt.Printf("But still test FlagB: %s \n", c.String("StringB"))
		return nil
	},
}
