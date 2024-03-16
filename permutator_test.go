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
			want: new(Permutator).New([]int{65, 61, 67, 63}, []byte{
				12, 132, 119, 175, 173, 178, 70, 65, 26, 103, 208, 69, 5, 96, 255, 43,
				252, 225, 149, 151, 80, 177, 101, 99, 205, 66, 16, 192, 232, 183, 105, 44,
				240, 144, 87, 216, 18, 222, 241, 113, 60, 89, 153, 131, 39, 223, 81, 79,
				104, 21, 213, 124, 34, 57, 165, 154, 193, 237, 194, 142, 95, 185, 169, 2,
				139, 15, 19, 162, 85, 231, 102, 92, 160, 214, 93, 8, 147, 61, 207, 203,
				33, 167, 163, 75, 187, 143, 164, 210, 191, 170, 172, 211, 125, 6, 251, 200,
				1, 111, 150, 121, 86, 59, 27, 74, 219, 157, 97, 166, 244, 112, 140, 24,
				155, 49, 41, 201, 73, 246, 30, 135, 176, 108, 129, 238, 83, 118, 38, 94,
				212, 206, 123, 235, 221, 67, 0, 202, 209, 116, 242, 31, 156, 37, 236, 196,
				42, 127, 133, 52, 48, 248, 17, 243, 227, 11, 218, 98, 114, 230, 130, 13,
				182, 174, 181, 56, 138, 82, 29, 25, 71, 168, 20, 23, 189, 198, 171, 226,
				90, 88, 195, 28, 32, 4, 136, 84, 190, 253, 51, 245, 68, 152, 158, 254,
				7, 247, 228, 122, 100, 120, 224, 128, 9, 141, 184, 22, 197, 54, 72, 134,
				14, 47, 250, 35, 186, 146, 50, 76, 249, 234, 137, 199, 109, 180, 78, 46,
				220, 188, 117, 62, 91, 40, 145, 204, 106, 179, 161, 159, 233, 58, 126, 217,
				63, 148, 215, 64, 115, 10, 55, 36, 77, 229, 53, 45, 110, 3, 107, 239}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cycles := make([]int, len(proFormPermutators[0].Cycles))
			for idx, val := range proFormPermutators[0].Cycles {
				cycles[idx] = val.Length
			}
			// p := new(Permutator).New(cycles, append([]byte(nil), proFormPermutators[0].Randp...))
			p := new(Permutator)
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
