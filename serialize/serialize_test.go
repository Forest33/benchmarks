package serialize

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/proto"

	"github.com/forest33/benchmarks/entity"
	serialize "github.com/forest33/benchmarks/proto"
)

const usersCount = 100

var (
	users        = make([]*entity.User, usersCount)
	protoUsers   = &serialize.Users{}
	jsonUsers    []byte
	msgpackUsers []byte
	pbUsers      []byte
)

func init() {
	pUsers := make([]*serialize.User, usersCount)
	for i := 0; i < usersCount; i++ {
		u := &entity.User{
			ID:        rand.Uint64(),
			FirstName: fmt.Sprintf("John-%d", i),
			LastName:  fmt.Sprintf("Doe-%d", i),
			Active:    func() bool { return rand.Intn(2) == 1 }(),
			Flags: func() []int32 {
				f := make([]int32, 0, 10)
				for i := 0; i < rand.Intn(10); i++ {
					f = append(f, 1<<rand.Intn(16))
				}
				return f
			}(),
		}
		users[i] = u
		pUsers[i] = &serialize.User{
			Id:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Active:    u.Active,
			Flags:     u.Flags,
		}
	}
	protoUsers = &serialize.Users{User: pUsers}
	jsonUsers, _ = json.Marshal(&users)
	msgpackUsers, _ = msgpack.Marshal(&users)
	pbUsers, _ = proto.Marshal(protoUsers)

	fmt.Printf("JSON:\t\t%d bytes\n", len(jsonUsers))
	fmt.Printf("Msgpack:\t%d bytes\n", len(msgpackUsers))
	fmt.Printf("Protobuf:\t%d bytes\n", len(pbUsers))
}

func BenchmarkMarshalJson(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := json.Marshal(users)
		_ = out
	}
}

func BenchmarkUnmarshalJson(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out := make([]*entity.User, 0, usersCount)
		_ = json.Unmarshal(jsonUsers, &out)
	}
}

func BenchmarkMarshalJsonIter(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := jsoniter.Marshal(users)
		_ = out
	}
}

func BenchmarkUnmarshalJsonIter(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out := make([]*entity.User, 0, usersCount)
		_ = jsoniter.Unmarshal(jsonUsers, &out)
	}
}

func BenchmarkMarshalMsgpack(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := msgpack.Marshal(users)
		_ = out
	}
}

func BenchmarkUnmarshalMsgpack(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out := make([]*entity.User, 0, usersCount)
		_ = msgpack.Unmarshal(msgpackUsers, &out)
	}
}

func BenchmarkMarshalProtobuf(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, _ := proto.Marshal(protoUsers)
		_ = out
	}
}

func BenchmarkUnmarshalProtobuf(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out := &serialize.Users{}
		_ = proto.Unmarshal(pbUsers, out)
		_ = out
	}
}
