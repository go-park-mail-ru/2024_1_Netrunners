```mermaid
---
title: netrunnerflix
---
erDiagram   
    USERS ||--o{ COMMENT : "author"
    USERS ||--o{ FILM : "comment"
    FILM ||--o{ FILM_ACTOR : "performed in"
    ACTOR ||--o{ FILM_ACTOR : "performed in"
    DIRECTOR ||--o{ FILM : "directed"
    FILM ||--o{ COMMENT : "commented on"
    ACTOR ||--|| FILM_ACTOR : "performed in"
```
