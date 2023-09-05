# tcc-alonsofritz
Repositorio para organização de TCC
Tema: Programação Concorrente com Go, e um comparativo entre GO, C++ e Java, com problemas classicos de concorrencia e um comparativo entre desempenho na multiplicação de matrizes (?), e/ou espectro de fourier ? transofrmação de fourier?

- Computação Concorrente
	> Processos e Threas, Kernel Theads, User space thread
	> Condicao de corrida
	> Sincronização por condicao logica
	> deadlockse starvations

- Linguagem Go
	> Especificações Gerais da Linguagem, como utiliza-la e etc
		* Packages, tipos de dados, controles de fluxo, interfaces e metodos, recursos para concorrencia, goroutines, escolanamento de goroutines, channels...

- Funcionamento de Threads e Concorrencia em Linguagem C

- Funcionamento de Threads e Concorrencia em Linguagem Java

- Exemplos do Uso de Recursos de concorrencia do Go implementados em problemas clássicos de concorrencia, fazendo um paralelo a sua implementação utiliando C++ e Java
	> Produtor Consumidor
	> Leitores e Escritores
	> Barbeiro Dorminhoco
	> Programação Distribuida
		* Aplicação Cliente/Servidor
		* Chat

- Avaliação de Desempenho em Go utilizando multiplicação de matrizes ou transformação de fourier(?)
	> Codigo em Go
	> comparação com C++ e Java
	> otimizações de acesso a cache
	> resultados dos timers
 
## Mais interessantes:
- https://journal.universitasbumigora.ac.id/index.php/matrik/article/view/1325/1154
- https://www.diva-portal.org/smash/get/diva2:1220262/FULLTEXT01.pdf
- https://www.sjsu.edu/people/robert.chun/courses/cs159/s3/S.pdf
- https://www.diva-portal.org/smash/get/diva2:824741/FULLTEXT03.pdf
- https://www.csc.kth.se/utbildning/kth/kurser/DD143X/dkand11/Group9Alexander/Joakim_Anneback_Johan_Stjernberg.finalreport.2.pdf
- https://www.csc.kth.se/utbildning/kandidatexjobb/datateknik/2011/rapport/anneback_joakim_OCH_stjernberg_johan_K11073.pdf
- https://stanford-cs242.github.io/f17/assets/projects/2017/gsfisher-chrisyeh.pdf

## Referencias
- https://dl.acm.org/doi/abs/10.1145/3297858.3304069
- https://ieeexplore.ieee.org/abstract/document/6920368
- https://ieeexplore.ieee.org/abstract/document/10145271
- 
- https://www.diva-portal.org/smash/get/diva2:1701961/FULLTEXT01.pdf
- https://books.google.com.br/books?hl=pt-BR&lr=&id=dulODwAAQBAJ&oi=fnd&pg=PP1&dq=golang+concurrency&ots=3x1Cnnnx4p&sig=LfR9lOJHrL_VsGU-7sF9wIfBKpQ#v=onepage&q=golang%20concurrency&f=false
- 
- https://www.diva-portal.org/smash/get/diva2:1326869/FULLTEXT01.pdf
- https://www.scaler.com/topics/golang/concurrency-in-golang/
- 
- https://uwspace.uwaterloo.ca/handle/10012/14706
- https://stanford-cs242.github.io/f17/assets/projects/2017/gsfisher-chrisyeh.pdf
- https://www.ingentaconnect.com/content/ben/rascs/2021/00000014/00000006/art00027
- https://medium.com/@isuruvihan/stateful-goroutines-efficient-concurrency-with-go-e6e0736b2342

