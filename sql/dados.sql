insert into usuarios (nome, nick, email, senha)
values
("Usuário 1", "usuario_1", "usuario1@gmail.com", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjU4OTE1MjEsInVzdWFyaW9JZCI6MTR9.BGYa-1kDs64OoY5U8OvERIB1Qxdwqed3kI5uLMG9ynY"), -- usuario1
("Usuário 2", "usuario_2", "usuario2@gmail.com", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjU4OTE1MjEsInVzdWFyaW9JZCI6MTR9.BGYa-1kDs64OoY5U8OvERIB1Qxdwqed3kI5uLMG9ynY"), -- usuario2
("Usuário 3", "usuario_3", "usuario3@gmail.com", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjU4OTE1MjEsInVzdWFyaW9JZCI6MTR9.BGYa-1kDs64OoY5U8OvERIB1Qxdwqed3kI5uLMG9ynY"), -- usuario3

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);
