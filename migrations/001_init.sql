CREATE TABLE IF NOT EXISTS wallets (
    id UUID PRIMARY KEY,
    balance BIGINT NOT NULL DEFAULT 0 CHECK (balance >= 0),
    created_at TIMESTAMP NOT NULL DEFAULT now()
    );