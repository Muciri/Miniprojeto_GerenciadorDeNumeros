package main
import (
    "fmt"
    "errors"
    "slices"
    "sort"
    "os"
    "strconv"
)

var lista []int

func adicionarNumero(numero int) error {
    if numero < 0 {
        return errors.New("numero tem que ser maior que 0")
    }
   
    lista = append(lista, numero)
    return nil
}

func listarNumeros() []int  {
    return lista
}

func listarNumerosPares() []int  {
    var listaPar []int
   
    for _, numero := range lista {
        if numero%2 == 0 {
            listaPar = append(listaPar, numero)
        }
    }
   
    return listaPar
}

func removerNumero(indice int) {
    if indice < 0 || indice >= len(lista) {
        return
    }
   
    lista = append(lista[:indice], lista[indice+1:]...)
}

func calcularEstatisticas() (int, int, float64) {
    if len(lista) == 0 {
        return 0, 0, 0.0
    }
   
    var minimo int = lista[0]
    var maximo int = lista[0]
    var media float64 = 0.0
   
    for _, numero := range lista {
        if numero < minimo {
            minimo = numero
        }
       
        if numero > maximo {
            maximo = numero
        }
       
        media += float64(numero)
    }
    media = media / float64(len(lista))


    return minimo, maximo, media
}

func dividir(numero1,numero2 float64) (float64, error) {
    if numero2 == 0 {
        return 0, errors.New("não pode dividir por 0")
    }

    return numero1 / numero2, nil
}

func limparLista() {
    lista = lista[:0]
}

func ordenarCrescente() {
    sort.Ints(lista)
}

func ordenarDecrescente() {
    sort.Ints(lista)
    slices.Reverse(lista)
}

func exportarArquivo() error {
    arquivo, err := os.Create("lista.txt")
    if err != nil {
        return err
    }
    
    defer arquivo.Close()

    fmt.Fprintf(arquivo, "lista: [")
    for indice, numero := range lista {
        _, err := fmt.Fprintf(arquivo, strconv.Itoa(numero))
        if indice != len(lista)-1 {
            fmt.Fprintf(arquivo, ", ")
        }
        if err != nil {
            return err
        }
    }
    fmt.Fprintf(arquivo, "]")

    return nil
}

func limparTerminal() {
    fmt.Print("\033[H\033[2J")
}

func main() {
    var escolha int

    for {
    fmt.Print(`
Gerenciador de Números
===== MENU =====
1  - Adicionar Número
2  - Listar Números
3  - Remover por Índice
4  - Estatísticas
5  - Divisão De Numeros
6  - Limpar Lista
7  - Listar Numeros Pares
8  - Ordenar a Lista em Ordem Crescente
9  - Ordenar a Lista em Ordem Decrescente
10 - Exportar em arquivo texto
0  - Sair
Escolha uma opção: `)

        fmt.Scan(&escolha)
        limparTerminal()
        fmt.Println("===================================")

        switch escolha {
            case 1:
                var numero int
                fmt.Println("escreva o número: ")
                fmt.Scan(&numero)
               
                err := adicionarNumero(numero)
           
                if err != nil {
                    fmt.Println(err)
                } else {
                    fmt.Println("Número adicionado: ", numero)
                }
   
            case 2:
                fmt.Println("lista: ",listarNumeros())
   
            case 3:
                var indice int
                fmt.Println("escreva o número: ")
                fmt.Scan(&indice)
                removerNumero(indice)
   
            case 4:
                minimo, maximo, media := calcularEstatisticas()
                fmt.Println("estatísticas:")
                fmt.Println("minimo: ",minimo)
                fmt.Println("maximo: ",maximo)
                fmt.Println("media: ",media)
   
            case 5:
                var numero1 float64
                var numero2 float64
   
                fmt.Println("escreva o primeiro número: ")
                fmt.Scan(&numero1)
                fmt.Println("escreva o segundo número: ")
                fmt.Scan(&numero2)
   
                resultado, erro := dividir(numero1, numero2)
   
                if erro == nil {
                fmt.Println(resultado)
                } else {
                    fmt.Println(erro)
                }
   
            case 6:
                limparLista()
            case 7:
                fmt.Println(listarNumerosPares())
               
            case 8:
                ordenarCrescente()
                fmt.Println("lista ordenada em ordem crescente")
               
            case 9:
                ordenarDecrescente()
                fmt.Println("lista ordenada em ordem decrescente")

            case 10:
                err := exportarArquivo()

                if err != nil {
                    fmt.Println("Erro ao exportar:", err)
                } else {
                    fmt.Println("Lista exportada com sucesso!")
                }
   
            case 0:
                fmt.Println("Saindo...")
                return
   
            default:
                fmt.Println("Opção inválida!")
        }
        fmt.Println("===================================")
    }
}
