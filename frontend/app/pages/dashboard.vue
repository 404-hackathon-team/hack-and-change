<script setup lang="ts">
import dayjs from 'dayjs'
import LastLessonCard from "~/components/dashboard/LastLessonCard.vue";
import CourseCard from "~/components/dashboard/CourseCard.vue";
import NotificationCard from "~/components/dashboard/NotificationCard.vue";
import StatisticsCard from "~/components/dashboard/StatisticsCard.vue";

const { user, isAuthenticated, fetchMe } = useAuth()

const displayName = computed(() => user.value?.lastName || user.value?.email || 'Error')

const greeting = ref('')
const navbarOffset = ref(0)
const search = ref('')

const updateGreeting = () => {
  const hour = dayjs().hour()
  if (hour >= 5 && hour < 12) greeting.value = 'Доброе утро'
  else if (hour >= 12 && hour < 18) greeting.value = 'Добрый день'
  else if (hour >= 18 && hour < 23) greeting.value = 'Добрый вечер'
  else greeting.value = 'Доброй ночи'
}

const notifications = [
  {
    type: 'deadline_future',
    date: '25.11',
    dayOfWeek: 'ВТ',
    title: 'Modern WebDeveloper 2025',
    description: 'Дедлайн по ДЗ по уроку 4'
  },
  {
    type: 'deadline_past',
    date: '24.11',
    dayOfWeek: 'ПН',
    title: 'JavaScript Advanced',
    description: 'Дедлайн по ДЗ по уроку 12 просрочен'
  },
  {
    type: 'homework_graded',
    date: '23.11',
    dayOfWeek: 'ВС',
    title: 'Vue.js Framework',
    description: 'Преподаватель оценил ваше ДЗ по уроку 4'
  },
  {
    type: 'new_lesson',
    date: '25.11',
    dayOfWeek: 'ВТ',
    title: 'Modern WebDeveloper 2025',
    description: 'Открылся новый урок: Урок 4: JS. DOM дерево.'
  },
  {
    type: 'deadline_future',
    date: '28.11',
    dayOfWeek: 'ПТ',
    title: 'UI/UX Design Basics',
    description: 'Сдача курсового проекта'
  },
  {
    type: 'homework_graded',
    date: '20.11',
    dayOfWeek: 'СР',
    title: 'Backend Python',
    description: 'Преподаватель оставил комментарий к вашему коду'
  }
] as const

const courses = [
  {
    id: 1,
    title: 'Modern WebDeveloper 2025',
    lessonsDone: 4,
    lessonsTotal: 32,
    deadlineDate: '30.11.25',
    notificationsCount: 3,
    lastUpdate: '25.11.25'
  },
  {
    id: 2,
    title: 'Vue.js 3 Deep Dive',
    lessonsDone: 28,
    lessonsTotal: 40,
    deadlineDate: null,
    notificationsCount: 0,
    lastUpdate: '24.11.25'
  },
  {
    id: 3,
    title: 'UI/UX Design Fundamentals',
    lessonsDone: 10,
    lessonsTotal: 15,
    deadlineDate: '01.12.25',
    notificationsCount: 1,
    lastUpdate: '20.11.25'
  },
  {
    id: 4,
    title: 'Backend with Node.js',
    lessonsDone: 0,
    lessonsTotal: 55,
    deadlineDate: null,
    notificationsCount: 5,
    lastUpdate: '10.11.25'
  },
  {
    id: 5,
    title: 'Docker & Kubernetes',
    lessonsDone: 12,
    lessonsTotal: 12,
    deadlineDate: null,
    notificationsCount: 0,
    lastUpdate: '01.11.25'
  }
]

const scrollToSection = (id: string) => {
  const element = document.getElementById(id)
  if (element) {
    const offset = window.innerWidth < 1024 ? 80 : navbarOffset.value + 60
    const bodyRect = document.body.getBoundingClientRect().top
    const elementRect = element.getBoundingClientRect().top
    const elementPosition = elementRect - bodyRect
    const offsetPosition = elementPosition - offset

    window.scrollTo({
      top: offsetPosition,
      behavior: 'smooth'
    })
  }
}

onMounted(() => {
  if (!user.value) fetchMe()
  if (!isAuthenticated.value) navigateTo('/login')
  updateGreeting()

  const nav = document.querySelector('nav')
  navbarOffset.value = nav ? nav.getBoundingClientRect().height : 45
})
</script>

