package handshake

import (
	"reflect"
	"testing"

	"github.com/mingyech/dtls/v2/internal/ciphersuite/types"
	"github.com/mingyech/dtls/v2/pkg/crypto/elliptic"
	"github.com/mingyech/dtls/v2/pkg/crypto/hash"
	"github.com/mingyech/dtls/v2/pkg/crypto/signature"
)

func TestHandshakeMessageServerKeyExchange(t *testing.T) {
	test := func(rawServerKeyExchange []byte, parsedServerKeyExchange *MessageServerKeyExchange) {
		c := &MessageServerKeyExchange{
			KeyExchangeAlgorithm: types.KeyExchangeAlgorithmEcdhe,
		}
		if err := c.Unmarshal(rawServerKeyExchange); err != nil {
			t.Error(err)
		} else if !reflect.DeepEqual(c, parsedServerKeyExchange) {
			t.Errorf("handshakeMessageServerKeyExchange unmarshal: got %#v, want %#v", c, parsedServerKeyExchange)
		}

		raw, err := c.Marshal()
		if err != nil {
			t.Error(err)
		} else if !reflect.DeepEqual(raw, rawServerKeyExchange) {
			t.Errorf("handshakeMessageServerKeyExchange marshal: got %#v, want %#v", raw, rawServerKeyExchange)
		}
	}

	t.Run("Hash+Signature", func(t *testing.T) {
		rawServerKeyExchange := []byte{
			0x03, 0x00, 0x1d, 0x41, 0x04, 0x0c, 0xb9, 0xa3, 0xb9, 0x90, 0x71, 0x35, 0x4a, 0x08, 0x66, 0xaf,
			0xd6, 0x88, 0x58, 0x29, 0x69, 0x98, 0xf1, 0x87, 0x0f, 0xb5, 0xa8, 0xcd, 0x92, 0xf6, 0x2b, 0x08,
			0x0c, 0xd4, 0x16, 0x5b, 0xcc, 0x81, 0xf2, 0x58, 0x91, 0x8e, 0x62, 0xdf, 0xc1, 0xec, 0x72, 0xe8,
			0x47, 0x24, 0x42, 0x96, 0xb8, 0x7b, 0xee, 0xe7, 0x0d, 0xdc, 0x44, 0xec, 0xf3, 0x97, 0x6b, 0x1b,
			0x45, 0x28, 0xac, 0x3f, 0x35, 0x02, 0x03, 0x00, 0x47, 0x30, 0x45, 0x02, 0x21, 0x00, 0xb2, 0x0b,
			0x22, 0x95, 0x3d, 0x56, 0x57, 0x6a, 0x3f, 0x85, 0x30, 0x6f, 0x55, 0xc3, 0xf4, 0x24, 0x1b, 0x21,
			0x07, 0xe5, 0xdf, 0xba, 0x24, 0x02, 0x68, 0x95, 0x1f, 0x6e, 0x13, 0xbd, 0x9f, 0xaa, 0x02, 0x20,
			0x49, 0x9c, 0x9d, 0xdf, 0x84, 0x60, 0x33, 0x27, 0x96, 0x9e, 0x58, 0x6d, 0x72, 0x13, 0xe7, 0x3a,
			0xe8, 0xdf, 0x43, 0x75, 0xc7, 0xb9, 0x37, 0x6e, 0x90, 0xe5, 0x3b, 0x81, 0xd4, 0xda, 0x68, 0xcd,
		}
		parsedServerKeyExchange := &MessageServerKeyExchange{
			EllipticCurveType:    elliptic.CurveTypeNamedCurve,
			NamedCurve:           elliptic.X25519,
			PublicKey:            rawServerKeyExchange[4:69],
			HashAlgorithm:        hash.SHA1,
			SignatureAlgorithm:   signature.ECDSA,
			Signature:            rawServerKeyExchange[73:144],
			KeyExchangeAlgorithm: types.KeyExchangeAlgorithmEcdhe,
		}

		test(rawServerKeyExchange, parsedServerKeyExchange)
	})

	t.Run("Anonymous", func(t *testing.T) {
		rawServerKeyExchange := []byte{
			0x03, 0x00, 0x1d, 0x41, 0x04, 0x0c, 0xb9, 0xa3, 0xb9, 0x90, 0x71, 0x35, 0x4a, 0x08, 0x66, 0xaf,
			0xd6, 0x88, 0x58, 0x29, 0x69, 0x98, 0xf1, 0x87, 0x0f, 0xb5, 0xa8, 0xcd, 0x92, 0xf6, 0x2b, 0x08,
			0x0c, 0xd4, 0x16, 0x5b, 0xcc, 0x81, 0xf2, 0x58, 0x91, 0x8e, 0x62, 0xdf, 0xc1, 0xec, 0x72, 0xe8,
			0x47, 0x24, 0x42, 0x96, 0xb8, 0x7b, 0xee, 0xe7, 0x0d, 0xdc, 0x44, 0xec, 0xf3, 0x97, 0x6b, 0x1b,
			0x45, 0x28, 0xac, 0x3f, 0x35,
		}
		parsedServerKeyExchange := &MessageServerKeyExchange{
			EllipticCurveType:    elliptic.CurveTypeNamedCurve,
			NamedCurve:           elliptic.X25519,
			PublicKey:            rawServerKeyExchange[4:69],
			HashAlgorithm:        hash.None,
			SignatureAlgorithm:   signature.Anonymous,
			KeyExchangeAlgorithm: types.KeyExchangeAlgorithmEcdhe,
		}

		test(rawServerKeyExchange, parsedServerKeyExchange)
	})
}
