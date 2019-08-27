CREATE TABLE publishers (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    website text NOT NULL,
    domain text NOT NULL,
    account_id text NOT NULL,
    relationship relationship NOT NULL,
    authority text,
    created_at timestamp with time zone NOT NULL
);
