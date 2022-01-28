package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/go-vgo/robotgo"
)

func main() {
	// robotgo.KeyTap("esc", os.Exit(0))
	fmt.Printf("Last Cloudia首抽機 version 1.0\n")
	dataMap := loadConfig("config.txt")

	fmt.Printf("config設定如下:\n")
	for k, v := range dataMap {
		fmt.Printf("%v\t\t%v\n", k, v)
	}

	luckyBallGo(1, dataMap)
}

func errHandler(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		panic(err)
	}
}

func loadConfig(csvPath string) map[string]float64 {

	dataMap := make(map[string]float64, 13)
	csvFile, err := os.Open(csvPath)
	errHandler(err, "ErrCode 003001. Open CSV file failed.")

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	errHandler(err, "ErrCode 003002. Read CSV file failed.")

	for _, line := range csvLines {
		tempFloat, err := strconv.ParseFloat(line[1], 64)
		errHandler(err, "ErrCode 003003. Convert string to float error.")
		dataMap[line[0]] = tempFloat
	}

	return dataMap
}

func luckyBallGo(round int, dataMap map[string]float64) {
	// init PNG path
	var (
		PNG_skip1            string = "img/skip.png"
		PNG_skip2            string = "img/skip2.png"
		PNG_reRoll           string = "img/reRoll.png"
		PNG_Char_ThundSevia  string = "img/Char_ThundSevia.png"
		PNG_Char_IceSevia    string = "img/Char_IceSevia.png"
		PNG_Ark_KadinaGreat  string = "img/Ark_KadinaGreat.png"
		PNG_Ark_Pirate       string = "img/Ark_Pirate.png"
		PNG_Ark_Skyship      string = "img/Ark_Skyship.png"
		PNG_Ark_Sandwyrm     string = "img/Ark_Sandwyrm.png"
		PNG_Ark_ForeignLand  string = "img/Ark_ForeignLand.png"
		PNG_Ark_SwordCorpses string = "img/Ark_SwordCorpses.png"
		PNG_Ark_IcyGuardian  string = "img/Ark_IcyGuardian.png"
	)
	for {
		var (
			niceRoll         bool = false
			Ark_KadinaGreat  bool = false
			Char_ThundSevia  bool = false
			Char_IceSevia    bool = false
			Ark_IcyGuardian  bool = false
			Ark_Sandwyrm     bool = false
			Ark_SwordCorpses bool = false
			Ark_ForeignLand  bool = false
			Ark_Pirate       bool = false
			Ark_Skyship      bool = false
			final_point      int  = 0
			min_point        int  = int(dataMap["PassingScore"])
			tmp_fx           int  = 0
			tmp_fy           int  = 0
		)

		bitmap_screen := robotgo.CaptureScreen(0, 0, int(dataMap["ScreenWidth"]), int(dataMap["ScreenLength"]))
		robotgo.FreeBitmap(bitmap_screen)
		defer robotgo.FreeBitmap(bitmap_screen)
		fmt.Printf("Round %d\n", round)
		for {
			// update screenshot
			bitmap_screen = robotgo.CaptureScreen(0, 0, int(dataMap["ScreenWidth"]), int(dataMap["ScreenLength"]))

			// 出現跳過動畫按鈕直接點
			fx, fy := robotgo.FindPic(PNG_skip1, bitmap_screen, dataMap["Tolerance_2"])
			// fmt.Printf("Skip button - result:%d,%d\n", fx, fy)
			if fx != -1 && fy != -1 {
				// robotgo.Move(fx, fy)
				robotgo.MoveClick(fx, fy, "left", false)
			}
			fx, fy = robotgo.FindPic(PNG_skip2, bitmap_screen, dataMap["Tolerance_2"])
			// fmt.Printf("Skip button - result:%d,%d\n", fx, fy)
			if fx != -1 && fy != -1 {
				// robotgo.Move(fx, fy)
				robotgo.MoveClick(fx, fy, "left", false)
			}

			fx2, fy2 := robotgo.FindPic(PNG_reRoll, bitmap_screen, dataMap["Tolerance_2"])
			// fmt.Printf("Reroll button - result:%d,%d\n", fx2, fy2)
			if fx2 != -1 && fy2 != -1 {
				// 該輪轉蛋結果已出
				fmt.Printf("偵測到「再抽一次」，進入評分階段\n")
				tmp_fx = fx2
				tmp_fy = fy2
				robotgo.FreeBitmap(bitmap_screen)
				break
			}
			robotgo.FreeBitmap(bitmap_screen)
		}
		for i := 0; i < 3; i++ {
			bitmap_screen = robotgo.CaptureScreen(0, 0, int(dataMap["ScreenWidth"]), int(dataMap["ScreenLength"]))
			// 判斷該輪是否為想要的結果
			fx, fy := robotgo.FindPic(PNG_Ark_KadinaGreat, bitmap_screen, dataMap["Tolerance"])
			if fx != -1 && fy != -1 {
				// fmt.Printf("Ark - result:%d,%d\n", fx, fy)
				Ark_KadinaGreat = true
			}
			fx, fy = robotgo.FindPic(PNG_Char_ThundSevia, bitmap_screen, dataMap["Tolerance"])
			if fx != -1 && fy != -1 {
				// fmt.Printf("Ark - result:%d,%d\n", fx, fy)
				Char_ThundSevia = true
			}
			fx, fy = robotgo.FindPic(PNG_Char_IceSevia, bitmap_screen, dataMap["Tolerance"])
			if fx != -1 && fy != -1 {
				// fmt.Printf("Ark - result:%d,%d\n", fx, fy)
				Char_IceSevia = true
			}
			fx, fy = robotgo.FindPic(PNG_Ark_IcyGuardian, bitmap_screen, dataMap["Tolerance"])
			if fx != -1 && fy != -1 {
				// fmt.Printf("Ark - result:%d,%d\n", fx, fy)
				Ark_IcyGuardian = true
			}
			fx, fy = robotgo.FindPic(PNG_Ark_Sandwyrm, bitmap_screen, dataMap["Tolerance"])
			if fx != -1 && fy != -1 {
				// fmt.Printf("Ark - result:%d,%d\n", fx, fy)
				Ark_Sandwyrm = true
			}
			fx, fy = robotgo.FindPic(PNG_Ark_SwordCorpses, bitmap_screen, dataMap["Tolerance"])
			if fx != -1 && fy != -1 {
				// fmt.Printf("Ark - result:%d,%d\n", fx, fy)
				Ark_SwordCorpses = true
			}
			fx, fy = robotgo.FindPic(PNG_Ark_ForeignLand, bitmap_screen, dataMap["Tolerance"])
			if fx != -1 && fy != -1 {
				// fmt.Printf("Ark - result:%d,%d\n", fx, fy)
				Ark_ForeignLand = true
			}
			fx, fy = robotgo.FindPic(PNG_Ark_Pirate, bitmap_screen, dataMap["Tolerance"])
			if fx != -1 && fy != -1 {
				// fmt.Printf("Ark - result:%d,%d\n", fx, fy)
				Ark_Pirate = true
			}
			fx, fy = robotgo.FindPic(PNG_Ark_Skyship, bitmap_screen, dataMap["Tolerance"])
			if fx != -1 && fy != -1 {
				// fmt.Printf("Ark - result:%d,%d\n", fx, fy)
				Ark_Skyship = true
			}
			robotgo.FreeBitmap(bitmap_screen)
		}

		// 計分區
		if Ark_KadinaGreat {
			final_point += int(dataMap["Ark_KadinaGreat"])
		}
		if Ark_Pirate {
			final_point += int(dataMap["Ark_Pirate"])
		}
		if Ark_Sandwyrm {
			final_point += int(dataMap["Ark_Sandwyrm"])
		}
		if Ark_Skyship {
			final_point += int(dataMap["Ark_Skyship"])
		}
		if Ark_ForeignLand {
			final_point += int(dataMap["Ark_ForeignLand"])
		}
		if Ark_IcyGuardian {
			final_point += int(dataMap["Ark_IcyGuardian"])
		}
		if Ark_SwordCorpses {
			final_point += int(dataMap["Ark_SwordCorpses"])
		}
		if Char_ThundSevia {
			final_point += int(dataMap["Char_ThundSevia"])
		}
		if Char_IceSevia {
			final_point += int(dataMap["Char_IceSevia"])
		}

		fmt.Printf("該輪得分:%d\n", final_point)
		fmt.Printf("Char_雷姬:%v\n", Char_ThundSevia)
		fmt.Printf("Char_冰姬:%v\n", Char_IceSevia)
		fmt.Printf("Ark_榮光 :%v\n", Ark_KadinaGreat)
		fmt.Printf("Ark_藍冰 :%v\n", Ark_IcyGuardian)
		fmt.Printf("Ark_沙蟲 :%v\n", Ark_Sandwyrm)
		fmt.Printf("Ark_屍劍 :%v\n", Ark_SwordCorpses)
		fmt.Printf("Ark_異國 :%v\n", Ark_ForeignLand)
		fmt.Printf("Ark_海盜 :%v\n", Ark_Pirate)
		fmt.Printf("Ark_空艇 :%v\n", Ark_Skyship)

		// 判斷是否要結束
		if final_point >= min_point {
			niceRoll = true
		}

		if niceRoll {
			// 該輪為想要結果
			break
		} else {
			bitmap_screen = robotgo.CaptureScreen(0, 0, int(dataMap["ScreenWidth"]), int(dataMap["ScreenLength"]))
			fx2, fy2 := robotgo.FindPic("img/reRoll.png", bitmap_screen, dataMap["Tolerance_2"])
			if fx2 != -1 && fy2 != -1 {
				robotgo.MoveClick(fx2, fy2, "left", false)
			} else {
				robotgo.MoveClick(tmp_fx, tmp_fy, "left", false)
			}
			// fmt.Printf("Reroll button - result:%d,%d\n", fx2, fy2)

			robotgo.FreeBitmap(bitmap_screen)
			round++
			robotgo.Sleep(1)
			// time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("Have a Nice Roll.")
}
