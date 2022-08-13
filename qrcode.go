package robi

import "github.com/skip2/go-qrcode"

type Qrcode struct {
}

func (q *Qrcode) CreateFile(content string, path string) error {
	return qrcode.WriteFile(content, qrcode.Medium, 256, path)
}
