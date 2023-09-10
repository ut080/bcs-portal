<script setup lang="ts">
</script>

<!-- TODO: Add responsive design breakpoints -->
<template>
  <div class="flex flex-col">
    <!-- top row of header -->
    <nav class="flex flex-row items-center">
      <div class="w-1/4"></div>
      <NuxtLink to="/">
        <img src="/img/logo-bcs.png" alt="Civil Air Patrol Blackhawk Cadet Squadron" class="h-20"/>
      </NuxtLink>
      <div class="grow"></div>
      <img src="/img/bcs-patch.png" alt="Hit 'em harder!" class="h-32"/>
      <div class="w-1/4"></div>
    </nav>

    <!-- bottom row of header/menu bar -->
    <div class="h-12 w-full bg-symbol-blue flex flex-row items-center place-content-around text-lg text-white ">
      <div class="w-1/4"></div>
      <NavButton path="/about" display-name="ABOUT"/>
      <NavButton path="/join" display-name="JOIN"/>
      <NavButton path="/calendar" display-name="CALENDAR" />
      <NavButton path="/ood" display-name="ORDERS OF THE DAY" />
      <NavButton path="/pubs" display-name="PUBLICATIONS AND FORMS" />
      <div class="w-1/4"></div>
    </div>
  </div>
</template>

<style scoped>
</style>
