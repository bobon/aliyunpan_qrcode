# qrcode2console

golang 将二维码图片输出到控制台显示 

- 为什么有这个库？

  - 写爬虫的时候，遇到需要扫二维码的场景，控制台可以输出，然后直接扫码。
- 使用【二维码识别区外不能有其他图案，否则会生成失败】

  ```
  go get github.com/yantao1995/qrcode2console 
  ```

#### 支持从base64中生成


```
qr, err := qrcode2console.NewQrcodeFromPath("base64")
if err != nil {
    log.Fatalln(err)
}
qr.SetBound(1) //不设置外边框将使用原外框的最小边的缩放比例
qr.PrintForConsole()
```

#### 支持从文件夹路径生成


```
qr, err := qrcode2console.NewQrcodeFromPath("filePath")
if err != nil {
    log.Fatalln(err)
}
qr.SetBound(1) //不设置外边框将使用原外框的最小边的缩放比例
qr.PrintForConsole()
```
