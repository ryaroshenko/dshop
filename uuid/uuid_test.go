package uuid

import (
 "testing"
)

func TestUUID(t *testing.T) {
  uid := New()

  if uid == EmptyUUID {
    t.Fatalf("GenUUID error %s", uid)
  }

  t.Logf("uuid[%s]\n", uid)
}

func TestDecode(t *testing.T) {
	uid := New()

  if uid == EmptyUUID {
    t.Fatalf("GenUUID error %s", uid)
  }

  _, err := uid.Decode()

  if err != nil {
  	t.Fatalf("Decode UUID error %s", err)
  }

  t.Logf("uuid[%s]\n", uid)
}

func BenchmarkUUID(b *testing.B) {
  m := make(map[UUID]int, 1000) 

  for i := 0; i < b.N; i++ {
    uid := New()

    if uid == EmptyUUID {
      b.Fatalf("GenUUID error %s", uid)
    } 

    b.StopTimer()
    c := m[uid]

    if c > 0 {
      b.Fatalf("duplicate uuid[%s] count %d", uid, c)
    } 

    m[uid] = c + 1
    b.StartTimer()
  }
}
