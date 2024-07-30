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
        (1, 2),
        (1, 3),
        (2, 3);