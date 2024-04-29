CREATE TABLE IF not EXISTS {{ index .Options "Namespace" }}.transactions (
    id UUID DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    account_id UUID NOT NULL,
    transaction_uuid UUID NOT NULL UNIQUE,
    amount INTEGER NOT NULL,

    PRIMARY KEY (id),

    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
            REFERENCES {{ index .Options "Namespace" }}.users(id)
            ON DELETE NO ACTION,

    CONSTRAINT fk_account_id
    FOREIGN KEY(account_id)
    REFERENCES {{ index .Options "Namespace" }}.accounts(id)
            ON DELETE NO ACTION
);

comment on table {{ index .Options "Namespace" }}.transactions is 'Transactions for a given account';
