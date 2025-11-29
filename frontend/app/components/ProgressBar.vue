<script setup lang="ts">
interface Props {
  percent: number
}

const props = withDefaults(defineProps<Props>(), {
  percent: 0,
})

const clampedPercent = computed(() =>
    Math.min(100, Math.max(0, props.percent)),
)
</script>

<template>
  <div class="w-full h-3 bg-disabled rounded-full overflow-hidden">
    <div
        class="h-full bg-gradient-to-r from-primary via-primary to-primaryHover rounded-full transition-[width] duration-500 ease-in-out"
        :style="{ width: `${clampedPercent}%` }"
    />
  </div>
</template>