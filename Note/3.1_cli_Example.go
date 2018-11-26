//cli的使用案例

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
