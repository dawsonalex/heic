package heic

/*
 * Copyright (c) 2021 Alexander F. Rødseth <xyproto@archlinux.org>
 *
 * heif is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * heif is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with heic.  If not, see <http://www.gnu.org/licenses/>.
 */

/*
#cgo pkg-config: libheif
#include <libheif/heif.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"

type MetadataID uint

func (h *ImageHandle) MetadataCount() int {
	n := int(C.heif_image_handle_get_number_of_metadata_blocks(h.handle, nil))
	keepAlive(h)
	return n
}

func (h *ImageHandle) MetadataIDs() []MetadataID {
	nMeta := h.MetadataCount()
	if nMeta == 0 {
		return []MetadataID{}
	}
	meta := make([]C.uint, nMeta)
	C.heif_image_handle_get_list_of_metadata_block_IDs(h.handle, nil, &meta[0], C.int(nMeta))
	metaDataIDs := make([]MetadataID, nMeta)
	for i := 0; i < nMeta; i++ {
		metaDataIDs[i] = MetadataID(meta[i])
	}
	keepAlive(h)
	return metaDataIDs
}

func (h *ImageHandle) ExifCount() int {
	filter := C.CString("Exif")
	n := int(C.heif_image_handle_get_number_of_metadata_blocks(h.handle, filter))
	keepAlive(h)
	return n
}

func (h *ImageHandle) ExifIDs() []MetadataID {
	nMeta := h.ExifCount()
	if nMeta == 0 {
		return []MetadataID{}
	}
	filter := C.CString("Exif")
	meta := make([]C.uint, nMeta)
	C.heif_image_handle_get_list_of_metadata_block_IDs(h.handle, filter, &meta[0], C.int(nMeta))
	metaDataIDs := make([]MetadataID, nMeta)
	for i := 0; i < nMeta; i++ {
		metaDataIDs[i] = MetadataID(meta[i])
	}
	keepAlive(h)
	return metaDataIDs
}
