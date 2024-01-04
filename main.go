package main

import (
	"fmt"
	"strconv"
	"strings"
)

var AlpMapping = map[int64][]string{
	1: {"a", "j", "s"},
	2: {"b", "k", "t"},
	3: {"c", "l", "u"},
	4: {"d", "m", "v"},
	5: {"e", "n", "w"},
	6: {"f", "o", "x"},
	7: {"g", "p", "y"},
	8: {"h", "q", "z"},
	9: {"i", "r"},
}

var (
	vowelList     = []string{"a", "e", "i", "o", "u"}
	consonantList = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "s", "t", "v", "w", "x", "z"}
)

func main() {
	var fullName, birthDay string

	fmt.Print("Enter full name: ")
	_, err := fmt.Scan(&fullName)
	if err != nil {
		fmt.Println("Error reading full name:", err)
		return
	}

	fullName = strings.ToLower(fullName)

	fmt.Print("Enter birthday: ")
	_, err = fmt.Scan(&birthDay)
	if err != nil {
		fmt.Println("Error reading birthday:", err)
		return
	}

	fmt.Println("--------------------------------------------------")

	lifePathNo := calcLifePathNo(birthDay)
	fmt.Println("Life Path:", lifePathNo)

	birthDayNo := calcBirthDayNo(birthDay)
	fmt.Println("Date of Birth:", birthDayNo)

	commissionNo := calcMissionNo(fullName)
	fmt.Println("Commission:", commissionNo)

	conn := calcLifePathCommission(lifePathNo, commissionNo)
	fmt.Println("Life Path - Commission:", conn)

	maturityNo := calcFinalNo(lifePathNo + commissionNo)
	fmt.Println("Maturity:", maturityNo)

	missingNo := calcMissingNo(fullName)
	fmt.Println("Missing:", missingNo)

	soulNo := calcSoulNo(fullName)
	fmt.Println("Soul:", soulNo)

	stages := calcStage(birthDay, lifePathNo)
	fmt.Println("Stages:", stages)

}

func calculateDigitSum(num int64) (sum int64) {
	for num != 0 {
		digit := num % 10
		sum += digit
		num /= 10
	}

	return
}

func calcFinalNo(n int64) int64 {
	for {
		if n <= 9 || n == 11 || n == 22 || n == 33 {
			return n
		}
		n = calculateDigitSum(n)
	}

}

func calcLifePathNo(birthDay string) int64 {
	splitBirthDay := strings.Split(birthDay, "/")

	var sum int64 = 0
	for _, value := range splitBirthDay {
		intVal, _ := strconv.ParseInt(value, 10, 64)
		sum += calcFinalNo(intVal)
	}

	return calcFinalNo(sum)
}

func calcBirthDayNo(birthDay string) int64 {
	splitBirthDay := strings.Split(birthDay, "/")

	day := splitBirthDay[0]
	intVal, _ := strconv.ParseInt(day, 10, 64)

	return calcFinalNo(intVal)
}

func calcMissionNo(fullName string) int64 {
	splitFullName := strings.Split(fullName, "/")

	var sum int64 = 0

	for _, value := range splitFullName {
		var specifyName int64 = 0
		for _, char := range value {
			for no, c := range AlpMapping {
				for _, v := range c {
					if fmt.Sprintf("%c", char) == v {
						specifyName += no
						break
					}
				}
			}

		}
		sum += calcFinalNo(specifyName)
	}

	return calcFinalNo(sum)
}

func calcMissingNo(fullName string) []int64 {
	var full []int64 = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var noList, missingList []int64

	splitFullName := strings.Split(fullName, "/")

	for _, value := range splitFullName {
		for _, char := range value {
			for no, c := range AlpMapping {
				for _, v := range c {
					if fmt.Sprintf("%c", char) == v {
						noList = append(noList, no)
						break
					}
				}
			}

		}
	}

	for _, no := range full {
		missed := true
		for _, n := range noList {
			if no == n {
				missed = false
				break
			}
		}

		if missed {
			missingList = append(missingList, no)
		}
	}

	return missingList
}

