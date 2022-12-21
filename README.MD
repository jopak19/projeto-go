## Corrida de revezamento com programação concorrente

O programa demonstra programação concorrente usando GO.

Sâo criadas quatro gorotinas representando cada um dos corredores. As gorotinas são iniciadas e ficam bloqueadas
até o bastão chegar a elas. O bastão é representado por um canal. Quando um valor é recebido o corredor inicia a sua corrida
e, no fim, passa o bastão caso exista algum corredor ainda não tenha corrido. No fim a última gorotina fecha o canal.


## Como executar

Para executar o programa é necessário ter o GO instalado. Para instalar o GO siga as instruções em https://golang.org/doc/install

Para executar o programa, abra um terminal e digite:

```bash
    go run revezamento.go
```