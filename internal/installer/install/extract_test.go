package install_test

import (
	"os"
	"testing"

	"github.com/Tom5521/EduTrack/internal/installer/install"
)

func TestExtract(t *testing.T) {
	os.Chdir("/tmp")
	f, err := install.Files.ReadFile("files/EduTrack-linux64.tar.xz")
	if err != nil {
		t.Fatal(err)
	}
	//os.Mkdir("tttEst", os.ModePerm)
	err = install.ExtractTarXz(f)
	if err != nil {
		t.Fail()
	}
}
