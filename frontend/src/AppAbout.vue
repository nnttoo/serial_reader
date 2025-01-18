<script setup lang="ts">
import { onMounted } from 'vue'; 
import { GetSerialNames, IsConnected, OpenAbout, PortClose, PortOpen } from '../wailsjs/go/main/App';

 
export type AppAboutProp = {
    aboutClose : ()=>void
}
const prop = defineProps<AppAboutProp>(); 
</script>
<template>
    <div class="frame">
        <div class="ctn">
            
            <h1>Serial Reader v3</h1>
            <p>
                Serial Reader Windows is an open-source application that allows users to read data from serial ports on Windows with a more attractive and user-friendly interface. This application simplifies interacting with devices connected via serial ports, such as microcontrollers, sensors, and other hardware.
            </p>
            <p>
                Haryanto <br/>
                Github link (nnttoo) <a href="#" v-on:click="OpenAbout()">Github</a>
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