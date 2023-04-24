# multithreading-go

## Como funciona

O sistema faz duas requisições para duas APIs distintas de consulta de CEP e exibe a
resposta da qual responder primeiro e descarta a do outro.

Caso não respondam dentro do intervalor de 1 segundo, exibe a resposta: "Timeout".

IMPORTANTE: Caso receba um erro de alguma das duas, exibi o erro no terminal e continua
a execução da qual não falhou (obdescendo o intervalo de 1s) e exibe a resposta dizendo
qual é e sua resposta.

Caso ambas falhem irá exibir no terminal o motivo pelo qual ambas falharam e um possível
timeout.

Caso o número de argumentos utilizados seja incorreto irá ser exibido um erro e não será
executado.

### Como executar

Basta digitar o comando abaixo seguido do cep de consulta desejado

```bash
go run main.go [cep]
```

