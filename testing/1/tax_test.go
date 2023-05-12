package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBacth(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}

	for _, test := range table {
		result := CalculateTax(test.amount)
		if result != test.expected {
			t.Errorf("Expected %f but got %f", test.expected, result)
		}
	}
}

/*
$ go test -v (verbose)

$ go test -coverprofile=coverage.out (cobertura de teste)
*/

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}
func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

/*
$ go test -bench=. (teste de benchmark padrão)

$ go test -bench=. -run=^# (teste de bench com parametro padrao)

$ go test -bench=. -run=^# -count=10 (executando 10 vezes de cada teste)

$ go test -bench=. -run=^# -count=10 -benchtime=3s (colocando o time de 3s)

$ go test -bench=. -run=^# -benchmem (benchmark de memoria)
*/

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1.0, -2.0, -2.5, 500.0, 1000.0, 10501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0 ", result)
		}
		if amount > 20000 && result != 20.0 {
			t.Errorf("Received %f but expected 20 ", result)
		}
	})
}

/*
 $ go test -fuzz=.
*/
