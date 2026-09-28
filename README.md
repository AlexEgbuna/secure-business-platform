# Secure Business Platform

A production-oriented business application designed to demonstrate modern backend engineering, platform security, identity management, 
API security, and secure infrastructure practices.

The platform is being developed alongside GoTLS Proxy to provide a realistic environment for implementing enterprise-grade authentication, 
authorization, API protection, and edge security.

---

## Project Goals

This project is designed to demonstrate practical implementation of:

- Secure authentication systems
- Identity and access management
- REST API development
- PostgreSQL integration
- Secure session management
- OWASP Top 10 protections
- OWASP API Security Top 10 protections
- Advanced HTTP security
- Audit logging
- Infrastructure security
- Secure deployment practices
- Cloud-ready architecture
- Stateful security integration with GoTLS Proxy
- Security telemetry and observability
- SIEM and detection engineering
- Vulnerability management
- DevSecOps security controls
- Threat modeling and security testing
- Incident response engineering

---

## Technology Stack

### Backend

- Go
- PostgreSQL
- Redis
- HTML Templates

### Security

- OAuth 2.0
- OpenID Connect (OIDC)
- JWT
- RBAC
- Secure Session Management
- OWASP Top 10
- OWASP API Security Top 10

### Infrastructure

- Linux
- Systemd
- Docker (Planned)
- AWS (Planned)
- CI/CD (Planned)

### Observability

- Prometheus (planned)
- Grafana (planned)
- Loki (planned)
- Alertmanager (planned)

### SIEM / Detection

- Wazuh (planned)
- Security event ingestion (planned)
- Detection engineering (planned)

---

## Project Structure

```text
secure-business-platform/

├── cmd/
│   └── server/

├── internal/
│   ├── auth/
│   ├── users/
│   ├── organizations/
│   ├── roles/
│   ├── dashboard/
│   ├── audit/
│   ├── middleware/
│   ├── database/
│   ├── config/
│   ├── services/
│   ├── api/
│   ├── security/
│   └── proxy/

├── web/
│   ├── templates/
│   ├── static/
│   ├── css/
│   ├── js/
│   └── images/

├── migrations/
├── scripts/
├── deployments/

├── docs/
│   ├── architecture/
│   ├── deployment/
│   ├── reports/
│   ├── security/
│   ├── testing/
│   ├── threat-models/
│   ├── detection/
│   ├── runbooks/
│   └── roadmap.md

├── tests/

├── .env.example
├── config.example.json
├── .gitignore
├── README.md
├── go.mod
└── go.sum
```

---

## Development Status

Current Stage:

**Stage 1 — Business Application Foundation**

Status:

🚧 In Active Development

Current Focus:

Application router implementation completed. Database layer implementation begins next.

---

## Planned Features

The implementation roadmap includes:

- Authentication
- User Management
- Organization Management
- Role Management
- Dashboard
- REST APIs
- PostgreSQL Integration
- Audit Logging
- Identity & Access Management
- API Security
- Edge Security Integration with GoTLS Proxy
- Detection & Response
- Cloud Deployment

---

## Integration

This application is designed to integrate with the GoTLS Proxy project.

The Business Platform is designed to operate behind GoTLS Proxy.

The Business Platform provides authenticated business/application functionality, while GoTLS provides the Internet-facing edge security layer, including:

- TLS termination
- traffic routing
- proxy enforcement
- trust scoring
- browser/device verification
- request inspection
- advanced traffic filtering

The two systems will progressively integrate through:

- identity context
- application context
- shared security state
- security telemetry
- controlled service-to-service communication
- security decisions

The Business Platform does not duplicate GoTLS edge-security functionality.

Internal TLS/re-encryption will be introduced when the deployment architecture reaches the appropriate traffic-routing and service-integration stage.

Together, the systems form a security-focused application platform while maintaining clear separation of responsibilities between application security and edge/network/stateful security.

---

## Documentation

Engineering documentation is maintained under:

```text
docs/
```

including:

- Architecture
- Deployment
- Security
- Testing
- Engineering Reports

---

## Author

**Alexander Egbuna**

Security-focused Backend & Platform Engineer

Current areas of focus:

- Go Backend Engineering
- Platform Security
- Linux Infrastructure
- Reverse Proxy Architecture
- Identity & Access Management
- API Security
- Cloud & DevSecOps

---

## License

Copyright © 2026 Alexander Egbuna.

This repository is provided for portfolio and evaluation purposes.

Permission is granted to view, clone, and execute the project for evaluation and educational purposes only.

No permission is granted to modify, redistribute, commercialize, or incorporate this code into other projects without prior written authorization from the author.
