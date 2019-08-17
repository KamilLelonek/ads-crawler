DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'relationship') THEN
        DROP TYPE relationship;
    END IF;
END$$;
