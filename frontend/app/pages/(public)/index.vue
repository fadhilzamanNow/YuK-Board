<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { GalleryVerticalEnd, CheckCircle, Users, ListTodo } from "lucide-vue-next";
import landingPic from "@/assets/image/landingpagepic.jpg";
</script>

<template>
  <div class="min-h-svh flex flex-col">
    <!-- Navbar -->
    <header class="border-b">
      <div class="container mx-auto flex items-center justify-between p-4">
        <a href="/" class="flex items-center gap-2 font-bold text-lg">
          <div class="bg-primary text-primary-foreground flex size-8 items-center justify-center rounded-md">
            <GalleryVerticalEnd class="size-5" />
          </div>
          YukBoard
        </a>
        <div class="flex gap-2">
          <ClientOnly>
            <ModeToggle />
          </ClientOnly>
          <Button variant="ghost" as-child>
            <NuxtLink to="/login">Masuk</NuxtLink>
          </Button>
          <Button as-child>
            <NuxtLink to="/register">Daftar</NuxtLink>
          </Button>
        </div>
      </div>
    </header>

    <!-- Hero -->
    <main class="flex-1">
      <section class="container mx-auto px-4 py-20 text-center">
        <h1 class="text-4xl md:text-6xl font-bold tracking-tight mb-6 animate-fade-in-up">
          Kolaborasi Tim Jadi
          <span class="text-primary">Lebih Mudah</span>
        </h1>
        <p class="text-muted-foreground text-lg md:text-xl max-w-2xl mx-auto mb-8 animate-fade-in-up animation-delay-100">
          Kelola tugas bersama tim dengan mudah. Buat daftar, undang anggota, dan selesaikan pekerjaan bersama-sama.
        </p>
        <div class="flex gap-4 justify-center animate-fade-in-up animation-delay-200">
          <Button size="lg" as-child>
            <NuxtLink to="/register">Mulai Gratis</NuxtLink>
          </Button>
          <Button size="lg" variant="outline" as-child>
            <NuxtLink to="/login">Masuk</NuxtLink>
          </Button>
        </div>
      </section>

      <div class="relative animate-fade-in-up animation-delay-300">
        <div class="absolute inset-0 top-1/2 bg-primary"></div>
        <div class="relative container mx-auto px-4">
          <img :src="landingPic" alt="YukBoard Preview" class="rounded-lg shadow-2xl mx-auto max-w-4xl w-full" />
        </div>
      </div>
      <div class="bg-primary h-20"></div>

      <!-- Features -->
      <section class="container mx-auto px-4 py-16">
        <h2 class="text-2xl md:text-3xl font-bold text-center mb-12">Fitur Unggulan</h2>

        <div class="relative max-w-4xl mx-auto">
          <!-- Vertical line -->
          <div class="absolute left-1/2 -translate-x-1/2 top-0 bottom-0 w-0.5 bg-primary/20 hidden md:block"></div>

          <!-- Feature 1 -->
          <div class="relative flex flex-col md:flex-row items-center gap-6 pb-32 animate-on-scroll">
            <div class="md:w-1/2 md:text-right md:pr-12">
              <h3 class="font-semibold text-xl mb-2">Buat Daftar Tugas</h3>
              <p class="text-muted-foreground">Organisir pekerjaan dengan daftar tugas yang rapi dan terstruktur.</p>
            </div>
            <div class="absolute left-1/2 -translate-x-1/2 bg-primary text-primary-foreground w-12 h-12 rounded-full flex items-center justify-center z-10 hidden md:flex">
              <ListTodo class="size-6" />
            </div>
            <div class="md:w-1/2 md:pl-12"></div>
          </div>

          <!-- Feature 2 -->
          <div class="relative flex flex-col md:flex-row items-center gap-6 pb-32 animate-on-scroll">
            <div class="md:w-1/2 md:pr-12"></div>
            <div class="absolute left-1/2 -translate-x-1/2 bg-primary text-primary-foreground w-12 h-12 rounded-full flex items-center justify-center z-10 hidden md:flex">
              <Users class="size-6" />
            </div>
            <div class="md:w-1/2 md:pl-12">
              <h3 class="font-semibold text-xl mb-2">Undang Tim</h3>
              <p class="text-muted-foreground">Undang anggota tim via email dan kolaborasi secara real-time.</p>
            </div>
          </div>

          <!-- Feature 3 -->
          <div class="relative flex flex-col md:flex-row items-center gap-6 animate-on-scroll">
            <div class="md:w-1/2 md:text-right md:pr-12">
              <h3 class="font-semibold text-xl mb-2">Lacak Progress</h3>
              <p class="text-muted-foreground">Pantau status tugas dari todo hingga selesai dengan mudah.</p>
            </div>
            <div class="absolute left-1/2 -translate-x-1/2 bg-primary text-primary-foreground w-12 h-12 rounded-full flex items-center justify-center z-10 hidden md:flex">
              <CheckCircle class="size-6" />
            </div>
            <div class="md:w-1/2 md:pl-12"></div>
          </div>
        </div>
      </section>
    </main>
  </div>
</template>

<style scoped>
@keyframes fade-in-up {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fade-in-up {
  animation: fade-in-up 0.6s ease-out forwards;
}

.animation-delay-100 { animation-delay: 0.1s; opacity: 0; }
.animation-delay-200 { animation-delay: 0.2s; opacity: 0; }
.animation-delay-300 { animation-delay: 0.3s; opacity: 0; }

.animate-on-scroll {
  opacity: 0;
  transform: translateY(30px);
  animation: fade-in-up 0.6s ease-out forwards;
  animation-timeline: view();
  animation-range: entry 0% entry 40%;
}
</style>
