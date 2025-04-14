import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const isLoggedIn = ref(false)
  const token = ref(localStorage.getItem('token') || '')

  const login = (newToken) => { 
    token.value = newToken
    localStorage.setItem('token', newToken)
    isLoggedIn.value = true 
  }

  const logout = () => { 
    token.value = ''
    localStorage.removeItem('token')
    isLoggedIn.value = false 
  }

  // Inicializar el estado de autenticación basado en el token
  if (token.value) {
    isLoggedIn.value = true
  }

  return { isLoggedIn, token, login, logout }
})