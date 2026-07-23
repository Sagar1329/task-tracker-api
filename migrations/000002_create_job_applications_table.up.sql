CREATE TYPE application_status AS ENUM (
    'Applied',
    'Interview',
    'Rejected',
    'Offer'
);

CREATE TYPE job_portal AS ENUM (
    'LinkedIn',
    'Naukri',
    'Indeed',
    'Wellfound',
    'Company Careers',
    'Other'
);

CREATE TABLE job_applications (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE,

    company_name VARCHAR(255) NOT NULL,

    job_title VARCHAR(255) NOT NULL,

    job_link TEXT,

    job_portal job_portal NOT NULL,

    location VARCHAR(255),

    status application_status NOT NULL DEFAULT 'Applied',

    applied_date DATE NOT NULL,

    notes TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_job_applications_user_id
ON job_applications(user_id);

CREATE INDEX idx_job_applications_status
ON job_applications(status);

CREATE INDEX idx_job_applications_company_name
ON job_applications(company_name);

CREATE INDEX idx_job_applications_applied_date
ON job_applications(applied_date);