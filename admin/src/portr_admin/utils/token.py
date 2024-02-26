import nanoid
from ulid import ULID


def generate_secret_key() -> str:
    return nanoid.generate(size=42)


def generate_oauth_state() -> str:
    return nanoid.generate(size=26)


def generate_session_token() -> str:
    return nanoid.generate(size=32)


def generate_connection_id() -> str:
    return str(ULID())