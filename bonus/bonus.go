package bonus

import "sort"
//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.
func CalculateTowers(barLengths []int) (int, int) {
	sort.Ints(barLengths)
	unicobarra := make(map[int]bool)
	for _, length := range barLengths {
		unicobarra[length] = true
	}
		caontador:= make(map[int]int)
		for _, length := range barLengths {
			caontador[length]++
		}
		numerotorre:=len(unicobarra)
		var maxaltura int
		for _, contar:= range caontador {
			if contar > maxaltura {
				maxaltura = contar
			}
		}
	return maxaltura, numerotorr
}
