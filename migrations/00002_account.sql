-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS account (
    id UUID PRIMARY KEY,
    client_id UUID NOT NULL REFERENCES client(id) ON DELETE CASCADE,
    balance BIGINT NOT NULL DEFAULT 0 CHECK (balance >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_account_client_id ON account(client_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account;
-- +goose StatementEnd

