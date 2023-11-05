# NataliaNatalia
<div style="text-align: center;">    
    <img src="public/icon-natalianatalia.png" alt="NataliaNatalia" width="128px"><br/>
    OSINT platform<br/>
    _README-SPANISH_
    <br/>
    <br/>
</div>

# ¿Qué es NataliaNatalia?

NataliaNatilia es una plataforma que permite la creación de perfiles para investigar con herramientas abiertas y aplicaciones quese utilicen para esos fines y almacene toda la información en archivos relacionados con esos perfiles.

* El fuerte de la plataforma será poder utilizar herramientas externas (algo así como los plugins de Maltego) y poder volcar esa información en un mismo lugar.

* El otro fuerte será tener todo almacenado en archivos (o en la base de datos) porque sabemos que todos los datos en internet son volátiles. Hoy están, mañana puede que no.

* Una plataforma pensada principalmente en español.

* Usos limpios. Para poder separar las investigaciones lo ideal será ejecutar un programa cada vez que se quiera trabajar con él.

# Instalación: (revisión)

    2. Ejecutar la aplicación y dirigirse a http://localhost:3007/api/setup
    

## Requisitos:

    - Python ([https://www.python.org])

# Ejecutar

    1. Ejecutar el archivo init.bat en windows o init.sh en linux
        // Esto ejecutar el _docker compose up_ para levantar la base de datos y la app
    
    2. Abrir el navegador en http://localhost:3007
    3. Al iniciar debe crearse una Tanga (Target)
    4. Luego se asocian metadatos con las tangas. Por defecto se incluyen los metadatos básicos, pero són opcionales. Se pueden agregar una cantidad ilimitada de metadatos para asociar.
    5. Agregar los scripts que tengas instalados en tu pc. Por el momento esto debe hacerse por fuera de esta aplicación.
    6. Una vez agregado un script natalianatalia ejecuta una consulta de ayuda para obtener toda la ayuda de la aplicación y de esa forma a) saber que está disponible b) conocer cuales son los parámetros para brindar ayuda al usuario.
    7. Ahora el script está disponible para asociar con los metadatos.
       1. Por el momento se pueden seleccionar dos modos AAO (All At Once) o 1B1 (One By One) en la que se ejecuta una sola petición al script con todos los parámetros juntos o que se le aplique 1 por uno en ejecuciones diferentes.
    8. 


# Todos:
    1. Crear la estructura de la bd del proyecto.
    2. Crear los CRUDS de estas tablas. (utilizar algún orm?)
    3. Scripts: Asociar scripts para ser ejecutados.
       1. Lo primero que debe hacer la app es reconocer el script: Realizando un NOMBRESCRIPT.PY --h
          1. De esta forma se sabe que parámetros acepta.
          2. poder asociar esos parámetros con los metadados del perfil.
       2. 
    
    Front-end
    4. Crear el frontend
    5. Exploración de perfiles.

# Tecnologías:
Rust,SQLITE, HTML, CSS, JAVASCRIPT, VUE framework