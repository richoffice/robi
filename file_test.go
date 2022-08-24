package robi

import "testing"

func TestFile_RelativePath(t *testing.T) {
	f := &File{}
	expected := "../../refs/certs/images/licence.jpg"
	actual := f.RelativePath("bids/sandbox/测试项目.yaml", "refs/certs/licence.md", "./images/licence.jpg")
	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
