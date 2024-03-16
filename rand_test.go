// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tnt2engine

import (
	"reflect"
	"testing"
)

func closeTntMachine(e *Tnt2Engine) {
	blk := new(CipherBlock)
	e.Left() <- *blk
	<-e.Right()
}

func TestNewRand(t *testing.T) {
	tntMachine := new(Tnt2Engine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer closeTntMachine(tntMachine)
	type args struct {
		src *Tnt2Engine
	}
	tests := []struct {
		name  string
		args  args
		want  *Rand
		wantK string
	}{
		{
			name:  "NewRandTest 1",
			args:  args{tntMachine},
			want:  &Rand{tntMachine, CipherBlockBytes, emptyBlk},
			wantK: "8MyZ1wEtrXp1/krHycfE7jnplAlELSBzAPkLsnNPwLo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.CounterKey() != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.CounterKey(), tt.wantK)
			}
			if got := NewRand(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tntMachine := new(Tnt2Engine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer closeTntMachine(tntMachine)
	type args struct {
		src *Tnt2Engine
	}
	tests := []struct {
		name  string
		args  args
		want  *Rand
		wantK string
	}{
		{
			name:  "NewTest 1",
			args:  args{tntMachine},
			want:  &Rand{tntMachine, CipherBlockBytes, emptyBlk},
			wantK: "8MyZ1wEtrXp1/krHycfE7jnplAlELSBzAPkLsnNPwLo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.CounterKey() != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.CounterKey(), tt.wantK)
			}
			if got := new(Rand).New(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRand_Intn(t *testing.T) {
	tntMachine := new(Tnt2Engine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer closeTntMachine(tntMachine)
	rnd := new(Rand).New(tntMachine)
	type args struct {
		max int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		wantK string
		wantR *Rand
	}{
		{
			name:  "Intn Test 1",
			args:  args{1000},
			want:  694,
			wantK: "8MyZ1wEtrXp1/krHycfE7jnplAlELSBzAPkLsnNPwLo",
			wantR: &Rand{tntMachine, CipherBlockBytes, emptyBlk},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.CounterKey() != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.CounterKey(), tt.wantK)
			}
			if !reflect.DeepEqual(rnd, tt.wantR) {
				t.Errorf("NewRand() = %v, want %v", rnd, tt.wantR)
			}
			if got := rnd.Intn(tt.args.max); got != tt.want {
				t.Errorf("Rand.Intn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRand_Int63n(t *testing.T) {
	tntMachine := new(Tnt2Engine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer closeTntMachine(tntMachine)
	rnd := new(Rand).New(tntMachine)
	type args struct {
		n int64
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		wantK string
		wantR *Rand
	}{
		{
			name:  "Int63n Test 1",
			args:  args{1000000000},
			want:  117654694,
			wantK: "8MyZ1wEtrXp1/krHycfE7jnplAlELSBzAPkLsnNPwLo",
			wantR: &Rand{tntMachine, CipherBlockBytes, emptyBlk},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.CounterKey() != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.CounterKey(), tt.wantK)
			}
			if !reflect.DeepEqual(rnd, tt.wantR) {
				t.Errorf("New() = %v, want %v", rnd, tt.wantR)
			}
			if got := rnd.Int63n(tt.args.n); got != tt.want {
				t.Errorf("Rand.Int63n() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRand_Perm(t *testing.T) {
	tntMachine := new(Tnt2Engine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer closeTntMachine(tntMachine)
	rnd := new(Rand).New(tntMachine)
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		wantK string
	}{
		{
			name:  "Prem Test 1",
			args:  args{10},
			want:  []int{5, 0, 4, 1, 9, 6, 7, 2, 3, 8},
			wantK: "8MyZ1wEtrXp1/krHycfE7jnplAlELSBzAPkLsnNPwLo",
		}, {
			name: "Prem Test 1",
			args: args{256},
			want: []int{
				134, 84, 98, 76, 50, 209, 141, 99, 89, 24, 68, 55, 181, 62, 78, 66,
				164, 233, 96, 57, 30, 242, 94, 168, 186, 187, 12, 104, 34, 142, 194, 153,
				3, 150, 179, 204, 43, 26, 212, 25, 126, 69, 2, 39, 102, 170, 23, 53,
				234, 61, 173, 111, 247, 122, 77, 211, 19, 152, 110, 129, 90, 120, 198, 200,
				27, 193, 243, 143, 37, 54, 42, 221, 210, 199, 156, 22, 125, 79, 133, 95,
				56, 145, 228, 112, 139, 171, 220, 51, 166, 189, 241, 63, 117, 216, 118, 254,
				213, 107, 132, 140, 154, 72, 36, 236, 114, 97, 188, 163, 1, 224, 227, 127,
				7, 158, 100, 21, 190, 207, 0, 109, 131, 108, 251, 29, 255, 80, 203, 192,
				41, 183, 174, 33, 230, 32, 149, 177, 138, 161, 9, 165, 155, 159, 101, 70,
				148, 182, 28, 201, 232, 178, 176, 14, 81, 60, 85, 71, 13, 8, 87, 137,
				64, 214, 169, 249, 246, 240, 128, 130, 73, 47, 65, 49, 180, 191, 208, 38,
				115, 226, 136, 74, 185, 217, 175, 75, 58, 219, 10, 172, 196, 5, 119, 35,
				105, 48, 4, 93, 151, 250, 44, 225, 160, 103, 88, 46, 215, 157, 184, 239,
				205, 15, 162, 91, 45, 116, 167, 83, 223, 16, 92, 59, 113, 252, 206, 222,
				123, 229, 238, 86, 106, 237, 218, 253, 40, 195, 197, 202, 11, 67, 20, 245,
				147, 248, 231, 17, 31, 82, 121, 146, 235, 135, 18, 6, 244, 144, 124, 52},
			wantK: "8MyZ1wEtrXp1/krHycfE7jnplAlELSBzAPkLsnNPwLo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.CounterKey() != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.CounterKey(), tt.wantK)
			}
			if got := rnd.Perm(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rand.Perm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRand_Read(t *testing.T) {
	tntMachine := new(Tnt2Engine)
	tntMachine.Init([]byte("SecretKey"), "")
	tntMachine.SetEngineType("E")
	tntMachine.SetIndex(BigZero)
	tntMachine.BuildCipherMachine()
	defer closeTntMachine(tntMachine)
	rnd := new(Rand).New(tntMachine)
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantN   int
		wantErr bool
		wantK   string
	}{
		{
			name: "Read test 1",
			args: args{make([]byte, 36)},
			want: []byte{
				10, 84, 10, 4, 202, 145, 128, 166, 213, 202, 0, 123, 58, 250, 108, 83,
				54, 175, 122, 26, 36, 72, 197, 185, 158, 47, 241, 64, 124, 176, 71, 230,
				243, 129, 174, 233},
			wantN:   36,
			wantErr: false,
			wantK:   "8MyZ1wEtrXp1/krHycfE7jnplAlELSBzAPkLsnNPwLo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.CounterKey() != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.CounterKey(), tt.wantK)
			}
			gotN, err := rnd.Read(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rand.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Rand.Read() count = %v, want %v", gotN, tt.wantN)
			}
			if !reflect.DeepEqual(tt.args.p, tt.want) {
				t.Errorf("Rand.Read() = %v, want %v", tt.args.p, tt.want)
			}
		})
	}
}
