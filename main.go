package main

import (
	"fmt"
)

type Listas struct {
	soliList  []int // Solicitações Lista
	horasList []int // Horas Lista
	pendList  []int // Pendencias Lista
	//alunoMap  map[int]int
}

func main() {
	l := new(Listas)
	var matricula int
	var codHoras int

	l.addSolitacao(matricula, codHoras)
	l.printSolicitacoes()
	//l.printTotalSolicitacoes()
	//l.validarSoli()
}

func (l *Listas) addSolitacao(matricula int, codHoras int) { // Tem como função adicionar solicitações e contar as horas por código digitado.

	countSolicitacoes := 0

	alunosMap := make(map[int]string)

	fmt.Println("\nInforme sua matricula: ")
	_, err := fmt.Scan(&matricula)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("001 - Monitoria de Ensino 2hs\n010 - Participação em Seminário 3hs\n322 - Postagem de artigo 5hs\n324 - Produção de software 20hs\n345 - Patente de produto 50hs\n456 - Participação em minicurso 20hs")
	fmt.Println("Digite o código referente as Horas: ")
	_, err = fmt.Scan(&codHoras)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch codHoras {
	case 1:

		countSolicitacoes += 1
		numHoras := 2

		l.soliList = append(l.soliList, countSolicitacoes)
		l.horasList = append(l.horasList, numHoras)

		formatNum := fmt.Sprintf("%03d", codHoras)

		alunosMap[matricula] = formatNum

		for m, h := range alunosMap {
			fmt.Printf("Matricula e horas do Aluno: %d - %s", m, h)
		}
		// Corrige a formatação do Print do número
		//fmt.Printf("%d - %s", matricula, formatNum)

	case 10:

		countSolicitacoes += 1
		numHoras := 3

		l.soliList = append(l.soliList, countSolicitacoes) // Adiciona +1 na fila de solicitações assim contando o total de Solicitações na fila.
		l.horasList = append(l.horasList, numHoras)        // Adiciona as Horas referente por Código.

		formatNum := fmt.Sprintf("%03d", codHoras) // Corrige a formatação do Print do número

		alunosMap[matricula] = formatNum

		for m, h := range alunosMap {
			fmt.Printf("Matricula e horas do Aluno: %d - %s", m, h)
		}

	case 322:

		countSolicitacoes += 1
		numHoras := 5

		l.soliList = append(l.soliList, countSolicitacoes) // Adiciona +1 na fila de solicitações assim contando o total de Solicitações na fila.
		l.horasList = append(l.horasList, numHoras)        // Adiciona as Horas referente por Código.

		fmt.Printf("%d - %d", matricula, codHoras)

	case 324:

		countSolicitacoes += 1
		numHoras := 20

		l.soliList = append(l.soliList, countSolicitacoes) // Adiciona +1 na fila de solicitações assim contando o total de Solicitações na fila.
		l.horasList = append(l.horasList, numHoras)        // Adiciona as Horas referente por Código.

		alunosMap[matricula] = string(rune(numHoras))

		for m, h := range alunosMap {
			fmt.Printf("Matricula e horas do Aluno: %d - %s", m, h)
		}

	case 345:

		countSolicitacoes += 1
		numHoras := 50

		l.soliList = append(l.soliList, countSolicitacoes) // Adiciona +1 na fila de solicitações assim contando o total de Solicitações na fila.
		l.horasList = append(l.horasList, numHoras)        // Adiciona as Horas referente por Código.

		alunosMap[matricula] = string(rune(numHoras))

		for m, h := range alunosMap {
			fmt.Printf("Matricula e horas do Aluno: %d - %s", m, h)
		}

	case 456:

		countSolicitacoes += 1
		numHoras := 20

		l.soliList = append(l.soliList, countSolicitacoes) // Adiciona +1 na fila de solicitações assim contando o total de Solicitações na fila.
		l.horasList = append(l.horasList, numHoras)        // Adiciona as Horas referente por Código.

		alunosMap[matricula] = string(rune(numHoras))

		for m, h := range alunosMap {
			fmt.Printf("Matricula e horas do Aluno: %d - %s", m, h)
		}

	default:
		fmt.Println("Digite um número valido!")
	}

}

//Valida as Solicitações que estão pendentes

func (l *Listas) validarSoli() {
	if len(l.soliList) > 0 {
		l.soliList = l.soliList[:len(l.soliList)-1] // Remove a contagem de solicitações na fila

		totalHoras := 0
		for _, soma := range l.horasList {
			totalHoras += soma // Soma total das horas já registradas
		}

		fmt.Printf("\nSolicitação validada com Sucesso! O total de horas é: %v", totalHoras)

	} else {
		fmt.Println("\nNão há solicitações para validar!")
	}
}

func (l *Listas) printSolicitacoes() {
	fmt.Printf("\nSolicitações para análise: %d", l.pendList) // Print na tela das Solicitações pendentes

}

func (l *Listas) printTotalSolicitacoes() {
	fmt.Printf("\nTotal de Solicitações para análise: %d", len(l.soliList))
}
