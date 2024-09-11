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
	// крата простых чисел
	prime map[int]bool
	// карта сумм чисел до 100, кроме сумм простых чисел
	// ключ - сумма. значение - слайс пар чисел, дающих эту сумму
	Sums map[int][]Pair
	// произведений по суммам
	// ключ - сумма. значение - слайс произведений пар чисел
	Muls map[int][]int
)

func main() {
	fmt.Println("Sages...")
	// prime - карта простых чисел
	prime = SieveOfEratosthenes(100)
	// карта сумм чисел до 100, кроме сумм простых чисел
	// ключ - сумма. значение слайс пар чисел, дающих эту сумму
	Sums = SumsNoPrime()
	// вывод ключей мапы в отсортированном слайсе
	fmt.Println("***Суммы и пары***")
	//	OutSortKeyMapPairs(Sums)
	OutSortKeyMapPairs(Sums)

	fmt.Println("***Произведения по суммам***")

	// построение такой же мапы произведений по парам сумм
	Muls = MulsByPairsSums(Sums)
	OutSortKeyMapInts(Muls)
	// удаление элементов слайсов с одинаковыми произведениями
	fmt.Println("***Удаление одинаковых произведений***")
	DelEquElems()
	OutSortKeyMapInts(Muls)
	SearchAnswer()
}

func remove(slice []int, pos int) []int {
	if pos < 0 || pos > len(slice)-1 {
		// ничего удалить нельзя
		return slice
	}
	if pos == 0 {
		// первый элемент
		return slice[1:]
	}
	if pos == len(slice)-1 {
		// последний элемент
		return slice[:len(slice)-1]
	}
	// средние элементы
	return append(slice[:pos], slice[pos+1:]...)
}

func SearchAnswer() {
	var (
		key    int
		s      []int
		ansMul int
	)

	// поиск ответа
	fmt.Println("***Поиск ответа***")
	// найти в Muls элемент с одним произведением
	for key, s = range Muls {
		if len(s) == 1 {
			ansMul = s[0]
			break
		}
	}
	// по ключу элемента найти в Sums пару c исходным произведением
	// выдать пару за ответ
	if pairs, ok := Sums[key]; ok {
		f := false
		for _, p := range pairs {
			if p.n1*p.n2 == ansMul {
				fmt.Println("Решение: ", p)
				f = true
			}
		}
		if !f {
			fmt.Println("Нет решения")
		}
	} else {
		fmt.Println("Нет решения")
	}

}

// удаление одинаковых элементов из слайсов карты произведений
func DelEquElems() {
	for k, _ := range Muls {
		for i := 0; i < len(Muls[k]); {
			f := false
			for k1, _ := range Muls {
				if k != k1 {
					for i1 := 0; i1 < len(Muls[k1]); {
						if Muls[k][i] == Muls[k1][i1] {
							// удаляем элемент i1
							Muls[k1] = remove(Muls[k1], i1)
							// фиксируем факт удаления
							f = true
						} else {
							i1++
						}
					}
				}
			}
			// если были совпадения то удаляем и сам элемент
			if f {
				Muls[k] = remove(Muls[k], i)
			} else {
				i++
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

func OutSortKeyMapPairs(mp map[int][]Pair) {
	ss := make([]int, 0)
	for k, _ := range mp {
		ss = append(ss, k)
	}
	sort.Ints(ss)
	// fmt.Println(ss)
	for _, v := range ss {
		fmt.Println(v, mp[v])
	}
}

func OutSortKeyMapInts(mp map[int][]int) {
	ss := make([]int, 0)
	for k, _ := range mp {
		ss = append(ss, k)
	}
	sort.Ints(ss)
	for _, v := range ss {
		fmt.Println(v, mp[v])
	}
}

func OutSortKeyMapAny(mp map[int][]any) {
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
	// мапа пустых слайсов пар всех возможных сумм от 4 до 100
	for i := 4; i < 100; i++ {
		mp[i] = []Pair{}
	}
	for i := 2; i < 100; i++ {
		for j := i; j < 100; j++ {
			if i+j < 100 {
				if IsPrime(j) && IsPrime(i) {
					// удаляем сумму если оба числа простые
					delete(mp, i+j)
				} else {
					// добавляем если еще не удалено
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
	// карта всех чисел от 2 до n
	for i := 2; i <= n; i++ {
		mapNum[i] = true
	}
	// сбрасываем в false элементы по алгоритму Эратосфена
	for i, v := range mapNum {
		if v {
			for j := i * i; j <= n; j += i {
				mapNum[j] = false
			}
		}
	}
	// удаляем все сброшенные элементы из карты
	for i, v := range mapNum {
		if !v {
			delete(mapNum, i)
		}
	}
	// возвращаем полученную карту
	return mapNum
}

// Функция простоты числа
func IsPrime(n int) bool {
	_, ok := prime[n]
	return ok
}
