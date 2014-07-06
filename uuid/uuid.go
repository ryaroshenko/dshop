/*
uuid.go : Ashish Banerjee, Sat, 12 May 12

Type 4 PRSG (Pseudo Random Sequence Geretaor) UUID

http://www.ietf.org/rfc/rfc4122.txt
*/
package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

type UUID []byte

func New() (*UUID, error) {
	bt := make([]byte, 16)
	n, err := rand.Read(bt)

	if n != len(bt) || err != nil {
		return nil, err
	}

	// TODO: verify the two lines implement RFC 4122 correctly
	bt[8] = 0x80 // variant bits see page 5
	bt[4] = 0x40 // version 4 Pseudo Random, see page 7

	uid := UUID(bt)

	return &uid, nil
}

func (uuid *UUID) String() string {
	bt := []byte(*uuid)
	return strings.ToUpper(hex.EncodeToString(bt[:4]) + "-" +
		hex.EncodeToString(bt[4:6]) + "-" +
		hex.EncodeToString(bt[6:8]) + "-" +
		hex.EncodeToString(bt[8:10]) + "-" +
		hex.EncodeToString(bt[10:]))
}
