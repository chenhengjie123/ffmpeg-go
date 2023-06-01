package main

import (
	"fmt"
	"os"

	"github.com/chenhengjie123/ffmpeg-go/ffcommon"
	"github.com/chenhengjie123/ffmpeg-go/libavutil"
)

func main() {
	os.Setenv("Path", os.Getenv("Path")+";./lib")
	ffcommon.SetAvcodecPath("/usr/local/ffmpeg/lib/libavcodec.dylib")
	ffcommon.SetAvutilPath("/usr/local/ffmpeg/lib/libavutil.dylib")
	ffcommon.SetAvdevicePath("/usr/local/ffmpeg/lib/libavdevice.dylib")
	ffcommon.SetAvfilterPath("/usr/local/ffmpeg/lib/libavfilter.dylib")
	ffcommon.SetAvformatPath("/usr/local/ffmpeg/lib/libavformat.dylib")
	ffcommon.SetAvpostprocPath("/usr/local/ffmpeg/lib/libpostproc.dylib")
	ffcommon.SetAvswresamplePath("/usr/local/ffmpeg/lib/libswresample.dylib")
	ffcommon.SetAvswscalePath("/usr/local/ffmpeg/lib/libswscale.dylib")
	if true {
		ret := libavutil.AvFrameAlloc()
		fmt.Println(ret)
		fmt.Println(libavutil.AV_NUM_DATA_POINTERS)
	}
	// if true {
	// 	libavutil.AvLog(0, 0, "a", "b")
	// }
	// if true {
	// 	ret := libavcodec.AvcodecVersion()
	// 	fmt.Println(ret)
	// }
	// if true {
	// 	fmt.Println(libavutil.AvVersionInfo())
	// }
	// if true {
	// 	fmt.Println(libavcodec.AvcodecLicense())
	// }
	// if true {
	// 	fmt.Println(libavcodec.AvcodecConfiguration())
	// }
	// if true {
	// 	ans := libavutil.AvAdler32Update(111, nil, 0)
	// 	fmt.Println(ans)
	// }
	// if true {
	// 	ans := libavutil.AvAesAlloc()
	// 	fmt.Println(ans)
	// }
	// if true {
	// 	fmt.Println(libavutil.AV_MATRIX_ENCODING_DOLBYHEADPHONE)
	// }
	// if true {
	// 	fmt.Println(libavutil.AvutilVersion())
	// }
	// if true {
	// 	fmt.Println(libavutil.AvVersionInfo())
	// }
	// if true {
	// 	fmt.Println(libavutil.AvutilConfiguration())
	// }
	// if true {
	// 	fmt.Println(libavutil.AvutilLicense())
	// }
	// if true {
	// 	fmt.Println(libavutil.AVMEDIA_TYPE_VIDEO)
	// 	fmt.Println(libavutil.AvGetMediaTypeString(libavutil.AVMEDIA_TYPE_AUDIO))
	// }
	// if true {
	// 	fmt.Println(libavutil.AvGetTimeBaseQ())
	// 	libavutil.AvGetTimeBaseQ()
	// }
}
