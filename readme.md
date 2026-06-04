# CLI TASK MANAGER

Este es mi primer proyecto de go, usa librerías estándar como encoding/json y fmt. 

## Creación del ejecutable
Generacion del ejecutable
```bash
./script.sh
```
* Ejecutable generado en /build/task-cli

## Comandos básicos del programa
| Command | Descript |
|---------|----------|
| task-cli add <name> | creando una tarea |
| task-cli list | listando todas las tareas |
| task-cli help | mostrar información |
| task-cli remove <id> | eliminar tarea por id |
| task-cli done <id> |marca la tarea como realizada |

## Ejecutando desde go

```bash
go run . <args...>
```
