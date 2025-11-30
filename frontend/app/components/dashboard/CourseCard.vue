<template>
  <article class="w-full h-full box-border border-main rounded bg-background p-3 flex flex-col justify-between relative overflow-hidden">

    <div class="flex h-full w-full flex-col justify-between">
      <h3 class="header4 text-text-primary">
        Modern WebDeveloper 2025
      </h3>

      <p class="pbody text-text-secondary">
        Уроков пройдено 4/32
      </p>

      <div>
        <ProgressBar :percent="12" />
      </div>
    </div>

    <div class="flex h-full w-full flex-col mt-8">
      <div class="flex flex-col gap-2">

        <div class="flex items-center gap-2 text-warning">
          <Icon name="material-symbols:release-alert-outline" class="w-5 h-5 shrink-0" />
          <span class="pbody">Сдайте ДЗ до 30.11.25</span>
        </div>

        <div class="flex items-center gap-2 text-warning">
          <Icon name="material-symbols:add-comment-rounded" class="w-5 h-5 shrink-0" />
          <span class="pbody">+3 уведомления</span>
        </div>

      </div>
      <div class="flex items-end justify-between">
      <span class="pcaption text-disabled mb-1 whitespace-nowrap">
        Последнее обновление: 25.11.25
      </span>

        <button class="btn-primary">
          Продолжить
        </button>
      </div>
    </div>

  </article>
</template>

<script setup lang="ts">
// Логика отсутствует, так как данные захардкожены
</script>