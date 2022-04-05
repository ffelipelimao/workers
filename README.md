### Treinando o conceito de workers

Nesse caso, estamos gerando em paralelo atraves das goroutines um worker, onde cada worker vai verificar a informacao que esta passada via channel, e imprimir seu valor

E o valor que esta sendo passado para o channel tambem esta sendo usado como uma goroutine com um loop infinito incrementando um valor de 0 a infinito

e depois o valor do channel de done esta sendo repassado ao vazio, para nao terminar o progama