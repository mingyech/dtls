package dtls

import (
	"reflect"
	"testing"
)

func TestFragmentBuffer(t *testing.T) {
	for _, test := range []struct {
		Name     string
		In       [][]byte
		Expected []byte
	}{
		{
			Name: "Single Fragment",
			In: [][]byte{
				[]byte{0x16, 0xfe, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0F, 0x03, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xfe, 0xff, 0x00},
			},
			Expected: []byte{0x03, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xfe, 0xff, 0x00},
		},
		{
			Name: "Multiple Fragments",
			In: [][]byte{
				[]byte{0x16, 0xfe, 0xfd, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x81, 0x0b, 0x00, 0x00, 0x0F, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x02, 0x03, 0x04},
				[]byte{0x16, 0xfe, 0xfd, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x81, 0x0b, 0x00, 0x00, 0x0F, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x05, 0x05, 0x06, 0x07, 0x08, 0x09},
				[]byte{0x16, 0xfe, 0xfd, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x81, 0x0b, 0x00, 0x00, 0x0F, 0x00, 0x00, 0x00, 0x00, 0x0A, 0x00, 0x00, 0x05, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E},
			},
			Expected: []byte{0x0b, 0x00, 0x00, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e},
		},
		{
			Name: "Multiple Unordered Fragments",
			In: [][]byte{
				[]byte{0x16, 0xfe, 0xfd, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x81, 0x0b, 0x00, 0x00, 0x0F, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x02, 0x03, 0x04},
				[]byte{0x16, 0xfe, 0xfd, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x81, 0x0b, 0x00, 0x00, 0x0F, 0x00, 0x00, 0x00, 0x00, 0x0A, 0x00, 0x00, 0x05, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E},
				[]byte{0x16, 0xfe, 0xfd, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x81, 0x0b, 0x00, 0x00, 0x0F, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x05, 0x05, 0x06, 0x07, 0x08, 0x09},
			},
			Expected: []byte{0x0b, 0x00, 0x00, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e},
		},
	} {
		fragmentBuffer := newFragmentBuffer()
		for _, frag := range test.In {
			status, err := fragmentBuffer.push(frag)
			if err != nil {
				t.Error(err)
			} else if !status {
				t.Errorf("fragmentBuffer didn't accept fragments for '%s'", test.Name)
			}
		}

		out, _ := fragmentBuffer.pop()
		if !reflect.DeepEqual(out, test.Expected) {
			t.Errorf("fragmentBuffer '%s' push/pop: got % 02x, want % 02x", test.Name, out, test.Expected)
		}

		if frag, _ := fragmentBuffer.pop(); frag != nil {
			t.Errorf("fragmentBuffer popped single buffer multiple times for '%s'", test.Name)
		}
	}
}
