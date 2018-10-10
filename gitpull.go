package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const (
	//配置文件
	CONFFILE string = "/home/wuxiaoyong/wwwroot/git/pullconf.txt"
)

func main() {

	//读取文件的信息
	bytes, err := ioutil.ReadFile(CONFFILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("****In execution, please waiting****")
	//按照换行符分割
	text := string(bytes)
	cmdarr := strings.Split(text, "\r\n")

	pwd, _ := os.Getwd()
	for _, val := range cmdarr {
		tmpval := strings.TrimSpace(val)
		//切换到当前项目目录
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

	fmt.Println("****push ok****")


}

//执行命令函数
//commandName 命名名称，如cat，ls，git等
//params 命令参数，如ls -l的-l，git log 的log等

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)
	_, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return false
	}
	//cmd.Start 与 cmd.Wait 必须一起使用。
	//cmd.Start 不用等命令执行完成， 就结束
	//cmd.Wait 等待命令结束
	cmd.Start()
	cmd.Wait()
	return true
}