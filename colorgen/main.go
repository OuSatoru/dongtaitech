package main

import (
	"bufio"
	"fmt"
	"os"
)

type RGB struct {
	Red   int
	Green int
	Blue  int
}

func main() {
	fmt.Println("依次输入RGB值，每输一个按一次回车。")
	stdin := bufio.NewReader(os.Stdin)
	var r, g, b int
	fmt.Fscan(stdin, &r)
	stdin.ReadString('\n')
	fmt.Fscan(stdin, &g)
	stdin.ReadString('\n')
	fmt.Fscan(stdin, &b)
	color := RGB{r, g, b}
	fmt.Println(toColorStr(color.traverse()))
	fmt.Scanln()
}

func (c RGB) traverse() RGB {
	return RGB{255 - c.Red, 255 - c.Green, 255 - c.Blue}
}

func toColorStr(c RGB) int {
	return 256*256*c.Red + 256*c.Green + c.Blue + 1
}
