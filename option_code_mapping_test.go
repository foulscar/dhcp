package dhcp

import (
	"testing"
)

func TestGlobalOptionCodeMapping(t *testing.T) {
	optMap := &GlobalOptionCodeMapping

	t.Run("len(ToDataType) == len(ToDataUnmarshaller) must be true", func(t *testing.T) {
		dataTypeLen := len(optMap.ToDataType)
		dataUnmarshallerLen := len(optMap.ToDataUnmarshaller)
		if dataTypeLen != dataUnmarshallerLen {
			t.Errorf(
				"len(GlobalOptionCodeMapping.ToDataType) = %d\n"+
					"len(GlobalOptionCodeMapping.ToDataUnmarshaller = %d)",
				dataTypeLen,
				dataUnmarshallerLen)
		}
	})

	t.Run("All values in ToDataType must not be nil", func(t *testing.T) {
		for code, dataType := range optMap.ToDataType {
			if dataType == nil {
				t.Errorf("GlobalOptionCodeMapping.ToDataType[%d] is nil", code)
			}
		}
	})

	t.Run("All entries in ToDataType must have a corresponding entry in ToDataUnmarshaller and ToString", func(t *testing.T) {
		for code := range optMap.ToDataType {
			_, stringOk := optMap.ToString[code]
			_, dataUnmarshallerOk := optMap.ToDataUnmarshaller[code]
			if !stringOk {
				t.Errorf("GlobalOptionCodeMapping.ToString[%d] does not exist", code)
			}
			if !dataUnmarshallerOk {
				t.Errorf("GlobalOptionCodeMapping.ToDataUnmarshaller[%d] does not exist", code)
			}
		}
	})

	t.Run("All entries in ToDataUnmarshaller must have a corresponding entry in ToDataType and ToString", func(t *testing.T) {
		for code := range optMap.ToDataUnmarshaller {
			_, stringOk := optMap.ToString[code]
			_, dataTypeOk := optMap.ToDataType[code]
			if !stringOk {
				t.Errorf("GlobalOptionCodeMapping.ToString[%d] does not exist", code)
			}
			if !dataTypeOk {
				t.Errorf("GlobalOptionCodeMapping.ToDataType[%d] does not exist", code)
			}
		}
	})
}
