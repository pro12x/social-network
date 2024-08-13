CREATE TABLE IF NOT EXISTS category_post (
    category_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    PRIMARY KEY (category_id, post_id),
    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE
);