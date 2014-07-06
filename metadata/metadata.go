// metadata - пакет содержит описания структур данных
package metadata

import (
	"dshop/datetime"
)

// Shop - структура "Магазин"
type Shop struct {
	UID         string
	Name        string
	Description string
}

// Tran - структура "Транзакция"
type Tran struct {
	ShopUID    string
	TranTime   datetime.DateTime
	PackageUID string
}

// Goods - структура "Товар"
type Goods struct {
	UID       string
	Article   string
	ShortName string
	Name      string
	TranInfo  Tran
}
