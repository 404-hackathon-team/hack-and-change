<script setup lang="ts">
import dayjs from 'dayjs'
import LastLessonCard from "~/components/dashboard/LastLessonCard.vue";
import CourseCard from "~/components/dashboard/CourseCard.vue";

const { user, isAuthenticated, fetchMe } = useAuth()

const displayName = computed(() => user.value?.lastName || user.value?.email || 'Error')

const greeting = ref('')

const navbarOffset = ref(0)

const updateGreeting = () => {
  const hour = dayjs().hour()

  if (hour >= 5 && hour < 12) {
    greeting.value = 'Доброе утро'
  } else if (hour >= 12 && hour < 18) {
    greeting.value = 'Добрый день'
  } else if (hour >= 18 && hour < 23) {
    greeting.value = 'Добрый вечер'
  } else {
    greeting.value = 'Доброй ночи'
  }
}

onMounted(() => {
  if (!user.value) {
    fetchMe()
  }

  if (!isAuthenticated.value) {
    navigateTo('/login')
  }

  updateGreeting()

  const nav = document.querySelector('nav')
  if (nav) {
    navbarOffset.value = nav.getBoundingClientRect().height + 45
  } else {
    navbarOffset.value = 45
  }
})
</script>

<template>
  <div
      class="w-full h-screen box-border px-12"
      :style="{ paddingTop: `${navbarOffset}px`, paddingBottom: '2rem' }"
  >
    <div class="grid grid-cols-12 grid-rows-[auto_2fr_3fr] gap-6 w-full h-full">

      <section class="col-span-8 flex items-center">
        <h1 class="header1 text-text-primary">
          {{ greeting }}, {{ displayName }}
        </h1>
      </section>

      <aside class="col-span-4 flex items-center">
        <input
            class="bg-background border border-gray-200 rounded-xl placeholder-gray-400 text-text-primary py-3 px-5 w-full focus:outline-none focus:ring-2 focus:ring-blue-500"
            type="text"
            placeholder="Поиск..."
            autocomplete="search"
        >
      </aside>

      <section class="col-span-8 row-start-2">
        <LastLessonCard />
      </section>

      <aside class="col-span-4 row-start-2 row-span-2 flex flex-col items-center border-large rounded-lg py-5 px-10">
        <h3 class="header2">Мои курсы</h3>
        <div class="w-full">
          <!-- course cards -->
          <CourseCard />
        </div>
      </aside>

      <section class="col-span-3 row-start-3">
        <!-- stat card -->
      </section>

      <section class="col-span-5 row-start-3">
        <!-- notifications card -->
      </section>

    </div>
  </div>
</template>

<style scoped>

</style>
