CREATE TABLE campaigns (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    coupon_remains INT NOT NULL DEFAULT 0 CHECK (coupon_remains >= 0),
    begin_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE coupons (
    id BIGSERIAL PRIMARY KEY,
    code CHAR(10) NOT NULL UNIQUE,
    campaign_id INT NOT NULL,
    issued_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_coupons_campaign
        FOREIGN KEY(campaign_id)
        REFERENCES campaigns(id)
        ON DELETE CASCADE
);
CREATE INDEX idx_coupons_campaign_id ON coupons(campaign_id);

CREATE TABLE campaign_schedules (
    id SERIAL PRIMARY KEY,
    campaign_id INT NOT NULL UNIQUE,
    status VARCHAR(10) NOT NULL DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'ACTIVATED', 'DONE')),
    begin_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_schedules_campaign
        FOREIGN KEY(campaign_id)
        REFERENCES campaigns(id)
        ON DELETE CASCADE
);
