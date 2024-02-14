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
			wantK: "ab2677fa2eecca36541ea85fd8d871203383b898bb025b8ec8fd5f24719eee1c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.cntrKey != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.cntrKey, tt.wantK)
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
			wantK: "ab2677fa2eecca36541ea85fd8d871203383b898bb025b8ec8fd5f24719eee1c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.cntrKey != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.cntrKey, tt.wantK)
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
			want:  855,
			wantK: "ab2677fa2eecca36541ea85fd8d871203383b898bb025b8ec8fd5f24719eee1c",
			wantR: &Rand{tntMachine, CipherBlockBytes, emptyBlk},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.cntrKey != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.cntrKey, tt.wantK)
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
			want:  156611855,
			wantK: "ab2677fa2eecca36541ea85fd8d871203383b898bb025b8ec8fd5f24719eee1c",
			wantR: &Rand{tntMachine, CipherBlockBytes, emptyBlk},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.cntrKey != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.cntrKey, tt.wantK)
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
			want:  []int{3, 0, 8, 6, 7, 2, 9, 5, 1, 4},
			wantK: "ab2677fa2eecca36541ea85fd8d871203383b898bb025b8ec8fd5f24719eee1c",
		}, {
			name: "Prem Test 1",
			args: args{256},
			want: []int{
				92, 214, 104, 48, 130, 202, 98, 35, 114, 213, 6, 75, 95, 155, 234, 117,
				58, 217, 121, 0, 27, 67, 143, 77, 38, 59, 99, 50, 254, 132, 20, 93,
				152, 139, 120, 15, 158, 53, 195, 60, 116, 63, 136, 200, 141, 40, 43, 145,
				102, 89, 246, 42, 71, 125, 82, 8, 64, 2, 198, 164, 204, 81, 245, 180,
				183, 127, 205, 224, 157, 122, 232, 1, 128, 182, 239, 225, 70, 11, 23, 162,
				24, 219, 28, 163, 49, 55, 110, 47, 44, 237, 192, 147, 185, 196, 193, 170,
				175, 222, 156, 74, 181, 4, 230, 253, 233, 207, 96, 194, 215, 161, 91, 167,
				31, 146, 220, 144, 100, 153, 112, 13, 142, 19, 251, 39, 79, 94, 148, 179,
				255, 131, 173, 84, 65, 150, 221, 149, 86, 72, 235, 223, 154, 203, 17, 16,
				140, 184, 240, 212, 33, 68, 41, 160, 9, 252, 37, 30, 14, 243, 248, 197,
				45, 165, 87, 90, 250, 216, 34, 36, 227, 177, 7, 76, 62, 5, 242, 206,
				106, 78, 236, 209, 231, 3, 54, 26, 52, 85, 171, 12, 109, 97, 124, 56,
				22, 176, 113, 244, 21, 119, 218, 188, 247, 178, 105, 151, 29, 10, 199, 174,
				115, 103, 168, 138, 134, 51, 123, 249, 201, 169, 69, 83, 73, 137, 211, 186,
				210, 80, 238, 108, 46, 189, 18, 107, 208, 159, 61, 25, 241, 190, 133, 118,
				191, 166, 129, 111, 66, 101, 126, 32, 172, 187, 57, 88, 228, 229, 226, 135},
			wantK: "ab2677fa2eecca36541ea85fd8d871203383b898bb025b8ec8fd5f24719eee1c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.cntrKey != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.cntrKey, tt.wantK)
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
				255, 1, 190, 15, 162, 173, 77, 15, 188, 20,
				132, 243, 237, 246, 4, 49, 234, 241, 119, 54,
				14, 6, 223, 131, 169, 200, 208, 236, 69, 16,
				126, 146, 233, 246, 104, 252},
			wantN:   36,
			wantErr: false,
			wantK:   "ab2677fa2eecca36541ea85fd8d871203383b898bb025b8ec8fd5f24719eee1c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tntMachine.cntrKey != tt.wantK {
				t.Errorf("tntMachine.cntrKey = %v, want %v", tntMachine.cntrKey, tt.wantK)
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
