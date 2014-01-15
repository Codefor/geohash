package geohash

import(
    "testing"
)

func TestEncode(t *testing.T){
    if Encode(40,116) != "wx47x9u8gumn"{
	t.Error("Encode error")
    }
}

func TestDecode(t *testing.T){
    lat,lon := Decode("wx47x9u8gumn")
    if lat[1] - 40 > 0.1 || lon[1] - 116 > 0.1 {
	t.Error("Decode error")
    }
}

func TestAdjacent(t *testing.T){
    if Adjacent("wx47x9u8gumn","left") != "wx47x9u8guky" {
	t.Error("Adjacent error")
    }
}

func BenchmarkEncode(b *testing.B) {
    for i := 0; i < b.N; i++ {
	Encode(40,116)
    }
}

func BenchmarkDecode(b *testing.B) {
    for i := 0; i < b.N; i++ {
	Decode("wx47x9u8gumn")
    }
}

func BenchmarkAdjacent(b *testing.B) {
    for i := 0; i < b.N; i++ {
	Adjacent("wx47x9u8gumn","left")
    }
}
