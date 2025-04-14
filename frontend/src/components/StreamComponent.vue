<template>
  <div class="stream-container">
    
    <div v-if="!isStreamActive" class="message">
      <p>
        El stream ha finalizado. Gracias por ver :)<br />
        Para ver más, <a href="https://www.zapping.com/?ref=zapping.cl" target="_blank">haz clic aquí</a>.<br /><br />
        O reinicia el stream 👇🏽
      </p>
    </div>
    <div v-else>
      <h1 class="live-title"><span class="blinking">🔴</span> En vivo</h1>
      <video ref="video" autoplay muted playsinline controls></video>
    </div>
    <div class="controls">
      <button v-if="isStreamActive" @click="previousStream">⏪ 10s atrás</button>
      <button @click="resetStream">🔄 Reiniciar Stream</button>
    </div>
    <button class="logout-btn" @click="logout">Cerrar Sesión</button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Hls from 'hls.js'
import { getResetStream, getPreviousSegment } from '../api/video.js'
import { useAuthStore } from '../stores/auth'
import { useRouter } from 'vue-router'

const isStreamActive = ref(true)
const video = ref(null)
let hls = null
const auth = useAuthStore()
const router = useRouter()


const loadStream = () => {
  const videoElement = video.value
  const hlsUrl = `${import.meta.env.VITE_API_URL}/playlist.m3u8`

  if (Hls.isSupported()) {
    if (hls) hls.destroy()
    hls = new Hls({ liveSyncDuration: 10, startPosition: 0, lowLatencyMode: true })
    hls.loadSource(hlsUrl)
    hls.attachMedia(videoElement)
    hls.on(Hls.Events.LEVEL_LOADED, (_, data) => {
      const details = data.details
      if (!details.live) {
        const checkInterval = setInterval(() => {
          if (videoElement.duration && videoElement.currentTime >= videoElement.duration - 0.1) {
            isStreamActive.value = false
            clearInterval(checkInterval)
          }
        }, 500)
      }
    })
  } else if (videoElement.canPlayType('application/vnd.apple.mpegurl')) {
    videoElement.src = hlsUrl
  }
}

const logout = () => {
  auth.logout()
  router.push('/login')
}

const resetStream = async () => {
  await getResetStream()
  isStreamActive.value = true
  loadStream()
}

const previousStream = async () => {
  await getPreviousSegment()
  loadStream()
}

onMounted(() => {
  if (auth.isLoggedIn) loadStream()
  if (!isStreamActive.value || !auth.isLoggedIn) {
    router.push('/login') // Redirige al login
  }
})
</script>

<style scoped>
.stream-container {
  max-width: 960px;
  margin: 3rem auto;
  padding: 2rem;
  background-color: #1e293b;
  border-radius: 20px;
  box-shadow: 0 8px 24px rgba(0,0,0,0.2);
  font-family: system-ui, sans-serif;
  text-align: center;
  color: #f8fafc;
}

.live-title {
  color: #38bdf8;
  font-size: 2rem;
  margin-bottom: 1.5rem;
}

video {
  width: 100%;
  max-height: 520px;
  border-radius: 12px;
  box-shadow: 0 0 15px rgba(0,0,0,0.4);
  margin-bottom: 1.5rem;
  background: #000;
}

.controls button,
.logout-btn {
  margin: 0.5rem;
  padding: 0.75rem 1.2rem;
  font-size: 1rem;
  font-weight: 600;
  background-color: #38bdf8;
  color: #0f172a;
  border: none;
  border-radius: 10px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.controls button:hover,
.logout-btn:hover {
  background-color: #0ea5e9;
}

.message {
  font-size: 1.1rem;
  color: #cbd5e1;
  margin-bottom: 2rem;
}

.message a {
  color: #38bdf8;
  text-decoration: underline;
}

.blinking {
  animation: blink 2s infinite;
}
@keyframes blink {
  0%, 100% { opacity: 1; }
  50% { opacity: 0; }
}
</style>
