```mermaid
---
title: netrunnerflix
---
erDiagram   
    USERS ||--o{ COMMENT : "author"
    USERS ||--o{ FILM : "comment"
    ACTOR }o--o{ FILM : "performed in"
    DIRECTOR ||--o{ FILM : "directed"
    FILM ||--o{ COMMENT : "commented on"
```
