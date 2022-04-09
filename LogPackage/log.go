package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.Println("Println")
	//Output: 2022/04/03 23:03:06 Println
	PrintBr()

	log.SetFlags(0)
	log.Println("SetFlags(0)")
	//Output: SetFlags(0)
	PrintBr()

	checkOtherFlages()
	PrintBr()
	log.SetFlags(3)

	//io.Discard doesn't show logs even log.Fatal:
	log.SetOutput(io.Discard)
	log.Println("SetOutput(io.Discard)")
	//Output:

	//os.Stdout show logs on the terminal:
	log.SetOutput(os.Stdout)
	PrintBr()

	//save log into file:
	//If the file doesn't exist, create it or append to the file
	file, _ := os.OpenFile("LogPackage\\log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 1)
	log.SetOutput(file)
	log.Println("loged info in file.")
	log.SetOutput(os.Stdout)

	//fatal:
	defer log.Println("defer log.Println with log.Fatal")
	//to hard exit without running the defer
	log.Fatal("Hard Exit")
	//Output: 19:11:43.062693 Hard Exit

}
func PrintBr() {
	fmt.Println("---------------------------------------------------------------------------------------------------")
}
func checkOtherFlages() {
	for i := 0; i <= 100; i++ {
		log.SetFlags(i)
		log.Println(i)
	}
	//Outputs:
	// 0
	// 2022/04/03 1
	// 23:11:39 2
	// 2022/04/03 23:11:39 3
	// 23:11:39.606977 4
	// 2022/04/03 23:11:39.606977 5
	// 23:11:39.606977 6
	// 2022/04/03 23:11:39.606977 7
	// D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 8
	// 2022/04/03 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 9
	// 23:11:39 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 10
	// 2022/04/03 23:11:39 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 11
	// 23:11:39.608978 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 12
	// 2022/04/03 23:11:39.608978 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 13
	// 23:11:39.608978 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 14
	// 2022/04/03 23:11:39.608978 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 15
	// log.go:24: 16
	// 2022/04/03 log.go:24: 17
	// 23:11:39 log.go:24: 18
	// 2022/04/03 23:11:39 log.go:24: 19
	// 23:11:39.610978 log.go:24: 20
	// 2022/04/03 23:11:39.610978 log.go:24: 21
	// 23:11:39.610978 log.go:24: 22
	// 2022/04/03 23:11:39.610978 log.go:24: 23
	// log.go:24: 24
	// 2022/04/03 log.go:24: 25
	// 23:11:39 log.go:24: 26
	// 2022/04/03 23:11:39 log.go:24: 27
	// 23:11:39.611979 log.go:24: 28
	// 2022/04/03 23:11:39.611979 log.go:24: 29
	// 23:11:39.612978 log.go:24: 30
	// 2022/04/03 23:11:39.612978 log.go:24: 31
	// 32
	// 2022/04/03 33
	// 18:41:39 34
	// 2022/04/03 18:41:39 35
	// 18:41:39.613978 36
	// 2022/04/03 18:41:39.613978 37
	// 18:41:39.613978 38
	// 2022/04/03 18:41:39.613978 39
	// D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 40
	// 2022/04/03 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 41
	// 18:41:39 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 42
	// 2022/04/03 18:41:39 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 43
	// 18:41:39.614978 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 44
	// 2022/04/03 18:41:39.614978 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 45
	// 18:41:39.614978 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 46
	// 2022/04/03 18:41:39.614978 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 47
	// log.go:24: 48
	// 2022/04/03 log.go:24: 49
	// 18:41:39 log.go:24: 50
	// 2022/04/03 18:41:39 log.go:24: 51
	// 18:41:39.615979 log.go:24: 52
	// 2022/04/03 18:41:39.615979 log.go:24: 53
	// 18:41:39.615979 log.go:24: 54
	// 2022/04/03 18:41:39.616978 log.go:24: 55
	// log.go:24: 56
	// 2022/04/03 log.go:24: 57
	// 18:41:39 log.go:24: 58
	// 2022/04/03 18:41:39 log.go:24: 59
	// 18:41:39.616978 log.go:24: 60
	// 2022/04/03 18:41:39.617979 log.go:24: 61
	// 18:41:39.617979 log.go:24: 62
	// 2022/04/03 18:41:39.617979 log.go:24: 63
	// 64
	// 2022/04/03 65
	// 23:11:39 66
	// 2022/04/03 23:11:39 67
	// 23:11:39.618978 68
	// 2022/04/03 23:11:39.618978 69
	// 23:11:39.618978 70
	// 2022/04/03 23:11:39.618978 71
	// D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 72
	// 2022/04/03 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 73
	// 23:11:39 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 74
	// 2022/04/03 23:11:39 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 75
	// 23:11:39.619980 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 76
	// 2022/04/03 23:11:39.619980 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 77
	// 23:11:39.619980 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 78
	// 2022/04/03 23:11:39.619980 D:/Go_App/Useful_commands_in_GoLang/Useful_commands_in_GoLang/logpakage/log.go:24: 79
	// log.go:24: 80
	// 2022/04/03 log.go:24: 81
	// 23:11:39 log.go:24: 82
	// 2022/04/03 23:11:39 log.go:24: 83
	// 23:11:39.620979 log.go:24: 84
	// 2022/04/03 23:11:39.620979 log.go:24: 85
	// 23:11:39.621979 log.go:24: 86
	// 2022/04/03 23:11:39.621979 log.go:24: 87
	// log.go:24: 88
	// 2022/04/03 log.go:24: 89
	// 23:11:39 log.go:24: 90
	// 2022/04/03 23:11:39 log.go:24: 91
	// 23:11:39.622980 log.go:24: 92
	// 2022/04/03 23:11:39.622980 log.go:24: 93
	// 23:11:39.622980 log.go:24: 94
	// 2022/04/03 23:11:39.622980 log.go:24: 95
	// 96
	// 2022/04/03 97
	// 18:41:39 98
	// 2022/04/03 18:41:39 99
	// 18:41:39.623979 100
}
