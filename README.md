# Flix360 Core API - Backend V1 🚀

Backend empresarial para el sistema de gestión de inventarios por RFID (Flix360), diseñado para soportar lecturas masivas de hardware en campo (Pistolas Android Zebra/Chainway) y operar en entornos con intermitencia de red.

## 🏗 Arquitectura
El proyecto está construido en **Go (Golang)** utilizando los principios de **Clean Architecture** y **Hexagonal Architecture**. 
Esto garantiza que la lógica de negocio esté completamente aislada de la base de datos y del framework HTTP.

*   **Framework HTTP:** Gin-Gonic
*   **Base de Datos:** PostgreSQL
*   **Gestor de BD:** `sqlx` (con queries preparadas para inyección segura de structs)
*   **Seguridad:** JSON Web Tokens (JWT)

### Estructura de Carpetas
*   `cmd/api/`: Capa de inicialización (Servidor, Router, Inyección de Dependencias).
*   `internal/core/domain/`: Entidades de negocio y estructuras (Structs puros).
*   `internal/core/ports/`: Contratos/Interfaces (Inbound para Servicios, Outbound para Repositorios).
*   `internal/core/services/`: Lógica de negocio (Casos de uso).
*   `internal/adapters/repositories/postgres/`: Implementación física de la BD.
*   `internal/adapters/handlers/http/`: Controladores Gin y Middlewares.
*   `internal/workers/`: Procesos en segundo plano (Goroutines).

---

## ✨ Características Principales (Features)

1.  **Inventario RFID Inmutable:** 
    *   Uso de `inventory_assets` para vincular un Producto físico, una Etiqueta RFID y una Locación.
    *   Trazabilidad garantizada mediante `inventory_movements` (Historial de auditoría tipo "append-only").
2.  **Catálogo Dinámico:**
    *   Categorías y Definición de Atributos (`attribute_definitions`).
    *   Valores dinámicos JSONB por producto, evitando alterar esquemas de BD por cada cliente.
3.  **Seguridad Multi-Tenant:**
    *   Middleware JWT que inyecta automáticamente el `company_id` y `user_id` en el contexto de cada petición.
    *   Evita fugas de datos entre empresas.
4.  **Soporte Masivo (Bulk):**
    *   Endpoints optimizados (ej. `POST /inventory-movements/bulk`) para procesar cientos de lecturas de la pistola RFID en una sola transacción de milisegundos.
5.  **Motor Offline (Background Worker):**
    *   Uso del patrón *Outbox/Sync Queue*. 
    *   La API recibe payloads en `POST /sync-queue` y responde inmediatamente. Un **Worker** en Goroutine procesa la cola cada 10 segundos, asegurando tolerancia a fallos de red.

---

## 🛣 Mapa de Rutas (Endpoints)

Todas las rutas operan bajo el prefijo `/api/v1`.

### 🔓 Rutas Públicas
| Método | Endpoint | Descripción |
| :--- | :--- | :--- |
| `POST` | `/login` | Autenticación de usuario. Retorna Token JWT. |
| `POST` | `/users` | Creación de usuarios (Admin). |

### 🔒 Rutas Privadas (Requieren Header `Authorization: Bearer <token>`)

#### Catálogo y Productos
| Método | Endpoint | Descripción |
| :--- | :--- | :--- |
| `GET` | `/products` | Lista productos con paginación (`?limit=20&offset=0&search=X`). |
| `POST` | `/products` | Crea un nuevo producto base. |
| `GET` | `/products/:id` | Detalle de un producto específico. |

#### Categorías y Atributos Dinámicos
| Método | Endpoint | Descripción |
| :--- | :--- | :--- |
| `POST` | `/categories` | Crea categoría (ej. Ropa Deportiva). |
| `GET` | `/categories/:id` | Detalle de categoría. |
| `POST` | `/attribute-definitions` | Define una regla (ej. "Talla" obligatoria). |
| `GET` | `/attribute-definitions/:id` | Detalle de regla de atributo. |
| `POST` | `/product-attribute-values` | Asigna un valor a un producto (ej. "Talla M"). |
| `GET` | `/product-attribute-values/:id` | Obtiene el valor del atributo. |

#### Físico y Hardware
| Método | Endpoint | Descripción |
| :--- | :--- | :--- |
| `POST` | `/rfid-tags` | Registra chips EPC en blanco. |
| `GET` | `/rfid-tags/:epc` | Busca un tag por su código EPC. |
| `POST` | `/locations` | Registra bodegas/tiendas. |
| `GET` | `/locations/:id` | Detalle de locación. |
| `POST` | `/devices` | Registra pistolas RFID (Hardware). |
| `GET` | `/devices/:id` | Detalle del dispositivo. |

#### Operación Inventario y Auditoría
| Método | Endpoint | Descripción |
| :--- | :--- | :--- |
| `POST` | `/inventory-assets` | Une Producto + RFID + Locación (El Activo Físico). |
| `GET` | `/inventory-assets` | Lista activos (`?location_id=XYZ&limit=50`). Para cruces de auditoría. |
| `GET` | `/inventory-assets/:id` | Detalle de un activo. |
| `POST` | `/inventory-movements/bulk` | **[CRÍTICO]** Recibe array de movimientos masivos de la pistola. |
| `POST` | `/inventory-movements` | Movimiento individual (Entrada/Salida/Transferencia). |
| `GET` | `/inventory-movements/:id` | Detalle de movimiento. |

#### Sincronización (Modo Offline)
| Método | Endpoint | Descripción |
| :--- | :--- | :--- |
| `POST` | `/sync-queue` | Recibe JSON genérico desde SQLite (Android) para procesar asíncronamente. |

---
*Desarrollado para la gestión inteligente de inventarios físicos.*