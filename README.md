Roman Yaroshenko (roma.yaroshenko@gmail.com)
---------------------------------------------

uuid.go - пакет для работы с уникальными идентификаторами
datetime.go - пакет для работы с датой в "паскалевском" формате

Installation
-------------

::

  $ go get github.com/ryaroshenko/rpkg

Example
--------

::

  package main

  import (
    "fmt"
    "github.com/ryaroshenko/rpkg/datetime"
    "github.com/ryaroshenko/rpkg/uuid"
    "time"
  )

  func main() {
    t = time.Now()
    dt := datetime.EncodeTime(t)
    fmt.Printf("Дата+Время = %s\n", dt)

    dt = "30.06.2014 14:56:43.024"
    t, err = dt.Decode()
    fmt.Printf("Время = %s\n", t)

    fmt.Printf("UUID = %s\n", uuid.New())
  }

See also uuid_test.go