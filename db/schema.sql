CREATE TABLE email (
    recipient VARCHAR(255) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    send_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    uuid UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid()
);

CREATE TABLE tracker (
    tracker_id SERIAL PRIMARY KEY,
    open_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    email_uuid UUID NOT NULL,
    ip_address VARCHAR(255),

    FOREIGN KEY (email_uuid) REFERENCES email(uuid)
);
