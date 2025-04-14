export async function userRegister(username, password, baseURL=import.meta.env.VITE_API_URL) {
    const res = await fetch(`${baseURL}/register`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ "username": username, "password": password })
  })
  return res
}

export async function userLogin(username, password, baseURL=import.meta.env.VITE_API_URL) {
    const res = await fetch(`${baseURL}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ "username": username, "password": password })
  })
  return res
}
