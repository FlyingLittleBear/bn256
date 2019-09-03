package bn256

import (
	"testing"

	"bytes"
	"math/rand"
	"strconv"
)

func TestKnownHashes(t *testing.T) {
	for i, mh := range marshaledHashes {
		g := Hash([]byte{byte(i)})
		if !bytes.Equal(mh[:], g.Marshal()) {
			t.Fatal("found a collision of hashes ")
		}
	}
}

func TestHashCollision(t *testing.T) {
	g := Hash([]byte(strconv.Itoa(rand.Int())))
	h := Hash([]byte(strconv.Itoa(rand.Int())))
	if *(g.p) == *(h.p) {
		t.Fatal("found a collision of hashes ")
	}
}

func TestHashTAICollision(t *testing.T) {
	g := HashTAI([]byte(strconv.Itoa(rand.Int())))
	h := HashTAI([]byte(strconv.Itoa(rand.Int())))
	if *(g.p) == *(h.p) {
		t.Fatal("found a collision of hashes ")
	}
}

func BenchmarkHash(b *testing.B) {
	data := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = []byte(strconv.Itoa(i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Hash(data[i])
	}
}

func BenchmarkHashTAI(b *testing.B) {
	data := make([][]byte, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = []byte(strconv.Itoa(i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		HashTAI(data[i])
	}
}

var marshaledHashes = [10][64]byte{
	[64]byte{86, 219, 66, 190, 229, 239, 145, 232, 200, 198, 99, 72, 64, 5, 23, 86, 50, 72, 45, 104, 205, 224, 152, 2, 85, 238, 89, 58, 169, 32, 43, 235, 93, 141, 205, 152, 81, 226, 40, 13, 165, 143, 105, 253, 197, 49, 104, 187, 1, 106, 232, 115, 133, 76, 54, 145, 84, 89, 225, 20, 29, 211, 218, 126},

	[64]byte{49, 118, 83, 24, 40, 34, 8, 163, 252, 254, 106, 40, 85, 131, 223, 153, 235, 47, 95, 71, 201, 187, 164, 154, 94, 43, 103, 155, 139, 71, 53, 149, 23, 10, 124, 69, 54, 225, 183, 113, 173, 84, 216, 238, 137, 112, 76, 33, 63, 176, 217, 220, 205, 133, 186, 245, 104, 185, 101, 223, 48, 186, 70, 190},

	[64]byte{50, 150, 250, 128, 223, 30, 30, 27, 63, 106, 236, 70, 0, 188, 138, 160, 110, 6, 194, 20, 249, 6, 186, 115, 74, 15, 28, 152, 24, 139, 112, 119, 5, 112, 175, 203, 12, 54, 236, 255, 144, 117, 24, 255, 53, 219, 49, 182, 10, 79, 171, 128, 2, 40, 22, 240, 37, 137, 201, 178, 46, 236, 122, 2},

	[64]byte{113, 61, 0, 173, 156, 240, 35, 206, 144, 114, 169, 60, 136, 223, 164, 115, 177, 96, 42, 226, 23, 7, 224, 234, 233, 104, 149, 144, 80, 12, 54, 247, 8, 51, 184, 114, 210, 30, 77, 184, 95, 119, 144, 171, 56, 193, 255, 41, 144, 70, 26, 119, 249, 150, 85, 161, 170, 55, 218, 229, 170, 95, 90, 52},

	[64]byte{72, 64, 95, 245, 200, 115, 80, 228, 13, 108, 7, 72, 223, 95, 221, 86, 164, 124, 250, 186, 72, 1, 195, 164, 226, 44, 106, 46, 163, 199, 46, 148, 80, 37, 0, 208, 251, 157, 239, 151, 196, 117, 136, 102, 59, 47, 23, 105, 14, 178, 90, 176, 120, 146, 159, 213, 1, 253, 55, 200, 225, 5, 116, 186},

	[64]byte{101, 188, 50, 231, 72, 66, 175, 42, 217, 140, 137, 237, 109, 189, 101, 137, 21, 189, 198, 253, 80, 19, 128, 214, 102, 173, 40, 249, 157, 186, 215, 193, 100, 186, 168, 69, 78, 233, 174, 120, 126, 87, 4, 3, 1, 184, 246, 60, 150, 242, 42, 25, 51, 127, 4, 91, 95, 246, 63, 189, 235, 22, 54, 246},

	[64]byte{67, 136, 51, 244, 48, 110, 136, 70, 225, 74, 212, 146, 201, 127, 227, 23, 16, 187, 163, 208, 184, 243, 103, 185, 145, 144, 138, 202, 108, 131, 159, 18, 110, 28, 181, 199, 232, 86, 177, 98, 91, 124, 39, 145, 175, 86, 165, 88, 120, 96, 215, 162, 209, 108, 42, 137, 116, 76, 169, 114, 131, 123, 217, 145},

	[64]byte{87, 121, 233, 202, 190, 188, 241, 164, 189, 65, 107, 142, 182, 240, 42, 71, 9, 65, 111, 53, 111, 185, 150, 143, 55, 209, 104, 216, 114, 97, 94, 233, 84, 126, 127, 46, 50, 25, 202, 118, 138, 86, 95, 179, 185, 39, 219, 165, 132, 145, 186, 64, 204, 66, 183, 140, 178, 231, 132, 252, 63, 22, 138, 27},

	[64]byte{32, 145, 60, 214, 169, 54, 69, 176, 82, 234, 77, 188, 178, 21, 52, 22, 108, 94, 252, 181, 129, 112, 79, 229, 223, 218, 65, 59, 242, 208, 99, 227, 95, 57, 54, 159, 78, 124, 107, 204, 248, 193, 82, 199, 148, 227, 35, 205, 76, 73, 249, 240, 210, 244, 13, 241, 209, 199, 255, 212, 245, 60, 37, 95},

	[64]byte{8, 157, 196, 16, 182, 221, 150, 94, 80, 158, 146, 57, 178, 104, 6, 225, 65, 242, 231, 152, 69, 157, 235, 99, 139, 79, 144, 254, 129, 244, 25, 108, 73, 141, 168, 36, 129, 235, 68, 165, 0, 21, 112, 129, 218, 175, 24, 100, 71, 157, 65, 67, 34, 194, 21, 26, 107, 190, 111, 227, 49, 140, 129, 132},
}
