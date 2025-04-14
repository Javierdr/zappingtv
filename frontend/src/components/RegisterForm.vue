<template>
  <div class="register-container">
    <h1>Crear Cuenta</h1>
    <form @submit.prevent="register" class="register-form">
      <input v-model="username" placeholder="Usuario" required />
      <input v-model="password" type="password" placeholder="Contraseña" required />
      <button type="submit">Registrarse</button>
    </form>
    <button class="secondary-btn" @click="goToLogin">¿Ya tienes cuenta? Inicia sesión</button>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { userRegister } from '../api/user.js'
import { useAuthStore } from '../stores/auth'


const username = ref("")
const password = ref("")
const router = useRouter()
const auth = useAuthStore()

const register = async () => {
  const res = await userRegister(username.value, password.value)
  if (res.ok) {
    auth.login()
    router.push('/stream')
  } else {
    alert("Registro fallido")
  }
}

const goToLogin = () => {
  router.push('/login')
}
</script>

<style scoped>
.register-container {
  max-width: 420px;
  margin: 2rem auto;
  padding: 2.5rem;
  background-color: #1e293b;
  border-radius: 16px;
  box-shadow: 0 8px 24px rgba(0,0,0,0.2);
  font-family: system-ui, sans-serif;
  color: #f8fafc;
  text-align: center;
}

h1 {
  font-size: 1.8rem;
  color: #38bdf8;
  margin-bottom: 1.8rem;
}

.register-form input {
  width: 100%;
  padding: 0.75rem;
  margin-bottom: 1.2rem;
  border: none;
  border-radius: 8px;
  background-color: #334155;
  color: #f8fafc;
  font-size: 1rem;
}

.register-form input::placeholder {
  color: #cbd5e1;
}

.register-form button {
  width: 100%;
  background-color: #38bdf8;
  color: #0f172a;
  border: none;
  padding: 0.8rem;
  font-size: 1rem;
  font-weight: 600;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.register-form button:hover {
  background-color: #0ea5e9;
}

.secondary-btn {
  margin-top: 1.2rem;
  background: none;
  border: none;
  color: #38bdf8;
  cursor: pointer;
  font-size: 0.95rem;
  text-decoration: underline;
}
</style>
