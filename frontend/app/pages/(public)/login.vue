<script setup lang="ts">
const login = useMutationLogin();

const handleLogin = () => {
    login.mutateAsync(
        {
            email: "email",
            password: "password",
        },
        {
            onSuccess: () => {
                console.log("success");
            },
            onError: () => {
                console.log("error");
            },
        },
    );
};
</script>

<template>
    <div>LOGIN</div>
    <UiButton @click="handleLogin">test</UiButton>
</template>
