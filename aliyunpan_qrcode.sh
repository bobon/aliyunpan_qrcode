#!/bin/bash

data=$(curl -sS 'https://passport.aliyundrive.com/newlogin/qrcode/generate.do?appName=aliyun_drive&fromSite=52&appEntrance=web' \
-H 'referer: https://passport.aliyundrive.com/mini_login.htm?lang=zh_cn&appName=aliyun_drive&appEntrance=web&styleType=auto&bizParams=&notLoadSsoView=false&notKeepLogin=false&isMobile=false&ad__pass__q__rememberLogin=true&ad__pass__q__rememberLoginDefaultValue=true&ad__pass__q__forgotPassword=true&ad__pass__q__licenseMargin=true&ad__pass__q__loginType=normal&hidePhoneCode=true&rnd=0.20099676922221987' \
| jq)
qrCodeCk=$(echo -e "$data" | jq -r '.content.data|("ck=" + .ck + "&t=" + (.t|tostring))')
codeContent=$(echo -e "$data" | jq -r '.content.data.codeContent')
#echo $codeContent
qrcode-terminal "$codeContent"

function pause(){
    read -n 1 -p "$*" INP
        if [ "$INP" != "" ] ; then
            echo -ne '\b \n'
        fi
}
pause 'After scanning the QR code, press any key to continue.'

qrCodeResponse=$(curl -sS 'https://passport.aliyundrive.com/newlogin/qrcode/query.do?appName=aliyun_drive&fromSite=52&_bx-v=2.0.31' -X POST \
-H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0' -H 'Accept: image/avif,image/webp,*/*' \
-H 'Accept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2' \
-H 'origin: https://passport.aliyundrive.com' \
-H 'referer: "https://passport.aliyundrive.com/mini_login.htm?lang=zh_cn&appName=aliyun_drive&appEntrance=web&styleType=auto&bizParams=&notLoadSsoView=false&notKeepLogin=false&isMobile=false&ad__pass__q__rememberLogin=true&ad__pass__q__rememberLoginDefaultValue=true&ad__pass__q__forgotPassword=true&ad__pass__q__licenseMargin=true&ad__pass__q__loginType=normal&hidePhoneCode=true&rnd=0.17778824737759047' \
-H 'content-type: application/x-www-form-urlencoded' \
--data-raw ''${qrCodeCk}'' | jq)

login_result=$(echo -e "$qrCodeResponse" | jq -r '.content.data.bizExt' | base64 -d)
echo -e "$login_result" | jq -r '"refreshToken: " + .pds_login_result.refreshToken'

