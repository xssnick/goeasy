package simply

import (
	"encoding/json"
	"log"
	"testing"
)

type TestType struct {
	One   string
	Two   int
	Three bool
}

// BenchmarkJson-8   	  825301	      1246 ns/op	     312 B/op	       8 allocs/op
func BenchmarkJson(b *testing.B) {
	var data = []byte(`{"One":"hello!","Two":777,"Three":true}`)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Json(TestType{}, data)
	}
}

func BenchmarkJsonF(b *testing.B) {
	var data = []byte(`{"One":"hello!","Two":777,"Three":true}`)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		JsonF(TestType{}, data)
	}
}

// BenchmarkJsonDefault-8   	 1118924	      1082 ns/op	     248 B/op	       6 allocs/op
func BenchmarkJsonDefault(b *testing.B) {
	var data = []byte(`{"One":"hello!","Two":777,"Three":true}`)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tp := TestType{}
		err := json.Unmarshal(data, &tp)
		if err != nil {
			log.Println(err)
		}
	}
}
