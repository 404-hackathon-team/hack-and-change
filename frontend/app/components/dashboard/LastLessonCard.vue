<script setup lang="ts">

// TODO backend not ready

</script>

<template>
  <div class="flex flex-col items-center border-large rounded-lg gap-4 p-4 w-full h-full">
    <h3 class="header2">Последний урок</h3>
    <div class="flex flex-col gap-5 ml-4 w-full">
      <h5 class="header4">Урок 4. JS: DOM дерево. из курса Modern WebDeveloper 2025</h5>
      <div class="flex flex-row gap-5 ml-2 items-center mr-12">
        <p class="whitespace-nowrap">Урок пройден на 76%</p>
        <ProgressBar :percent="76" />
      </div>
      <div class="flex flex-row gap-5 ml-2 items-center mr-12">
        <p class="whitespace-nowrap">Курс пройден на 76%</p>
        <ProgressBar :percent="76" />
      </div>
      <div class="flex flex-row justify-between mr4">
        <button class="btn-primary pbody">Продолжить урок</button>
        <button class="btn-primary pbody">Открыть страницу курса</button>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>