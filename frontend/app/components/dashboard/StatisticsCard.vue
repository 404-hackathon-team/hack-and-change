<script setup lang="ts">
import SemiCircleChart from '~/components/dashboard/SemiCircleChart.vue'

const percent = 74
const homeworksNeeded = 3
const overdueDeadlines = 1
</script>

<template>
  <article class="w-full h-full box-border border-main rounded bg-surface p-6 flex flex-col justify-between items-center text-center">

    <div class="w-full flex flex-col items-center pb-4 border-b border-border">
      <div class="w-40 mb-2">
        <SemiCircleChart :percent="percent" />
      </div>
      <h4 class="header4 text-text-primary leading-tight">
        Открытых уроков <br> пройдено
      </h4>
    </div>

    <div class="w-full flex flex-col items-center py-4 border-b border-border">
      <span class="header1 text-text-primary leading-none mb-1">
        {{ homeworksNeeded }}
      </span>
      <h4 class="header4 text-text-primary leading-tight">
        ДЗ необходимо <br> сдать
      </h4>
    </div>

    <div class="w-full flex flex-col items-center pt-4">
      <span class="header1 text-text-primary leading-none mb-1">
        {{ overdueDeadlines }}
      </span>
      <h4 class="header4 text-text-primary leading-tight">
        Просрочено <br> дедлайнов
      </h4>
    </div>

  </article>
</template>

<style scoped>
/* Дополнительные отступы не нужны, все управляется через UnoCSS классы */
</style>