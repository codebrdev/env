package env

import (
	"reflect"
	"testing"
)

func TestGetAsSlice(t *testing.T) {
	const key = "TEST_GET_AS_SLICE"
	def := []string{"default"}

	tests := []struct {
		name  string
		value *string
		sep   string
		want  []string
	}{
		{"variável ausente retorna o default", nil, ",", def},
		{"valor simples", ptr("a,b,c"), ",", []string{"a", "b", "c"}},
		{"remove espaços ao redor dos itens", ptr(" a , b ,c "), ",", []string{"a", "b", "c"}},
		{"ignora itens vazios", ptr("a,,b, ,c,"), ",", []string{"a", "b", "c"}},
		{"separador customizado", ptr("a;b;c"), ";", []string{"a", "b", "c"}},
		{"item único", ptr("a"), ",", []string{"a"}},
		{"valor vazio retorna slice vazio, não o default", ptr(""), ",", []string{}},
		{"apenas separadores retorna slice vazio", ptr(",,,"), ",", []string{}},
		{"separador ausente no valor mantém item único", ptr("a;b"), ",", []string{"a;b"}},
		{"preserva espaços internos", ptr("a b, c d"), ",", []string{"a b", "c d"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != nil {
				t.Setenv(key, *tt.value)
			}

			got := GetAsSlice(key, def, tt.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAsSlice() = %#v, esperado %#v", got, tt.want)
			}
		})
	}
}

func ptr(s string) *string { return &s }
