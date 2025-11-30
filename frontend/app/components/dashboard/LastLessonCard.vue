<script setup lang="ts">
const cardTitle = 'Последний урок'
const lessonName = 'Урок 4. JS: DOM дерево. из курса Modern WebDeveloper 2025'
const lessonProgressLabel = 'Урок пройден на 76%'
const lessonProgressValue = 76
const courseProgressLabel = 'Курс пройден на 76%'
const courseProgressValue = 76
const btnContinueText = 'Продолжить урок'
const btnOpenCourseText = 'Открыть страницу курса'
</script>

<template>
  <div class="flex flex-col items-center border-large round gap-4 p-6 w-full h-full bg-surface box-border">
    <h3 class="header2 text-text-primary text-center">{{ cardTitle }}</h3>

    <div class="flex flex-col gap-6 w-full h-full justify-center">
      <h5 class="header4 text-text-primary text-center lg:text-left">
        {{ lessonName }}
      </h5>

      <div class="flex flex-col gap-4 w-full px-1 lg:px-4">
        <div class="flex flex-col lg:flex-row gap-2 lg:gap-5 items-start lg:items-center w-full">
          <p class="pbody text-text-primary whitespace-nowrap min-w-[160px]">
            {{ lessonProgressLabel }}
          </p>
          <div class="w-full">
            <ProgressBar :percent="lessonProgressValue" />
          </div>
        </div>

        <div class="flex flex-col lg:flex-row gap-2 lg:gap-5 items-start lg:items-center w-full">
          <p class="pbody text-text-primary whitespace-nowrap min-w-[160px]">
            {{ courseProgressLabel }}
          </p>
          <div class="w-full">
            <ProgressBar :percent="courseProgressValue" />
          </div>
        </div>
      </div>

      <div class="flex flex-col sm:flex-row gap-4 justify-between mt-2 lg:mt-4 w-full">
        <button class="btn-primary pbody w-full sm:w-auto text-center justify-center shadow-lg shadow-primary/20">
          {{ btnContinueText }}
        </button>
        <button class="btn-primary pbody w-full sm:w-auto text-center justify-center shadow-lg shadow-primary/20">
          {{ btnOpenCourseText }}
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>