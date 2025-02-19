CREATE TABLE oracle_node_status (
    id SERIAL PRIMARY KEY,
    status_name VARCHAR(50) UNIQUE NOT NULL
);

INSERT INTO oracle_node_status (status_name) VALUES ('active'), ('inactive'), ('slashed');

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) UNIQUE NOT NULL
);

INSERT INTO roles (role_name) VALUES ('admin'), ('leader'), ('secondary leader'), ('regular');

CREATE TABLE oracle_nodes (
    node_address VARCHAR(255) PRIMARY KEY,
    oracle_node_public_key VARCHAR(255) NOT NULL,
    reward_wallet_address VARCHAR(255) NOT NULL,
    proxy_wallet_address VARCHAR(255),
    status_id INT NOT NULL,
    role_id INT NOT NULL,
    FOREIGN KEY (status_id) REFERENCES oracle_node_status(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

CREATE TABLE leaders (
    id SERIAL PRIMARY KEY,
    node_address VARCHAR(255) UNIQUE NOT NULL,
    assigned_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (node_address) REFERENCES oracle_nodes(node_address) ON DELETE CASCADE
);

CREATE TABLE rounds (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    round_number INT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE batches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    round_id UUID NOT NULL,
    batch_number INT NOT NULL,
    message_hash TEXT NOT NULL,
    message_json JSON NOT NULL,
    total_signatures INT DEFAULT 0,
    sent_signatures INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (round_id) REFERENCES rounds(id) ON DELETE CASCADE
);

CREATE TABLE leader_transactions (
    tx_id SERIAL PRIMARY KEY,
    leader_id VARCHAR(255) NOT NULL,
    round_id UUID NOT NULL,
    batch_id UUID NOT NULL,
    signatures_sent INT NOT NULL,
    tx_hash VARCHAR(100) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (round_id) REFERENCES rounds(id) ON DELETE CASCADE,
    FOREIGN KEY (batch_id) REFERENCES batches(id) ON DELETE CASCADE,
    FOREIGN KEY (leader_id) REFERENCES leaders(node_address) ON DELETE CASCADE
);

CREATE TABLE signatures (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    batch_id UUID NOT NULL,
    node_id UUID NOT NULL,
    oracle_node_public_key VARCHAR(255) NOT NULL,
    signature TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (batch_id) REFERENCES batches(id) ON DELETE CASCADE
);

CREATE TABLE jwt_issuance (
    id SERIAL PRIMARY KEY,
    node_address VARCHAR(255) NOT NULL,
    jwt_token TEXT NOT NULL,
    issued_at TIMESTAMP DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,
    blacklisted BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (node_address) REFERENCES oracle_nodes(node_address) ON DELETE CASCADE
);
