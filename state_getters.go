package dtls

import "github.com/mingyech/dtls/v2/pkg/protocol/handshake"

func (s *State) RemoteRandomBytes() [handshake.RandomBytesLength]byte {
	return s.remoteRandom.RandomBytes
}
