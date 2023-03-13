package speed

import (
	"context"
	"crypto/rand"
	"io"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/juju/ratelimit"

	"github.com/miaokobot/miaospeed/interfaces"
	"github.com/miaokobot/miaospeed/preconfigs"
	"github.com/miaokobot/miaospeed/utils"
	"github.com/miaokobot/miaospeed/utils/structs"
	"github.com/miaokobot/miaospeed/vendors"
)

func Once(speed *Speed, proxy interfaces.Vendor, cfg *interfaces.SlaveRequestConfigs) {
	speed.Speeds = make([]uint64, cfg.DownloadDuration)

	downloadFiles := RefetchDownloadFiles(proxy, cfg.DownloadURL)
	utils.DLogf("Speed Prefetch | Using files arr=%v", downloadFiles)

	th := int(cfg.DownloadThreading)
	wcGroups := []*WriteCounter{}
	ctxCancels := []context.CancelFunc{}

	initWG := sync.WaitGroup{}
	writingLock := sync.Mutex{}
	for i := 0; i < th; i++ {
		initWG.Add(1)
		go func() {
			wc := WriteCounter{
				RateLimit: int64(utils.GCFG.SpeedLimit) / int64(th),
			}
			cancelFunc := SingleThread(downloadFiles, proxy, cfg.DownloadDuration, &wc)

			writingLock.Lock()
			ctxCancels = append(ctxCancels, cancelFunc)
			wcGroups = append(wcGroups, &wc)
			writingLock.Unlock()

			initWG.Done()
		}()
	}
	initWG.Wait()

	// normalization
	for i := 0; i < th; i++ {
		wcGroups[i].Take()
	}

	for t := 0; t < int(cfg.DownloadDuration); t++ {
		time.Sleep(time.Second - time.Millisecond*10)
		byteLen := uint64(0)
		for i := 0; i < th; i++ {
			threadLen := wcGroups[i].Take()
			utils.DLogf("Task Thread | time=%d thread=%d speed=%d", t+1, i+1, threadLen)
			byteLen += threadLen
		}
		speed.Speeds[t] = byteLen
		speed.TotalSize += byteLen
		speed.MaxSpeed = structs.Max(speed.MaxSpeed, byteLen)
	}
	speed.AvgSpeed = speed.TotalSize / uint64(cfg.DownloadDuration)

	for i := 0; i < th; i++ {
		ctxCancels[i]()
	}
}

func SingleThread(downloadFiles []string, proxy interfaces.Vendor, timeoutSeconds int64, wc *WriteCounter) context.CancelFunc {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds+1)*time.Second)
	isCancelled := false

	downloadFilesCopy := downloadFiles[:]
	fileLen := len(downloadFilesCopy)
	readyChan := make(chan bool)

	go func() {
		isReady := false
		defer func() {
			if !isReady {
				close(readyChan)
			}
		}()

		// 100 only for safty
		for i := 0; i < 100; i++ {
			// if outside cancel or deadline meet(either by time or by hand)
			if isCancelled || ctx.Err() != nil {
				return
			}
			// download file
			file := downloadFilesCopy[i%fileLen]
			resp, _, err := vendors.RequestUnsafe(ctx, proxy, &interfaces.RequestOptions{
				URL: file,
			})

			if !isReady {
				isReady = true
				close(readyChan)
			}
			if err == nil {
				var bodyReader io.Reader = nil
				if wc.RateLimit >= 1024 {
					bucket := ratelimit.NewBucketWithRate(float64(wc.RateLimit)*0.95, wc.RateLimit)
					bodyReader = ratelimit.Reader(resp.Body, bucket)
				} else {
					bodyReader = resp.Body
				}
				io.Copy(io.Discard, io.TeeReader(bodyReader, wc))
			}
			// close body
			if resp != nil && resp.Body != nil {
				resp.Body.Close()
			}
		}
	}()

	<-readyChan
	return func() {
		isCancelled = true
		cancel()
	}
}

func RefetchDownloadFiles(proxy interfaces.Vendor, file string) []string {
	defaultList := []string{preconfigs.SPEED_DEFAULT_LARGE_FILE_STATIC_MSFT}
	if proxy == nil || proxy.Status() == interfaces.VStatusNotReady {
		return defaultList
	}

	switch file {
	case preconfigs.SPEED_DEFAULT_LARGE_FILE_DYN_INTL:
		body, _, _ := vendors.RequestWithRetry(proxy, 1, 1000, &interfaces.RequestOptions{
			URL:     "https://ipinfo.io",
			NoRedir: true,
		})

		if strings.Contains(string(body), "Microsoft") {
			return []string{preconfigs.SPEED_DEFAULT_LARGE_FILE_STATIC_MSFT}
		} else if strings.Contains(string(body), "Google") {
			return []string{preconfigs.SPEED_DEFAULT_LARGE_FILE_STATIC_GOOGLE}
		} else {
			return []string{preconfigs.SPEED_DEFAULT_LARGE_FILE_STATIC_APPLE}
		}

	case preconfigs.SPEED_DEFAULT_LARGE_FILE_DYN_ALL_INTL:
		FileUrl := getRandomUrl(preconfigs.SPEED_DEFAULT_LARGE_FILE_DYNAMIC)
		return []string{FileUrl}

	case preconfigs.SPEED_DEFAULT_LARGE_FILE_DYN_FAST:
		body, _, _ := vendors.RequestWithRetry(proxy, 3, 1000, &interfaces.RequestOptions{
			URL:     "https://api.fast.com/netflix/speedtest/v2?https=false&token=YXNkZmFzZGxmbnNkYWZoYXNkZmhrYWxm&urlCount=5",
			NoRedir: true,
		})
		url := jsoniter.Get(body, "targets", 0, "url").ToString()
		if url != "" {
			return []string{url}
		} else {
			return defaultList
		}
	}
	return []string{file}
}
func getRandomUrl(url string) string {
	SpeedUrl := strings.Split(strings.ReplaceAll(url, "\r\n", "\n"), "\n")
	Random, _ := rand.Int(rand.Reader, big.NewInt(int64(len(SpeedUrl))))
	RandomNum, _ := strconv.Atoi(Random.String())
	return SpeedUrl[RandomNum]
}
