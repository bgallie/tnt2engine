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
			want: proFormPermutators[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cycles := make([]int, len(proFormPermutators[0].Cycles))
			for idx, val := range proFormPermutators[0].Cycles {
				cycles[idx] = val.Length
			}
			if got := new(Permutator).New(cycles, append([]byte(nil), proFormPermutators[0].Randp...)); !reflect.DeepEqual(got, tt.want) {
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
			want: new(Permutator).New([]int{83, 43, 57, 73}, []byte{
				162, 10, 51, 151, 78, 119, 191, 128, 69, 192, 179, 63, 100, 221, 154, 23,
				152, 127, 82, 196, 36, 172, 190, 198, 28, 147, 249, 113, 168, 171, 135, 98,
				209, 150, 39, 145, 200, 62, 216, 252, 83, 146, 109, 238, 64, 174, 12, 195,
				56, 133, 55, 218, 85, 235, 226, 163, 58, 255, 180, 25, 242, 53, 15, 236,
				217, 11, 223, 173, 52, 107, 229, 74, 110, 8, 241, 65, 233, 50, 203, 81,
				9, 16, 5, 89, 228, 148, 189, 240, 186, 167, 234, 211, 253, 126, 169, 230,
				207, 77, 194, 26, 61, 99, 116, 222, 41, 247, 239, 42, 130, 176, 79, 160,
				22, 46, 19, 35, 205, 14, 95, 138, 117, 215, 134, 254, 48, 108, 165, 177,
				49, 153, 244, 245, 47, 246, 45, 219, 73, 199, 84, 214, 181, 123, 251, 27,
				248, 75, 224, 140, 88, 184, 104, 139, 0, 161, 164, 166, 202, 156, 131, 21,
				125, 20, 137, 33, 2, 97, 93, 213, 91, 29, 43, 44, 193, 132, 3, 114,
				32, 68, 105, 243, 96, 67, 1, 70, 144, 31, 187, 6, 210, 60, 149, 111,
				158, 157, 231, 121, 183, 66, 122, 220, 201, 54, 129, 94, 101, 30, 250, 212,
				87, 38, 86, 112, 120, 57, 118, 92, 142, 106, 159, 76, 175, 7, 225, 103,
				204, 59, 170, 17, 182, 71, 197, 90, 18, 34, 40, 237, 80, 37, 115, 208,
				72, 136, 155, 227, 102, 232, 4, 188, 141, 24, 178, 206, 13, 124, 185, 143}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cycles := make([]int, len(proFormPermutators[0].Cycles))
			for idx, val := range proFormPermutators[0].Cycles {
				cycles[idx] = val.Length
			}
			p := new(Permutator).New(cycles, append([]byte(nil), proFormPermutators[0].Randp...))
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
			cycles := make([]int, len(proFormPermutators[0].Cycles))
			for idx, val := range proFormPermutators[0].Cycles {
				cycles[idx] = val.Length
			}
			p := new(Permutator).New(cycles, append([]byte(nil), proFormPermutators[0].Randp...))
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
			cycles := make([]int, len(proFormPermutators[0].Cycles))
			for idx, val := range proFormPermutators[0].Cycles {
				cycles[idx] = val.Length
			}
			p := new(Permutator).New(cycles, append([]byte(nil), proFormPermutators[0].Randp...))
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
			cycles := make([]int, len(proFormPermutators[0].Cycles))
			for idx, val := range proFormPermutators[0].Cycles {
				cycles[idx] = val.Length
			}
			p := new(Permutator).New(cycles, append([]byte(nil), proFormPermutators[0].Randp...))
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
			cycles := make([]int, len(proFormPermutators[0].Cycles))
			for idx, val := range proFormPermutators[0].Cycles {
				cycles[idx] = val.Length
			}
			p := new(Permutator).New(cycles, append([]byte(nil), proFormPermutators[0].Randp...))
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
			cycles := make([]int, len(proFormPermutators[0].Cycles))
			for idx, val := range proFormPermutators[0].Cycles {
				cycles[idx] = val.Length
			}
			p := new(Permutator).New(cycles, append([]byte(nil), proFormPermutators[0].Randp...))
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
			cycles := make([]int, len(proFormPermutators[0].Cycles))
			for idx, val := range proFormPermutators[0].Cycles {
				cycles[idx] = val.Length
			}
			p := new(Permutator).New(cycles, append([]byte(nil), proFormPermutators[0].Randp...))
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
			p := proFormPermutators[0]
			if got := p.String(); got != tt.want {
				t.Errorf("Permutator.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
