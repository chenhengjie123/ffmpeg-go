# github.com/chenhengjie123/ffmpeg-go

golang binding for ffmpeg

如果想在 windows 上使用，建议直接使用原版：<https://github.com/moonfdd/ffmpeg-go>

以下说明均针对 mac 上使用

# 环境准备

## ffmpeg 编译

此项目通过 cgo 调用 ffmpeg ，因此运行环境需要事先完成 ffmpeg 的编译

参考命令如下

```sh
# 下载编译 x264 包，编码 h264 时需要使用
wget https://code.videolan.org/videolan/x264/-/archive/master/x264-master.zip
unzip x264-master.zip
cd x264-master
./configure --prefix=/usr/local/libx264 --enable-shared --enable-static
make 
sudo make install

# 下载编译 ffmpeg ，将其安装至 /usr/local/ffmpeg 路径下
# 并启用共享库（--enable-shared）、关联libx264（--enable-libx264 --enable-gpl --extra-cflags=-I/usr/local/libx264/include --extra-ldflags=-L/usr/local/libx264/）
wget  https://github.com/FFmpeg/FFmpeg/archive/refs/tags/n5.1.3.zip
unzip n5.1.3.zip
cd FFmpeg-n5.1.3
./configure --prefix=/usr/local/ffmpeg --enable-shared --enable-libx264 --enable-gpl --extra-cflags=-I/usr/local/libx264/include --extra-ldflags=-L/usr/local/libx264/
make -j8
sudo make install
```

## 【可选】修改项目中加载动态库相关代码，保障调用的是上一步编译的 ffmpeg 

如果编译安装命令和上一步示例命令完全一致，即 mac 上 ffmpeg 的安装路径为 `/usr/local/ffmpeg` ，此问题可忽略。

如果不是，需要全局搜索 `/usr/local/ffmpeg` ，并将其替换为实际路径。



# fast start

git clone https://github.com/chenhengjie123/ffmpeg-go.git

cd ffmpeg-go

go run ./examples/apitest/main.go


