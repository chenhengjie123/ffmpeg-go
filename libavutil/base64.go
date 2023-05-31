package libavutil

import (
	"unsafe"

	"github.com/chenhengjie123/ffmpeg-go/ffcommon"
)

/*
 * Copyright (c) 2006 Ryan Martell. (rdm4@martellventures.com)
 *
 * This file is part of FFmpeg.
 *
 * FFmpeg is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * FFmpeg is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with FFmpeg; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

//#ifndef AVUTIL_BASE64_H
//#define AVUTIL_BASE64_H
//
//#include <stdint.h>

/**
 * @defgroup lavu_base64 Base64
 * @ingroup lavu_crypto
 * @{
 */

/**
 * Decode a base64-encoded string.
 *
 * @param out      buffer for decoded data
 * @param in       null-terminated input string
 * @param out_size size in bytes of the out buffer, must be at
 *                 least 3/4 of the length of in, that is AV_BASE64_DECODE_SIZE(strlen(in))
 * @return         number of bytes written, or a negative value in case of
 *                 invalid input
 */
//int av_base64_decode(uint8_t *out, const char *in, int out_size);
func AvBase64Decode(out *ffcommon.FUint8T, in ffcommon.FConstCharP, out_size ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_base64_decode").Call(
		uintptr(unsafe.Pointer(out)),
		ffcommon.UintPtrFromString(in),
		uintptr(out_size),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Calculate the output size in bytes needed to decode a base64 string
 * with length x to a data buffer.
 */
//#define AV_BASE64_DECODE_SIZE(x) ((x) * 3LL / 4)
func AV_BASE64_DECODE_SIZE(x int64) (res int64) {
	res = x * 3 / 4
	return
}

/**
 * Encode data to base64 and null-terminate.
 *
 * @param out      buffer for encoded data
 * @param out_size size in bytes of the out buffer (including the
 *                 null terminator), must be at least AV_BASE64_SIZE(in_size)
 * @param in       input buffer containing the data to encode
 * @param in_size  size in bytes of the in buffer
 * @return         out or NULL in case of error
 */
//char *av_base64_encode(char *out, int out_size, const uint8_t *in, int in_size);
func AvBase64Encode(out ffcommon.FCharP, out_size ffcommon.FInt, in *ffcommon.FUint8T, in_size ffcommon.FInt) (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_base64_encode").Call(
		ffcommon.UintPtrFromString(out),
		uintptr(out_size),
		uintptr(unsafe.Pointer(in)),
		uintptr(in_size),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Calculate the output size needed to base64-encode x bytes to a
 * null-terminated string.
 */
//#define AV_BASE64_SIZE(x)  (((x)+2) / 3 * 4 + 1)
func AV_BASE64_SIZE(x int64) (res int64) {
	res = (x+2)/3*4 + 1
	return
}

/**
 * @}
 */

//#endif /* AVUTIL_BASE64_H */
