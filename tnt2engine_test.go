// This is free and unencumbered software released into the public domain.
// See the UNLICENSE file for details.

package tnt2engine

import (
	"io"
	"math/big"
	"reflect"
	"testing"
)

func TestTnt2Engine_Left(t *testing.T) {
	var tnt2Machine Tnt2Engine
	tnt2Machine.Init([]byte("SecretKey"), "")
	tnt2Machine.SetEngineType("E")
	tnt2Machine.BuildCipherMachine()
	tests := []struct {
		name string
		want chan CipherBlock
	}{
		{
			name: "ttel1",
			want: tnt2Machine.left,
		},
	}
	for _, tt := range tests {
		e := tnt2Machine
		t.Run(tt.name, func(t *testing.T) {
			if got := e.Left(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tnt2Engine.Left() = %v, want %v", got, tt.want)
			}
		})
	}
	tnt2Machine.CloseCipherMachine()
}

func TestTnt2Engine_Right(t *testing.T) {
	var tnt2Machine Tnt2Engine
	tnt2Machine.Init([]byte("SecretKey"), "")
	tnt2Machine.SetEngineType("E")
	tnt2Machine.BuildCipherMachine()
	tests := []struct {
		name string
		want chan CipherBlock
	}{
		{
			name: "tter1",
			want: tnt2Machine.right,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tnt2Machine
			if got := e.Right(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tnt2Engine.Right() = %v, want %v", got, tt.want)
			}
		})
	}
	var blk CipherBlock
	tnt2Machine.left <- blk
	<-tnt2Machine.right
}

func TestTnt2Engine_CounterKey(t *testing.T) {
	var tnt2Machine Tnt2Engine
	tests := []struct {
		name             string
		key              string
		proFormaFileName string
		want             string
	}{
		{
			name:             "ttec1",
			key:              "SecretKey",
			proFormaFileName: "",
			want:             "ab2677fa2eecca36541ea85fd8d871203383b898bb025b8ec8fd5f24719eee1c",
		},
		{
			name:             "ttec2",
			key:              "SecretKey",
			proFormaFileName: "test.proforma.json",
			want:             "9d468a888bf287c0cdc3008569b76c0cb7091b062f0b6209461436534392f95c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tnt2Machine.Init([]byte(tt.key), tt.proFormaFileName)
			if got := tnt2Machine.CounterKey(); got != tt.want {
				t.Errorf("Tnt2Engine.CounterKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTnt2Engine_Index(t *testing.T) {
	var tnt2Machine Tnt2Engine
	tnt2Machine.Init([]byte("SecretKey"), "")
	iCnt, _ := new(big.Int).SetString("1234567890", 10)
	tnt2Machine.SetIndex(iCnt)
	tests := []struct {
		name     string
		wantCntr *big.Int
	}{
		{
			name:     "ttei1",
			wantCntr: iCnt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tnt2Machine
			if gotCntr := e.Index(); !reflect.DeepEqual(gotCntr, tt.wantCntr) {
				t.Errorf("Tnt2Engine.Index() = %v, want %v", gotCntr, tt.wantCntr)
			}
		})
	}
}

func TestTnt2Engine_SetIndex(t *testing.T) {
	var tnt2Machine Tnt2Engine
	tnt2Machine.Init([]byte("SecretKey"), "")
	iCnt, _ := new(big.Int).SetString("1234567890", 10)
	type args struct {
		iCnt *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "ttesi1",
			args: args{iCnt},
			want: iCnt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tnt2Machine
			e.SetIndex(tt.args.iCnt)
			if got := e.Index(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tnt2Engine.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTnt2Engine_SetEngineType(t *testing.T) {
	var tnt2Machine Tnt2Engine
	tnt2Machine.Init([]byte("SecretKey"), "")
	type args struct {
		engineType string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "tteset1",
			args: args{"E"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tnt2Machine
			e.SetEngineType(tt.args.engineType)
			if got := e.engineType; got != tt.args.engineType {
				t.Errorf("Tnt2Engine.SetEngineType() = %v, want %v", got, tt.args.engineType)
			}
		})
	}
}

func TestTnt2Engine_Engine(t *testing.T) {
	var tnt2Machine Tnt2Engine
	tnt2Machine.Init([]byte("SecretKey"), "")
	tests := []struct {
		name string
		want []Crypter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tnt2Machine
			if got := e.Engine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tnt2Engine.Engine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTnt2Engine_EngineType(t *testing.T) {
	var tnt2Machine Tnt2Engine
	tnt2Machine.Init([]byte("SecretKey"), "")
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tnt2Machine
			if got := e.EngineType(); got != tt.want {
				t.Errorf("Tnt2Engine.EngineType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTnt2Engine_MaximalStates(t *testing.T) {
	var tnt2Machine Tnt2Engine
	tnt2Machine.Init([]byte("SecretKey"), "")
	want, _ := new(big.Int).SetString("49101257188406090296051850365430624307", 10)
	tests := []struct {
		name string
		want *big.Int
	}{
		{
			name: "tteset1",
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tnt2Machine
			if got := e.MaximalStates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tnt2Engine.MaximalStates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTnt2Engine_Init(t *testing.T) {
	var tnt2Machine Tnt2Engine
	tnt2Machine.Init([]byte("SecretKey"), "")
	type args struct {
		secret           []byte
		proFormaFileName string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tnt2Machine
			e.Init(tt.args.secret, tt.args.proFormaFileName)
		})
	}
}

func TestTnt2Engine_BuildCipherMachine(t *testing.T) {
	var tnt2Machine Tnt2Engine
	tnt2Machine.Init([]byte("SecretKey"), "")
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tnt2Machine
			e.BuildCipherMachine()
		})
	}
}

func Test_createProFormaMachine(t *testing.T) {
	type args struct {
		pfmReader io.Reader
	}
	tests := []struct {
		name string
		args args
		want *[]Crypter
	}{
		{
			name: "tcpfm1",
			args: args{pfmReader: nil},
			want: &[]Crypter{Rotor1, Rotor2, Permutator1, Rotor3, Rotor4, Permutator2, Rotor5, Rotor6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createProFormaMachine(tt.args.pfmReader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createProFormaMachine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateRotor(t *testing.T) {
	type args struct {
		r      *Rotor
		random *Rand
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.r.Update(tt.args.random)
		})
	}
}

func Test_updatePermutator(t *testing.T) {
	type args struct {
		p      *Permutator
		random *Rand
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.p.Update(tt.args.random)
		})
	}
}
