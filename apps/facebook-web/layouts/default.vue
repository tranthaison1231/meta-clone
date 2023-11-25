<script setup lang="ts">
  import { request } from '~/lib/request';
  useHead({
    title: 'Facebook'
  })

  definePageMeta({
    middleware: 'auth'
  });

  const { data } = useQuery({
    queryKey: ['me'],
    queryFn: () => request.get('/me'),
  })
</script>

<template>
  <Header :user="data?.data.user" />
  <slot :user="data?.data.user"/>
</template>
