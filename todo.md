# TODO 

# NataliaNatalia

Es un proyecto que sirve para cargar datos sobre un objetivo y a la cual se la asocia a scripts o programas para obtener más información.
NataliaNatalia es un automatizador de procesos de recabación de información.

# Estructura del proyecto

## 1.0  Tanga

La tanga es el objetivo, el nombre de la entidad a investigar.

Estructura:

- Nombre
- Comentario

### [x] 1.2 Tanga Fields (Campos de la tanga)

Los campos de la tanga es la información de la tanga. Estos campos pueden ser email, teléfono, nicks, redes sociales. Cualquier dato que se tenga de la tanga para luego usar con los scapps (scripts/apps)

Estos campos se definen así:

- TANGA_ID
- Nombre del campo
- Valor del campo

TODO

Arreglar en la edición que muestre la tanga actual

## 2.0 SCAPP (scripts / Apps)

Son scripts o aplicaciones que serán ejecutadas: nmap, dig, y otros. Son programas de líneas de comandos (cli) que se usan para obtener información.

Estructura:

- Nombre del programa
- Ubicación o Alias (ej: nmap o $HOME/scripts/gettwitter.py)
- withSudo (si la ejecución es con SUDO)
- Comentario (Estas son anotaciones útiles para el usuario)

### 2.1 Paramétros de scapp

SOn los parámetros y flags que tiene los programas para funcionar.

Estructura:

- SCAPP_ID *
- orden
- Identificador: Nombre del parámetro ej: "-sn" ó "-u"
- Valor (?)
- IsFlag: Si el parámetro es un flag entonces no tiene que tener un valor asociado.
- Comentario
- category_id


* SCAPP_ID debería ser un array de IDS y en base al óden de esos parámetros es como se utilizan por los INEX.

## 3.0 Instrucciones

Es la asociación de los params de scap con los datos de una tanga.
Estas instrucciones se ejecutan o mejor dicho sus resultados se muestran en el navegador o en una nueva consola.

Estas instrucciones pueden copiarse para cambiar el valor del dato de la tanga o para hacerlas más complejas

Una Instrucción tiene la lista de parámetros y una lista de campos de tanga Cada uno de estos está asociado por un índice, es decir que son dos arrays

Estructura:

- PARAM_ID
- TANGA_FIELD_ID
- EXECUTEIN: Donde se ejecuta (navegador u consola)

### 3.1 Paquetes de Instrucciones.

Los paquetes es instrucciones son varias instrucciones juntas.

**COMING SOON**
