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

//#ifndef AVUTIL_TX_H
//#define AVUTIL_TX_H
//
//#include <stdint.h>
//#include <stddef.h>

//typedef struct AVTXContext AVTXContext;
type AVTXContext struct {
}
type AVComplexFloat struct {
	Re, Im ffcommon.FFloat
}

type AVComplexDouble struct {
	Re, Im ffcommon.FDouble
}

type AVComplexInt32 struct {
	Re, Im ffcommon.FInt32T
}
type AVTXType int32

const (
	/**
	 * Standard complex to complex FFT with sample data type AVComplexFloat.
	 * Output is not 1/len normalized. Scaling currently unsupported.
	 * The stride parameter is ignored.
	 */
	AV_TX_FLOAT_FFT = 0

	/**
	 * Standard MDCT with sample data type of float and a scale type of
	 * float. Length is the frame size, not the window size (which is 2x frame)
	 * For forward transforms, the stride specifies the spacing between each
	 * sample in the output array in bytes. The input must be a flat array.
	 * For inverse transforms, the stride specifies the spacing between each
	 * sample in the input array in bytes. The output will be a flat array.
	 * Stride must be a non-zero multiple of sizeof(float).
	 * NOTE: the inverse transform is half-length, meaning the output will not
	 * contain redundant data. This is what most codecs work with.
	 */
	AV_TX_FLOAT_MDCT = 1

	/**
	 * Same as AV_TX_FLOAT_FFT with a data type of AVComplexDouble.
	 */
	AV_TX_DOUBLE_FFT = 2

	/**
	 * Same as AV_TX_FLOAT_MDCT with data and scale type of double.
	 * Stride must be a non-zero multiple of sizeof(double).
	 */
	AV_TX_DOUBLE_MDCT = 3

	/**
	 * Same as AV_TX_FLOAT_FFT with a data type of AVComplexInt32.
	 */
	AV_TX_INT32_FFT = 4

	/**
	 * Same as AV_TX_FLOAT_MDCT with data type of int32_t and scale type of float.
	 * Only scale values less than or equal to 1.0 are supported.
	 * Stride must be a non-zero multiple of sizeof(int32_t).
	 */
	AV_TX_INT32_MDCT = 5
)

/**
 * Function pointer to a function to perform the transform.
 *
 * @note Using a different context than the one allocated during av_tx_init()
 * is not allowed.
 *
 * @param s the transform context
 * @param out the output array
 * @param in the input array
 * @param stride the input or output stride in bytes
 *
 * The out and in arrays must be aligned to the maximum required by the CPU
 * architecture.
 * The stride must follow the constraints the transform type has specified.
 */
//typedef void (*av_tx_fn)(AVTXContext *s, void *out, void *in, ptrdiff_t stride);
type AvTxFn = func(s *AVTXContext, out ffcommon.FVoidP, in ffcommon.FVoidP, stride ffcommon.FPtrdiffT) uintptr

/**
 * Flags for av_tx_init()
 */
type AVTXFlags int32

const (
	/**
	 * Performs an in-place transformation on the input. The output argument
	 * of av_tn_fn() MUST match the input. May be unsupported or slower for some
	 * transform types.
	 */
	AV_TX_INPLACE = 1 << 0
)

/**
 * Initialize a transform context with the given configuration
 * (i)MDCTs with an odd length are currently not supported.
 *
 * @param ctx the context to allocate, will be NULL on error
 * @param tx pointer to the transform function pointer to set
 * @param type type the type of transform
 * @param inv whether to do an inverse or a forward transform
 * @param len the size of the transform in samples
 * @param scale pointer to the value to scale the output if supported by type
 * @param flags a bitmask of AVTXFlags or 0
 *
 * @return 0 on success, negative error code on failure
 */
//int av_tx_init(AVTXContext **ctx, av_tx_fn *tx, enum AVTXType type,
//int inv, int len, const void *scale, uint64_t flags);
func AvTxInit(ctx **AVTXContext, tx AvTxFn, type0 AVTXType,
	inv, len0 ffcommon.FInt, scale ffcommon.FVoidP, flags ffcommon.FUint64T) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_tx_init").Call(
		uintptr(unsafe.Pointer(ctx)),
		ffcommon.NewCallback(tx),
		uintptr(type0),
		uintptr(inv),
		uintptr(len0),
		scale,
		uintptr(flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Frees a context and sets ctx to NULL, does nothing when ctx == NULL
 */
//void av_tx_uninit(AVTXContext **ctx);
func AvTxUninit(ctx **AVTXContext) {
	ffcommon.GetAvutilDll().NewProc("av_tx_uninit").Call(
		uintptr(unsafe.Pointer(ctx)),
	)
}

//#endif /* AVUTIL_TX_H */
