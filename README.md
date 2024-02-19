# JWT

JWT（JSON Web Token）是一种开放的行业标准 RFC 7519 方法，用于在双方之间安全地表示请求。

# 认证和授权

认证（Authentication）：You are who you say you are。

授权（Authorization）: You have permission to do what you are requesting to do。

# OPA
OPA（Open Policy Agent）：OPA允许你定义和执行策略。在授权场景中，OPA可以用来决定一个已经通过认证的用户是否有权限执行特定的操作。OPA可以接收用户的身份信息（如从JWT中获取），以及请求的上下文信息，然后根据预定义的策略来决定是否授权。

# 引用

- [Serivce](https://github.com/ardanlabs/service)
