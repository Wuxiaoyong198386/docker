package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const (
	//配置文件
	CONFFILE string = "/home/wuxiaoyong/wwwroot/git/pushconf.txt"
)

func main() {

		//读取文件的信息+
		bytes, err := ioutil.ReadFile(CONFFILE)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//按照换行符分割
		text := string(bytes)
		cmdarr := strings.Split(text, "\r\n")

		pwd, _ := os.Getwd()
		for _, val := range cmdarr {
			tmpval := strings.TrimSpace(val)

			//如果是新命令开始，那么是切换目录操作

			os.Chdir(pwd)
			if tmpval != "" {
				//分割命令
				cmdarr := strings.Split(tmpval, " ")
				//添加commit说明，如果为空，默认为“提交说明”
				command := cmdarr[0]
				if cmdarr[1]=="commit" {
					for idx, args := range os.Args {
						if idx==1 && args!= "" {
							cmdarr[3]=args
						}
					}
				}
				//命令参数
				params := cmdarr[1:]
				//执行cmd命令
				execCommand(command, params)
			}
		}
}

//执行命令函数
//commandName 命名名称，如cat，ls，git等
//params 命令参数，如ls -l的-l，git log 的log等

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true
}
