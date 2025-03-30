package dhcp

import (
	"os"
	"path/filepath"
	"testing"
)

type msgSample struct {
	raw     []byte
	msg     *Message
	fileExt string
}

var msgSamples = make(map[string]msgSample)

func initMsgSamples(t *testing.T) {
	if len(msgSamples) > 0 {
		return
	}

	files, err := os.ReadDir("testdata/samples")
	if err != nil {
		t.Fatal("could not read directory 'testdata/samples'.", err)
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".raw" {
			continue
		}

		data, err := os.ReadFile("testdata/samples/" + file.Name())
		if err != nil {
			t.Fatal("could not read 'testdata/samples/"+file.Name()+"'", err)
		}

		if !IsEncodedMessage(data) {
			t.Fatal("'testdata/samples/"+file.Name()+"'", "is not an encoded message")
		}

		msgSamples[file.Name()] = msgSample{
			raw:     data,
			msg:     nil,
			fileExt: ".raw",
		}
	}
}

func TestIfNewMessageIsValid(t *testing.T) {
	err := NewMessage().IsValid()
	if err != nil {
		t.Error(err.JSON())
	}
}
