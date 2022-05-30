package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n) // int 타입값을 입력받음
	if err != nil {
		stdin.ReadString('\n') // 에러 발생 시 입력 스트림을 비움
	}
	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(100) // 랜덤값 생성
	count := 1

	for {
		fmt.Printf("숫자값을 입력하세요 : ")
		n, err := InputIntValue() // 숫자값 입력
		if err != nil {
			fmt.Println("숫자만 입력하세요")
		} else {
			if n > r { // 숫자값 비교
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if n < r {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Println("숫자를 맞췄습니다. 시도한 횟수 : ", count)
				// 같을 경우 메세지 출력 후 break
				break
			}
			count++
		}
	}
}