func buildNameNoList(fullName string) [][][]interface{} {
	splitFullName := strings.Split(fullName, "/")

	var nameNoList, nameCheckedNoList [][][]interface{}
	for _, name := range splitFullName {
		var charNoList [][]interface{}
		for _, char := range name {
			for no, c := range AlpMapping {
				for _, v := range c {
					if charStr := fmt.Sprintf("%c", char); charStr == v {
						charNoList = append(charNoList, []interface{}{charStr, no})
						break
					}
				}
			}
		}
		nameNoList = append(nameNoList, charNoList)
	}

	for _, nameList := range nameNoList {
		var l [][]interface{}
		for idx, chars := range nameList {
			var prev, next []interface{}
			if idx > 0 {
				prev = nameList[idx-1]
			}
			if idx < len(nameList)-1 {
				next = nameList[idx+1]
			}
			if isVowel(chars, prev, next) {
				l = append(l, []interface{}{chars[0], chars[1], 0})
			} else {
				l = append(l, []interface{}{chars[0], chars[1], 1})
			}
		}

		nameCheckedNoList = append(nameCheckedNoList, l)
	}

	return nameCheckedNoList
}

func calcSoulNo(fullName string) int64 {
	var result int64 = 0
	nameNoList := buildNameNoList(fullName)
	for _, name := range nameNoList {
		var vowelSum int64 = 0
		for _, chars := range name {
			if chars[2] == 0 {
				vowelSum += chars[1].(int64)
			}
		}

		result += calcFinalNo(vowelSum)
	}

	return calcFinalNo(result)
}

func isVowel(chars, prev, next []interface{}) bool {
	if chars[0] == "y" {
		if len(prev) == 0 {
			return false
		}

		if isVowel(prev, []interface{}{}, []interface{}{}) {
			return false
		}

		if len(next) != 0 && isVowel(next, []interface{}{}, []interface{}{}) {
			return false
		}

		return true
	}

	for _, vowel := range vowelList {
		if chars[0] == vowel {
			return true
		}
	}

	return false
}

func calcStage(birthDay string, lifePath int64) interface{} {
	splitBirthDay := strings.Split(birthDay, "/")
	date := splitBirthDay[0]
	month := splitBirthDay[1]
	year := splitBirthDay[2]

	dateInt, _ := strconv.ParseInt(date, 10, 64)
	monthInt, _ := strconv.ParseInt(month, 10, 64)
	yearInt, _ := strconv.ParseInt(year, 10, 64)

	firstStage := calcFinalNo(calcFinalNo(dateInt) + calcFinalNo(monthInt))
	firstStage = calculateDigitSum(firstStage)
	firstStageTo := 36 - lifePath

	secondStage := calcFinalNo(calcFinalNo(dateInt) + calcFinalNo(yearInt))
	secondStage = calculateDigitSum(secondStage)
	secondStageFrom := firstStageTo + 1
	secondStageTo := firstStageTo + 9

	thirdStage := calcFinalNo(firstStage + secondStage)
	thirdStage = calculateDigitSum(thirdStage)
	thirdStageFrom := secondStageTo + 1
	thirdStageTo := secondStageTo + 9

	fourthStage := calcFinalNo(calcFinalNo(monthInt) + calcFinalNo(yearInt))
	fourthStageFrom := thirdStageTo + 1

	return [][]interface{}{
		{fmt.Sprintf("0->%v", firstStageTo), firstStage},
		{fmt.Sprintf("%v->%v", secondStageFrom, secondStageTo), secondStage},
		{fmt.Sprintf("%v->%v", thirdStageFrom, thirdStageTo), thirdStage},
		{fmt.Sprintf("%v->...", fourthStageFrom), fourthStage},
	}
}

func calcLifePathCommission(lifePath, commission int64) int64 {
	conn := calculateDigitSum(lifePath) - calculateDigitSum(commission)
	if conn < 0 {
		return -conn
	}

	return conn
}
