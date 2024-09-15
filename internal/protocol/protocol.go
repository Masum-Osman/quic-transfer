package protocol

import (
	"encoding/binary"
	"io"
)

type FileInfo struct {
	Name string
	Size int64
}

func SendFileInfo(w io.Writer, info FileInfo) error {
	if err := binary.Write(w, binary.BigEndian, int64(len(info.Name))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(info.Name)); err != nil {
		return err
	}
	return binary.Write(w, binary.BigEndian, info.Size)
}

func ReceiveFileInfo(r io.Reader) (FileInfo, error) {
	var nameLen int64
	if err := binary.Read(r, binary.BigEndian, &nameLen); err != nil {
		return FileInfo{}, err
	}

	nameBuf := make([]byte, nameLen)
	if _, err := io.ReadFull(r, nameBuf); err != nil {
		return FileInfo{}, err
	}

	var size int64
	if err := binary.Read(r, binary.BigEndian, &size); err != nil {
		return FileInfo{}, err
	}

	return FileInfo{
		Name: string(nameBuf),
		Size: size,
	}, nil
}
