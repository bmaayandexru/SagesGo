package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	n1, n2 int
}

//type APairs []Pair

var (
	prime map[int]bool
	Sums  map[int][]Pair
	Muls  map[int][]int
)

func main() {
	fmt.Println("Sages...")
	// prime - карта простых чисел
	prime = SieveOfEratosthenes(100)
	//	fmt.Println(prime)

	// карта сумм, кроме сумм простых чисел
	Sums = SumsNoPrime()
	// вывод ключей мапы в отсортированном слайсе
	OutSortKeyMap(Sums)
	OutSortKeyMap1(Sums)

	fmt.Println("*****************")

	// построение такой же мапы произведений по парам сумм
	Muls = MulsByPairsSums(Sums)
	OutSortKeyMap2(Muls)
	DeleteEquElems()
	OutSortKeyMap2(Muls)
}

func DeleteEquElems() {
	for k, v := range Muls {
		for i, s := range v {
			f := false
			for k1, v1 := range Muls {
				if k != k1 {
					for i1, s1 := range v1 {
						if s == s1 {
							// удаляем элементы s1
							v1[i1] = 0
							// фиксируем факт удаления
							f = true
						}
					}
				}
			}
			// если были совпадения то удаляем и сам элемент
			if f {
				v[i] = 0
			}
		}
	}
}

func MulsByPairsSums(Sums map[int][]Pair) map[int][]int {
	mp := make(map[int][]int)
	for _, pairs := range Sums {
		for _, p := range pairs {
			mp[p.n1+p.n2] = append(mp[p.n1+p.n2], p.n1*p.n2)
		}
	}
	return mp
}

func OutSortKeyMap(mp map[int][]Pair) {
	ss := make([]int, 0)
	for k, _ := range mp {
		ss = append(ss, k)
	}
	sort.Ints(ss)
	fmt.Println(ss)
}

func OutSortKeyMap1(mp map[int][]Pair) {
	ss := make([]int, 0)
	for k, _ := range mp {
		ss = append(ss, k)
	}
	sort.Ints(ss)
	for _, v := range ss {
		fmt.Println(v, mp[v])
	}

}

func OutSortKeyMap2(mp map[int][]int) {
	ss := make([]int, 0)
	for k, _ := range mp {
		ss = append(ss, k)
	}
	sort.Ints(ss)
	for _, v := range ss {
		fmt.Println(v, mp[v])
	}

}

func DeleteSumsDoubleMul() {
	for k, v := range Sums {
		// проход по сумам
		// k значение суммы
		// v слайс пар
		found := false
		for _, p := range v {
			// проход по парам
			mul := p.n1 * p.n2
			// нужно искать mul в остальных суммах кроме этой
			for k1, v1 := range Sums {
				if k1 != k {
					for _, p1 := range v1 {
						// идем по вектору и ищем mul
						if p1.n1*p1.n2 == mul {
							// удаляем текущую сумму и ставим флаг
							delete(Sums, k1)
							found = true
						}
					}
				}
			}
		}

		if found {
			// нужно удалить и эту сумму
			delete(Sums, k)
		}

	}
}

func SumsNoPrime() map[int][]Pair {
	mp := make(map[int][]Pair)
	// мапа пустых слайсов
	for i := 4; i < 100; i++ {
		mp[i] = []Pair{}
	}
	for i := 2; i < 100; i++ {
		for j := i; j < 100; j++ {
			if i+j < 100 {
				if IsPrime(j) && IsPrime(i) {
					// удаляем
					delete(mp, i+j)
				} else {
					// добавляем если есть
					if _, ok := mp[i+j]; ok {
						mp[i+j] = append(mp[i+j], Pair{i, j})
					}
				}
			}
		}
	}
	return mp
}

func NoDoublePrime(i, j int) bool {
	pi := IsPrime(i)
	pj := IsPrime(j)
	return (!pi || !pj)
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
// func IsPrime(n int, m map[int]bool) bool {
func IsPrime(n int) bool {
	_, ok := prime[n]
	// fmt.Println(n, k, ok)
	return ok
}
