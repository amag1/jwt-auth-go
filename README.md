# ES - Servicio de Autenticación y Autorización de Usuarios
## Introducción

Este proyecto proporciona un servicio de autenticación y autorización utilizando JWT (JSON Web Token). El servicio permite a los usuarios registrarse, iniciar sesión, refrescar los tokens de sesión y cerrar sesión de forma segura. Un aspecto clave de este sistema es el uso de tokens de sesión (access tokens) y tokens de refresco, garantizando una autenticación segura y escalable para aplicaciones web modernas.

## Importancia de la Autenticación y la Autorización

La autenticación y la autorización son fundamentales para mantener la seguridad y privacidad de los datos de los usuarios. La autenticación garantiza que el usuario es quien dice ser, mientras que la autorización determina qué acciones puede realizar el usuario autenticado. En cualquier aplicación que maneje datos sensibles o información personal, es crucial tener mecanismos robustos de autenticación y autorización.
Cómo Funciona el Mecanismo de Tokens de Sesión y de Refresco con JWT

Este proyecto utiliza JWT (JSON Web Token) para gestionar las sesiones de usuario. JWT es una forma compacta y autónoma de transmitir información de manera segura entre partes. Se utiliza ampliamente para la autenticación sin estado (stateless authentication).

1. Inicio de Sesión (Login): Cuando un usuario inicia sesión, se le proporcionan dos tokens:
    - Token de Acceso (Access Token): Un token de corta duración (ej., 15 minutos) que se usa para autenticar al usuario y permitir el acceso a recursos protegidos.
    - Token de Refresco (Refresh Token): Un token de larga duración (ej., 7 días) que se usa para obtener nuevos tokens de acceso sin necesidad de iniciar sesión nuevamente.

2. Refresco de Token: Cuando el token de acceso caduca, el usuario puede enviar el token de refresco al servidor para obtener un nuevo token de acceso. Esto asegura que el usuario permanezca autenticado sin necesidad de iniciar sesión repetidamente.

3. Cerrar Sesión (Logout): El token de refresco se invalida al cerrar sesión, evitando su uso para generar nuevos tokens de acceso.

## Endpoints Disponibles

Todas las rutas disponibles para la autenticación de usuarios están bajo el endpoint `/api/user`. A continuación, se detallan las rutas disponibles:

- POST /api/user/login
    - Inicia sesión de un usuario utilizando su correo electrónico y contraseña.
    - Devuelve un token de acceso y un token de refresco.

- POST /api/user/register
    - Registra un nuevo usuario.
    - Requiere detalles del usuario como el correo electrónico y la contraseña, pero no lo inicia sesión automáticamente.

- POST /api/user/refresh
    - Toma un token de refresco válido y devuelve un nuevo token de acceso, extendiendo la sesión del usuario.

POST /api/user/logout
    - Invalida el token de refresco, cerrando la sesión actual del usuario.

GET /api/user/protected
    - Una ruta protegida de ejemplo, accesible solo con un token de acceso válido.
    - Demuestra el uso de middleware de autenticación para restringir el acceso solo a usuarios autorizados.
        
# EN -User Authentication and Authorization Service
## Introduction

This project provides an authentication and authorization service using JWT (JSON Web Token). The service allows users to register, log in, refresh session tokens, and log out securely. A key aspect of this system is the use of session tokens (access tokens) and refresh tokens, ensuring secure and scalable authentication for modern web applications.
Why Authentication and Authorization Are Important

Authentication and authorization are essential for maintaining the security and privacy of user data. Authentication ensures that the user is who they claim to be, while authorization determines what the authenticated user is allowed to do. In any application that handles sensitive data or personal information, robust authentication and authorization mechanisms are critical.

## How the JWT Session and Refresh Token Mechanism Works

This project uses JWT (JSON Web Token) to manage user sessions. JWT is a compact, self-contained way of securely transmitting information between parties. It’s widely used for stateless authentication.

1. Login: When a user logs in, they are provided with two tokens:
    - Access Token: A short-lived token (e.g., 15 minutes) used to authenticate the user and allow access to protected resources.
    - Refresh Token: A long-lived token (e.g., 7 days) used to obtain new access tokens without requiring the user to log in again.

2. Token Refresh: When the access token expires, the user can send the refresh token to the server to obtain a new access token. This ensures the user remains authenticated without needing to log in repeatedly.

3. Logout: The refresh token is invalidated upon logout, preventing further use of that token to generate new access tokens.

## Available Endpoints

All the available routes for user authentication are under the `/api/user` endpoint. Below are the available routes:

- POST /api/user/login
    - Logs in a user using their email and password.
    - Returns an access token and refresh token.

- POST /api/user/register
    - Registers a new user.
    - Requires user details such as email and password but does not log them in automatically.

- POST /api/user/refresh
    - Takes a valid refresh token and returns a new access token, extending the user’s session.

- POST /api/user/logout
    - Invalidates the refresh token, logging the user out of the current session.

- GET /api/user/protected
    - A sample protected route, only accessible with a valid access token.
    - Demonstrates the use of authentication middleware for restricting access to authorized users only.

