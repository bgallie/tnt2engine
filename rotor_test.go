// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tnt2engine

import (
	"math/big"
	"reflect"
	"testing"
)

func TestNewRotor(t *testing.T) {
	type args struct {
		size  int
		start int
		step  int
		rotor []byte
	}
	tests := []struct {
		name string
		args args
		want *Rotor
	}{
		{
			name: "tnr",
			args: args{
				size:  1783,
				start: 863,
				step:  1033,
				rotor: []byte{
					184, 25, 190, 250, 35, 11, 111, 218, 111, 1, 44, 59, 137, 12, 184, 22,
					154, 226, 101, 88, 167, 109, 45, 92, 19, 164, 132, 233, 34, 133, 138, 222,
					59, 49, 123, 208, 179, 248, 61, 216, 55, 59, 235, 57, 67, 172, 233, 232,
					87, 236, 189, 170, 196, 124, 216, 109, 4, 106, 207, 150, 166, 164, 99, 57,
					131, 27, 1, 236, 168, 78, 122, 81, 165, 26, 32, 56, 129, 105, 35, 26,
					247, 208, 56, 235, 91, 183, 67, 150, 112, 103, 173, 197, 69, 13, 115, 14,
					129, 206, 74, 46, 119, 208, 95, 67, 119, 7, 191, 210, 128, 117, 140, 245,
					41, 168, 63, 203, 53, 241, 221, 28, 158, 40, 89, 76, 126, 58, 33, 40,
					78, 130, 93, 116, 206, 66, 4, 10, 109, 86, 150, 53, 200, 34, 26, 37,
					232, 185, 214, 47, 131, 18, 241, 210, 18, 81, 107, 161, 97, 65, 238, 250,
					81, 133, 54, 158, 54, 10, 254, 135, 110, 162, 175, 250, 117, 66, 232, 66,
					50, 102, 70, 76, 185, 249, 57, 59, 247, 195, 101, 8, 157, 235, 24, 94,
					204, 74, 100, 196, 93, 24, 179, 27, 118, 168, 29, 10, 38, 204, 210, 123,
					111, 247, 225, 171, 60, 166, 239, 124, 43, 180, 223, 240, 66, 2, 68, 220,
					12, 95, 253, 145, 133, 55, 237, 183, 0, 150, 157, 68, 6, 92, 11, 77,
					241, 50, 172, 211, 182, 22, 174, 9, 82, 194, 116, 145, 66, 69, 111, 0},
			},
			want: Rotor1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := new(Rotor).New(tt.args.size, tt.args.start, tt.args.step, tt.args.rotor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRotor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotor_Update(t *testing.T) {
	tnt2Machine := new(Tnt2Engine)
	tnt2Machine.Init([]byte("SecretKey"), "")
	tnt2Machine.SetEngineType("E")
	tnt2Machine.SetIndex(BigZero)
	tnt2Machine.BuildCipherMachine()
	rnd := new(Rand).New(tnt2Machine)
	defer tnt2Machine.CloseCipherMachine()
	tests := []struct {
		name string
		want *Rotor
	}{
		{
			name: "tur1",
			want: new(Rotor).New(7823, 1646, 4330, []byte{
				146, 1, 45, 232, 126, 13, 178, 164, 161, 20, 175, 177, 230, 170, 89, 163,
				45, 126, 175, 71, 189, 0, 131, 248, 223, 168, 185, 51, 240, 12, 192, 78,
				145, 82, 231, 96, 26, 28, 57, 62, 134, 76, 249, 154, 37, 232, 58, 99,
				81, 200, 48, 250, 112, 17, 115, 195, 88, 188, 120, 215, 227, 109, 149, 61,
				153, 133, 192, 126, 117, 69, 189, 125, 71, 128, 121, 142, 250, 125, 141, 211,
				52, 159, 100, 46, 205, 190, 205, 116, 18, 249, 38, 254, 179, 128, 29, 121,
				32, 87, 76, 172, 212, 24, 94, 62, 19, 67, 91, 8, 7, 51, 180, 116,
				247, 189, 251, 8, 27, 179, 95, 61, 173, 101, 204, 214, 168, 137, 73, 242,
				245, 63, 210, 81, 36, 16, 109, 246, 241, 204, 200, 111, 207, 136, 208, 219,
				148, 165, 85, 161, 64, 134, 195, 8, 227, 198, 68, 119, 99, 253, 115, 64,
				50, 32, 205, 179, 68, 70, 218, 72, 235, 129, 173, 194, 159, 45, 173, 23,
				222, 124, 144, 168, 142, 212, 59, 147, 233, 163, 211, 131, 22, 110, 132, 123,
				151, 181, 148, 40, 162, 107, 109, 206, 154, 240, 144, 57, 125, 186, 13, 138,
				136, 29, 92, 33, 1, 1, 227, 211, 91, 211, 244, 51, 69, 218, 118, 24,
				31, 51, 47, 219, 127, 6, 252, 36, 67, 165, 17, 147, 27, 73, 247, 42,
				251, 198, 124, 142, 136, 110, 57, 202, 232, 135, 155, 63, 245, 39, 43, 50,
				158, 57, 103, 82, 190, 95, 183, 173, 29, 216, 59, 233, 153, 91, 14, 254,
				37, 132, 39, 60, 73, 115, 119, 11, 133, 64, 154, 252, 124, 235, 12, 99,
				136, 177, 98, 250, 17, 247, 73, 158, 236, 28, 147, 214, 239, 84, 16, 156,
				234, 107, 90, 96, 198, 182, 33, 199, 73, 85, 22, 215, 113, 230, 60, 130,
				52, 126, 48, 191, 32, 199, 231, 144, 151, 192, 98, 97, 236, 154, 83, 255,
				5, 205, 123, 14, 255, 247, 201, 165, 202, 194, 118, 45, 152, 155, 125, 32,
				192, 113, 208, 65, 152, 128, 98, 43, 167, 200, 80, 121, 154, 170, 89, 91,
				91, 223, 132, 165, 0, 223, 138, 220, 180, 128, 24, 241, 134, 129, 101, 23,
				185, 2, 76, 3, 164, 88, 16, 229, 115, 64, 123, 197, 136, 132, 115, 53,
				203, 215, 152, 212, 114, 109, 59, 100, 61, 0, 42, 213, 132, 115, 206, 238,
				13, 231, 157, 250, 80, 20, 251, 140, 47, 63, 116, 35, 142, 245, 112, 7,
				150, 155, 243, 25, 196, 139, 102, 82, 217, 34, 75, 57, 26, 99, 6, 35,
				46, 81, 132, 210, 91, 215, 70, 14, 236, 192, 131, 114, 83, 196, 169, 71,
				182, 90, 216, 128, 196, 67, 119, 231, 241, 143, 165, 20, 13, 157, 125, 202,
				193, 23, 84, 164, 113, 44, 116, 74, 178, 53, 58, 43, 78, 231, 37, 87,
				43, 245, 164, 150, 140, 79, 191, 6, 216, 83, 157, 5, 76, 216, 145, 113,
				39, 113, 243, 106, 176, 213, 46, 216, 150, 144, 122, 63, 254, 246, 181, 48,
				252, 95, 94, 127, 205, 89, 91, 180, 97, 1, 254, 116, 239, 115, 117, 95,
				222, 172, 245, 54, 41, 19, 23, 29, 159, 211, 242, 162, 232, 44, 71, 15,
				250, 209, 115, 77, 187, 199, 57, 10, 108, 199, 7, 186, 177, 77, 63, 5,
				41, 100, 33, 141, 116, 174, 174, 196, 209, 184, 225, 245, 121, 139, 78, 147,
				157, 235, 56, 28, 81, 8, 100, 12, 232, 235, 38, 4, 163, 177, 114, 85,
				93, 218, 212, 165, 95, 58, 70, 187, 82, 75, 161, 181, 42, 162, 91, 225,
				86, 227, 91, 186, 249, 125, 41, 3, 154, 28, 81, 252, 149, 186, 8, 250,
				49, 22, 170, 132, 41, 130, 160, 105, 26, 16, 31, 67, 230, 143, 152, 2,
				123, 222, 45, 6, 229, 102, 251, 66, 93, 125, 239, 49, 119, 205, 97, 253,
				5, 187, 128, 106, 30, 82, 25, 219, 253, 212, 143, 68, 173, 146, 29, 192,
				159, 234, 140, 135, 230, 202, 139, 125, 186, 176, 34, 175, 170, 2, 18, 169,
				174, 236, 126, 11, 32, 112, 64, 71, 84, 240, 164, 252, 102, 132, 29, 204,
				13, 69, 56, 186, 207, 191, 152, 0, 141, 181, 132, 101, 181, 116, 217, 52,
				55, 192, 18, 18, 121, 125, 61, 143, 195, 150, 23, 183, 239, 199, 88, 157,
				147, 255, 172, 218, 253, 187, 239, 124, 250, 78, 89, 253, 225, 253, 140, 114,
				167, 208, 86, 228, 255, 42, 78, 41, 249, 101, 169, 235, 137, 18, 147, 154,
				146, 33, 198, 17, 14, 14, 205, 73, 139, 196, 183, 50, 225, 182, 88, 12,
				20, 126, 180, 50, 55, 226, 8, 213, 207, 96, 69, 240, 221, 225, 129, 42,
				31, 196, 167, 12, 71, 68, 156, 225, 171, 24, 92, 181, 84, 52, 1, 163,
				227, 55, 194, 55, 15, 247, 72, 219, 67, 6, 20, 2, 127, 184, 64, 120,
				34, 214, 78, 122, 114, 243, 126, 252, 206, 165, 93, 186, 138, 152, 242, 102,
				185, 119, 252, 253, 132, 195, 198, 44, 232, 212, 27, 70, 46, 234, 165, 71,
				179, 68, 255, 35, 144, 88, 178, 73, 25, 137, 109, 61, 100, 107, 215, 153,
				237, 55, 34, 183, 245, 208, 43, 238, 203, 239, 141, 101, 120, 199, 79, 249,
				7, 54, 222, 144, 140, 21, 142, 211, 114, 69, 251, 5, 254, 99, 183, 217,
				5, 231, 26, 72, 254, 47, 163, 224, 171, 177, 244, 151, 199, 135, 20, 121,
				199, 198, 17, 6, 149, 209, 84, 113, 73, 217, 130, 77, 110, 93, 127, 76,
				192, 225, 63, 160, 33, 17, 27, 50, 47, 164, 185, 224, 18, 243, 193, 51,
				41, 24, 201, 128, 22, 116, 191, 6, 89, 210, 80, 138, 215, 88, 115, 213,
				172, 209, 22, 191, 215, 163, 94, 128, 65, 252, 111, 212, 220, 25, 120, 6,
				96, 39}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(Rotor).New(Rotor1.Size, Rotor1.Start, Rotor1.Size, append([]byte(nil), Rotor1.Rotor...))
			r.Update(rnd)
			if !reflect.DeepEqual(r, tt.want) {
				t.Errorf("Updated Rotor() = %v, want %v", r, tt.want)
			}
		})
	}
}

func TestRotor_sliceRotor(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "trsr1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(Rotor).New(1783, 863, 1033, []byte{
				184, 25, 190, 250, 35, 11, 111, 218, 111, 1, 44, 59, 137, 12, 184, 22,
				154, 226, 101, 88, 167, 109, 45, 92, 19, 164, 132, 233, 34, 133, 138, 222,
				59, 49, 123, 208, 179, 248, 61, 216, 55, 59, 235, 57, 67, 172, 233, 232,
				87, 236, 189, 170, 196, 124, 216, 109, 4, 106, 207, 150, 166, 164, 99, 57,
				131, 27, 1, 236, 168, 78, 122, 81, 165, 26, 32, 56, 129, 105, 35, 26,
				247, 208, 56, 235, 91, 183, 67, 150, 112, 103, 173, 197, 69, 13, 115, 14,
				129, 206, 74, 46, 119, 208, 95, 67, 119, 7, 191, 210, 128, 117, 140, 245,
				41, 168, 63, 203, 53, 241, 221, 28, 158, 40, 89, 76, 126, 58, 33, 40,
				78, 130, 93, 116, 206, 66, 4, 10, 109, 86, 150, 53, 200, 34, 26, 37,
				232, 185, 214, 47, 131, 18, 241, 210, 18, 81, 107, 161, 97, 65, 238, 250,
				81, 133, 54, 158, 54, 10, 254, 135, 110, 162, 175, 250, 117, 66, 232, 66,
				50, 102, 70, 76, 185, 249, 57, 59, 247, 195, 101, 8, 157, 235, 24, 94,
				204, 74, 100, 196, 93, 24, 179, 27, 118, 168, 29, 10, 38, 204, 210, 123,
				111, 247, 225, 171, 60, 166, 239, 124, 43, 180, 223, 240, 66, 2, 68, 220,
				12, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
			r.sliceRotor()
			if !reflect.DeepEqual(r, Rotor1) {
				t.Errorf("Sliced Rotor() = %v, want %v", r, Rotor1)
			}
		})
	}
}

func TestRotor_SetIndex(t *testing.T) {
	type args struct {
		idx *big.Int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "trsi1",
			args: args{
				idx: big.NewInt(10000),
			},
			want: 161,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rotor1
			r.SetIndex(tt.args.idx)
			if r.Current != tt.want {
				t.Errorf("r.Current = %v, want %v", r.Current, tt.want)
			}
		})
	}
}

