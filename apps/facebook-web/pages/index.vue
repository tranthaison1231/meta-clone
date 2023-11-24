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
    <div grid="~ cols-7" class="gap-20">
      <div class="col-span-2 p-2 flex items-center gap-2">
        <div> Logo </div>
        <Input class="h-10 w-60 rounded-3xl bg-gray-100" placeholder="Search Facebook" />
      </div>
      <div class="col-span-3 flex justify-between"> 
        <div> Home </div>
        <div> Video </div>
        <div> Marketplace </div>
        <div> Events </div>
        <div> Gaming </div>
      </div>
      <div class="col-span-2 flex items-center justify-end gap-2"> 
        <div> Menu </div>
        <div> Message </div>
        <div> Notification </div>
        <img class="h-10 w-10 rounded-full" :src="data?.data.user.avatar" />
      </div>
    </div>
  </header>
  <div grid="~ cols-7" class="gap-20 bg-gray-200 h-[calc(100vh-3.5rem)] p-4">
    <div class="col-span-2">
      <div class="flex gap-2 items-center">
        <img class="h-10 w-10 rounded-full" :src="data?.data.user.avatar" />
        <p> {{ data?.data.user.email }} </p>
      </div>
    </div>
    <div class="col-span-3"> 
      <div> Stories </div>
      <div class="p-4 bg-white rounded-xl mt-4">
        <div class="flex items-center gap-2">
          <img class="h-10 w-10 rounded-full" :src="data?.data.user.avatar" /> 
          <Input class="h-10 w-60 rounded-3xl bg-gray-100 w-full" placeholder="What's on your mind" />  
        </div>
        <hr class="mt-4 mb-6"/>
        <div class="flex justify-around">
          <div> Live video </div>
          <div> Photo/video </div>
          <div> Feeling/Activity </div>
        </div>
      </div>
      <div class="mt-4"> Posts </div>
    </div>
    <div class="col-span-2"> 
      <div>Your Pages and profiles</div>
      <hr />
      <div>
        <div>
          <h2> Contacts </h2>  
        </div>
      </div>
    </div>
  </div>
</template>
