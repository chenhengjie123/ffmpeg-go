package libavutil

import (
	"unsafe"

	"github.com/chenhengjie123/ffmpeg-go/ffcommon"
)

/*
 * copyright (c) 2006 Michael Niedermayer <michaelni@gmx.at>
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

/**
 * @file
 * @ingroup lavu_md5
 * Public header for MD5 hash function implementation.
 */

//#ifndef AVUTIL_MD5_H
//#define AVUTIL_MD5_H
//
//#include <stddef.h>
//#include <stdint.h>
//
//#include "attributes.h"
//#include "version.h"

/**
 * @defgroup lavu_md5 MD5
 * @ingroup lavu_hash
 * MD5 hash function implementation.
 *
 * @{
 */

//extern const int av_md5_size;
//
//struct AVMD5;
type AVMD5 struct {
}

/**
 * Allocate an AVMD5 context.
 */
//struct AVMD5 *av_md5_alloc(void);
func AvMd5Alloc() (res *AVMD5) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_md5_alloc").Call()
	res = (*AVMD5)(unsafe.Pointer(t))
	return
}

/**
 * Initialize MD5 hashing.
 *
 * @param ctx pointer to the function context (of size av_md5_size)
 */
//void av_md5_init(struct AVMD5 *ctx);
func (ctx *AVMD5) AvMd5Init() {
	ffcommon.GetAvutilDll().NewProc("av_md5_init").Call(
		uintptr(unsafe.Pointer(ctx)),
	)
}

/**
 * Update hash value.
 *
 * @param ctx hash function context
 * @param src input data to update hash with
 * @param len input data length
 */
//#if FF_API_CRYPTO_SIZE_T
//void av_md5_update(struct AVMD5 *ctx, const uint8_t *src, int len);
//#else
//void av_md5_update(struct AVMD5 *ctx, const uint8_t *src, size_t len);
//#endif
func (ctx *AVMD5) AvMd5Update(src *ffcommon.FUint8T, len0 ffcommon.FUnsignedIntOrSizeT) {
	ffcommon.GetAvutilDll().NewProc("av_md5_update").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(src)),
		uintptr(len0),
	)
}

/**
 * Finish hashing and output digest value.
 *
 * @param ctx hash function context
 * @param dst buffer where output digest value is stored
 */
//void av_md5_final(struct AVMD5 *ctx, uint8_t *dst);
func (ctx *AVMD5) AvMd5Final(dst *ffcommon.FUint8T) {
	ffcommon.GetAvutilDll().NewProc("av_md5_final").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(dst)),
	)
}

/**
 * Hash an array of data.
 *
 * @param dst The output buffer to write the digest into
 * @param src The data to hash
 * @param len The length of the data, in bytes
 */
//#if FF_API_CRYPTO_SIZE_T
//void av_md5_sum(uint8_t *dst, const uint8_t *src, const int len);
//#else
//void av_md5_sum(uint8_t *dst, const uint8_t *src, size_t len);
//#endif
func AvMd5Sum(dst, src *ffcommon.FUint8T, len0 ffcommon.FUnsignedIntOrSizeT) {
	ffcommon.GetAvutilDll().NewProc("av_md5_sum").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(len0),
	)
}

/**
 * @}
 */

//#endif /* AVUTIL_MD5_H */
