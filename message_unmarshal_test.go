package dhcp

import "testing"

func TestUnmarshalMessage(t *testing.T) {
	initMsgSamples(t)

	tester := func(t *testing.T, data []byte) {
		msg, err := UnmarshalMessage(data)
		if err != nil {
			t.Error(err.JSON())
		}

		if validErr := msg.IsValid(); validErr != nil {
			t.Error(validErr.JSON())
		}
	}

	for fileName, sample := range msgSamples {
		if sample.fileExt != ".raw" {
			continue
		}

		t.Run("Unmarshal message from testdata/samples/"+fileName, func(t *testing.T) {
			tester(t, sample.raw)
		})
	}
}
