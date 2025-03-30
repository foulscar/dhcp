package dhcp

import "testing"

func TestUnmarshalMessage(t *testing.T) {
	initMsgSamples(t)

	tester := func(t *testing.T, data []byte) {
		msg, err := UnmarshalMessage(data)
		if err != nil {
			t.Errorf("Error unmarshalling. %s", err)
		}

		valid, reason := msg.IsValid()
		if !valid {
			t.Errorf("Unmarshalled message is invalid. reason: %s", reason)
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
