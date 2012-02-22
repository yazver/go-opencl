/*
 * Copyright © 2012 go-opencl authors
 *
 * This file is part of go-opencl.
 *
 * go-opencl is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * go-opencl is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with go-opencl.  If not, see <http://www.gnu.org/licenses/>.
 */

package cl

/*
#cgo CFLAGS: -I CL
#cgo LDFLAGS: -lOpenCL

#include "CL/opencl.h"
*/
import "C"

import "unsafe"

type channelOrder C.cl_channel_order

const (
	R         channelOrder = C.CL_R
	A         channelOrder = C.CL_A
	RG        channelOrder = C.CL_RG
	RA        channelOrder = C.CL_RA
	RGB       channelOrder = C.CL_RGB
	RGBA      channelOrder = C.CL_RGBA
	BGRA      channelOrder = C.CL_BGRA
	ARGB      channelOrder = C.CL_ARGB
	INTENSITY channelOrder = C.CL_INTENSITY
	LUMINANCE channelOrder = C.CL_LUMINANCE
	Rx        channelOrder = C.CL_Rx
	RGx       channelOrder = C.CL_RGx
	RGBx      channelOrder = C.CL_RGBx
)

type channelType C.cl_channel_type

const (
	FLOAT            channelType = C.CL_FLOAT
	HALF_FLOAT       channelType = C.CL_HALF_FLOAT
	SIGNED_INT8      channelType = C.CL_SIGNED_INT8
	SIGNED_INT16     channelType = C.CL_SIGNED_INT16
	SIGNED_INT32     channelType = C.CL_SIGNED_INT32
	SNORM_INT8       channelType = C.CL_SNORM_INT8
	SNORM_INT16      channelType = C.CL_SNORM_INT16
	UNORM_INT8       channelType = C.CL_UNORM_INT8
	UNORM_INT16      channelType = C.CL_UNORM_INT16
	UNORM_SHORT_565  channelType = C.CL_UNORM_SHORT_565
	UNORM_SHORT_555  channelType = C.CL_UNORM_SHORT_555
	UNORM_INT_101010 channelType = C.CL_UNORM_INT_101010
	UNSIGNED_INT8    channelType = C.CL_UNSIGNED_INT8
	UNSIGNED_INT16   channelType = C.CL_UNSIGNED_INT16
	UNSIGNED_INT32   channelType = C.CL_UNSIGNED_INT32
)

type ImageFormat struct {
	ChannelOrder    channelOrder
	ChannelDataType channelType
}

type AddressingMode C.cl_addressing_mode

const (
	ADDRESS_NONE            AddressingMode = C.CL_ADDRESS_NONE
	ADDRESS_CLAMP_TO_EDGE   AddressingMode = C.CL_ADDRESS_CLAMP_TO_EDGE
	ADDRESS_CLAMP           AddressingMode = C.CL_ADDRESS_CLAMP
	ADDRESS_REPEAT          AddressingMode = C.CL_ADDRESS_REPEAT
	ADDRESS_MIRRORED_REPEAT AddressingMode = C.CL_ADDRESS_MIRRORED_REPEAT
)

/* cl_filter_mode */
type FilterMode C.cl_filter_mode

const (
	FILTER_NEAREST = C.CL_FILTER_NEAREST
	FILTER_LINEAR  = C.CL_FILTER_LINEAR
)

type Image struct {
	id      C.cl_mem
	w, h, d uint32
}

func (im *Image) release() error {
	err := releaseMemObject(im.id)
	im.id = nil
	return err
}

/* cl_image_info */
type ImageInfo C.cl_image_info

const (
	IMAGE_FORMAT       ImageInfo = C.CL_IMAGE_FORMAT
	IMAGE_ELEMENT_SIZE ImageInfo = C.CL_IMAGE_ELEMENT_SIZE
	IMAGE_ROW_PITCH    ImageInfo = C.CL_IMAGE_ROW_PITCH
	IMAGE_SLICE_PITCH  ImageInfo = C.CL_IMAGE_SLICE_PITCH
	IMAGE_WIDTH        ImageInfo = C.CL_IMAGE_WIDTH
	IMAGE_HEIGHT       ImageInfo = C.CL_IMAGE_HEIGHT
	IMAGE_DEPTH        ImageInfo = C.CL_IMAGE_DEPTH
)

func (im *Image) Info(param ImageInfo) (uint64, error) {
	var ret uint64
	if err := C.clGetImageInfo(im.id, C.cl_image_info(param), C.size_t(8), unsafe.Pointer(&ret), nil); err != C.CL_SUCCESS {
		return 0, Cl_error(err)
	}
	return ret, nil
}

type Sampler struct {
	id C.cl_sampler
}

func (s *Sampler) release() error {
	if s.id != nil {
		if err := C.clReleaseSampler(C.cl_sampler(s.id)); err != C.CL_SUCCESS {
			return Cl_error(err)
		}
	}
	return nil
}

type SamplerInfo C.cl_sampler_info

const (
	SAMPLER_REFERENCE_COUNT   SamplerInfo = C.CL_SAMPLER_REFERENCE_COUNT
	SAMPLER_CONTEXT           SamplerInfo = C.CL_SAMPLER_CONTEXT
	SAMPLER_NORMALIZED_COORDS SamplerInfo = C.CL_SAMPLER_NORMALIZED_COORDS
	SAMPLER_ADDRESSING_MODE   SamplerInfo = C.CL_SAMPLER_ADDRESSING_MODE
	SAMPLER_FILTER_MODE       SamplerInfo = C.CL_SAMPLER_FILTER_MODE
)

func (s *Sampler) Info(param SamplerInfo) (uint32, error) {
	var ret uint32
	var c_size_ret C.size_t
	if err := C.clGetSamplerInfo(s.id, C.cl_sampler_info(param), C.size_t(4), unsafe.Pointer(&ret), &c_size_ret); err != C.CL_SUCCESS {
		return 0, Cl_error(err)
	}
	return ret, nil
}
