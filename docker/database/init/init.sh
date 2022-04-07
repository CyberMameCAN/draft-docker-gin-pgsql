set -e
psql -U postgres postgres << EOSQL
CREATE TABLE Users(
  account_id    SERIAL PRIMARY KEY,
  account_name  VARCHAR(20),
  email         VARCHAR(100),
  password      CHAR(64)
);

CREATE TABLE UserStatus(
  status        VARCHAR(20) PRIMARY KEY
);
EOSQL