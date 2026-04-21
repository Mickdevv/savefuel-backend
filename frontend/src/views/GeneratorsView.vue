<script setup lang="ts">
import { useCurrentPageStore } from '@/stores/current-page';
import Button from 'primevue/button'
import { useRouter } from 'vue-router';

const router = useRouter()

const currentPageStore = useCurrentPageStore()
currentPageStore.setCurrentPage('generators')
</script>
<template>
  <div class="page-container">
    <div class="left-container">
      <h2>{{ $t('pages.generators.direct-cost-benefit.title') }}</h2>
      <p>{{ $t('pages.generators.direct-cost-benefit.p1') }}</p>
      <p>{{ $t('pages.generators.direct-cost-benefit.p2') }}</p>
      <p>{{ $t('pages.generators.direct-cost-benefit.p3') }}</p>

      <h2>{{ $t('pages.generators.other-benefits.title') }}</h2>
      <p>{{ $t('pages.generators.other-benefits.p1_1') }}
        <a target="_blank" :href="$t('links.documents.fo-white-paper')">{{ $t('pages.generators.other-benefits.p1_2')
          }}</a>
        {{ $t('pages.generators.other-benefits.p1_3') }}
      </p>
      <ul class="custom-list">
        <li v-for="(p, i) in $tm('pages.generators.other-benefits.bullet-points')">{{ p }}</li>
      </ul>
      <h2>{{ $t('pages.generators.trial-reports.title') }}</h2>
      <p>
        <a target="_blank" :href="$t('links.documents.peruvian-mine')">{{ $t('pages.generators.trial-reports.p1')
          }}</a>, <a target="_blank" :href="$t('links.documents.')">{{
            $t('pages.generators.trial-reports.p2') }}</a>
      </p>
    </div>
    <div class="right-container">
      <div class="images">
        <img src="@/assets/7_kW_generator.png" alt="">
        <img src="@/assets/Generator_photo.png" alt="">

      </div>
      <Button @click="router.push('/four-guarantees')" class="primary-button"
        :label="$t('pages.generators.four-guarantees-button')"></Button>
    </div>
  </div>
</template>
<style scoped>
.page-container {
  display: block;
}

.images {
  display: none;
}

img {
  width: 100%;
  height: auto;
  margin-bottom: 1rem;
}

.right-container .left-container {
  width: 100%;
}

@media (min-width: 1080px) {
  .images {
    display: block;
  }

  .page-container {
    display: flex;
  }

  .right-container {
    max-width: 25%;
  }

  .left-container {
    max-width: 75%;
  }
}

.custom-list {
  list-style: none;
  padding-left: 0;
}

.custom-list li {
  position: relative;
  padding-left: 2rem;
  /* space for logo */
  margin-bottom: 0.5rem;
}

.custom-list li::before {
  content: "";
  padding: 0;
  position: absolute;
  left: 0;

  width: 1.5rem;
  height: 1.5rem;

  background-image: url('@/assets/SFE_Logo.png');
  /* ✅ your logo */
  background-size: contain;
  background-repeat: no-repeat;
}
</style>
