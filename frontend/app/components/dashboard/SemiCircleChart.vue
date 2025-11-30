<script setup lang="ts">
interface Props {
  percent: number
}

const props = defineProps<Props>()

const radius = 40
const circumference = Math.PI * radius
const strokeDashoffset = computed(() => {
  return circumference - (circumference * props.percent) / 100
})
</script>

<template>
  <div class="relative w-full flex flex-col items-center justify-end overflow-hidden" style="aspect-ratio: 2/1;">
    <svg
        viewBox="0 0 100 50"
        class="w-full h-full transform overflow-visible"
    >
      <path
          d="M 10 50 A 40 40 0 0 1 90 50"
          fill="none"
          stroke="#9CA3AF"
          stroke-width="12"
          stroke-linecap="butt"
          class="opacity-50"
      />


      <path
          d="M 10 50 A 40 40 0 0 1 90 50"
          fill="none"
          class="text-primary transition-all duration-1000 ease-out"
          stroke="currentColor"
          stroke-width="12"
          stroke-linecap="butt"
          :stroke-dasharray="circumference"
          :stroke-dashoffset="strokeDashoffset"
      />
    </svg>

    <div class="absolute bottom-0 flex items-baseline">
      <span class="header2 text-text-primary leading-none">
        {{ percent }}
      </span>
      <span class="header2 text-text-primary leading-none ml-1">
        %
      </span>
    </div>
  </div>
</template>

<style scoped>
</style>