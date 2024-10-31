# Permissions - GoLang
## Standards
- Project Structure: https://github.com/golang-standards/project-layout

## Run locally
- go run ./src/permissions/cmd/server/main.go

## Required Permissions
- canViewPermissionGroups
- canEditPermissoinGroups
- canViewPermissions
- canEditPermissions

## Environmnet Variables
- `JWT_SIGNING_METHOD`
    - Controls which signing method is used
    - Available Values
        - `ECDSA384` ECDSA 384 (default)
    - Planned Values:
        - `HS256` HMAC-SHA 256
        - `HS256` HMAC-SHA 256
        - `HS256` HMAC-SHA 256
        - `RS256` RSA 256
        
- `SESSION_EXPIRATION`
    - Controls how long a session persists in Minutes. 
    - Users are forced to log in again after this expires
    - Values: 10 to 10080 (7 days)

- `JWT_PRIVATE_KEY`
    - Used to sign session keys
    - By default, uses:
```
-----BEGIN EC PRIVATE KEY-----
MIGkAgEBBDC4czoxahGqOAy2eCbsNjyEfFCsRItQ+G00whfrCbJQfsEDFN3HiSO5InXH8ZqjfmGgBwYFK4EEACKhZANiAATrXPwqQbsF+yKhRyYwxNNtnSEdHyTMhcg9hymgueps48dc9Izg9gKwtuFpPO7DSwBIMxx1IRmrAXDeSudfAcoSncgPmiXa+PiqnEPNl2XhPR029Z5EwIYtkYA9XPrM4Pg=
-----END EC PRIVATE KEY-----
```
- `JWT_PUBLIC_KEY`
    - used to verify session key signatures
    - If set to "" or not added, it will generate a new public key based on the private key
    - The default public key should be 
```
-----BEGIN PUBLIC KEY-----
MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAE61z8KkG7BfsioUcmMMTTbZ0hHR8kzIXI
PYcpoLnqbOPHXPSM4PYCsLbhaTzuw0sASDMcdSEZqwFw3krnXwHKEp3ID5ol2vj4
qpxDzZdl4T0dNvWeRMCGLZGAPVz6zOD4
-----END PUBLIC KEY-----
``` 
- `SESSION_KEY_HEADER`
    - The name of the header containing the session key
    - default to `X-API-Key`

### THIS SHOULD BE FOR TESTING ONLY!!!
- failure to set this value will enable unverified users to spoof as super admins