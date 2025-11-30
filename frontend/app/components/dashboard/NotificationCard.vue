<script setup lang="ts">
interface Props {
  type: 'deadline_future' | 'deadline_past' | 'homework_graded' | 'new_lesson'
  date: string
  dayOfWeek: string
  title: string
  description: string
}

const props = defineProps<Props>()

const config = computed(() => {
  switch (props.type) {
    case 'deadline_past':
      return {
        icon: 'material-symbols:error-circle-rounded',
        colorClass: 'text-error/40'
      }
    case 'homework_graded':
      return {
        icon: 'material-symbols:star-outline',
        colorClass: 'text-text-secondary/20'
      }
    case 'new_lesson':
      return {
        icon: 'material-symbols:add-circle-outline',
        colorClass: 'text-text-secondary/20'
      }
    case 'deadline_future':
    default:
      return {
        icon: 'material-symbols:schedule-outline',
        colorClass: 'text-text-secondary/20'
      }
  }
})
</script>

<template>
  <article
      class="relative w-full overflow-hidden border-main rounded bg-background p-4 pr-0 flex items-center gap-5 shrink-0"
  >
    <div class="flex flex-col items-center justify-center shrink-0 pl-2 w-14">
      <span class="header3 text-text-primary leading-none mb-1">
        {{ date }}
      </span>
      <span class="psubtitle font-bold text-text-primary uppercase leading-none">
        {{ dayOfWeek }}
      </span>
    </div>

    <div class="flex flex-col gap-1 z-10 mr-4 relative">
      <h4 class="header4 text-text-primary line-clamp-1">
        {{ title }}
      </h4>
      <p class="pbody text-text-secondary text-sm line-clamp-2">
        {{ description }}
      </p>
    </div>

    <div
        class="absolute -right-5 -bottom-7 pointer-events-none"
        :class="config.colorClass"
    >
      <Icon
          :name="config.icon"
          size="110px"
          class="shrink-0"
      />
    </div>
  </article>
</template>

<style scoped>
</style>