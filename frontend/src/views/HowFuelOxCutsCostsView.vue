<script setup lang="ts">
import CostBenefitCalculator from '@/components/CostBenefitCalculator.vue';
import { useCurrentPageStore } from '@/stores/current-page';
import Button from 'primevue/button'
import { useRouter } from 'vue-router';

const router = useRouter()

const currentPageStore = useCurrentPageStore()
currentPageStore.setCurrentPage('how-fo-cuts-costs')
</script>



<template>
  <div class="page-container">

    <div class="text-container">
      <p>{{ $t('pages.how-fo-cuts-costs.p1') }}</p>
      <p>{{ $t('pages.how-fo-cuts-costs.p2') }}</p>
      <ol>
        <li>{{ $t('pages.how-fo-cuts-costs.p3') }}</li>
        <li>{{ $t('pages.how-fo-cuts-costs.p4') }}</li>
      </ol>
      <p>{{ $t('pages.how-fo-cuts-costs.p5') }}</p>
      <p>{{ $t('pages.how-fo-cuts-costs.p6_1') }}<a target="_blank"
          :href="$t('links.documents.land-cruiser-22pc-report')">{{
            $t('pages.how-fo-cuts-costs.p6_2') }}</a></p>
      <p>{{ $t('pages.how-fo-cuts-costs.p7') }}</p>
      <p>{{ $t('pages.how-fo-cuts-costs.p8') }}</p>
      <p>{{ $t('pages.how-fo-cuts-costs.p9_1') }}<a target="_blank" :href="$t('links.documents.fo-white-paper')">{{
        $t('pages.how-fo-cuts-costs.p9_2') }}</a></p>
    </div>

    <div class="cost-calculator-container">
      <CostBenefitCalculator class="cost-benefit-calculator" />
      <Button @click="router.push('/free-trial-procedure')" class="primary-button"
        :label="$t('pages.how-fo-cuts-costs.explore-a-free-trial-button')" severity="warning" />
    </div>
  </div>
</template>

<style scoped>
.page-container {
  gap: 1rem;
  display: block;
}

@media (min-width: 1080px) {
  .page-container {
    display: flex;
  }
}

p {
  font-size: 1rem;
  margin-top: 2rem;
}

.learn-more-button {
  background-color: gold;
  border: none !important;
  width: fit-content;
}

.learn-more-button:hover {
  background-color: yellow !important;
  transition: 0.2s;
}

.cost-benefit-calculator {
  margin-bottom: 1rem;
  border-radius: 5px;
  border: solid;
}

.cost-calculator-container {
  margin: 1rem, 0rem;
  justify-content: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  min-width: 45%;
}
</style>
