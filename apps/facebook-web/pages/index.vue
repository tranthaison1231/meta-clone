<template>
  <div class="bg-gray-200 h-screen">
    <div class="max-w-5xl w-full mx-auto flex justify-between items-center h-screen">
      <div class="w-1/2">
        <h1 class="text-primary text-6xl font-bold">facebook</h1>
        <p class="text-3xl mt-4">Facebook helps you connect and share with the people in your life.</p>
      </div>
      <div class="p-4 bg-white rounded-xl w-104 h-fit">
        <Input placeholder="Email address or phone number"/>
        <Input placeholder="Password" class="my-3" />
        <Button class="w-full text-xl">Log in</Button>
        <a class="text-primary hover:underline mt-3 block text-center"> Forgotten password? </a>
        <hr class="my-4"/>
        <Button variant="success" class="w-full text-xl my-2">Create new account</Button>
      </div>
    </div>
  </div>
</template>
