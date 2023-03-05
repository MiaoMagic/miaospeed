package preconfigs

import (
	"strings"
)

const PROXY_DEFAULT_STUN_SERVER = "udp://stun.msl.la:3478"

const NETCAT_HTTP_PAYLOAD = `GET %s HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Host: %s
User-Agent: HTTPie/3.0.2 MiaoSpeed/%s

`

const (
	SPEED_DEFAULT_DURATION  int64 = 3
	SPEED_DEFAULT_THREADING uint  = 1

	SPEED_DEFAULT_LARGE_FILE_STATIC_APPLE    string = "https://secure-appldnld.apple.com/itunes12/041-56365-20210413-79A5AC6D-6B0A-46EC-B399-5A381183E3BE/iTunes12.8.3.dmg"
	SPEED_DEFAULT_LARGE_FILE_STATIC_MSFT     string = "https://download.microsoft.com/download/9/B/1/9B1EAE94-B328-451E-AF9D-89BD73CDA9E0/WPSDK-7.1.1-KB2669187-x86.exe"
	SPEED_DEFAULT_LARGE_FILE_STATIC_GOOGLE   string = "https://dl.google.com/developers/android/sc/images/factory/bramble-s3b1.220218.004-factory-a2a7cafb.zip"
	SPEED_DEFAULT_LARGE_FILE_STATIC_CACHEFLY string = "http://cachefly.cachefly.net/200mb.test"

	SPEED_DEFAULT_LARGE_FILE_DYN_INTL     string = "DYNAMIC:INTL"
	SPEED_DEFAULT_LARGE_FILE_DYN_FAST     string = "DYNAMIC:FAST"
	SPEED_DEFAULT_LARGE_FILE_DYN_ALL_INTL string = "DYNAMIC:ALL"

	SPEED_DEFAULT_LARGE_FILE_DEFAULT = SPEED_DEFAULT_LARGE_FILE_DYN_INTL

	SLAVE_DEFAULT_PING         = "https://www.gstatic.com/generate_204"
	SLAVE_DEFAULT_RETRY   uint = 3
	SLAVE_DEFAULT_TIMEOUT uint = 5000
)

var (
	SPEED_DEFAULT_LARGE_FILE_DYNAMIC = strings.Split("https://dl.google.com/developers/android/sc/images/factory/bramble-s3b1.220218.004-factory-a2a7cafb.zip,https://secure-appldnld.apple.com/itunes12/041-56365-20210413-79A5AC6D-6B0A-46EC-B399-5A381183E3BE/iTunes12.8.3.dmg,https://secure-appldnld.apple.com/itunes12/032-11199-20221212-7680817F-9CEC-4DD3-9191-8D0C20E8A548/iTunes64Setup.exe,https://download.microsoft.com/download/9/B/1/9B1EAE94-B328-451E-AF9D-89BD73CDA9E0/WPSDK-7.1.1-KB2669187-x86.exe", ",")
)
