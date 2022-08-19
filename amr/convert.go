package amr

import (
	"log"
	"os/exec"
	"strings"
)

func ConvertMP3(src, dst string) {
	cmd := exec.Command("/bin/bash", "amr/converter.sh", src, dst, "mp3")
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Printf("cmd.StdoutPipe产生的错误:%f", err)
	}
	if err = cmd.Start(); err != nil {
		log.Printf("cmd.Run产生的错误:%f", err)
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		//log.Printf("正在处理第 %d/%d 个文件: %s\n", index+1, total, file)
		s := string(tmp)
		s = strings.Replace(s, "\u0000", "", -1)
		//log.Println(s)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Println("命令执行中有错误产生", err)
	}
}
