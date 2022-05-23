package main

import (
	"testing"
)

func Test_imprimeArquivo(t *testing.T) {
	type args struct {
		nomeDoArquivo string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			imprimeArquivo(tt.args.nomeDoArquivo)
		})
	}
}

func Test_registraProdutoComprado(t *testing.T) {
	type args struct {
		p Product
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registraProdutoComprado(tt.args.p)
		})
	}
}
