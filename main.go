package main

import "fmt"

type Pair struct {
	n1, n2 int
}

type APairs []Pair

var (
	prime map[int]bool
	Sums  map[int]APairs
)

func main() {
	fmt.Println("Sages...")
	// prime - карта простых чисел
	prime = make(map[int]bool)
	prime := SieveOfEratosthenes(100)
	fmt.Println(prime)
	for i := 1; i < 100; i++ {
		//k, v := prime[i]
		//fmt.Println(i, k, v)
		IsPrime(i)
	}
	// карта сумм, кроме сумм простых чисел
	// Sums = SumsNoPrime()
	// fmt.Println(Sums)
	// удаление однозначных сумм
	// удаление сумм, произведения которых встречаются и в других суммах
	// поиск сумм с одним вариантом произведения
}

func SumsNoPrime() map[int]APairs {
	mp := make(map[int]APairs)
	for i := 2; i < 100; i++ {
		for j := i; j < 100; j++ {
			ki, vi := prime[i]
			kj, vj := prime[j]
			fmt.Println(ki, kj, vi, vj)
			if (i+j < 100) && (i != j) && !(IsPrime(j) && IsPrime(i)) {
				if _, ok := mp[i+j]; !ok {
					mp[i+j] = make(APairs, 0)
				}
				mp[i+j] = append(mp[i+j], Pair{i, j})
			}
		}
	}
	// fmt.Println(mp)
	for k, v := range mp {
		fmt.Println(k, v)
	}
	return mp
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
func IsPrime(n int) bool {
	k, ok := prime[n]
	fmt.Println(n, k, ok)
	return k
}
