# capgemini-brennoirvine

Esse é o meu projeto seguindo os padrões pedidos no desafio, desde já agradeço a oportunidade e espero fazer
parte desse time para eu poder agregar e aprender muito sobre todos esses assuntos que me despertam muito interesse
e vontade de aprender.

Como executar a API Golang que valida sequências dada a matriz NxN e busca a proporção de sequências válidas em relação ao total:

- É importante verificar se a sua máquina não está rodando algum serviço utilizando as portas do projeto: 8080 e 5432
- Faça o download ou o clone do repositório
- Entre na pasta raiz do projeto e execute o seguinte comando usando o docker-compose

```bash
$ docker-compose up -d
```
Isso fará com que o projeto execute em modo background

rotas: 

    localhost:8080/sequence
        ex req:
            {
                "letters": ["buhdhb", "dbhuhd", "uubuhu", "bhbdhh", "hdhudb", "udbduh"]
            }

    localhost:8080/stats
        ex preview:
            {
                "count_valid": 8,
                "count_invalid": 1,
                "ratio": 0.89
            }


OBS: 
- O projeto foi desenvolvido em uma modelagem de software que permite escalar com facilidade.
- Levando em consideração a entrada de exemplo do pdf do desafio, a rota de validar a matriz trata os dados 
e ignora possíveis espaços em branco nas strings passadas.
- Os testes foram realizados na camada de aplicação testando o usecase e todas as validações feitas com a matriz passada
- Nível 1 - A API foi construída em Golang utilizando o GIN para servir as rotas.
- Nível 2 - O banco escolhido para salvar os dados foi o postgres.
- Nível 3 - O projeto está orquestrado com docker compose para ser executado em qualquer ambiente com o docker instalado,
a api está rodando em um container separadamente do banco e os dois estão na mesma rede.
- Nível 4 - Foi desenhado um modelo de solução para que o projeto tenha a capacidade de atender um grande fluxo de 
requisições onde o app serve como um intermediador para popular o redis com as requisições que serão consumidas pelos workers.
