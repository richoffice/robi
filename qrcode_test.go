package robi

import "testing"

func TestQrcode_CreateFile(t *testing.T) {
	content := "http://yylives.cc/abc123/test/xxx"
	q := &Qrcode{}
	err := q.CreateFile(content, "testfiles/qrcode.png")
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
}
