<script setup lang="ts">
definePageMeta({
  middleware: 'auth',
})
</script>

<template>
  <h1>{{ $route.params.slug }}</h1>
</template>
