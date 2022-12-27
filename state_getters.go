package dtls

import "github.com/mingyech/dtls/v2/pkg/protocol/handshake"

// RemoteRandomBytes returns the random bytes from the client or server hello
func (s *State) RemoteRandomBytes() [handshake.RandomBytesLength]byte {
	return s.remoteRandom.RandomBytes
}
