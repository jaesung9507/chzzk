# Chzzk
Unofficial CHZZK API client for Go

## Example
Get HLS URL from the live detail endpoint:
```go
package main

import (
	"fmt"

	"github.com/jaesung9507/chzzk"
)

func main() {
	client := chzzk.NewClient(nil)

	liveDetail, err := client.GetLiveDetail("CHANNEL_ID")
	if err != nil {
		panic(err)
	}

	playbackInfo, err := liveDetail.GetLivePlayback()
	if err != nil {
		panic(err)
	}

	fmt.Println("HLS URL:", playbackInfo.GetHLSPath())
}
```

## Disclaimer
This Go client uses unofficial APIs of NAVER CHZZK and is not affiliated with them in any way.

CHZZK may change or disable these APIs at any time, which may cause this client to stop working without notice.

## License
MIT License
