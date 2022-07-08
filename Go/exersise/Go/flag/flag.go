package main

import (
	"flag"
	"fmt"
)

func main() {
	backup_dir := flag.String("b", "/home/default_dir", "backup path")
	debug_mode := flag.Bool("d", false, "debug mode")

	flag.Parse()

	fmt.Println("backup_dir: ", *backup_dir)
	fmt.Println("debug_mode: ", *debug_mode)
}

/**
  启动命令：go run main.go -b /home/backup
  输出结果：
      backup_dir:  /home/backup  // 因为指定了路径，所以覆盖了默认路径
      debug_mode:  false         // 启动命令里不带 -d 参数，所以启用默认值
*/
