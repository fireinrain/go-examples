package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

//获取命令行参数
var infile *string = flag.String("i", "infile", "File contarins values for sorting")
var outfile *string = flag.String("o", "outfile", "File to recieve sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

//读取文件
func readValuesInFile(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open infile: ", infile)
		return
	}
	//及时关闭文件回收资源
	defer file.Close()

	reader := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPrefix, err2 := reader.ReadLine()
		if err2 != nil {
			if err2 != io.EOF {
				err = err2

			}
			break
		}
		if isPrefix {
			fmt.Println("a too long line,seems unexpected")
			return
		}

		str := string(line)

		value, error2 := strconv.Atoi(str)
		if error2 != nil {
			err = error2
			return
		}

		values = append(values, value)

	}
	return values, err

}

//写入文件
func WriteValuesToFile(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("failed to create output file: ", outfile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

//获取当前文件运行时的路径
/**
path:  C:\Users\sunrise\AppData\Local\Temp\___go_build_sorter_go.exe
*/
func getRuntimeFilePath() (path string, err error) {
	executable, error := os.Executable()
	if error != nil {
		fmt.Println("error: ", error)

	}
	fmt.Println("path: ", executable)
	return
}

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}
	//fmt.Println(*infile)

	//getRuntimeFilePath()

	//读取文件到数组
	values, err := readValuesInFile(*infile)
	//values, err := readValuesInFile("infile.txt")

	if err == nil {
		fmt.Println("values: ", values)
	} else {
		fmt.Println("error: ", err)
	}
}
