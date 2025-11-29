<script setup lang="ts">

</script>

<template>
  <div class="flex flex-col items-center min-h-screen justify-center">
    <div class="border-large bg-surface flex flex-col gap-16 items-center px-50 py-15">
      <h1 class="header1">Вход</h1>
      <form class="flex flex-col gap-12 items-center">
        <div class="flex flex-col gap-6 items-center">
          <input class="bg-background border-medium text-text-secondary py-3 px-5" type="text" placeholder="Логин">
          <input class="bg-background border-medium text-text-secondary py-3 px-5" type="text" placeholder="Пароль">
        </div>
        <button class="btn-primary-big header4">Войти</button>
      </form>
    </div>
  </div>
</template>

<style scoped>

</style>