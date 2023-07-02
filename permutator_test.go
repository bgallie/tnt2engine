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
			want: new(Permutator).New([]int{83, 73, 47, 53}, []byte{
				175, 89, 238, 236, 121, 247, 253, 161, 139, 57, 14, 189, 60, 191, 244, 51,
				120, 107, 153, 46, 1, 135, 113, 8, 117, 243, 184, 149, 252, 209, 93, 116,
				19, 146, 132, 234, 168, 157, 208, 45, 212, 254, 219, 162, 15, 214, 196, 154,
				105, 83, 140, 86, 118, 125, 224, 80, 29, 221, 143, 20, 52, 227, 26, 90,
				79, 75, 65, 141, 216, 48, 104, 18, 108, 235, 25, 74, 134, 167, 130, 200,
				220, 76, 61, 115, 16, 39, 195, 240, 127, 225, 5, 106, 12, 163, 31, 136,
				38, 102, 176, 173, 250, 245, 30, 204, 7, 124, 223, 24, 0, 129, 201, 180,
				237, 123, 133, 62, 137, 217, 197, 228, 42, 112, 192, 40, 36, 114, 159, 222,
				44, 229, 11, 72, 188, 142, 211, 64, 85, 202, 59, 190, 172, 198, 103, 152,
				178, 182, 131, 98, 147, 165, 232, 109, 179, 13, 215, 151, 47, 2, 249, 239,
				23, 77, 63, 81, 218, 4, 145, 99, 82, 94, 187, 174, 183, 43, 156, 50,
				34, 164, 21, 54, 56, 158, 185, 177, 199, 226, 66, 128, 68, 41, 233, 84,
				210, 242, 171, 100, 58, 32, 53, 255, 92, 67, 186, 144, 248, 35, 194, 251,
				49, 69, 246, 71, 126, 138, 22, 17, 170, 213, 205, 207, 96, 150, 87, 27,
				193, 203, 28, 148, 160, 55, 110, 155, 166, 101, 3, 78, 97, 122, 230, 119,
				73, 181, 95, 37, 241, 70, 111, 169, 33, 6, 88, 231, 91, 206, 10, 9}),
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
