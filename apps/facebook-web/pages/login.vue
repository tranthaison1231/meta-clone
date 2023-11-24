<script setup lang="ts">
  import { loginSchema } from '~/lib/schema'
  import { request } from '~/lib/request'
  import { configure } from 'vee-validate';
  import type { z } from 'zod';

  const { mutateAsync } = useMutation({
    mutationFn: (data: z.infer<typeof loginSchema>) => request.post('/sign-in', data),
  })

  configure({
    validateOnBlur: true
  })

  const { handleSubmit, errors, useFieldModel } = useForm({
    validationSchema: toTypedSchema(loginSchema),
  })

  const [email, password] = useFieldModel(['email', 'password']);

  const onSubmit = handleSubmit(async (values) => {
    await mutateAsync(values)
    navigateTo('/')
  })
</script>

<template>
  <div class="bg-gray-200 h-screen">
    <div class="max-w-5xl w-full mx-auto flex justify-between items-center h-screen">
      <div class="w-1/2">
        <h1 class="text-primary text-6xl font-bold">facebook</h1>
        <p class="text-3xl mt-4">Facebook helps you connect and share with the people in your life.</p>
      </div>
      <form class="p-4 bg-white rounded-xl w-104 h-fit" @submit="onSubmit">
        <Input placeholder="Email address or phone number" v-model="email" />
        <span class="mt-1 w-full text-red-500">{{ errors.email }}</span>
        <Input placeholder="Password" class="my-3" v-model="password" type="password" />
        <span class="mt-1 w-full text-red-500">{{ errors.password }}</span>
        <Button class="w-full text-xl" type="submit">Log in</Button>
        <a class="text-primary hover:underline mt-3 block text-center"> Forgotten password? </a>
        <hr class="my-4" />
        <Button variant="success" class="w-full text-xl my-2" type="button">Create new account</Button>
      </form>
    </div>
  </div>
</template>
