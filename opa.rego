package swisscom

default allow = false

allow {
    input.path = ["api", "users"]
    input.method = "GET"
    input.authenticated
}

allow {
    input.path = ["api", "users"]
    input.method = "POST"
    input.authenticated
    input.role == "admin"
}