<template>
  <div
      class="w-full box-border px-4 lg:px-12 h-auto min-h-screen lg:h-screen lg:overflow-hidden pb-8 lg:pb-8 bg-background pt-[var(--navbar-padding-mob)] lg:pt-[var(--navbar-padding)]"
      :style="{ '--navbar-padding': `${navbarOffset + 45}px`, '--navbar-padding-mob': `${navbarOffset + 5}px` }"
  >
    <div class="lg:hidden relative mb-4 sticky top-[var(--navbar-padding-mob)] z-30">

      <div class="flex gap-2 overflow-x-auto pb-2 no-scrollbar px-6">
        <button @click="scrollToSection('last-lesson')" class="whitespace-nowrap bg-surface border-medium round px-4 py-2 text-sm font-bold text-text-primary shadow-sm shrink-0">
          Последний урок
        </button>
        <button @click="scrollToSection('courses')" class="whitespace-nowrap bg-surface border-medium round px-4 py-2 text-sm font-bold text-text-primary shadow-sm shrink-0">
          Курсы
        </button>
        <button @click="scrollToSection('statistics')" class="whitespace-nowrap bg-surface border-medium round px-4 py-2 text-sm font-bold text-text-primary shadow-sm shrink-0">
          Статистика
        </button>
        <button @click="scrollToSection('notifications')" class="whitespace-nowrap bg-surface border-medium round px-4 py-2 text-sm font-bold text-text-primary shadow-sm shrink-0">
          Уведомления
        </button>
      </div>

      <div class="absolute top-0 left-0 h-full w-8 bg-gradient-to-r from-background to-transparent pointer-events-none rounded-l-lg"></div>
      <div class="absolute top-0 right-0 h-full w-8 bg-gradient-to-l from-background to-transparent pointer-events-none rounded-r-lg"></div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-12 lg:grid-rows-[auto_2fr_3fr] gap-6 w-full h-full">

      <section class="col-span-1 lg:col-span-8 flex items-center justify-start">
        <h1 class="header1 text-text-primary text-center lg:text-left">
          {{ greeting }}, {{ displayName }}
        </h1>
      </section>

      <aside class="col-span-1 lg:col-span-4 flex items-center">
        <input
            v-model="search"
            class="bg-background border-medium round placeholder-gray-400 text-text-primary py-3 px-5 w-full focus:outline-none focus:ring-2 focus:ring-primary/50 transition"
            type="text"
            placeholder="Поиск..."
            autocomplete="search"
        >
      </aside>

      <section id="last-lesson" class="col-span-1 lg:col-span-8 lg:row-start-2 scroll-mt-24">
        <LastLessonCard />
      </section>

      <aside
          id="courses"
          class="col-span-1 lg:col-span-4 lg:row-start-2 lg:row-span-2 flex flex-col border-large round bg-surface overflow-hidden shadow-sm h-[450px] lg:h-auto scroll-mt-24"
      >
        <div class="pt-6 px-1">
          <h3 class="header2 text-center text-text-primary mb-4 px-4">
            Мои курсы
          </h3>
        </div>

        <div class="relative w-full flex-1 min-h-0">

          <div class="w-full h-full overflow-y-auto flex flex-col gap-3 pb-12 px-6 lg:px-8 custom-scrollbar">
            <div v-for="course in courses" :key="course.id" class="h-auto shrink-0 first:mt-0 last:mb-4">
              <CourseCard
                  :title="course.title"
                  :lessons-done="course.lessonsDone"
                  :lessons-total="course.lessonsTotal"
                  :deadline-date="course.deadlineDate"
                  :notifications-count="course.notificationsCount"
                  :last-update="course.lastUpdate"
              />
            </div>
          </div>

          <div
              class="absolute bottom-0 left-0 w-full h-20 bg-gradient-to-t from-surface via-surface/90 to-transparent pointer-events-none z-10"
          />
        </div>
      </aside>

      <section id="statistics" class="col-span-1 lg:col-span-3 lg:row-start-3 scroll-mt-24">
        <StatisticsCard />
      </section>

      <section
          id="notifications"
          class="col-span-1 lg:col-span-5 lg:row-start-3 border-large round bg-surface flex flex-col overflow-hidden h-[450px] lg:h-auto scroll-mt-24"
      >
        <div class="pt-6 px-1">
          <h2 class="header2 text-center mb-4">Уведомления</h2>
        </div>

        <div class="relative w-full flex-1 min-h-0">

          <div class="w-full h-full overflow-y-auto flex flex-col gap-3 px-6 pb-12 custom-scrollbar">
            <NotificationCard
                v-for="(notification, index) in notifications"
                :key="index"
                :type="notification.type"
                :date="notification.date"
                :day-of-week="notification.dayOfWeek"
                :title="notification.title"
                :description="notification.description"
            />
          </div>

          <div
              class="absolute bottom-0 left-0 w-full h-20 bg-gradient-to-t from-surface via-surface/90 to-transparent pointer-events-none z-10"
          />
        </div>
      </section>

    </div>
  </div>
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background-color: #E0E4EC;
  border-radius: 4px;
}
.no-scrollbar::-webkit-scrollbar {
  display: none;
}
.no-scrollbar {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
</style>