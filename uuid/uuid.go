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

// UUID - тип данных для хранения уникального идентификатора в
// формате "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"
type UUID string

const (
	EmptyUUID = UUID("00000000-0000-0000-0000-000000000000")
)

// New - генерирует новое значение типа UUID
func New() UUID {
	bt := make([]byte, 16)
	n, err := rand.Read(bt)

	if n != len(bt) || err != nil {
		return EmptyUUID
	}

	// TODO: verify the two lines implement RFC 4122 correctly
	bt[8] = 0x80 // variant bits see page 5
	bt[4] = 0x40 // version 4 Pseudo Random, see page 7

	s := strings.ToUpper(hex.EncodeToString(bt[:4]) + "-" +
		hex.EncodeToString(bt[4:6]) + "-" +
		hex.EncodeToString(bt[6:8]) + "-" +
		hex.EncodeToString(bt[8:10]) + "-" +
		hex.EncodeToString(bt[10:]))

	return UUID(s)
}

// String - поддержка для функций форматирования типа fmt.Printf()
func (uuid UUID) String() string {
	return string(uuid)
}

// Decode - функция декодирования значение типа UUID в значение типа []byte
func (uuid UUID) Decode() ([]byte, error) {
	return hex.DecodeString(strings.Replace(strings.ToLower(uuid.String()), "-", "", -1))
}