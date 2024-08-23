package main

import "fmt"

type Pair struct {
	n1, n2 int
}

//type APairs []Pair

var (
	prime map[int]bool
	Sums  map[int][]Pair
)

func main() {
	fmt.Println("Sages...")
	// prime - карта простых чисел
	prime = SieveOfEratosthenes(100)
	fmt.Println(prime)

	// карта сумм, кроме сумм простых чисел
	Sums = SumsNoPrime()
	// fmt.Println(mp)
	for k, v := range Sums {
		fmt.Println(k, len(v))
	}
	fmt.Println("*****************")
	DeleteSingleSums()
	for k, v := range Sums {
		fmt.Println(k, len(v))
	}
	// удаление однозначных сумм
	// удаление сумм, произведения которых встречаются и в других суммах
	// поиск сумм с одним вариантом произведения
}

func SumsNoPrime() map[int][]Pair {
	mp := make(map[int][]Pair)
	for i := 2; i < 100; i++ {
		for j := i; j < 100; j++ {
			if (i+j < 100) && (i != j) && !(IsPrime(j) && IsPrime(i)) {
				if _, ok := mp[i+j]; !ok {
					mp[i+j] = make([]Pair, 0)
				}
				mp[i+j] = append(mp[i+j], Pair{i, j})
			}
		}
	}
	return mp
}

func DeleteSingleSums() {
	for k, v := range Sums {
		if len(v) == 1 {
			delete(Sums, k)
		}
	}
}

// Решето Эратосфена. Полушаем карту простых чисел
func SieveOfEratosthenes(n int) map[int]bool {
	mapNum := make(map[int]bool)
	for i := 2; i <= n; i++ {
		mapNum[i] = true
	}
	for i, v := range mapNum {
		if v {
			for j := i * i; j <= n; j += i {
				mapNum[j] = false
			}
		}
	}
	for i, v := range mapNum {
		if !v {
			delete(mapNum, i)
		}
	}
	return mapNum
}

// Функция простоты числа
//func IsPrime(n int, m map[int]bool) bool {
func IsPrime(n int) bool {
	_, ok := prime[n]
	// fmt.Println(n, k, ok)
	return ok
}
