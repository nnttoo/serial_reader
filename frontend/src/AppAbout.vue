<script setup lang="ts">
import { onMounted } from 'vue';


export type AppAboutProp = {
    aboutClose : ()=>void
}
const prop = defineProps<AppAboutProp>(); 
</script>
<template>
    <div class="frame">
        <div class="ctn">
            
            <h1>Serial Reader v1</h1>
            <p>
                Github link <a href="https://github.com/nnttoo/serial_reader">Github</a>
            </p>
            
        </div>
        <div style="text-align: center;">
            <button class="btnclose" v-on:click="prop.aboutClose()">Close</button>
        </div>
        
    </div>
</template>
<style scoped>
.frame{
    display: flex;
    flex-direction: column;
    height: 100%;
}
.ctn{
    flex: 1;
}
.btnclose{ 
  background-color: #0525de;
  border: none; 
  min-width: 100px;
  width: 100%;
  border-radius: 5px;
  cursor: pointer;
  color: white;
  padding: 5px;
}
.btnclose:active { 
  transform: scale(0.97);
}
</style>