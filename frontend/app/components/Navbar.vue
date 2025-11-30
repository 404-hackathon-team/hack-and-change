<template>
  <div>
    <nav class="flex flex-row justify-around border-large-bottom fixed w-full items-center bg-surface z-36">
      <NuxtLink to="/">
        <h2 class="header1 text-primary">
          EduApp
        </h2>
      </NuxtLink>

      <div v-if="isAuthenticated" class="flex flex-row items-center gap-6">
        <NuxtLink class="header3" to="/dashboard">
          {{ displayName }}
        </NuxtLink>
      </div>

      <NuxtLink
          v-else
          class="header3"
          to="/login"
      >
        Войти
      </NuxtLink>
    </nav>
  </div>
</template>

<script setup lang="ts">
const { user, isAuthenticated, fetchMe } = useAuth()

const displayName = computed(() => user.value?.lastName || user.value?.email || 'Error')


onMounted(() => {
  if (!user.value) {
    fetchMe()
  }
})
</script>

<style scoped>

</style>
