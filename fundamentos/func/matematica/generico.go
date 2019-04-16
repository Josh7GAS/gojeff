package matematica

/*
Executa qualquer tipo de calculo
*/
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplica o numero x*y
func Multiplicador(x int, y int) int {

	return x * y
}

//Divisor efetua a divisão do numeroA pelo numeroB
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

//Divisor efetua a divisão do numeroA pelo numeroB e devolve o resto
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}
