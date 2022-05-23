# Manhã

## Exercício 1 - Guardar arquivo

Uma empresa que vende produtos de limpeza precisa de:
1. Implementar uma funcionalidade para guardar um arquivo de texto, com a informação
de produtos comprados, separados por ponto e vírgula (csv).
2. Deve possuir o ID do produto, preço e a quantidade.
3. Estes valores podem ser hardcodeados ou escritos manualmente em uma variável.

## Exercício 2 - Ler arquivo

A mesma empresa precisa ler o arquivo armazenado, para isso exige que:
- Seja impresso na tela os valores tabelados, com título ( à esquerda para o ID e à
direita para o Preço e Quantidade), o preço, a quantidade e abaixo do preço o total
deve ser exibido (somando preço por quantidade)

Exemplo:
ID Preco Quantidade
111223 30012.00 1
444321 1000000.00 4
434321 50.50 1

4030062.50

# Tarde

## Exercício 1 - Rede social

Uma empresa de mídia social precisa implementar uma estrutura de usuários com funções
que acrescentem informações à estrutura. Para otimizar e economizar memória, eles exigem
que a estrutura de usuários ocupe o mesmo lugar na memória para o programa principal e
para as funções:
- A estrutura deve possuir os seguintes campos: Nome, Sobrenome, idade, email e
senha
E devem implementar as seguintes funcionalidades:
- mudar o nome: me permite mudar o nome e sobrenome
- mudar a idade: me permite mudar a idade
- mudar e-mail: me permite mudar o e-mail
- mudar a senha: me permite mudar a senha

## Exercício 2 - E-commerce (Parte II)
Uma grande empresa de vendas na web precisa adicionar funcionalidades para adicionar
produtos aos usuários. Para fazer isso, eles exigem que usuários e produtos tenham o
mesmo endereço de memória no main do programa e nas funções.

Estruturas necessárias:
- Usuário: Nome, Sobrenome, E-mail, Produtos (array de produtos).
- Produto: Nome, preço, quantidade.
Algumas funções necessárias:
- Novo produto: recebe nome e preço, e retorna um produto.
- Adicionar produto: recebe usuário, produto e quantidade, não retorna nada, adiciona
o produto ao usuário.
- Deletar produtos: recebe um usuário, apaga os produtos do usuário.