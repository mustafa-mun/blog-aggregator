-- +goose Up
CREATE TABLE feedFollows (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  user_id UUID NOT NULL,
  feed_id UUID NOT NULL, 
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(feed_id) REFERENCES feeds(id),
  UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feedFollows;