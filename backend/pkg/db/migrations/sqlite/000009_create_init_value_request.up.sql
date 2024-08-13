INSERT INTO users (email, password, firstname, lastname, date_of_birth, avatar, nickname, about_me)
    VALUES
        ('johndoe@user.com', '$2y$10$ot3bVA9D4C9kumObaVFjruEQMgVBu7kO.2Z1bf5xv0kGS8eFzUL0G', 'John', 'Doe', '1990-01-01', 'https://cdn.pixabay.com/photo/2020/07/01/12/58/icon-5359553_960_720.png', 'johndoe', 'I am John Doe'),
        ('janedoe@user.com', '$2y$10$ot3bVA9D4C9kumObaVFjruEQMgVBu7kO.2Z1bf5xv0kGS8eFzUL0G', 'Jane', 'Doe', '1990-01-01', 'https://cdn.pixabay.com/photo/2020/07/01/12/58/icon-5359553_960_720.png', 'janedoe', 'I am Jane Doe'),
        ('marcusdoe@user.com', '$2y$10$ot3bVA9D4C9kumObaVFjruEQMgVBu7kO.2Z1bf5xv0kGS8eFzUL0G', 'Marcus', 'Doe', '1990-01-01', 'https://cdn.pixabay.com/photo/2020/07/01/12/58/icon-5359553_960_720.png', 'marcusdoe', 'I am Marcus Doe'),
        ('jackdoe@user.com', '$2y$10$ot3bVA9D4C9kumObaVFjruEQMgVBu7kO.2Z1bf5xv0kGS8eFzUL0G', 'Jack', 'Doe', '1990-01-01', 'https://cdn.pixabay.com/photo/2020/07/01/12/58/icon-5359553_960_720.png', 'jackdoe', 'I am Jack Doe');

INSERT INTO categories (name)
    VALUES
        ('General'),
        ('Technology'),
        ('Science'),
        ('Health'),
        ('Sports'),
        ('Entertainment'),
        ('Business'),
        ('Education'),
        ('Travel'),
        ('Food');

INSERT INTO follows (follower_id, followee_id)
    VALUES
        (2, 1),
        (3, 1),
        (2, 3),
        (4, 1),
        (4, 2),
        (4, 3);

INSERT INTO posts (title, content, image, privacy, user_id)
    VALUES
        ('Hello World', 'This is my first post', 'https://cdn.pixabay.com/photo/2024/06/10/16/17/ai-generated-8820987_960_720.png', 'public', 1),
        ('Hello World', 'This is my first post', 'https://cdn.pixabay.com/photo/2024/06/10/16/17/ai-generated-8820987_960_720.png', 'public', 2),
        ('Hello World', 'This is my first post', 'https://cdn.pixabay.com/photo/2024/06/10/16/17/ai-generated-8820987_960_720.png', 'public', 3),
        ('Hello World', 'This is my first post', 'https://cdn.pixabay.com/photo/2024/06/10/16/17/ai-generated-8820987_960_720.png', 'public', 4),
        ('Hello World', 'This is my first post', 'https://cdn.pixabay.com/photo/2024/06/10/16/17/ai-generated-8820987_960_720.png', 'private', 1),
        ('Hello World', 'This is my first post', 'https://cdn.pixabay.com/photo/2024/06/10/16/17/ai-generated-8820987_960_720.png', 'private', 2),
        ('Hello World', 'This is my first post', 'https://cdn.pixabay.com/photo/2024/06/10/16/17/ai-generated-8820987_960_720.png', 'private', 3),
        ('Hello World', 'This is my first post', 'https://cdn.pixabay.com/photo/2024/06/10/16/17/ai-generated-8820987_960_720.png', 'private', 4);


INSERT INTO category_post (post_id, category_id)
    VALUES
        (1, 1),
        (2, 2),
        (3, 3),
        (4, 4),
        (5, 5),
        (6, 6),
        (7, 7),
        (8, 8);
