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
			want: proFormaRotors[0],
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
			want: new(Rotor).New(8123, 4576, 5998, []byte{
				233, 246, 104, 252, 60, 172, 12, 41, 5, 73, 28, 77, 179, 45, 2, 182,
				240, 177, 98, 241, 80, 44, 246, 162, 169, 182, 196, 241, 248, 12, 76, 1,
				164, 19, 39, 62, 164, 29, 56, 195, 242, 176, 41, 229, 30, 100, 18, 192,
				179, 35, 110, 180, 162, 150, 204, 26, 227, 106, 121, 95, 86, 99, 204, 238,
				167, 0, 106, 255, 210, 222, 87, 176, 88, 252, 155, 195, 210, 244, 163, 48,
				80, 17, 242, 71, 203, 193, 137, 99, 0, 236, 25, 185, 64, 248, 148, 154,
				131, 109, 219, 85, 0, 163, 83, 132, 192, 116, 111, 173, 123, 78, 249, 117,
				0, 3, 243, 254, 105, 160, 18, 24, 76, 241, 62, 104, 93, 101, 108, 46,
				69, 238, 204, 7, 16, 53, 30, 222, 155, 28, 21, 248, 34, 87, 14, 27,
				226, 140, 32, 10, 233, 170, 229, 213, 78, 79, 239, 1, 106, 212, 65, 217,
				51, 84, 142, 124, 21, 96, 120, 63, 31, 3, 244, 196, 209, 61, 131, 30,
				231, 49, 229, 98, 111, 43, 53, 93, 32, 153, 171, 233, 62, 188, 228, 128,
				247, 207, 23, 143, 95, 246, 47, 198, 121, 80, 123, 141, 104, 240, 30, 71,
				102, 161, 140, 152, 213, 203, 175, 205, 153, 246, 158, 220, 56, 114, 190, 145,
				117, 110, 166, 217, 93, 129, 44, 184, 66, 77, 145, 138, 230, 69, 174, 219,
				96, 92, 153, 203, 191, 28, 124, 246, 167, 169, 142, 250, 192, 96, 114, 236,
				191, 164, 105, 175, 33, 85, 130, 122, 195, 54, 158, 234, 163, 70, 223, 64,
				198, 131, 116, 66, 133, 223, 18, 115, 210, 124, 134, 101, 55, 27, 95, 182,
				195, 154, 228, 133, 115, 121, 127, 210, 182, 61, 94, 91, 235, 171, 111, 98,
				73, 229, 212, 235, 108, 226, 83, 41, 223, 182, 144, 29, 128, 239, 246, 145,
				238, 13, 46, 78, 34, 144, 27, 241, 224, 54, 4, 230, 218, 223, 210, 50,
				194, 173, 191, 32, 38, 203, 74, 135, 223, 188, 110, 132, 136, 48, 131, 190,
				171, 10, 214, 188, 222, 163, 114, 153, 134, 225, 120, 252, 2, 226, 54, 117,
				219, 47, 207, 112, 248, 188, 109, 81, 73, 250, 39, 145, 13, 137, 20, 175,
				38, 51, 228, 9, 117, 255, 130, 110, 0, 75, 207, 245, 89, 188, 131, 1,
				20, 176, 184, 251, 191, 24, 242, 70, 249, 208, 169, 178, 158, 181, 37, 50,
				48, 83, 156, 140, 165, 4, 37, 131, 234, 79, 0, 198, 149, 168, 201, 143,
				146, 117, 223, 255, 179, 111, 8, 65, 216, 224, 79, 110, 128, 113, 186, 117,
				225, 52, 96, 82, 96, 184, 157, 81, 153, 212, 25, 62, 186, 41, 136, 255,
				29, 199, 168, 184, 222, 65, 74, 223, 3, 9, 230, 234, 230, 88, 23, 249,
				116, 37, 133, 179, 223, 39, 89, 136, 58, 84, 199, 30, 151, 170, 76, 240,
				96, 197, 230, 181, 22, 102, 90, 135, 42, 158, 139, 103, 39, 169, 209, 49,
				123, 165, 152, 243, 192, 180, 20, 109, 245, 128, 117, 118, 194, 220, 1, 41,
				167, 9, 252, 233, 19, 148, 20, 183, 62, 58, 149, 62, 216, 80, 83, 252,
				122, 46, 5, 52, 163, 43, 252, 86, 188, 83, 118, 19, 228, 166, 112, 98,
				205, 98, 224, 215, 25, 4, 215, 246, 230, 53, 198, 116, 20, 45, 197, 146,
				9, 241, 81, 103, 72, 81, 95, 5, 180, 251, 176, 110, 213, 251, 206, 96,
				183, 10, 38, 91, 253, 120, 226, 179, 52, 90, 55, 238, 68, 209, 171, 167,
				70, 174, 218, 69, 156, 90, 253, 36, 146, 46, 247, 101, 48, 194, 236, 198,
				56, 23, 99, 12, 165, 251, 173, 92, 146, 150, 44, 106, 235, 237, 245, 56,
				9, 72, 249, 74, 82, 28, 6, 187, 253, 249, 78, 28, 224, 28, 162, 195,
				98, 13, 90, 80, 102, 78, 49, 166, 23, 224, 46, 37, 162, 147, 198, 210,
				87, 241, 1, 80, 211, 104, 117, 106, 97, 66, 59, 124, 108, 96, 230, 19,
				118, 149, 232, 247, 103, 132, 45, 141, 234, 192, 91, 165, 153, 231, 187, 37,
				253, 242, 6, 175, 67, 143, 246, 142, 169, 81, 3, 149, 201, 199, 186, 197,
				32, 77, 64, 226, 129, 162, 73, 179, 134, 242, 200, 156, 102, 130, 65, 107,
				196, 237, 236, 69, 177, 253, 89, 164, 81, 220, 128, 15, 212, 9, 232, 233,
				135, 193, 107, 210, 15, 22, 251, 232, 232, 246, 158, 207, 40, 55, 96, 43,
				217, 105, 86, 210, 1, 36, 116, 25, 126, 171, 91, 113, 238, 20, 249, 219,
				139, 206, 148, 135, 57, 0, 24, 223, 44, 149, 246, 245, 58, 127, 102, 93,
				236, 131, 162, 136, 224, 29, 155, 205, 17, 225, 38, 160, 92, 242, 17, 231,
				82, 156, 230, 178, 138, 191, 37, 206, 201, 94, 66, 53, 238, 58, 15, 223,
				192, 198, 107, 26, 203, 232, 8, 106, 209, 184, 130, 248, 3, 209, 125, 210,
				215, 203, 211, 178, 188, 135, 196, 127, 225, 200, 0, 215, 56, 96, 67, 131,
				192, 13, 1, 40, 152, 189, 206, 241, 234, 71, 56, 221, 134, 54, 181, 199,
				7, 83, 200, 200, 184, 7, 131, 9, 92, 57, 184, 1, 190, 157, 109, 12,
				207, 213, 36, 9, 21, 61, 165, 13, 208, 20, 59, 180, 37, 232, 73, 53,
				95, 202, 71, 37, 170, 115, 150, 214, 209, 113, 220, 165, 86, 73, 253, 76,
				51, 121, 158, 201, 39, 110, 18, 12, 177, 29, 98, 19, 211, 170, 213, 83,
				133, 206, 72, 51, 118, 204, 215, 213, 31, 211, 194, 233, 223, 141, 46, 223,
				30, 255, 184, 114, 40, 182, 28, 149, 4, 153, 180, 37, 167, 60, 22, 223,
				19, 3, 22, 210, 116, 63, 183, 15, 50, 228, 209, 94, 252, 166, 181, 163,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 72, 183, 71, 227, 231, 97, 101, 72, 41,
				72, 226, 104, 154, 109, 17, 176, 133, 143, 21, 139, 135, 98, 177, 23, 77,
				181, 37, 142, 199, 103, 96, 10, 0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(Rotor).New(proFormaRotors[0].Size, proFormaRotors[0].Start, proFormaRotors[0].Size, append([]byte(nil), proFormaRotors[0].Rotor...))
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
			if !reflect.DeepEqual(r, proFormaRotors[0]) {
				t.Errorf("Sliced Rotor() = %v, want %v", r, proFormaRotors[0])
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
			r := proFormaRotors[0]
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
			r := proFormaRotors[0]
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
				193, 46, 58, 103,
			},
		},
	}
	r := new(Rotor).New(proFormaRotors[0].Size, proFormaRotors[0].Start,
		proFormaRotors[0].Step, append([]byte(nil), proFormaRotors[0].Rotor...))
	r.Current = 0
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
					193, 46, 58, 103,
				},
			},
			want: []byte{
				0, 0, 0, 0,
			},
		},
	}
	r := new(Rotor).New(proFormaRotors[0].Size, proFormaRotors[0].Start,
		proFormaRotors[0].Step, append([]byte(nil), proFormaRotors[0].Rotor...))
	r.Current = 0
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
				"\t184, 25, 190, 250, 35, 11, 111, 218, 111, 1, 44, 59, 137, 12, 184, 22,\n" +
				"\t154, 226, 101, 88, 167, 109, 45, 92, 19, 164, 132, 233, 34, 133, 138, 222,\n" +
				"\t59, 49, 123, 208, 179, 248, 61, 216, 55, 59, 235, 57, 67, 172, 233, 232,\n" +
				"\t87, 236, 189, 170, 196, 124, 216, 109, 4, 106, 207, 150, 166, 164, 99, 57,\n" +
				"\t131, 27, 1, 236, 168, 78, 122, 81, 165, 26, 32, 56, 129, 105, 35, 26,\n" +
				"\t247, 208, 56, 235, 91, 183, 67, 150, 112, 103, 173, 197, 69, 13, 115, 14,\n" +
				"\t129, 206, 74, 46, 119, 208, 95, 67, 119, 7, 191, 210, 128, 117, 140, 245,\n" +
				"\t41, 168, 63, 203, 53, 241, 221, 28, 158, 40, 89, 76, 126, 58, 33, 40,\n" +
				"\t78, 130, 93, 116, 206, 66, 4, 10, 109, 86, 150, 53, 200, 34, 26, 37,\n" +
				"\t232, 185, 214, 47, 131, 18, 241, 210, 18, 81, 107, 161, 97, 65, 238, 250,\n" +
				"\t81, 133, 54, 158, 54, 10, 254, 135, 110, 162, 175, 250, 117, 66, 232, 66,\n" +
				"\t50, 102, 70, 76, 185, 249, 57, 59, 247, 195, 101, 8, 157, 235, 24, 94,\n" +
				"\t204, 74, 100, 196, 93, 24, 179, 27, 118, 168, 29, 10, 38, 204, 210, 123,\n" +
				"\t111, 247, 225, 171, 60, 166, 239, 124, 43, 180, 223, 240, 66, 2, 68, 220,\n" +
				"\t12, 95, 253, 145, 133, 55, 237, 183, 0, 150, 157, 68, 6, 92, 11, 77,\n" +
				"\t241, 50, 172, 211, 182, 22, 174, 9, 82, 194, 116, 145, 66, 69, 111, 0})\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := proFormaRotors[0]
			if got := r.String(); got != tt.want {
				t.Errorf("Rotor.String() = %v, want = %v", got, tt.want)
			}
		})
	}
}
