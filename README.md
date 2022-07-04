# qrcode2console

golang 将二维码图片输出到控制台显示 

Parse QR code pictures or Base64 data, and then output to the console.


- 为什么有这个库？
  - 由于阿里云盘referrer的限制，必须使用移动端token，使用桌面web端token会导致无法下载与预览。使用此程序，可通过扫描二维码，获取阿里云盘的登陆refreshToken。

- 使用方法
  git clone 
  sudo apt-get install -y curl jq
  go get github.com/bobon/aliyunpan_qrcode  
  bash aliyunpan_qrcode.sh

