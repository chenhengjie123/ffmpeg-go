package main

import (
	"fmt"
	"os"

	"github.com/chenhengjie123/ffmpeg-go/ffcommon"
	"github.com/chenhengjie123/ffmpeg-go/libavformat"
	"github.com/chenhengjie123/ffmpeg-go/libavutil"
)

func main() {

	os.Setenv("Path", os.Getenv("Path")+";./lib")
	ffcommon.SetAvcodecPath("./lib_mac/libavcodec.dylib")
	ffcommon.SetAvutilPath("./lib_mac/libavutil.dylib")
	ffcommon.SetAvdevicePath("./lib_mac/libavdevice.dylib")
	ffcommon.SetAvfilterPath("./lib_mac/libavfilter.dylib")
	ffcommon.SetAvformatPath("./lib_mac/libavformat.dylib")
	ffcommon.SetAvpostprocPath("./lib_mac/libpostproc.dylib")
	ffcommon.SetAvswresamplePath("./lib_mac/libswresample.dylib")
	ffcommon.SetAvswscalePath("./lib_mac/libswscale.dylib")

	genDir := "./out"
	_, err := os.Stat(genDir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(genDir, 0777) //  Everyone can read write and execute
		}
	}
	main0()
}

func main0() (ret ffcommon.FInt) {
	var fmt_ctx *libavformat.AVFormatContext
	var tag *libavutil.AVDictionaryEntry

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <input_file>\nexample program to demonstrate the use of the libavformat metadata API.\n\n", os.Args[0])
		return 1
	}
	ret = libavformat.AvformatOpenInput(&fmt_ctx, os.Args[1], nil, nil)
	if ret != 0 {
		return ret
	}

	ret = fmt_ctx.AvformatFindStreamInfo(nil)
	if ret < 0 {
		libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Cannot find stream information\n")
		return ret
	}

	tag = fmt_ctx.Metadata.AvDictGet("", tag, libavutil.AV_DICT_IGNORE_SUFFIX)
	for tag != nil {
		fmt.Printf("%s=%s\n", ffcommon.StringFromPtr(tag.Key), ffcommon.StringFromPtr(tag.Value))
		tag = fmt_ctx.Metadata.AvDictGet("", tag, libavutil.AV_DICT_IGNORE_SUFFIX)
	}

	libavformat.AvformatCloseInput(&fmt_ctx)
	return 0
}
