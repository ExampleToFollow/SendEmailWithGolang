# 📧 Sistema de Envío de Correos Electrónicos en Go

Un sistema simple y eficiente para envío de correos electrónicos utilizando Go y el protocolo SMTP.

## 📁 Estructura del Archivo

### `main.go`

El archivo principal contiene toda la lógica del programa organizanda en las siguientes secciones:

```go
package main

import (
    "fmt"
    "log"
    "net/smtp"
    "os"
    "github.com/joho/godotenv"
)

func main() {
    // 1. Carga de variables de entorno
    // 2. Configuración SMTP
    // 3. Construcción del mensaje
    // 4. Autenticación
    // 5. Envío del correo
}
```

**Motivo de la estructura:**
- **Simplicidad**: Todo en un archivo para facilitar comprensión y mantenimiento
- **Separación lógica**: Cada sección tiene una responsabilidad específica
- **Flujo secuencial**: El código sigue el proceso natural de envío de email
- **Manejo de errores**: Verificaciones en cada paso crítico

## 📦 Dependencias

### Dependencias Estándar de Go
```go
"fmt"      // Formateo e impresión de salida
"log"      // Logging y manejo de errores fatales
"net/smtp" // Cliente SMTP nativo para envío de correos
"os"       // Acceso a variables de entorno del sistema
```

### Dependencias Externas
```go
"github.com/joho/godotenv" // Carga de archivos .env
```

**Instalación:**
```bash
go mod init awesomeProject
go get github.com/joho/godotenv
```

**¿Por qué estas dependencias?**
- **net/smtp**: Biblioteca nativa de Go, eficiente y confiable para SMTP
- **godotenv**: Estándar de facto para manejo de variables de entorno en Go
- **Mínimas dependencias**: Reduce superficie de ataque y complejidad

## 🌍 Variables de Entorno

### Configuración del archivo `.env`

El sistema utiliza variables de entorno para mantener la configuración sensible separada del código fuente:

```bash
# Configuración de correo electrónico
EMAIL_FROM=tu-email@gmail.com
EMAIL_PASS=tu-contraseña-de-aplicacion
EMAIL_TO=destinatario@ejemplo.com
```

### Implementación en el código

```go
// Cargar archivo .env
err := godotenv.Load()
if err != nil {
    log.Fatal("Error al cargar el archivo .env")
}

// Leer variables
from := os.Getenv("EMAIL_FROM")
password := os.Getenv("EMAIL_PASS")
to := []string{os.Getenv("EMAIL_TO")}
```

**Ventajas de este enfoque:**
- **Seguridad**: Credenciales no están hardcodeadas en el código
- **Flexibilidad**: Diferentes configuraciones por entorno (dev, prod)
- **Portabilidad**: Fácil despliegue en diferentes sistemas
- **Control de versiones**: `.env` no se incluye en Git

### Archivo `.gitignore` requerido
```gitignore
.env
*.log
```

## 🔐 Contraseñas de Aplicación de Gmail

### ¿Qué son las Contraseñas de Aplicación?

Las contraseñas de aplicación son contraseñas de 16 caracteres generadas por Google que permiten a aplicaciones menos seguras acceder a tu cuenta de Gmail sin usar tu contraseña principal.

### Configuración paso a paso

#### 1. Prerrequisitos
- Cuenta de Gmail activa
- Autenticación de 2 factores **ACTIVADA** (obligatorio)

#### 2. Generar Contraseña de Aplicación

1. Ve a [Configuración de Google Account](https://myaccount.google.com/)
2. Navega a **Seguridad**
3. En "Acceder a Google", selecciona **Verificación en 2 pasos**
4. Desplázate hasta **Contraseñas de aplicaciones**
5. Selecciona **Correo** como aplicación
6. Selecciona **Otro** como dispositivo
7. Ingresa "Go SMTP Client" como nombre
8. Copia la contraseña generada (formato: `xxxx xxxx xxxx xxxx`)

#### 3. Configurar en `.env`
```bash
EMAIL_FROM=tu-email@gmail.com
EMAIL_PASS=xxxx xxxx xxxx xxxx  # Contraseña de aplicación (sin espacios)
EMAIL_TO=destinatario@ejemplo.com
```

### Configuración SMTP para Gmail

```go
smtpHost := "smtp.gmail.com"
smtpPort := "587"  // Puerto TLS
```

**Especificaciones técnicas:**
- **Servidor**: smtp.gmail.com
- **Puerto**: 587 (STARTTLS)
- **Seguridad**: TLS obligatorio
- **Autenticación**: PLAIN AUTH

### Autenticación en el código

```go
// Configurar autenticación SMTP
auth := smtp.PlainAuth("", from, password, smtpHost)

// Enviar correo con autenticación
err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
```

## 🚀 Ejecución

### Comando de ejecución
```bash
go run main.go
```

### Salida esperada
```
Correo enviado exitosamente
```

## 🛡️ Consideraciones de Seguridad

### ⚠️ Buenas Prácticas

1. **NUNCA** incluir `.env` en control de versiones
2. **Usar solo** contraseñas de aplicación, nunca tu contraseña personal
3. **Rotar** contraseñas de aplicación periódicamente
4. **Limitar** acceso al archivo `.env` (permisos 600)

### Seguridad del archivo `.env`
```bash
# Establecer permisos restrictivos
chmod 600 .env
```

## 🐛 Solución de Problemas Comunes

### Error: "Authentication failed"
**Causa**: Contraseña incorrecta o autenticación 2FA no activada
**Solución**: 
- Verificar que usas contraseña de aplicación
- Confirmar que 2FA está activado
- Generar nueva contraseña de aplicación

### Error: "Error al cargar el archivo .env"
**Causa**: Archivo `.env` no existe o está mal ubicado
**Solución**: 
- Crear `.env` en la raíz del proyecto
- Verificar formato de variables

### Error: "Connection refused"
**Causa**: Firewall o problemas de red
**Solución**: 
- Verificar conexión a internet
- Comprobar que puerto 587 no esté bloqueado
