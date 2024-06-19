package main

import (
	"bufio"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/zkfy/log"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	file, err := os.Open("./huiming.csv")
	if err != nil {
		log.Info("打开文件出错")
		return
	}
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(1, func(i interface{}) {
		downloan(i)
		wg.Done()
	})
	defer p.Release()
	buf := bufio.NewReader(file)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Info("读取文件失败")
		}

		str := string(line)
		strArr := strings.Split(str, ",")
		//fmt.Println(strArr)
		wg.Add(1)
		_ = p.Invoke(strArr)
		//downloan(strings.Trim(strArr[4],"\""))
		//break
	}
	wg.Wait()
	fmt.Println("end")
}

func downloan(i interface{}) {
	/*time.Sleep(3 * time.Second);
	fmt.Println(i)
	return*/
	strArr := i.([]string)
	//fmt.Println(strArr)
	fileName := fmt.Sprintf("%s_%s_%s_%s.pdf",
		strings.Trim(strArr[5], "\""),
		strings.Trim(strArr[1], "\""),
		strings.Trim(strArr[2], "\""),
		strings.Trim(strArr[3], "\""))
	fmt.Println(fileName)
	url := strings.Trim(strArr[4], "\"")
	fmt.Println(url)
	client := http.Client{}
	res, err := client.Get(url)
	if err != nil {
		log.Info("下载文件出错")
		log.Info(err.Error())
		return
	}
	defer res.Body.Close()
	buf := bufio.NewReader(res.Body)
	file, err := os.Create("./huiming/" + fileName)
	if err != nil {
		log.Info("打开文件出错")
		return
	}
	defer file.Close()
	io.Copy(file, buf)
}
