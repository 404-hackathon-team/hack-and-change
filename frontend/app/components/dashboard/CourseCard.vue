<script setup lang="ts">
interface Props {
  title: string
  lessonsDone: number
  lessonsTotal: number
  deadlineDate?: string
  notificationsCount?: number
  lastUpdate: string
}

const props = defineProps<Props>()

const progressPercent = computed(() => {
  if (props.lessonsTotal === 0) return 0
  return Math.round((props.lessonsDone / props.lessonsTotal) * 100)
})
</script>

<template>
  <article class="w-full h-full border-main round bg-background p-4 lg:p-5 flex flex-col justify-between relative overflow-hidden box-border">

    <div class="flex flex-col w-full">
      <h3 class="header4 text-text-primary mb-2 line-clamp-2">
        {{ title }}
      </h3>

      <p class="pbody text-text-secondary mb-3">
        Уроков пройдено {{ lessonsDone }}/{{ lessonsTotal }}
      </p>

      <div class="w-full">
        <ProgressBar :percent="progressPercent" />
      </div>
    </div>

    <div class="flex flex-col w-full mt-4 lg:mt-6">

      <div class="flex flex-col gap-2 mb-4">

        <div v-if="deadlineDate" class="flex items-center gap-2 text-warning">
          <Icon name="material-symbols:release-alert-outline" class="w-5 h-5 shrink-0" />
          <span class="pbody">Сдать ДЗ до {{ deadlineDate }}</span>
        </div>

        <div v-if="notificationsCount && notificationsCount > 0" class="flex items-center gap-2 text-warning">
          <Icon name="material-symbols:add-comment-rounded" class="w-5 h-5 shrink-0" />
          <span class="pbody">+{{ notificationsCount }} уведомления</span>
        </div>

      </div>

      <div class="flex flex-col sm:flex-row items-start sm:items-end justify-between gap-3 sm:gap-0">
        <span class="pcaption text-disabled mb-1 whitespace-nowrap">
          Обновлено: {{ lastUpdate }}
        </span>

        <button class="btn-primary py-2 px-4 text-sm shadow-lg shadow-primary/20 transition hover:bg-primaryHover w-full sm:w-auto text-center justify-center">
          Продолжить
        </button>
      </div>
    </div>

  </article>
</template>

<style scoped>
</style>