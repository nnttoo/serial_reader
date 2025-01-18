<script setup lang="ts">
import { ref } from 'vue';
import AppMain from './AppMain.vue';
import AppAbout from './AppAbout.vue';


const pageType = ref<string>("main")

</script>

<template>
    <AppMain :openAbout="()=>{ pageType='about'}" v-if="pageType =='main'"/>
    <AppAbout :aboutClose="()=>{ pageType='main'}" v-else />
</template>