DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'relationship') THEN
        CREATE TYPE relationship AS ENUM
        (
            'RESELLER',
            'DIRECT'
        );
    END IF;
END$$;
