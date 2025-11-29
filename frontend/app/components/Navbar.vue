<template>
  <div>
    <nav class="flex flex-row justify-around border-large-bottom fixed w-full items-center">
      <NuxtLink to="/">
        <h2 class="header1 text-primary">
          EduApp
        </h2>
      </NuxtLink>

      <!-- Если авторизован — показываем имя и ссылку на дашборд -->
      <div v-if="isAuthenticated" class="flex flex-row items-center gap-6">
        <NuxtLink class="header3" to="/dashboard">
          {{ displayName }}
        </NuxtLink>
      </div>

      <!-- Если не авторизован — показываем Войти -->
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

console.log(user.value)

const displayName = computed(() => user.value?.lastName || user.value?.email || 'Error')


onMounted(() => {
  if (!user.value) {
    fetchMe()
  }
})
</script>

<style scoped>

</style>
