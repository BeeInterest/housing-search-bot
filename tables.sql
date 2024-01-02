CREATE TABLE IF NOT EXISTS telegram_users (
    tg_id            SERIAL        PRIMARY KEY,
    tg_nick          VARCHAR(512)  NOT NULL,
    tg_phone         VARCHAR(12),
    tg_name          VARCHAR(512)
)

CREATE TABLE IF NOT EXISTS housing_ads_type (
    ads_type_id      SERIAL        PRIMARY KEY,
    ads_type_name    VARCHAR(64)   NOT NULL
)

CREATE TABLE IF NOT EXISTS housing_ads (
    ads_id           SERIAL        PRIMARY KEY,
    ads_name         VARCHAR(1024),
    ads_type_id      INTEGER       NOT NULL,
    FOREIGN KEY (ads_type_id) REFERENCES housing_ads_type (ads_type_id) 
)

CREATE TABLE IF NOT EXISTS housing_ads_params (
    ads_param_id     SERIAL        PRIMARY KEY,
    ads_param_name   VARCHAR(512)  NOT NULL,
    ads_type_id      INTEGER       NOT NULL,
    ads_param_type   VARCHAR(4)    NOT NULL,
    FOREIGN KEY (ads_type_id) REFERENCES housing_ads_type (ads_type_id) 
)

CREATE TABLE IF NOT EXISTS housing_ads_info (
    ads_info_id     SERIAL         PRIMARY KEY,
    ads_text        TEXT,
    ads_param_id    INTEGER        NOT NULL,
    ads_id          INTEGER        NOT NULL,
    FOREIGN KEY (ads_param_id) REFERENCES housing_ads_params (ads_param_id), 
    FOREIGN KEY (ads_id) REFERENCES housing_ads (ads_id) 
)

CREATE TABLE IF NOT EXISTS users_x_ads (
    rel_id          SERIAL         PRIMARY KEY,
    tg_id           INTEGER        NOT NULL,
    ads_id          INTEGER        NOT NULL,
    FOREIGN KEY (ads_id) REFERENCES housing_ads (ads_id),
    FOREIGN KEY (tg_id) REFERENCES telegram_users (tg_id)
)

CREATE TABLE IF NOT EXISTS preferences_block (
    block_id        SERIAL         PRIMARY KEY,
    tg_id           INTEGER        NOT NULL,
    ads_type_id     INTEGER        NOT NULL,
    FOREIGN KEY (tg_id) REFERENCES telegram_users (tg_id),
    FOREIGN KEY (ads_type_id) REFERENCES housing_ads_type (ads_type_id) 
)

CREATE TABLE IF NOT EXISTS preferences_users (
    pref_id         SERIAL         PRIMARY KEY,
    pref_text       TEXT,
    ads_param_id    INTEGER        NOT NULL,
    block_id        INTEGER        NOT NULL,
    FOREIGN KEY (ads_param_id) REFERENCES housing_ads_params (ads_param_id), 
    FOREIGN KEY (block_id) REFERENCES preferences_block (block_id) 
);


