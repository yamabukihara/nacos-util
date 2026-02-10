package main

import (
	"log"
	"os"

	"github.com/yamabuki/nacos-config-tool/cmd"
)

func main() {
	setupLogger()
	cmd.Execute()
}

func setupLogger() {
	// logFileName := "./nacos-util-" + time.Now().Format("2006-01-02") + ".log"
	// logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// if err != nil {
	// 	log.Fatalf("无法创建日志文件：%s : %v", logFileName, err)
	// }
	// log.SetOutput(logFile)
	log.SetOutput(os.Stdout)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
