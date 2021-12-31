# Test

Guarda y consulta productos en memoria

## Construir la imagen docker

``` shell
docker build --build-arg Dockerfile -t product:1.0 .
```

## Archivo de configuración
Debe crear el archivo **configuration.json**
```json
{
  "port": 8000,
  "database": {
    "name": "beers"
  },
  "api_key_currency_layer": "eaad87921sdfge9ba55ae6e98f2500f9"
}
```

## Archivo de pruebas de API
El archivo **test-api-rest.http** tiene configuradas algunas pruebas del API.

## Consideraciones
- Para la conversion de moneda se esta usando https://currencylayer.com/
- Se esta usando un API KEY gratuito
