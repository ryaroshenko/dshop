// metadata - пакет содержит описания структур данных
package metadata

import (
	"dshop/datetime"
	"dshop/uuid"
	"encoding/gob"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

func LoadJSON(data *os.File, elem interface{}) error {
	decoder := json.NewDecoder(data)
	return decoder.Decode(elem)
}

func LoadGOB(data *os.File, elem interface{}) error {
	decoder := gob.NewDecoder(data)
	return decoder.Decode(elem)
}

func LoadData(filename string, elem interface{}) error {
	var data *os.File
	var err error

	if filename == "" {
		return errors.New("LoadData: Не задано имя файла")
	}	

	if _, errs := os.Stat(filename); errs != nil {
		return errs
	}	

	if data, err = os.Open(filename); err != nil {
		return err
	}

	fn := strings.ToLower(filename)

	if strings.HasSuffix(fn, ".json") {
		err = LoadJSON(data, elem)
	} else if strings.HasSuffix(fn, ".gob") {
		err = LoadGOB(data, elem)
	}

	if err != nil {
		return err
	}

	return data.Close()
}

func SaveJSON(data *os.File, elem interface{}) error {
	encoder := json.NewEncoder(data)
	return encoder.Encode(elem)
}

func SaveGOB(data *os.File, elem interface{}) error {
	encoder := gob.NewEncoder(data)
	return encoder.Encode(elem)
}

func SaveData(filename string, elem interface{}) error {
	var filename_old string
	var data *os.File
	var err error

	if filename == "" {
		return errors.New("SaveFile: Не задано имя файла")
	}	

	if _, errs := os.Stat(filename); errs == nil {
		i := strings.LastIndex(filename, ".")
		filename_old = filename[:i] + "_old" + filename[i:]

		if errs = os.Rename(filename, filename_old); errs != nil {
			return errs
		}
	}	

	if data, err = os.Create(filename); err != nil {
		return err
	}

	fn := strings.ToLower(filename)

	if strings.HasSuffix(fn, ".json") {
		err = SaveJSON(data, elem)
	} else if strings.HasSuffix(fn, ".gob") {
		err = SaveGOB(data, elem)
	}

	if err != nil {
		return err
	}

	if err = data.Close(); err == nil && filename_old != "" {
		err = os.Remove(filename_old)
	}

	return err
}

const (
	MetaVersion = int16(1)
)

// HeaderInfo - тип для хранения информации о заголовке данных
type HeaderInfo struct {
	DataName  string
	Version   int16
	ItemCount int64
	MaxUpdate int64
}

func (head *HeaderInfo) Copy() *HeaderInfo {
	new_head := new(HeaderInfo)

	new_head.DataName  = head.DataName
	new_head.Version   = head.Version
	new_head.ItemCount = head.ItemCount
	new_head.MaxUpdate = head.MaxUpdate

	return new_head
}

func (head *HeaderInfo) UpdateVersion() *HeaderInfo {
	head.Version = MetaVersion
	return head
}

// ItemInfo - тип для хранения информации о элементе данных
type ItemInfo struct {
	UID      uuid.UUID
	Name     string
	ShopUID  uuid.UUID
	Shared   bool
	TranTime datetime.DateTime
	Update   int64
}

func (item *ItemInfo) Copy() *ItemInfo {
	new_item := new(ItemInfo)

	new_item.UID      = item.UID
	new_item.Name     = item.Name
	new_item.ShopUID  = item.ShopUID
	new_item.Shared   = item.Shared
	new_item.TranTime = item.TranTime
	new_item.Update   = item.Update

	return new_item
}

func (item *ItemInfo) GenUID() *ItemInfo {
	item.UID = uuid.New()
	return item
}

func (item *ItemInfo) UpdateTranTime() *ItemInfo {
	item.TranTime = datetime.EncodeTime(time.Now())
	return item
}

func (item *ItemInfo) IncUpdate() *ItemInfo {
	item.Update++
	return item
}

// Shop - структура "Магазин"
type Shop struct {
	ItemInfo
	Description string
}

// ShopData - данные о магазинах
type ShopData struct {
	HeaderInfo
	Items []Shop
}

// Goods - структура "Товар"
type Goods struct {
	ItemInfo
	Article   string
	ShortName string
}

// GoodsData - данные о товарах
type GoodsList struct {
	HeaderInfo
	Items []Goods
}
