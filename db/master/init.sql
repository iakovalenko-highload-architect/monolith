CREATE TABLE users (
   id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
   password BYTEA NOT NULL,
   first_name TEXT NOT NULL,
   second_name TEXT NOT NULL,
   birthday DATE NOT NULL,
   city TEXT NOT NULL,
   biography TEXT NOT NULL,
   created_at DATE NOT NULL DEFAULT NOW(),
   updated_at DATE NOT NULL DEFAULT NOW()
);

CREATE TABLE friendships (
   user_id UUID NOT NULL,
   friend_id UUID NOT NULL,
   created_at DATE NOT NULL DEFAULT NOW(),
   updated_at DATE NOT NULL DEFAULT NOW(),

   FOREIGN KEY (user_id) REFERENCES users(id),
   FOREIGN KEY (friend_id) REFERENCES users(id)
);

CREATE TABLE posts (
   id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
   user_id UUID NOT NULL,
   text_ TEXT NOT NULL,
   created_at DATE NOT NULL DEFAULT NOW(),
   updated_at DATE NOT NULL DEFAULT NOW(),

   FOREIGN KEY (user_id) REFERENCES users(id)
);