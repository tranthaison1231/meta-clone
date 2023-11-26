import { getToken } from '~/lib/storage'

export default defineNuxtRouteMiddleware((to) => {
  const token = getToken()

  if (token && to?.name === 'login')
    return navigateTo('/')

  if (!token && to?.name !== 'login')
    return navigateTo('/login')
})
