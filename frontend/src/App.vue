<template>
  <div class="container">
    <img src="./logo.png" alt="Logo" class="logo" />
    <router-view />
    
  </div>
</template>

<script setup>
</script>

<style>
body {
  margin: 0;
  background-color: #1e293b;
  font-family: system-ui, sans-serif;
  color: #f8fafc;
}
#app {
  min-height: 100vh;
}

.container {
  text-align: center;
  padding-top: 2rem;
}

.logo {
  width: 120px;
}
</style>
