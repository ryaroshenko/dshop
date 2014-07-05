/**
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

func CreateUUID() ([]byte, error) {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)

	if n != len(uuid) || err != nil {
		return uuid, err
	}

	// TODO: verify the two lines implement RFC 4122 correctly
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7

	return uuid, nil
}

func GenUUID() (string, error) {
	uuid, err := CreateUUID()

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(uuid), nil
}

func GenGUID() string {
	uuid, err := CreateUUID()

	if err != nil {
		return ""
	}

	guid := hex.EncodeToString(uuid[:4]) + "-" +
		hex.EncodeToString(uuid[4:6]) + "-" +
		hex.EncodeToString(uuid[6:8]) + "-" +
		hex.EncodeToString(uuid[8:10]) + "-" +
		hex.EncodeToString(uuid[10:])

	return strings.ToUpper(guid)
}
