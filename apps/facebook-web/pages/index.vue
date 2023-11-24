<script setup lang="ts">
  import { request } from '~/lib/request';

  definePageMeta({
      middleware: 'auth'
  });

  const { data } = useQuery({
    queryKey: ['me'],
    queryFn: () => request.get('/me'),
  })
</script>

<template>
  
  <header class="h-14 shadow-md px-4">
    <div grid="~ cols-7">
      <div class="col-span-2 p-2 flex items-centergap-2">
        <div> Logo </div>
        <Input class="h-10 w-60 rounded-3xl bg-gray-100" placeholder="Search Facebook" />
      </div>
      <div class="col-span-3"> Menu </div>
      <div class="col-span-2 flex items-center"> 
        <img class="h-10 w-10 rounded-full" :src="data?.data.user.avatar" />
      </div>
    </div>
  </header>
  <div grid="~ cols-7" class="bg-gray-200 h-[calc(100vh-3.5rem)] p-4">
    <div class="col-span-2">Sidebar</div>
    <div class="col-span-3"> New feeds </div>
    <div class="col-span-2"> Right side</div>
  </div>
</template>
