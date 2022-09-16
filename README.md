# ProyectoMercadoLibre


Proyecto Prueba de API REST para indicar si es un mutante o no dependiendo de su patron genetico


# Como Iniciar

Windows:

Una vez instalado GO en su computadora, inicie con el comando:

go mod download

Para asegurar que las dependencias esten listas, despues usando el comando:

go run .

Iniciara el servidor del API y podra realizar peticiones a travez de:
localhost:3000

# Rutas existentes:

- localhost:3000 GET 


- localhost:3000 POST - en el body tipo "aplication/json" el codigo genetico a evaluar


# Testeo:

El reporte de testing de go se encuentra dentro del archivo:
coverage.out

Para poder saber que porcentaje del programa esta testeado ejecute el comando con el API ejecutandose:

go test --cover

Para poder vizualizar de forma comoda y como reporte en una pestaña del navegador web, solo ejecute el comando: 

go  tool cover --html=coverage.out


# Accesivilidad online:

El proyecto se encuentra actualmente motando en AWS con EC2 En el DNS publico:

ec2-3-87-52-4.compute-1.amazonaws.com

