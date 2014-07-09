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
)

func LoadJSON(data *os.File, elem interface{}) error {
	decoder := json.NewDecoder(data)
	return decoder.Decode(elem)
}

func LoadGOB(data *os.File, elem interface{}) error {
	decoder := gob.NewDecoder(data)
	return decoder.Decode(elem)
}

func LoadFile(filename string, elem interface{}) error {
	if filename == "" {
		return errors.New("LoadFile: Не задано имя файла")
	}	

	_, errs := os.Stat(filename)

	if errs != nil {
		return errs
	}	

	data, err := os.Open(filename)

	if err != nil {
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

	err = data.Close()

	return err
}

func SaveJSON(data *os.File, elem interface{}) error {
	encoder := json.NewEncoder(data)
	return encoder.Encode(elem)
}

func SaveGOB(data *os.File, elem interface{}) error {
	encoder := gob.NewEncoder(data)
	return encoder.Encode(elem)
}

func SaveFile(filename string, elem interface{}) error {
	if filename == "" {
		return errors.New("SaveFile: Не задано имя файла")
	}	

	_, errs := os.Stat(filename)

	if errs == nil {
		errs = os.Remove(filename)

		if errs != nil {
			return errs
		}
	}	

	data, err := os.Create(filename)

	if err != nil {
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

	err = data.Close()

	return err
}

type Lister interface {
	GetMeta() *MetaInfo
	Load(filename string) error
	Save(filename string) error
	IncLast()
	ReCalc()
}

const (
	MetaVersion = int16(1)
)

// MetaInfo - тип для хранения метаинформации о списке значений
type MetaInfo struct {
	Name    string
	Version int16
	Count   int64
	Last    int64
}

// MetaList - список значений типа MetaInfo
type MetaList struct {
	Meta  MetaInfo
	Items []MetaInfo
}

func (ml *MetaList) GetMeta() *MetaInfo {
	return &ml.Meta
}

func (ml *MetaList) Load(filename string) error {
	return LoadFile(filename, ml)
}

func (ml *MetaList) Save(filename string) error {
	return SaveFile(filename, ml)
}

func (ml *MetaList) IncLast() {
	ml.Meta.Last++
}

func (ml *MetaList) ReCalc() {
	ml.Meta.Version = MetaVersion
	ml.Meta.Count = int64(len(ml.Items))
}

// Shop - структура "Магазин"
type Shop struct {
	UID         uuid.UUID
	Name        string
	Description string
}

// ShopList - список значений типа Shop
type ShopList struct {
	Meta  MetaInfo
	Items []Shop
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

// GoodsList - список значений типа Goods
type GoodsList struct {
	Meta  MetaInfo
	Items []Goods
}
