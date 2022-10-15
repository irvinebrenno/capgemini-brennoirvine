# capgemini-brennoirvine

Esse é o meu projeto seguindo os padrões pedidos no desafio, desde já agradeço a oportunidade e espero fazer
parte desse time para que eu possa agregar e aprender muito sobre todos esses assuntos que me despertam muito interesse
e vontade de aprender.

Como executar a API Golang que busca sequencias dada a matriz NxN

```bash
$ docker-compose up -d
```

rotas: 

    localhost:8080/sequence
        payload:
            {
                "letters": ["buhdhb", "dbhuhd", "uubuhu", "bhbdhh", "hdhudb", "udbduh"]
            }

    localhost:8080/stats
        response:
            {
                "count_valid": 8,
                "count_invalid": 1,
                "ratio": 0.89
            }
