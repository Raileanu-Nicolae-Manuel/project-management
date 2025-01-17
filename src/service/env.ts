const ENV = {
    AUTH_TRUST_HOST: import.meta.env.AUTH_TRUST_HOST,
    AUTH_SECRET: import.meta.env.AUTH_SECRET,
    GITHUB_CLIENT_ID: import.meta.env.GITHUB_CLIENT_ID,
    GITHUB_CLIENT_SECRET: import.meta.env.GITHUB_CLIENT_SECRET,
    POSTGRES_USER: import.meta.env.POSTGRES_USER,
    POSTGRES_PASSWORD: import.meta.env.POSTGRES_PASSWORD,
    POSTGRES_DB: import.meta.env.POSTGRES_DB,
    POSTGRES_PORT: import.meta.env.POSTGRES_PORT,
    POSTGRES_HOST: import.meta.env.POSTGRES_HOST,
    POSTGRES_TEST_PORT: import.meta.env.POSTGRES_TEST_PORT,
}

export default ENV;
