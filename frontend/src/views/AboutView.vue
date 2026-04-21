<script setup lang="ts">
import { useCurrentPageStore } from '@/stores/current-page';
import Button from 'primevue/button'
import { useRouter } from 'vue-router';

const router = useRouter()
const currentPageStore = useCurrentPageStore()
currentPageStore.setCurrentPage('about')
</script>
<template>
  <div class="page-container">
    <h2>{{ $t('pages.about.title') }}</h2>
    <p>{{ $t('pages.about.p1') }}</p>
    <p>{{ $t('pages.about.p2') }}</p>
    <p>{{ $t('pages.about.p3') }}</p>
    <p>{{ $t('pages.about.p4_1') }} <a target="_blank" :href="$t('links.documents.trial-procedure')">{{
      $t('pages.about.p4_2') }}</a> {{ $t('pages.about.p4_3') }} <a target="_blank"
        :href="$t('links.documents.land-cruiser-22pc-report')">{{ $t('pages.about.p4_4') }}</a> {{
          $t('pages.about.p4_5') }} </p>
    <p>{{ $t('pages.about.p5_1') }} <a target="_blank" :href="$t('links.documents.laguilhon-trial')">{{
      $t('pages.about.p5_2') }}</a> {{ $t('pages.about.p5_3') }} </p>
    <p>{{ $t('pages.about.p6_1') }} <a target="_blank" :href="$t('links.documents.land-rover-trial')">{{
      $t('pages.about.p6_2') }}</a> {{ $t('pages.about.p6_3') }} </p>
    <p>{{ $t('pages.about.p7') }}</p>
    <p>{{ $t('pages.about.p8_1') }} <a target="_blank" :href="$t('links.documents.intertek-EN590-test')">{{
      $t('pages.about.p8_2') }}</a> {{ $t('pages.about.p8_3') }} </p>
    <p>{{ $t('pages.about.p9') }}</p>
    <p>{{ $t('pages.about.p10') }}</p>
  </div>
</template>
<style scoped>
.page-container {}

.right-container .left-container {
  width: 100%;
}

@media (min-width: 1080px) {
  .page-container {}

  .right-container {
    max-width: 25%;
  }

  .left-container {
    max-width: 75%;
  }
}
</style>
