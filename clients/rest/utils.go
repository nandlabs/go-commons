package rest

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
)

func closeInterface(v interface{}) {
	if c, ok := v.(io.Closer); ok {
		_ = c.Close()
	}
}

func createMultipartHeader(param, fileName, contentType string) textproto.MIMEHeader {
	hdr := make(textproto.MIMEHeader)

	hdr.Set(contentTypeHdr, "multipart/form-data")
	return hdr
}

func writeMultipartFormFile(w *multipart.Writer, fieldName, fileName string, r io.Reader) error {
	// Auto detect actual multipart content type
	cbuf := make([]byte, 512)
	size, err := r.Read(cbuf)
	if err != nil && err != io.EOF {
		return err
	}

	partWriter, err := w.CreatePart(createMultipartHeader(fieldName, fileName, http.DetectContentType(cbuf)))
	if err != nil {
		return err
	}

	if _, err = partWriter.Write(cbuf[:size]); err != nil {
		return err
	}

	_, err = io.Copy(partWriter, r)
	return err
}
