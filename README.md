## WebSecScanner

Um scanner de segurança web em Go que verifica headers HTTP e outras configurações de segurança em sites para identificar rapidamente potenciais falhas de configuração de segurança


# Funcionalidades Atuais

- Verifica headers HTTP de segurança comuns:

Server → identifica o servidor e tecnologia usada
X-Powered-By → indica o framework ou linguagem
Content-Security-Policy (CSP) → prevenção contra XSS
X-Frame-Options → proteção contra clickjacking
X-Content-Type-Options → evita MIME sniffing
Referrer-Policy → controle de informações do Referer

Estrutura modular que permite adicionar novas análises facilmente.


## Próximos Recursos Planejados

<ol>
<li>Verificação de cookies seguros (HttpOnly, Secure, SameSite);</li>
<li>Análise de CORS e políticas de compartilhamento de recursos;</li>
<li>Checagem de HSTS e configuração HTTPS;</li>
<li>Validação de TLS/SSL básico;</li>
<li>Alertas automáticos para configurações inseguras;</li>
<li>Frontend para mostrar visualmente as análises</li>
</ol>
