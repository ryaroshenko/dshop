// metadata - пакет содержит описания структур данных
package metadata

import (
	"dshop/datetime"
	"dshop/uuid"
)

const (
	MetaVersion = uint16(1)
)

type Meta struct {
	Name      string
	Version   uint16
	ItemCount uint32
}

// Shop - структура "Магазин"
type Shop struct {
	UID         uuid.UUID
	Name        string
	Description string
}

type ShopList struct {
	MetaInfo Meta
	Items    []Shop
}

// Tran - структура "Транзакция"
type Tran struct {
	ShopUID    uuid.UUID
	TranTime   datetime.DateTime
	PackageUID uuid.UUID
}

// Goods - структура "Товар"
type Goods struct {
	UID       uuid.UUID
	Article   string
	ShortName string
	Name      string
	Shared    bool
	TranInfo  Tran
}

type GoodsList struct {
	MetaInfo Meta
	Items    []Goods
}
