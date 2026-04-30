<script setup lang="ts">
import type { NavBarItem } from '@/types/navbar-item';
import { computed, ref } from 'vue';
import PanelMenu from 'primevue/panelmenu';
import 'primeicons/primeicons.css'
import Menu from 'primevue/menu';
import Button from 'primevue/button';
import Select from 'primevue/select';
import { useLocaleStore } from '@/stores/selected-language';
import { LocaleOptions } from '@/types/enums/locales.enum';
import { useCurrentPageStore } from '@/stores/current-page';
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';

const localeStore = useLocaleStore()
const currentPageStore = useCurrentPageStore()
const menu = ref();

const { t } = useI18n()

const router = useRouter()

const items = computed(() => [
  {
    label: t('menubar.home'),
    // icon: 'pi pi-eraser',
    command: () => router.push('/')
  },

  {
    label: t('menubar.how-fo-cuts-costs'),
    // icon: 'pi pi-eraser',
    command: () => router.push('/how-fuel-ox-cuts-costs')
  },

  {
    label: t('menubar.free-trial-procedure'),
    // icon: 'pi pi-eraser',
    command: () => router.push('/free-trial-procedure')
  },

  {
    label: t('menubar.four-guarantees'),
    // icon: 'pi pi-eraser',
    command: () => router.push('/four-guarantees')
  },

  {
    label: t('menubar.technical'),
    // icon: 'pi pi-eraser',
    command: () => router.push('/technical')
  },

  {
    label: t('menubar.vehicles'),
    // icon: 'pi pi-eraser',
    command: () => router.push('/vehicles')
  },

  {
    label: t('menubar.generators'),
    // icon: 'pi pi-eraser',
    command: () => router.push('/generators')
  },
  {
    label: t('menubar.about'),
    // icon: 'pi pi-eraser',
    command: () => router.push('/about')
  },
  {
    separator: true
  },

  {
    label: t('menubar.price-list'),
    // icon: 'pi pi-eraser',
    command: () => window.open(t('links.documents.price-list'))
  },

  {
    label: t('menubar.legal-details'),
    // icon: 'pi pi-heart',
    command: () => window.open(t('links.documents.legal-details'))
  },
  {
    label: t('menubar.gdpr'),
    // icon: 'pi pi-eraser',
    command: () => window.open(t('links.documents.gdpr'))
  },
  {
    separator: true
  },
  {
    label: 'contact@save-fuel.eu',
    // icon: 'pi pi-eraser',  
    command: () => window.location.href = `mailto:contact@save-fuel.eu${t('links.email.website-enquiry')}`

  },
])
const toggle = (event: any) => {
  menu.value.toggle(event);
};


</script>


<template>
  <div class="page-top-container">
    <div class="menubar">
      <a class="logo-link" href="/">
        <div style="font-weight: bold; display: flex;">
          <img src="../assets/SFE_Logo.png" style="max-width: 3rem;" />
          <p class="company-name">Save Fuel Europe SAS</p>
        </div>
      </a>
      <div>
        <h2>{{ $t(`pages.${currentPageStore.currentPage}.page-title`) }}</h2>
      </div>
      <div class="links">
        <span class="icons">
          <a target="_blank" class="sm-link linkedin-link" href="https://www.linkedin.com/company/save-fuel-europe/?viewAsMember=true ">
            <i class="pi pi-linkedin"></i>
          </a>
          <a target="_blank" class="sm-link youtube-link" href="https://www.youtube.com/@SaveFuelEurope">
            <i class="pi pi-youtube"></i>
          </a>
        </span>
        <div class="border border-right menu">
          <Button class="menu-internal-element" type="button" label="Menu" icon="pi pi-ellipsis-v" @click="toggle"
            aria-haspopup="true" aria-controls="overlay_menu" />
          <Menu class="menu-internal-element" ref="menu" id="overlay_menu" :model="items" :popup="true" />
        </div>
        <div class="language-selector">
          <Select class="language-selector-select" v-model="$i18n.locale" :options="$i18n.availableLocales"
            @change="localeStore.selectLanguage($i18n.locale)" />
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.company-name {
  display: none;
}


.youtube-link {
  display: none;
  color: hsla(10, 100%, 37%, 1);
}

.linkedin-link {
  display: none;
  color: hsla(195, 100%, 37%, 1);
}

@media (min-width: 1000px) {

  .youtube-link,
  .linkedin-link {
    display: inline;
  }

  .company-name {
    display: flex;
  }
}

.page-top-container {
  background: linear-gradient(to bottom,
      rgba(44, 57, 245, 0.2) 0%,
      /* rgba(44, 57, 205, 0.5) 50%, */
      rgba(44, 57, 245, 0) 100%);
}

.language-selector-select:focus {
  padding: 0rem 0.5rem;
  border-color: black !important;

}



.language-selector {
  padding: 0rem 0.5rem;
  border-color: black !important;
}

.sm-link {
  border-radius: 5px;
  margin: 0.5rem
}

.menu-internal-element {
  margin: 0.2rem !important;
  padding: 0.4rem !important;
  background-color: gold !important;

  border: none !important;
}

.menu-internal-element:focus {
  background-color: gold !important;
  border: none !important;
}

.menubar {
  max-width: 1280px;
  margin: 0 auto;
  align-items: center;
  display: flex;
  justify-content: space-between;
  padding: 1rem;
}

.links {
  display: flex;
  align-items: center;
}

.active {
  background-color: white;
}

.border-left {
  border-left: solid;
  padding: 0rem 0.5rem;
}

.border-right {
  border-right: solid;
  padding: 0rem 0.5rem;
}


.logo-link {
  text-decoration: none;
  color: inherit;
  transition: 0.2s;
}

.logo-link:hover {
  background: none;
}

.menubar-item {
  margin: 1rem;
}

.border {
  border-width: 1px;
}
</style>
