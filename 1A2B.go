package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var usrName string
	var YN string
	fmt.Println("hello please enter your name : ")
	fmt.Scanln(&usrName)
	fmt.Printf("would you want to play game right now? (Y/N) : ")
	fmt.Scanln(&YN)
	YN = validYN(YN)
	if YN == "Y" {
		startGame(usrName)
	} else {
		exitGame()
	}
}

// 亂數產生
func randomNumber() string {
	var result string
	rand.Seed(time.Now().UnixNano())
	for {
		tmpnum := strconv.Itoa(rand.Intn(9))
		if len(result) >= 4 {
			break
		} else {
			if !strings.Contains(result, tmpnum) {
				result = result + tmpnum
			}
		}
	}
	return result
}

// 輸入
func inputModel(content string) string {
	var tmp string
	fmt.Println(content)
	fmt.Scanln(&tmp)
	return tmp
}

// 開始遊戲
func startGame(usrName string) {
	var spentTimes int = 1
	var costT time.Duration
	tn := time.Now()
	solution := randomNumber()
	for {
		tmpNum := inputModel("please enter number : ")
		if solution == tmpNum {
			costT = time.Since(tn)
			break
		} else {
			spentTimes++
			fmt.Println(validNum(solution, tmpNum))
		}
	}
	var retry string
	fmt.Println("great!! ", usrName, " cost time : ", costT, " spent times : ", spentTimes, ", Do you want to try again?? (Y/N)")
	fmt.Scanln(&retry)
	retry = validYN(retry)
	if retry == "Y" {
		startGame(usrName)
	} else {
		exitGame()
	}
}

// 離開遊戲
func exitGame() {
	var input string
	fmt.Println("enter any key to exit")
	fmt.Scanln(&input)
	os.Exit(9999)
}

// 驗證數字
func validNum(sol, num string) string {
	var aNum int
	var bNum int
	num = validLen(num)
	num = validduplicate(num)
	for i, k := range num {
		for j, l := range sol {
			if k == l {
				if i == j {
					aNum++
				} else {
					bNum++
				}
			}
		}
	}
	return strconv.Itoa(aNum) + "A" + strconv.Itoa(bNum) + "B"
}

// 判斷輸入是否符合
func validYN(input string) string {
	for {
		if input == "Y" || input == "N" {
			break
		} else {
			fmt.Printf("please enter (Y/N)")
			fmt.Scanln(&input)
		}
	}
	return input
}

// 判斷輸入數字是否重複
func validduplicate(input string) string {
	for {
		tag := true
		for _, i := range input {
			times := 0
			for _, j := range input {
				if i == j {
					times++
				}
			}
			if times > 1 {
				tag = false
				break
			}
		}
		if tag {
			break
		}
		input = inputModel("do not duplicat number content, enter number again: ")
	}
	return input
}

// 檢查長度
func validLen(input string) string {
	for {
		if len(input) == 4 {
			break
		}
		input = inputModel("please enter four digits number, enter number again : ")
	}
	return input
}
