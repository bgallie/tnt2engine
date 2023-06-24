// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tnt2engine

import (
	"math/big"
	"reflect"
	"testing"
)

func TestNewPermutator(t *testing.T) {
	tests := []struct {
		name string
		want *Permutator
	}{
		{
			name: "tnp1",
			want: Permutator1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cycles := make([]int, len(Permutator1.Cycles))
			for idx, val := range Permutator1.Cycles {
				cycles[idx] = val.Length
			}
			if got := new(Permutator).New(cycles, append([]byte(nil), Permutator1.Randp...)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPermutator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermutator_Update(t *testing.T) {
	tnt2Machine := new(Tnt2Engine)
	tnt2Machine.Init([]byte("SecretKey"), "")
	tnt2Machine.SetEngineType("E")
	tnt2Machine.SetIndex(BigZero)
	tnt2Machine.BuildCipherMachine()
	rnd := new(Rand).New(tnt2Machine)
	defer tnt2Machine.CloseCipherMachine()
	tests := []struct {
		name string
		want *Permutator
	}{
		{
			name: "tpu1",
			want: new(Permutator).New([]int{73, 83, 53, 47}, []byte{
				93, 103, 99, 171, 14, 194, 90, 169, 124, 172, 219, 228, 47, 193, 110, 87,
				54, 31, 141, 109, 3, 26, 16, 249, 89, 208, 95, 227, 66, 229, 137, 23,
				220, 148, 40, 139, 50, 239, 100, 243, 127, 107, 238, 160, 18, 33, 122, 231,
				45, 13, 204, 118, 130, 173, 25, 43, 15, 247, 88, 143, 234, 49, 74, 1,
				244, 134, 190, 61, 131, 92, 9, 191, 112, 165, 158, 214, 177, 35, 17, 81,
				248, 115, 222, 140, 37, 205, 230, 179, 38, 113, 5, 123, 97, 63, 84, 138,
				98, 59, 197, 156, 105, 29, 69, 200, 106, 236, 159, 83, 255, 202, 176, 24,
				126, 166, 53, 41, 253, 136, 147, 254, 128, 28, 245, 129, 144, 196, 198, 116,
				76, 58, 192, 32, 48, 152, 30, 94, 21, 51, 135, 186, 161, 19, 62, 0,
				195, 44, 73, 250, 85, 117, 72, 218, 91, 203, 241, 145, 162, 108, 96, 8,
				46, 10, 154, 20, 101, 215, 34, 153, 120, 216, 155, 174, 55, 188, 65, 114,
				67, 146, 111, 75, 64, 119, 125, 170, 82, 232, 167, 224, 142, 233, 150, 209,
				182, 39, 184, 240, 187, 149, 70, 251, 2, 221, 235, 80, 213, 225, 226, 183,
				4, 246, 181, 22, 12, 57, 217, 68, 199, 175, 189, 168, 207, 7, 121, 42,
				78, 206, 211, 132, 252, 212, 79, 71, 11, 223, 104, 180, 185, 178, 201, 52,
				157, 36, 133, 164, 242, 60, 56, 151, 210, 102, 86, 237, 27, 6, 163, 77}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cycles := make([]int, len(Permutator1.Cycles))
			for idx, val := range Permutator1.Cycles {
				cycles[idx] = val.Length
			}
			p := new(Permutator).New(cycles, append([]byte(nil), Permutator1.Randp...))
			p.Update(rnd)
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Updated Permutator() = %v, want %v", p, tt.want)
			}
		})
	}
}

func TestPermutator_nextState(t *testing.T) {
	// This tests both Permutator.nextState() and Permutator.cycle()
	tests := []struct {
		name string
		want [256]byte
	}{
		{
			name: "tpns1",
			want: [256]byte{
				170, 88, 100, 228, 144, 177, 131, 225, 161, 205, 33, 124, 191, 107, 79, 230,
				179, 185, 44, 95, 7, 27, 40, 136, 77, 173, 149, 86, 65, 237, 92, 175,
				122, 182, 54, 159, 71, 43, 106, 114, 84, 93, 139, 184, 48, 11, 252, 130,
				137, 112, 53, 247, 108, 140, 25, 231, 28, 193, 165, 250, 15, 201, 147, 115,
				246, 210, 206, 181, 102, 156, 195, 17, 29, 202, 133, 0, 14, 39, 222, 221,
				162, 90, 89, 236, 20, 18, 13, 42, 120, 220, 169, 229, 243, 227, 219, 211,
				163, 176, 45, 22, 207, 157, 56, 132, 74, 251, 87, 160, 215, 117, 150, 148,
				187, 2, 153, 141, 249, 204, 128, 21, 123, 203, 73, 103, 94, 60, 49, 126,
				253, 254, 119, 168, 99, 5, 105, 68, 4, 199, 47, 72, 125, 213, 226, 36,
				41, 85, 138, 183, 127, 135, 26, 180, 51, 166, 30, 238, 35, 50, 97, 232,
				214, 37, 171, 155, 10, 174, 218, 78, 62, 12, 248, 241, 188, 16, 34, 61,
				240, 212, 83, 101, 151, 189, 91, 57, 52, 196, 110, 31, 208, 111, 82, 245,
				19, 200, 104, 118, 209, 98, 76, 192, 66, 6, 190, 58, 223, 55, 80, 64,
				24, 242, 3, 239, 23, 69, 158, 1, 152, 63, 9, 216, 167, 172, 164, 59,
				38, 142, 154, 146, 198, 129, 178, 75, 197, 224, 81, 235, 121, 46, 32, 113,
				217, 234, 194, 145, 8, 143, 244, 109, 67, 233, 70, 96, 186, 255, 116, 134},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cycles := make([]int, len(Permutator1.Cycles))
			for idx, val := range Permutator1.Cycles {
				cycles[idx] = val.Length
			}
			p := new(Permutator).New(cycles, append([]byte(nil), Permutator1.Randp...))
			p.nextState()
			if p.bitPerm != tt.want {
				t.Errorf("p.bitPerm = %v, want %v", p.bitPerm, tt.want)
			}
		})
	}
}

func TestPermutator_SetIndex(t *testing.T) {
	type args struct {
		idx *big.Int
	}
	tests := []struct {
		name string
		args args
		want [256]byte
	}{
		{
			name: "tpsi1",
			args: args{
				idx: BigOne,
			},
			want: [256]byte{
				170, 88, 100, 228, 144, 177, 131, 225, 161, 205, 33, 124, 191, 107, 79, 230,
				179, 185, 44, 95, 7, 27, 40, 136, 77, 173, 149, 86, 65, 237, 92, 175,
				122, 182, 54, 159, 71, 43, 106, 114, 84, 93, 139, 184, 48, 11, 252, 130,
				137, 112, 53, 247, 108, 140, 25, 231, 28, 193, 165, 250, 15, 201, 147, 115,
				246, 210, 206, 181, 102, 156, 195, 17, 29, 202, 133, 0, 14, 39, 222, 221,
				162, 90, 89, 236, 20, 18, 13, 42, 120, 220, 169, 229, 243, 227, 219, 211,
				163, 176, 45, 22, 207, 157, 56, 132, 74, 251, 87, 160, 215, 117, 150, 148,
				187, 2, 153, 141, 249, 204, 128, 21, 123, 203, 73, 103, 94, 60, 49, 126,
				253, 254, 119, 168, 99, 5, 105, 68, 4, 199, 47, 72, 125, 213, 226, 36,
				41, 85, 138, 183, 127, 135, 26, 180, 51, 166, 30, 238, 35, 50, 97, 232,
				214, 37, 171, 155, 10, 174, 218, 78, 62, 12, 248, 241, 188, 16, 34, 61,
				240, 212, 83, 101, 151, 189, 91, 57, 52, 196, 110, 31, 208, 111, 82, 245,
				19, 200, 104, 118, 209, 98, 76, 192, 66, 6, 190, 58, 223, 55, 80, 64,
				24, 242, 3, 239, 23, 69, 158, 1, 152, 63, 9, 216, 167, 172, 164, 59,
				38, 142, 154, 146, 198, 129, 178, 75, 197, 224, 81, 235, 121, 46, 32, 113,
				217, 234, 194, 145, 8, 143, 244, 109, 67, 233, 70, 96, 186, 255, 116, 134},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cycles := make([]int, len(Permutator1.Cycles))
			for idx, val := range Permutator1.Cycles {
				cycles[idx] = val.Length
			}
			p := new(Permutator).New(cycles, append([]byte(nil), Permutator1.Randp...))
			p.SetIndex(tt.args.idx)
			if p.bitPerm != tt.want {
				t.Errorf("p.bitPerm = %v, want %v", p.bitPerm, tt.want)
			}
		})
	}
}

func TestPermutator_Index(t *testing.T) {
	tests := []struct {
		name string
		want *big.Int
	}{
		{
			name: "tpi1",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cycles := make([]int, len(Permutator1.Cycles))
			for idx, val := range Permutator1.Cycles {
				cycles[idx] = val.Length
			}
			p := new(Permutator).New(cycles, append([]byte(nil), Permutator1.Randp...))
			if got := p.Index(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutator.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermutator_MaximalStates(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "tpi1",
			want: 14850609,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cycles := make([]int, len(Permutator1.Cycles))
			for idx, val := range Permutator1.Cycles {
				cycles[idx] = val.Length
			}
			p := new(Permutator).New(cycles, append([]byte(nil), Permutator1.Randp...))
			if p.MaximalStates != tt.want {
				t.Errorf("Permutator.MaximalStates = %d, want %d", p.MaximalStates, tt.want)
			}
		})
	}
}

func TestPermutator_ApplyF(t *testing.T) {
	type args struct {
		blk CipherBlock
	}
	tests := []struct {
		name string
		args args
		want CipherBlock
	}{
		{
			name: "tpaf1",
			args: args{
				[]byte{
					1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
					17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
				},
			},
			want: []byte{
				209, 217, 128, 24, 115, 4, 114, 33, 6, 18, 17, 204, 16, 160, 173, 86,
				133, 128, 48, 33, 152, 233, 34, 224, 3, 136, 162, 192, 32, 5, 2, 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cycles := make([]int, len(Permutator1.Cycles))
			for idx, val := range Permutator1.Cycles {
				cycles[idx] = val.Length
			}
			p := new(Permutator).New(cycles, append([]byte(nil), Permutator1.Randp...))
			if got := p.ApplyF(tt.args.blk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutator.ApplyF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermutator_ApplyG(t *testing.T) {
	type args struct {
		blk CipherBlock
	}
	tests := []struct {
		name string
		args args
		want CipherBlock
	}{
		{
			name: "tpaf1",
			args: args{
				[]byte{
					209, 217, 128, 24, 115, 4, 114, 33, 6, 18, 17, 204, 16, 160, 173, 86,
					133, 128, 48, 33, 152, 233, 34, 224, 3, 136, 162, 192, 32, 5, 2, 4,
				},
			},
			want: []byte{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
				17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cycles := make([]int, len(Permutator1.Cycles))
			for idx, val := range Permutator1.Cycles {
				cycles[idx] = val.Length
			}
			p := new(Permutator).New(cycles, append([]byte(nil), Permutator1.Randp...))
			if got := p.ApplyG(tt.args.blk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutator.ApplyF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermutator_String(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "tps1",
			want: "new(Permutator).New([]int{43, 57, 73, 83}, []byte{\n" +
				"	207, 252, 142, 205, 239, 35, 230, 62, 69, 94, 166, 89, 184, 81, 144, 120, \n" +
				"	27, 167, 39, 224, 75, 243, 87, 99, 47, 105, 163, 123, 129, 225, 2, 242, \n" +
				"	65, 43, 12, 113, 30, 102, 240, 78, 137, 109, 112, 210, 214, 118, 106, 22, \n" +
				"	232, 181, 164, 255, 70, 198, 160, 44, 231, 20, 228, 53, 85, 238, 178, 133, \n" +
				"	95, 194, 245, 234, 13, 147, 134, 25, 244, 91, 176, 38, 46, 1, 217, 249, \n" +
				"	250, 52, 182, 73, 206, 140, 216, 145, 60, 218, 213, 8, 151, 101, 156, 5, \n" +
				"	241, 67, 49, 42, 212, 180, 92, 21, 16, 130, 128, 126, 98, 199, 162, 188, \n" +
				"	117, 191, 66, 84, 57, 208, 158, 247, 41, 131, 227, 155, 61, 165, 253, 51, \n" +
				"	119, 103, 179, 93, 122, 83, 183, 116, 79, 222, 50, 59, 80, 110, 186, 141, \n" +
				"	90, 152, 127, 107, 54, 71, 185, 161, 169, 34, 148, 146, 157, 138, 24, 237, \n" +
				"	76, 196, 192, 251, 189, 201, 219, 86, 68, 37, 33, 82, 11, 170, 246, 72, \n" +
				"	229, 28, 32, 132, 23, 197, 108, 236, 220, 17, 150, 190, 171, 96, 26, 204, \n" +
				"	209, 31, 211, 4, 14, 136, 195, 45, 172, 111, 154, 36, 149, 226, 202, 187, \n" +
				"	193, 223, 139, 175, 124, 9, 3, 58, 125, 88, 15, 6, 121, 235, 221, 200, \n" +
				"	114, 254, 135, 168, 7, 29, 159, 48, 40, 115, 143, 203, 215, 77, 18, 55, \n" +
				"	56, 177, 100, 0, 173, 104, 248, 97, 74, 63, 233, 19, 64, 174, 153, 10})\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Permutator1
			if got := p.String(); got != tt.want {
				t.Errorf("Permutator.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
