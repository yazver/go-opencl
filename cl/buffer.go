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
#cgo CFLAGS: -I . -Wno-error=deprecated-declarations
#cgo linux LDFLAGS: -lOpenCL
#cgo windows LDFLAGS: -lOpenCL
#cgo darwin LDFLAGS: -framework OpenCL

#ifdef MAC
	#include "OpenCL/cl.h"
#else
	#include "CL/opencl.h"
#endif //MAC

*/
import "C"

type Buffer struct {
	id C.cl_mem
}

func (b *Buffer) release() error {
	err := releaseMemObject(b.id)
	b.id = nil
	return err
}
