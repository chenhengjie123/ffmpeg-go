package libavutil

import (
	"unsafe"

	"github.com/chenhengjie123/ffmpeg-go/ffcommon"
)

/*
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

//#ifndef AVUTIL_PIXELUTILS_H
//#define AVUTIL_PIXELUTILS_H
//
//#include <stddef.h>
//#include <stdint.h>
//#include "common.h"

/**
 * Sum of abs(src1[x] - src2[x])
 */
//typedef int (*av_pixelutils_sad_fn)(const uint8_t *src1, ptrdiff_t stride1,
//const uint8_t *src2, ptrdiff_t stride2);
type AvPixelutilsSadFn = func(src1 *ffcommon.FUint8T, stride1 ffcommon.FPtrdiffT, src2 *ffcommon.FUint8T, stride2 ffcommon.FPtrdiffT) uintptr

/**
 * Get a potentially optimized pointer to a Sum-of-absolute-differences
 * function (see the av_pixelutils_sad_fn prototype).
 *
 * @param w_bits  1<<w_bits is the requested width of the block size
 * @param h_bits  1<<h_bits is the requested height of the block size
 * @param aligned If set to 2, the returned sad function will assume src1 and
 *                src2 addresses are aligned on the block size.
 *                If set to 1, the returned sad function will assume src1 is
 *                aligned on the block size.
 *                If set to 0, the returned sad function assume no particular
 *                alignment.
 * @param log_ctx context used for logging, can be NULL
 *
 * @return a pointer to the SAD function or NULL in case of error (because of
 *         invalid parameters)
 */
//av_pixelutils_sad_fn av_pixelutils_get_sad_fn(int w_bits, int h_bits,
//int aligned, void *log_ctx);
func AvPixelutilsGetSadFn(w_bits, h_bits,
	aligned ffcommon.FInt, log_ctx ffcommon.FVoidP) (res AvPixelutilsSadFn) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_pixelutils_get_sad_fn").Call(
		uintptr(w_bits),
		uintptr(h_bits),
		uintptr(aligned),
		log_ctx,
	)
	res = *(*AvPixelutilsSadFn)(unsafe.Pointer(t))
	return
}

//#endif /* AVUTIL_PIXELUTILS_H */
