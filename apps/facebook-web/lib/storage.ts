export function getToken() {
  const token = useCookie('token')
  return token.value
}

export function setToken(value: string) {
  const token = useCookie('token')
  token.value = value
}

export function removeToken() {
  const token = useCookie('token')
  token.value = ''
}
