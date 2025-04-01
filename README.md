# FFI (Foreign Function Interface)

O Foreign Function Interface (FFI) é um mecanismo que permite que programas escritos em uma linguagem chamem funções ou utilizem bibliotecas escritas em outra linguagem. Isso pode ser útil quando se deseja integrar código nativo (como C, C++, Rust ou Go) com linguagens de mais alto nível, como Python, Java e C#.

Com FFI podemos reaproveitar bibliotecas existentes e facilitar a interoperabilidade entre diferentes tecnologias. No entanto, ele pode introduzir desafios como gerenciamento de memória, compatibilidade de tipos e diferenças entre convenções de chamada das linguagens envolvidas.

Para saber mais sobre FFI, [clique aqui]()

# Requisitos 

Testar se o conceito de interoperabilidade atende os seguintes requisitos:

- Ter apenas um unico codebase que seja reutiliado em N linguagens.
- Agnostico (um code base que se integra com AWS ou Azure)
- Baixa latencia.
- Facilitar o reuso.
- Performatico.
- Simples utilização

# Proposta de solução

![proposta](./doc/ffi.drawio.png)




  

