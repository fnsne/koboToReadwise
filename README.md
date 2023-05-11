# koboToReadwise
 使用Kobo APP的預設位置：`%userpath%/AppData/Local/Kobo/Kobo Desktop Edition/` 裡面的`Kobo.sqlite` 來解析
 最終輸出Amazon Kindle的註記格式檔案：`output.txt` 。
 可以在[Readwise](https://readwise.io/welcome/sync)使用`My Clippings.txt`方式來import highlight和notes
 
# How to use
1. 下載[tool](）
2. 執行tool，就會在同層的資料夾看到`output.txt`檔案，裡面就是轉換成kindle clipping的格式。
## Todos
- [x] 可以轉換出用於[Readwise](https://readwise.io/welcome/sync)的Kindle MyClipping.txt 所需的格式。
- [x] 可以查到作者名字
- [x] 可以轉換highlights
- [x] 可以轉換筆記到對應的highlight
- [x] 支援macOS
- [x] 使用env來指定使用的系統
- [ ] 發佈Release
- [ ] Add how to use
- [ ] 能直接上傳到readwise
- [ ] 能上傳書藉封面


## 參考來源
https://github.com/mollykannn/kobo-book-exporter-go/blob/main/models/bookList/bookList.go
https://www.vixual.net/blog/archives/117
