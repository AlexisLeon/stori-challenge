CREATE TABLE IF not EXISTS {{ index .Options "Namespace" }}.accounts (
    id UUID DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    currency VARCHAR(255) NOT NULL,
    balance INTEGER NOT NULL,

    PRIMARY KEY (id),

    CONSTRAINT fk_user_id
    FOREIGN KEY(user_id)
    REFERENCES {{ index .Options "Namespace" }}.users(id)
    ON DELETE NO ACTION

    -- For accounting-type ledgers, we could have balances for
    -- * cash
    -- * in escrow
    -- * receivables
    -- * confirmed
    -- * equity
    -- * etc
);

comment on table {{ index .Options "Namespace" }}.accounts is 'Account that holds funds for a given user';
