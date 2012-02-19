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

type MemFlags C.cl_mem_flags

const (
	MEM_READ_WRITE     MemFlags = C.CL_MEM_READ_WRITE
	MEM_WRITE_ONLY     MemFlags = C.CL_MEM_WRITE_ONLY
	MEM_READ_ONLY      MemFlags = C.CL_MEM_READ_ONLY
	MEM_USE_HOST_PTR   MemFlags = C.CL_MEM_USE_HOST_PTR
	MEM_ALLOC_HOST_PTR MemFlags = C.CL_MEM_ALLOC_HOST_PTR
	MEM_COPY_HOST_PTR  MemFlags = C.CL_MEM_COPY_HOST_PTR
)

const (
	MEM_OBJECT_BUFFER  = C.CL_MEM_OBJECT_BUFFER
	MEM_OBJECT_IMAGE2D = C.CL_MEM_OBJECT_IMAGE2D
	MEM_OBJECT_IMAGE3D = C.CL_MEM_OBJECT_IMAGE3D
)

func releaseMemObject(p C.cl_mem) error {
	if p != nil {
		if err := C.clReleaseMemObject(p); err != C.CL_SUCCESS {
			return Cl_error(err)
		}
	}
	return nil
}
