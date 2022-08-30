package builder

import (
	"fmt"
	"os"
)

func TaskForSimpleIcons() {
	fmt.Println("开始复制 Simple Icon 资源")
	_PrepareDirectory("pkg/simpleicon/icons")
	err := _CopyDirectory("embed/assets/vendor/simple-icons", "pkg/simpleicon/icons")
	if err != nil {
		fmt.Println("复制 Simple Icon 错误: ", err)
		os.Exit(-1)
	}
	fmt.Println("复制 Simple Icon 资源完成")
}
