<template>
  <div class="flex flex-col items-center min-h-screen justify-center h-full gap-8 w-full">
    <h1 class="header1 w-full max-w-[840px] text-center">
      EduApp - цифровая образовательная платформа
    </h1>

    <NuxtLink
        to="/login"
        class="btn-primary-big header4"
    >
      Войти
    </NuxtLink>
  </div>
</template>

<script setup lang="ts">
useSeoMeta({
  title: 'Главная - EduApp',
  description: 'Добро пожаловать на цифровую образовательную платформу EduApp',
})
</script>

<style scoped></style>
