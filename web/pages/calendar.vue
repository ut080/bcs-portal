<script setup lang="ts">
const { data } = await useAsyncData('calendar', () => queryContent('/calendar').findOne())

</script>

<template>
    <Calendar />
</template>

<style scoped>

</style>