<script setup lang="ts">
import Button from 'primevue/button'
import { useCurrentPageStore } from '@/stores/current-page';
import { useRouter } from 'vue-router';

const router = useRouter()
const currentPageStore = useCurrentPageStore()
currentPageStore.setCurrentPage("technical")

</script>

<template>
  <div class="page-container">

    <div class="paragraphs">
      <h2>{{ $t('pages.technical.history.title') }}</h2>
      <p>{{ $t('pages.technical.history.p1') }}</p>
      <p>{{ $t('pages.technical.history.p2') }}</p>
      <p>{{ $t('pages.technical.history.p3') }}</p>
      <p>{{ $t('pages.technical.history.p4') }}</p>
      <h2>{{ $t('pages.technical.technical-documents.title') }}</h2>
      <p> <a target="_blank" :href="$t('links.documents.product-information')"> {{
        $t('pages.technical.technical-documents.p1_1') }}</a> - {{
            $t('pages.technical.technical-documents.p1_2') }}</p>
      <p> <a target="_blank" :href="$t('links.documents.fo-white-paper')"> {{
        $t('pages.technical.technical-documents.p2_1') }}</a> - {{
            $t('pages.technical.technical-documents.p2_2') }}</p>
      <p> <a target="_blank" :href="$t('links.documents.faq')"> {{ $t('pages.technical.technical-documents.p3_1') }}</a>
        - {{
          $t('pages.technical.technical-documents.p3_2') }} <a href="mailto:contact@save-fuel.eu">contact@save-fuel.eu</a>
      </p>
      <p> <a target="_blank" :href="$t('links.documents.gnr-compared-to-EN590')"> {{
        $t('pages.technical.technical-documents.p4_1') }}</a> - {{
            $t('pages.technical.technical-documents.p4_2') }}</p>
      <p> <a target="_blank" :href="$t('links.documents.fo-MSDS')"> {{ $t('pages.technical.technical-documents.p5_1')
      }}</a> - {{
            $t('pages.technical.technical-documents.p5_2') }}</p>
      <p> <a target="_blank" :href="$t('links.documents.intertek-EN590-test')"> {{
        $t('pages.technical.technical-documents.p6_1') }}</a> - {{
            $t('pages.technical.technical-documents.p6_2') }}</p>
      <p> <a target="_blank" :href="$t('links.documents.ASTM-D975-test')"> {{
        $t('pages.technical.technical-documents.p7_1') }}</a> - {{
            $t('pages.technical.technical-documents.p7_2') }}</p>
      <p> <a target="_blank" :href="$t('links.documents.fo-guarantee')"> {{
        $t('pages.technical.technical-documents.p8_1') }}</a> - {{
            $t('pages.technical.technical-documents.p8_2') }}</p>
    </div>
    <div class="right-column">
      <div class="images">

        <img src="../assets/Fuel_molecules_polymerised.png" style="border-radius: 5px; width: 100%" alt="">
        <img src="../assets/Fuel_molecules_dispersed.png" style="border-radius: 5px; width: 100%" alt="">
      </div>
      <div class="buttons-bar">
        <Button class="primary-button" @click="router.push('/about')"
          :label="$t('pages.technical.about-button')"></Button>
      </div>
    </div>

  </div>
</template>

<style scoped>
.right-column {
  margin-left: 1rem;
}

.buttons-bar {
  display: flex;
  justify-content: center;
}

.page-container {
  display: flex;
  flex-wrap: wrap;
  align-items: flex-start;
}

.paragraphs {
  width: 100%;
}

.images {
  display: none;
}

.primary-button {
  margin: 1rem;
}

@media (min-width: 1080px) {
  .images {
    display: block;
  }

  .right-column {
    display: flex;
    flex-direction: column;
    justify-content: center;
    margin-left: 1rem;
    width: 25%;
  }

  .buttons-bar {
    display: flex;
    justify-content: center;
  }

  .page-container {
    display: flex;
    flex-wrap: nowrap;
  }

  .paragraphs {
    max-width: 75%;
  }

  .primary-button {
    margin: 1rem;
  }
}
</style>
