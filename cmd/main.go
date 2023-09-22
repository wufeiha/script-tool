package main

import (
	"fmt"
	"os"
	"script-tool/pkg"
	"script-tool/script"
)

//	func main() {
//		dir := "C:\\Users\\admin\\Desktop\\tool"
//		output := "output"
//		err := pkg.CreateOrReplaceDir(dir + "\\" + output)
//		if err != nil {
//			fmt.Println("create dir error:", err)
//			return
//		}
//		script.GenFiles(dir+"\\data.xlsx", dir+"\\template.xlsx", dir+"\\"+output+"\\")
//		fmt.Println("success")
//	}
func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Current directory:", dir)
	output := "output"
	err = pkg.CreateOrReplaceDir(dir + "\\" + output)
	if err != nil {
		fmt.Println("create dir error:", err)
		return
	}
	script.GenFiles(dir+"\\data.xlsx", dir+"\\template.xlsx", dir+"\\"+output+"\\")
	fmt.Println("success")
}