func TestRotor_Index(t *testing.T) {
	tests := []struct {
		name string
		want *big.Int
	}{
		{
			name: "trsi1",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rotor1
			if got := r.Index(); got != tt.want {
				t.Errorf("Rotor.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotor_ApplyF(t *testing.T) {
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
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				},
			},
			want: []byte{
				184, 25, 190, 250, 35, 11, 111, 218, 111, 1, 44, 59, 137, 12, 184, 22,
				154, 226, 101, 88, 167, 109, 45, 92, 19, 164, 132, 233, 34, 133, 138, 222,
			},
		},
		{
			name: "tpaf2",
			args: args{
				[]byte{
					0, 0, 0, 0,
				},
			},
			want: []byte{
				184, 25, 190, 250,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rotor1
			r.Current = 0
			if got := r.ApplyF(tt.args.blk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotor.ApplyF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotor_ApplyG(t *testing.T) {
	type args struct {
		blk CipherBlock
	}
	tests := []struct {
		name string
		args args
		want CipherBlock
	}{
		{
			name: "tpag1",
			args: args{
				[]byte{
					184, 25, 190, 250, 35, 11, 111, 218, 111, 1, 44, 59, 137, 12, 184, 22,
					154, 226, 101, 88, 167, 109, 45, 92, 19, 164, 132, 233, 34, 133, 138, 222,
				},
			},
			want: []byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			name: "tpag2",
			args: args{
				[]byte{
					184, 25, 190, 250,
				},
			},
			want: []byte{
				0, 0, 0, 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rotor1
			r.Current = 0
			if got := r.ApplyG(tt.args.blk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotor.ApplyG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotor_String(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "trs1",
			want: "new(Rotor).New(1783, 863, 1033, []byte{\n" +
				"\t184, 25, 190, 250, 35, 11, 111, 218, 111, 1, 44, 59, 137, 12, 184, 22, \n" +
				"\t154, 226, 101, 88, 167, 109, 45, 92, 19, 164, 132, 233, 34, 133, 138, 222, \n" +
				"\t59, 49, 123, 208, 179, 248, 61, 216, 55, 59, 235, 57, 67, 172, 233, 232, \n" +
				"\t87, 236, 189, 170, 196, 124, 216, 109, 4, 106, 207, 150, 166, 164, 99, 57, \n" +
				"\t131, 27, 1, 236, 168, 78, 122, 81, 165, 26, 32, 56, 129, 105, 35, 26, \n" +
				"\t247, 208, 56, 235, 91, 183, 67, 150, 112, 103, 173, 197, 69, 13, 115, 14, \n" +
				"\t129, 206, 74, 46, 119, 208, 95, 67, 119, 7, 191, 210, 128, 117, 140, 245, \n" +
				"\t41, 168, 63, 203, 53, 241, 221, 28, 158, 40, 89, 76, 126, 58, 33, 40, \n" +
				"\t78, 130, 93, 116, 206, 66, 4, 10, 109, 86, 150, 53, 200, 34, 26, 37, \n" +
				"\t232, 185, 214, 47, 131, 18, 241, 210, 18, 81, 107, 161, 97, 65, 238, 250, \n" +
				"\t81, 133, 54, 158, 54, 10, 254, 135, 110, 162, 175, 250, 117, 66, 232, 66, \n" +
				"\t50, 102, 70, 76, 185, 249, 57, 59, 247, 195, 101, 8, 157, 235, 24, 94, \n" +
				"\t204, 74, 100, 196, 93, 24, 179, 27, 118, 168, 29, 10, 38, 204, 210, 123, \n" +
				"\t111, 247, 225, 171, 60, 166, 239, 124, 43, 180, 223, 240, 66, 2, 68, 220, \n" +
				"\t12, 95, 253, 145, 133, 55, 237, 183, 0, 150, 157, 68, 6, 92, 11, 77, \n" +
				"\t241, 50, 172, 211, 182, 22, 174, 9, 82, 194, 116, 145, 66, 69, 111, 0})\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rotor1
			if got := r.String(); got != tt.want {
				t.Errorf("Rotor.String() = %v, want = %v", got, tt.want)
			}
		})
	}
}
