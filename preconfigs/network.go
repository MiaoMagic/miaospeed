package preconfigs

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

	SPEED_DEFAULT_LARGE_FILE_STATIC_APPLE    string = "https://updates.cdn-apple.com/2023WinterFCS/fullrestores/032-36511/4D001B1A-729F-4CEB-A08E-AF40220CAADD/iPhone14,7_16.3_20D47_Restore.ipsw"
	SPEED_DEFAULT_LARGE_FILE_STATIC_MSFT     string = "https://download.microsoft.com/download/9/B/1/9B1EAE94-B328-451E-AF9D-89BD73CDA9E0/WPSDK-7.1.1-KB2669187-x86.exe"
	SPEED_DEFAULT_LARGE_FILE_STATIC_GOOGLE   string = "https://dl.google.com/developers/android/sc/images/factory/bramble-s3b1.220218.004-factory-a2a7cafb.zip"
	SPEED_DEFAULT_LARGE_FILE_STATIC_CACHEFLY string = "http://cachefly.cachefly.net/200mb.test"

	SPEED_DEFAULT_LARGE_FILE_DYN_INTL string = "DYNAMIC:INTL"
	SPEED_DEFAULT_LARGE_FILE_DYN_FAST string = "DYNAMIC:FAST"

	SPEED_DEFAULT_LARGE_FILE_DEFAULT = SPEED_DEFAULT_LARGE_FILE_DYN_INTL

	SLAVE_DEFAULT_PING         = "http://gstatic.com/generate_204"
	SLAVE_DEFAULT_RETRY   uint = 3
	SLAVE_DEFAULT_TIMEOUT uint = 5000
)
