# qrcode2console

golang 将二维码图片输出到控制台显示 

Parse QR code pictures or Base64 data, and then output to the console.


- 为什么有这个库？
- Why is there this library?
- 
  - 写爬虫的时候，遇到需要扫二维码的场景，控制台可以输出，然后直接扫码。
  - When writing a crawler, you need to scan the QR code. The console can output it, and then scan the code directly.
  - 使用【二维码识别区外不能有其他图案，否则会生成失败】
  - Use [no other patterns outside the QR code recognition area, otherwise the generation will fail]

  ```
  go get github.com/yantao1995/qrcode2console 
  ```

#### 支持从base64中生成
#### Support generation from Base64


```
qr, err := qrcode2console.NewQrcodeFromBase64("base64")
if err != nil {
    log.Fatalln(err)
}
qr.SetBound(1) //不设置外边框将使用原外框的最小边的缩放比例 //If the outer frame is not set, the scale of the smallest edge of the original outer frame will be used
qr.PrintForConsole()
```

#### 支持从文件夹路径生成
#### Support generating from folder path

```
qr, err := qrcode2console.NewQrcodeFromPath("filePath")
if err != nil {
    log.Fatalln(err)
}
qr.SetBound(1) //不设置外边框将使用原外框的最小边的缩放比例 //If the outer frame is not set, the scale of the smallest edge of the original outer frame will be used
qr.PrintForConsole()
```
