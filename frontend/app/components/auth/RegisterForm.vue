<script setup lang="ts">
import type { HTMLAttributes } from "vue";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
    Field,
    FieldDescription,
    FieldGroup,
    FieldLabel,
} from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { toast } from "vue-sonner";
import { Eye, EyeOff } from "lucide-vue-next";

const props = defineProps<{
    class?: HTMLAttributes["class"];
}>();

const name = ref("");
const email = ref("");
const password = ref("");
const errors = ref<Record<string, string>>({});
const showPassword = ref(false);

const { mutateAsync: register, isPending } = useMutationRegister();

const handleRegister = () => {
    errors.value = {};
    register(
        { name: name.value, email: email.value, password: password.value },
        {
            onSuccess: (response) => {
                toast.success(response.message || "Registrasi berhasil");
                localStorage.setItem("authToken", response.token);
                navigateTo("/home");
            },
            // @ts-ignore
            onError: (error: CustomErrorResponse) => {
                toast.error(error.message || "Registrasi gagal");
                if (error.errors) errors.value = error.errors;
            },
        },
    );
};
</script>

<template>
    <form :class="cn('flex flex-col gap-6', props.class)">
        <FieldGroup>
            <div class="flex flex-col items-center gap-1 text-center">
                <h1 class="text-2xl font-bold">Daftar Akun</h1>
                <p class="text-muted-foreground text-sm text-balance">
                    Buat akun untuk mulai berkolaborasi
                </p>
            </div>
            <Field>
                <FieldLabel for="name">Nama</FieldLabel>
                <Input
                    id="name"
                    type="text"
                    v-model="name"
                    placeholder="Fadhil"
                    required
                />
                <span v-if="errors.name" class="text-red-500 text-[10px]"
                    >*{{ errors.name }}</span
                >
            </Field>
            <Field>
                <FieldLabel for="email">Email</FieldLabel>
                <Input
                    id="email"
                    type="email"
                    v-model="email"
                    placeholder="m@example.com"
                    required
                />
                <span v-if="errors.email" class="text-red-500 text-[10px]"
                    >*{{ errors.email }}</span
                >
            </Field>
            <Field>
                <FieldLabel for="password">Password</FieldLabel>
                <div class="relative">
                    <Input
                        id="password"
                        :type="showPassword ? 'text' : 'password'"
                        v-model="password"
                        placeholder="*****"
                        required
                    />
                    <button
                        type="button"
                        class="absolute right-3 top-1/2 -translate-y-1/2"
                        @click="showPassword = !showPassword"
                    >
                        <EyeOff
                            v-if="showPassword"
                            class="size-4 text-muted-foreground"
                        />
                        <Eye v-else class="size-4 text-muted-foreground" />
                    </button>
                </div>
                <span v-if="errors.password" class="text-red-500 text-[10px]"
                    >*{{ errors.password }}</span
                >
            </Field>
            <Field>
                <Button
                    type="submit"
                    :disabled="isPending"
                    @click.prevent="handleRegister"
                >
                    {{ isPending ? "Loading..." : "Daftar" }}
                </Button>
            </Field>
            <Field>
                <FieldDescription class="text-center">
                    Sudah memiliki akun?
                    <NuxtLink to="/login">Masuk</NuxtLink>
                </FieldDescription>
            </Field>
        </FieldGroup>
    </form>
</template>
