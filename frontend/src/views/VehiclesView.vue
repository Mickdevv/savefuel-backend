<script setup lang="ts">
import { useCurrentPageStore } from '@/stores/current-page';
import Button from 'primevue/button'
import { useRouter } from 'vue-router';

const router = useRouter()
const currentPageStore = useCurrentPageStore()
currentPageStore.setCurrentPage('vehicles')
</script>
<template>
  <div class="page-container">
    <div class="paragraphs-column">
      <div>

        <h2>{{ $t('pages.vehicles.direct-cost-benefit.title') }}</h2>
        <p>{{ $t('pages.vehicles.direct-cost-benefit.p1') }}</p>
        <p>{{ $t('pages.vehicles.direct-cost-benefit.p2') }}</p>
        <h2>{{ $t('pages.vehicles.other-benefits.title') }}</h2>
        <p>{{ $t('pages.vehicles.other-benefits.p1_1') }} <a target="_blank"
            :href="$t('links.documents.fuel-ox-white-paper')">{{
              $t('pages.vehicles.other-benefits.p1_2') }}</a>
          {{ $t('pages.vehicles.other-benefits.p1_3') }}</p>
        <ul class="custom-list">
          <li v-for="(p, i) in $tm('pages.vehicles.other-benefits.bullet-points')">
            {{ p }}
          </li>
        </ul>
      </div>
      <div class="documents-container">
        <div class="left-column">
          <h2>{{ $t('pages.vehicles.trial-reports.title') }}</h2>
          <ul>
            <li>{{ $t('pages.vehicles.trial-reports.t1_1') }} <a target="_blank"
                :href="$t('links.documents.laguilhon-trial')">{{
                  $t('pages.vehicles.trial-reports.t1_2')
                }}</a>
            </li>
            <li>{{ $t('pages.vehicles.trial-reports.t2_1') }} <a target="_blank" :href="$t('links.documents.')">{{
              $t('pages.vehicles.trial-reports.t2_2')
                }}</a>
            </li>
            <li>{{ $t('pages.vehicles.trial-reports.t3_1') }} <a target="_blank"
                :href="$t('links.documents.land-cruiser-22pc-report')">{{ $t('pages.vehicles.trial-reports.t3_2')
                }}</a>
            </li>
            <li>{{ $t('pages.vehicles.trial-reports.t4_1') }} <a target="_blank" href="">{{
              $t('pages.vehicles.trial-reports.t4_2')
                }}</a>
            </li>
            <li>{{ $t('pages.vehicles.trial-reports.t5_1') }} <a target="_blank" href="">{{
              $t('pages.vehicles.trial-reports.t5_2')
                }}</a>
            </li>
          </ul>
        </div>
        <div class="right-column">
          <h2>{{ $t('pages.vehicles.other-vehicle-documents.title') }}</h2>
          <ul>
            <li>{{ $t('pages.vehicles.other-vehicle-documents.t1_1') }} <a
                :href="$t('documents.links.fuel-ox-white-paper')" target="_blank">{{
                  $t('pages.vehicles.other-vehicle-documents.t1_2') }}</a></li>
            <li>{{ $t('pages.vehicles.other-vehicle-documents.t2') }}</li>
            <li>{{ $t('pages.vehicles.other-vehicle-documents.t3') }}</li>
            <li>{{ $t('pages.vehicles.other-vehicle-documents.t4') }}</li>
            <li>{{ $t('pages.vehicles.other-vehicle-documents.t5_1') }} <a :href="$t('documents.links.')">{{
              $t('pages.vehicles.other-vehicle-documents.t5_2') }}</a></li>
            <li>{{ $t('pages.vehicles.other-vehicle-documents.t6_1') }} <a :href="$t('documents.links.')">{{
              $t('pages.vehicles.other-vehicle-documents.t6_2') }}</a></li>
          </ul>
        </div>
      </div>
    </div>
    <div class="images-column-container">
      <div class="images">
        <img src="@/assets/shutterstock_Trucks_cropped.jpg" alt="">
        <img src="@/assets/LAND_ROVER_DISCOVERY.png" alt="">
        <img src="@/assets/Laguilhon_van.png" alt="">
      </div>
      <Button class="primary-button" @click="router.push('/four-guarantees')"
        :label="$t('pages.vehicles.four-guarantees-button')"></Button>
    </div>
  </div>
</template>
<style scoped>
.page-container {
  display: flex;
  align-items: flex-start;
  flex-wrap: wrap;
}

.images {
  display: none;
}

img {
  width: 100%;
  height: auto;
  margin-bottom: 1rem;
}

.documents-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  min-width: 200px
}

.custom-list {
  list-style: none;
  padding-left: 0;
}

.paragraphs-column {
  width: 100%
}

@media (min-width: 1080px) {
  .paragraphs-column {
    max-width: 75%
  }

  .images-column-container {
    max-width: 25%;
    display: flex;
    justify-content: center;
    flex-direction: column;
  }

  .images {
    display: block;
  }
}

Button {
  background-color: gold;
  border: none !important;
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
