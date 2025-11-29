<script setup lang="ts">
const email = ref('')
const password = ref('')
const loading = ref(false)
const errorMessage = ref('')

useSeoMeta({
  title: 'Вход - EduApp',
  description: 'Страница входа в цифровую образовательную платформу EduApp',
})

const api = useApiClient()

const onSubmit = async () => {
  // простая валидация на фронте
  if (!email.value || !password.value) {
    errorMessage.value = 'Введите логин и пароль'
    return
  }

  loading.value = true
  errorMessage.value = ''

  try {
    await api('/login', {
      method: 'POST',
      body: {
        email: email.value,
        password: password.value,
      },
    })
    // если дошли сюда — статус 201, кука уже поставлена бекендом
    await navigateTo('/dashboard')
  } catch (err: any) {
    const status = err?.status || err?.response?.status

    if (status === 400) {
      errorMessage.value = 'Неверный логин или пароль'
    } else {
      errorMessage.value = 'Произошла ошибка при входе. Попробуйте ещё раз'
    }
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="flex flex-col items-center min-h-screen justify-center">
    <div class="border-large bg-surface flex flex-col gap-16 items-center px-50 py-15">
      <h1 class="header1">Вход</h1>

      <form class="flex flex-col gap-12 items-center" @submit.prevent="onSubmit">
        <div class="flex flex-col gap-6 items-center w-full">
          <input
              v-model="email"
              class="bg-background border-medium placeholder-text-text-secondary text-text-primary py-3 px-5 w-full"
              type="text"
              placeholder="Логин"
              autocomplete="username"
          />
          <input
              v-model="password"
              class="bg-background border-medium placeholder-text-text-secondary text-text-primary py-3 px-5 w-full"
              type="password"
              placeholder="Пароль"
              autocomplete="current-password"
          />
        </div>

        <p v-if="errorMessage" class="text-error text-sm">
          {{ errorMessage }}
        </p>

        <button
            class="btn-primary-big header4"
            type="submit"
            :disabled="loading"
        >
          <span v-if="!loading">Войти</span>
          <span v-else>Входим…</span>
        </button>
      </form>
    </div>
  </div>
</template>

<style scoped>

</style>
