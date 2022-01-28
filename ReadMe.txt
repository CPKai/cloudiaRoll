!!!!!該程式會跟你搶滑鼠，使用請留意!!!!!
!!!!!模擬器的的長寬請使用270(寬)480(長)，240(DPI，這個不知會不會影響)，因img資料夾的比對圖是在這長寬下截取!!!!!
註1：要使用不同長寬的人請自行替換img資料夾下的所有圖片至目標長寬

使用前請先至config.txt檔中設定數值
以下為參數說明
Ark_KadinaGreat,100     抽中榮光，最終得分加100分
Ark_Pirate,10           抽中海盜船，最終得分加10分
Ark_Sandwyrm,10         抽中沙蟲
Ark_Skyship,10          抽中空艇
Ark_ForeignLand,10      抽中異國
Ark_IcyGuardian,10      抽中藍冰
Ark_SwordCorpses,10     抽中屍劍山
Char_ThundSevia,10      抽中角色-雷姬
Char_IceSevia,0         抽中角色-冰姬
PassingScore,120        最終計分要大於「或」等於120才算及格
ScreenWidth,1920        截圖寬度，建議填螢幕的寬度
ScreenLength,1080       截圖長度，建議填螢幕的長度
Tolerance,0.4           評分階段使用的容忍值，SSR轉出來會bling bling閃，跟img資料夾裡的對照圖不會完全一樣，所以調了一些容忍值
Tolerance_2,0.3         「跳過」與「再轉一次」判斷時使用的容忍值

程式概念是進行螢幕截圖(截圖大小照config中的參數)，再比對img資料夾中各圖片是否存在於剛剛的截圖中，最後進行計分
註2：多螢幕使用者請把目標模擬器放至主螢幕