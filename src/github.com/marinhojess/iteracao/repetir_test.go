package iteracao

import "testing"

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", 8)
	esperado := "aaaaaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s', mas obteve '%s'", esperado, repeticoes)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 8)
	}
}
