<script setup lang="ts">
import { useCurrentPageStore } from '@/stores/current-page';
import Button from 'primevue/button'
import { useRouter } from 'vue-router';

const router = useRouter()
const currentPageStore = useCurrentPageStore()
currentPageStore.setCurrentPage('four-guarantees')
</script>
<template>
  <div class="page-container">
    <div class="left-container">
      <h2>{{ $t('pages.four-guarantees.manufacturers-guarantee.title') }}</h2>
      <p>{{ $t('pages.four-guarantees.manufacturers-guarantee.p1_1') }} <a href="">{{
        $t('pages.four-guarantees.manufacturers-guarantee.p1_2') }}</a></p>

      <h2>{{ $t('pages.four-guarantees.EN590-ASTM-950.title') }}</h2>
      <p>{{ $t('pages.four-guarantees.EN590-ASTM-950.p1_1') }} <a href="">{{
        $t('pages.four-guarantees.EN590-ASTM-950.p1_2') }}</a></p>
      <p>{{ $t('pages.four-guarantees.EN590-ASTM-950.p2') }}</p>

      <h2>{{ $t('pages.four-guarantees.save-fuel-europe-guarantee.title') }}</h2>
      <p>{{ $t('pages.four-guarantees.save-fuel-europe-guarantee.p1') }} </p>

      <h2>{{ $t('pages.four-guarantees.historical-guarantee.title') }}</h2>
      <p>{{ $t('pages.four-guarantees.historical-guarantee.p1') }} </p>

      <Button class="primary-button" @click="router.push('/technical')"
        :label="$t('pages.four-guarantees.technical-button')"></Button>
    </div>
    <div class="right-container">
      <div class="img-container">
        <img src="@/assets/Fuel_Ox_logo.png" alt="">
        <a target="blank" href="https://fuelox.com">fuelox.com</a>
      </div>
      <div class="img-container">
        <img src="@/assets/Intertek_logo.png" alt="">
        <a target="blank" href="https://intertek-france.com">intertek-france.com</a>
      </div>
      <div class="img-container">
        <img src="@/assets/SFE_Logo.png" alt="">
      </div>
    </div>
  </div>
</template>
<style scoped>
.page-container {
  display: block;
}

.primary-button {
  margin: 1rem 0
}

.img-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

img {
  object-fit: contain;
  max-width: 100%;
}

.right-container .left-container {
  flex-direction: row;
  width: 100%;
}

.right-container {
  display: flex;
  justify-content: space-around;
}

@media (min-width: 1080px) {
  .page-container {
    display: flex;
  }

  .right-container {
    max-width: 15%;
    flex-direction: column;
    justify-content: space-around;
    flex-wrap: wrap;
  }

  .left-container {
    max-width: 85%;
  }
}
</style>
