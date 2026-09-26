# oGetProxy

Proxy HTTP ligero escrito en Go para desplegar servicios que necesitan realizar
solicitudes HTTP salientes, por ejemplo en [Render](https://render.com).

> ⚠️ **Estado del proyecto:** prototipo. Antes de usarlo en producción, revisa
> las consideraciones de seguridad y las limitaciones indicadas más abajo.

## 📋 Requisitos

- Go 1.27 o posterior.
- Un entorno con una variable `PAGE_PORT` disponible. Si no se define, el
	servidor escucha en el puerto `8028`.

## 🚀 Ejecución local

```bash
go run .
```

Para usar un puerto específico:

```bash
PAGE_PORT=8080 go run .
```

## 🔌 API

El proxy acepta una solicitud HTTP y usa estas cabeceras para construir la
solicitud de destino:

| Cabecera | Obligatoria | Descripción | Ejemplo |
| --- | --- | --- | --- |
| `URI` | Sí | Host y ruta de destino, sin el esquema. | `google.com` |
| `scheme` | No | Esquema de la solicitud. Por defecto es `https`. | `https` |
| `headers` | No | Objeto JSON cuyos valores son textos y que contiene las cabeceras que se enviarán al destino. | `{"Accept":"text/html"}` |

Ejemplo:

```bash
curl -X GET "http://localhost:8028" \
	-H 'URI: google.com' \
	-H 'scheme: https' \
	-H 'headers: {"Accept":"text/html"}'
```

La solicitud conserva el método y el cuerpo recibidos por el proxy. El destino
se construye con el formato `scheme://URI`.

## ⚠️ Limitaciones y seguridad

- El proyecto funciona actualmente como un proxy abierto: un cliente puede
	solicitar direcciones internas o servicios de terceros mediante `URI`.
- No hay autenticación, autorización ni lista de destinos permitidos.
- La validación del esquema, del host y de las cabeceras es limitada.
- Los errores no tienen un formato HTTP consistente y algunos se devuelven con
	código `200`.
- No se documenta soporte real para `ws`; el código usa `net/http` y crea
	solicitudes HTTP.

No lo expongas públicamente sin añadir controles de acceso, validación de URL,
una lista de hosts permitidos y límites de tamaño y tiempo.

## 🧪 Verificación

```bash
go test ./...
```

Actualmente no hay pruebas automatizadas en el repositorio. Añadir pruebas para
la construcción de la URL, las cabeceras, los códigos de estado y los errores
sería el siguiente paso recomendado.

## 📁 Estructura

```text
.
├── go.mod       # Definición del módulo y versión de Go
├── proxy.go     # Servidor HTTP y lógica del proxy
└── readme.md    # Documentación del proyecto
```

## 📄 Licencia

No se ha definido una licencia para este proyecto.
