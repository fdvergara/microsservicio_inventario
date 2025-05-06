# Microservicio de Inventario

Este microservicio es responsable de la gestión de inventario para una cadena de restaurantes. Permite validar la disponibilidad de ingredientes, actualizar existencias, y manejar recetas.

---

## 🧱 Arquitectura

Este servicio está construido con **Go (Golang)** y sigue el patrón de **Arquitectura Hexagonal**, separando claramente las siguientes capas:

- **Dominio**: Entidades de negocio como Ingrediente y Receta y los servicios relacionados con la logica de negocio.
- **Aplicación**: Controladores HTTP para los casos de uso.
- **Infraestructura**: Implementaciones de repositorios y próximamente mensajería con Kafka.

Igualmente se está utilizando la base de datos con MongoDB.

## 🐇 Configuración de RabbitMQ

Este proyecto utiliza **RabbitMQ** como sistema de mensajería para comunicación asíncrona basada en eventos. Contiene un consumidor que recibe las novedades para actualizar el inventario con los siguientes parametros:

```text
{
"ingrediente_id":abf63e1a-bbae-45bc-ac2a-7fb2a3f36ca5,
"cantidad":10
}
```

### 🔧 Parámetros de conexión

- **Host:** `localhost`
- **Puerto AMQP:** `5672`
- **Puerto UI (RabbitMQ Management):** `15672`
- **Usuario:** `guest`
- **Contraseña:** `guest`

### 🧩 Configuración del exchange

Se utiliza un exchange tipo `topic` llamado:

```text
Topic: inventario_topic
Routing Key: inventario.#
```

---

## ▶️ Ejecucióndel Proyecto

Para ejecutar este proyecto ejecute el comando:
**docker-compose up -d**
Dentro de la carpeta del proyecto. Aqui se le disponibilizará la aplicación y su base de datos, asegurese que cuente con el directorio **./data**, donde persistirá la información.

## 📄 Documentacion del Proyecto

En el directorio **docs/** del proyecto se encuentra la documentación swagger del microservicio.

```

```
