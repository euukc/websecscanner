## WebSecScanner

Um scanner de segurança web em Go que verifica headers HTTP e outras configurações de segurança em sites para identificar rapidamente potenciais falhas de configuração de segurança


# Funcionalidades Atuais

- Verifica headers HTTP de segurança comuns:

-- Server → identifica o servidor e tecnologia usada

-- X-Powered-By → indica o framework ou linguagem

-- Content-Security-Policy (CSP) → prevenção contra XSS

-- X-Frame-Options → proteção contra clickjacking

-- X-Content-Type-Options → evita MIME sniffing

-- Referrer-Policy → controle de informações do Referer

Estrutura modular que permite adicionar novas análises facilmente.


## Próximos Recursos Planejados

Verificação de cookies seguros (HttpOnly, Secure, SameSite);
Análise de CORS e políticas de compartilhamento de recursos;
Checagem de HSTS e configuração HTTPS;
Validação de TLS/SSL básico;
Alertas automáticos para configurações inseguras;
Frontend para mostrar visualmente as análises
