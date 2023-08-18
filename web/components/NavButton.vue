<script setup lang="ts">
defineProps<{
  path: string
  displayName: string
}>()

</script>

<template>
  <NuxtLink :to="path" class="h-12 grow hover:bg-scarlet-red flex items-center place-content-center">
    {{displayName}}
  </NuxtLink>
</template>

<style scoped>

</style>