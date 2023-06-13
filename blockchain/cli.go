package blockchain

import (
	"flag"
	"fmt"
	"os"
)

type CommandLine struct {
	bc *Blockchain
}

func NewCommandLine(bc *Blockchain) *CommandLine {
	return &CommandLine{bc: bc}
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  add -data DATA - 添加区块包含 DATA 数据")
	fmt.Println("  print - 打印区块链")
	fmt.Println("  help - 显示帮助信息")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CommandLine) Run() {
	cli.validateArgs()

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printCmd := flag.NewFlagSet("print", flag.ExitOnError)
	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)

	addData := addCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addCmd.Parse(os.Args[2:])
		if err != nil {
			cli.printUsage()
			os.Exit(1)
		}
	case "print":
		err := printCmd.Parse(os.Args[2:])
		if err != nil {
			cli.printUsage()
			os.Exit(1)
		}
	case "help":
		err := helpCmd.Parse(os.Args[2:])
		if err != nil {
			cli.printUsage()
			os.Exit(1)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addCmd.Parsed() {
		if *addData == "" {
			addCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addData)
	}
	if printCmd.Parsed() {
		cli.printChain()
	}
	if helpCmd.Parsed() {
		cli.printUsage()
	}
}

func (cli *CommandLine) addBlock(data string) {
	block,_ := cli.bc.GetBlock(int64(len(cli.bc.Blocks)-1))
	coinbase := NewCoinbaseTransaction("NewBlockPublicKey")
	transactions := append([]Transaction{*coinbase}, Transaction{ID: data})
	index := int64(len(cli.bc.Blocks))
	newBlock := NewBlock(index, block.PrevHash, transactions, 0)
	err := cli.bc.AddBlock(newBlock)
	if err != nil {
		fmt.Println("添加新区块失败：", err)
	} else {
		fmt.Println("添加新区块成功！")
	}
}


func (cli *CommandLine) printChain() {
	for _, block := range cli.bc.Blocks {
		fmt.Printf("============ 区块 %d ============\n", block.Index)
		fmt.Printf("前一个区块哈希：%x\n", block.PrevHash)
		fmt.Printf("当前区块哈希：%x\n", block.Hash)
		fmt.Printf("数据：")
		fmt.Printf("%v\n", block.Transactions)
	}
}
