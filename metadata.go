package heic

/*
#cgo pkg-config: libheif
#include <stdlib.h>
#include <string.h> // We use 'memcpy'
#include <libheif/heif.h>
*/
import "C"

func (h *ImageHandle) MetaDataCount() int {
	n := int(C.heif_image_handle_get_number_of_metadata_blocks(h.handle, nil))
	keepAlive(h)
	return n
}

func (h *ImageHandle) MetaDataIDs() []uint {
	nMeta := h.MetaDataCount()
	meta := make([]C.uint, nMeta)
	C.heif_image_handle_get_list_of_metadata_block_IDs(h.handle, nil, &meta[0], C.int(nMeta))
	metaDataIDs := make([]uint, nMeta)
	for i := 0; i < nMeta; i++ {
		metaDataIDs[i] = uint(meta[i])
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