## Notes
Go defende tornar a criação de threads fácil e leve e usar a passagem de mensagens pela memória compartilhada para comunicação entre threads. Na verdade, vimos mais goroutines criadas em programas Go do que threads tradicionais e há usos significativos do canal Go e de outros mecanismos de passagem de mensagens. No entanto, nosso estudo mostra que, se não forem usadas corretamente, essas duas práticas de programação podem causar erros de simultaneidade. Memória compartilhada versus passagem de mensagens. Nosso estudo descobriu que a passagem de mensagens não torna necessariamente os programas multithread menos propensos a erros do que a memória compartilhada. Na verdade, a passagem de mensagens é a principal causa do bloqueio de bugs. Para piorar a situação, quando combinada com primitivas de sincronização tradicionais ou com outros novos recursos e bibliotecas de linguagem, a passagem de mensagens pode causar bugs de bloqueio que são muito difíceis de detectar. A passagem de mensagens causa menos bugs sem bloqueio do que a sincronização de memória compartilhada e, surpreendentemente, foi usada até para corrigir bugs causados por sincronização incorreta de memória compartilhada. Acreditamos que a passagem de mensagens oferece uma forma limpa de comunicação entre threads e pode ser útil na passagem de dados e sinais. Mas eles só são úteis se usados ​​corretamente, o que exige que os programadores não apenas entendam bem os mecanismos de passagem de mensagens, mas também outros mecanismos de sincronização do Go. Implicação na detecção de bugs. Nosso estudo revela muitos padrões de código com bugs que podem ser aproveitados para conduzir a detecção de bugs de simultaneidade. Como esforço preliminar, construímos um detector direcionado a bugs não bloqueadores causados por funções anônimas (por exemplo, Figura 8). Nosso detector já descobriu alguns novos bugs, um dos quais foi confirmado por desenvolvedores de aplicativos reais [12].
De forma mais geral, acreditamos que a análise estática mais os algoritmos de detecção de deadlock anteriores ainda serão úteis na detecção da maioria dos bugs de bloqueio de Go causados por erros na sincronização de memória compartilhada. As tecnologias estáticas também podem ajudar na detecção de bugs causados pela combinação de canal e bloqueios, como o da Figura 7. O uso indevido de bibliotecas Go pode causar bugs bloqueadores e não bloqueadores. Resumimos vários padrões sobre o uso indevido de bibliotecas Go em nosso estudo. Os detectores podem aproveitar os padrões que aprendemos para revelar bugs anteriormente desconhecidos. Nosso estudo também descobriu que a violação das regras que Go impõe com suas primitivas de simultaneidade é um dos principais motivos para bugs de simultaneidade. Uma nova técnica dinâmica pode tentar impor tais regras e detectar violações em tempo de execução.

Estudando Bugs do Mundo Real. Existem muitos estudos empíricos sobre bugs do mundo real [9, 24, 25, 29, 40, 44, 45]. Esses estudos orientaram com sucesso o projeto de várias técnicas de combate a bugs. Até onde sabemos, nosso trabalho é o primeiro estudo focado em bugs de simultaneidade em Go e o primeiro a comparar bugs causados por erros ao acessar memória compartilhada e erros ao passar mensagens.

Combatendo Bugs de Bloqueio. Como um problema tradicional, existem muitos trabalhos de pesquisa combatendo impasses em C e Java [7, 28, 33–35, 51, 54, 55, 58, 59]. Embora útil, nosso estudo mostra que existem muitos bugs de bloqueio sem deadlock no Go, que não são o objetivo dessas técnicas. Algumas técnicas são propostas para detectar bugs de bloqueio causados pelo uso indevido do canal [38, 39, 49, 56]. No entanto, erros de bloqueio podem ser causados por outras primitivas. Nosso estudo revela muitos padrões de código para bloqueio de bugs que podem servir de base para futuras técnicas de bloqueio de detecção de bugs.

Combatendo Bugs Não Bloqueadores. Muitos trabalhos de pesquisa anteriores são conduzidos para detectar, diagnosticar e corrigir bugs não relacionados a deadlock, causados pela falha na sincronização de acessos à memória compartilhada [4, 5, 8, 14, 16, 17, 30 32, 43, 46, 47, 52, 62 –64]. Eles prometem ser aplicados a bugs de simultaneidade do Go. No entanto, nosso estudo descobriu que há uma parcela não negligenciável de bugs não bloqueadores causados por erros durante a passagem de mensagens, e esses bugs não são cobertos por trabalhos anteriores. Nosso estudo enfatiza a necessidade de novas técnicas para combater erros durante a passagem de mensagens
