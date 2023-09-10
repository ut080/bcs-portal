<script lang="ts">
import {defineComponent} from 'vue'
import FullCalendar from '@fullcalendar/vue3'
import dayGridPlugin from '@fullcalendar/daygrid'

export default defineComponent({
  name: "Calendar",
  components: {
    FullCalendar,
  },
  data() {
    return {
      calendarOptions: {
        plugins: [dayGridPlugin],
        initialView: 'dayGridMonth',
        handleWindowResize: true,
        height: 'auto',
      }
    }
  }
})
</script>

<template>
  <div id="calendar">
    <FullCalendar :options="calendarOptions"/>
  </div>
</template>

<style scoped>
#calendar {
  width: 50%
}
</style>