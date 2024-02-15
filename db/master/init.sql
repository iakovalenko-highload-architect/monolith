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
