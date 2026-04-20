<script setup lang="ts">
import { useCurrentPageStore } from '@/stores/current-page';
import Button from 'primevue/button'
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';

const router = useRouter()
const currentPageStore = useCurrentPageStore()

const { tm, t } = useI18n()
currentPageStore.setCurrentPage('home')
</script>



<template>
  <p v-for="(p, i) in tm('pages.home.paragraphs')">{{ p }}</p>
  <div class="button-container">
    <Button @click="router.push('/how-fuel-ox-cuts-costs')" class="primary-button"
      :label="$t('pages.home.learn-more-button')" severity="warning" />
  </div>
</template>

<style scoped>
p {
  font-size: 2rem;
  margin-top: 4rem;
  font-weight: bold;
}

.learn-more-button {
  background-color: gold;
  border: none !important;
}

.learn-more-button:hover {
  background-color: yellow !important;
  transition: 0.2s;
}


.button-container {
  margin-top: 10rem;
  display: flex;
  justify-content: center;
}
</style>
