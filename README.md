# Tango

Tango es un SDK que incluye Go + Templ + Htmlx
Ideal para crear APIs

## Diseño RFM (Routes Features Models)

### Routes

Las rutas se definen en la carpeta routes dentro del archivo setuproutes allí se agregan las funciones. Para poder acceder a las funcionalidades de Fiber o Gorm o la configuración de la app se debe pasar una variable

    tapp *webcore.TangoApp

### Features

Las "features" son los controladores de las rutas y reciben 2 parámetros: 1) el contexto de Fiber y otra variable de *webcore.TangoApp para poder acceder a la db y configuración.

### Models

Los modelos deben recibir también la variable que apunte a *webcore.TangoApp

## Public

los archivos publicos están dentro de la carpeta public y por defecto tienen varios accesos. Esto se puede ver en webcore_features/routes.go



# Generar un build

TODO

# Auth

TODO

Librería de autenticación